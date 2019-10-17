package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	casbinProto "github.com/paysuper/casbin-server/pkg/generated/api/proto/casbinpb"
	"github.com/paysuper/paysuper-billing-server/pkg"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/billing"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/grpc"
	postmarkSdrPkg "github.com/paysuper/postmark-sender/pkg"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

const (
	defaultCompanyName = "PaySuper"

	claimType   = "type"
	claimEmail  = "email"
	claimRoleId = "role_id"
	claimExpire = "exp"

	casbinMerchantUserMask = "%s_%s"

	roleNameMerchantOwner      = "Owner"
	roleNameMerchantDeveloper  = "Developer"
	roleNameMerchantAccounting = "Accounting"
	roleNameMerchantSupport    = "Support"
	roleNameMerchantViewOnly   = "View only"
	roleNameSystemAdmin        = "Admin"
	roleNameSystemRiskManager  = "Risk manager"
	roleNameSystemFinancial    = "Financial"
	roleNameSystemSupport      = "Support"
	roleNameSystemViewOnly     = "View only"
)

var (
	usersDbInternalError              = newBillingServerErrorMsg("uu000001", "unknown database error")
	errorUserAlreadyExist             = newBillingServerErrorMsg("uu000002", "user already exist")
	errorUserUnableToAdd              = newBillingServerErrorMsg("uu000003", "unable to add user")
	errorUserNotFound                 = newBillingServerErrorMsg("uu000004", "user not found")
	errorUserInviteAlreadyAccepted    = newBillingServerErrorMsg("uu000005", "user already accepted invite")
	errorUserMerchantNotFound         = newBillingServerErrorMsg("uu000006", "merchant not found")
	errorUserUnableToSendInvite       = newBillingServerErrorMsg("uu000007", "unable to send invite email")
	errorUserNoHavePermission         = newBillingServerErrorMsg("uu000008", "user not have permission to invite users")
	errorUserUnableToCreateToken      = newBillingServerErrorMsg("uu000009", "unable to create invite token")
	errorUserInvalidToken             = newBillingServerErrorMsg("uu000010", "invalid token string")
	errorUserInvalidInviteEmail       = newBillingServerErrorMsg("uu000011", "email in request and token are not equal")
	errorUserAlreadyHasRole           = newBillingServerErrorMsg("uu000012", "user already has role")
	errorUserUnableToSave             = newBillingServerErrorMsg("uu000013", "can't update user in db")
	errorUserUnableToAddToCasbin      = newBillingServerErrorMsg("uu000014", "unable to add user to the casbin server")
	errorUserUnsupportedRoleType      = newBillingServerErrorMsg("uu000015", "unsupported role type")
	errorUserUnableToDelete           = newBillingServerErrorMsg("uu000016", "unable to delete user")
	errorUserUnableToDeleteFromCasbin = newBillingServerErrorMsg("uu000017", "unable to delete user from the casbin server")

	merchantUserRoles = map[string][]*billing.RoleListItem{
		pkg.RoleTypeMerchant: {
			{Id: pkg.RoleMerchantOwner, Name: roleNameMerchantOwner},
			{Id: pkg.RoleMerchantOwner, Name: roleNameMerchantDeveloper},
			{Id: pkg.RoleMerchantOwner, Name: roleNameMerchantAccounting},
			{Id: pkg.RoleMerchantOwner, Name: roleNameMerchantSupport},
			{Id: pkg.RoleMerchantOwner, Name: roleNameMerchantSupport},
			{Id: pkg.RoleMerchantOwner, Name: roleNameMerchantViewOnly},
		},
		pkg.RoleTypeSystem: {
			{Id: pkg.RoleMerchantOwner, Name: roleNameSystemAdmin},
			{Id: pkg.RoleMerchantOwner, Name: roleNameSystemRiskManager},
			{Id: pkg.RoleMerchantOwner, Name: roleNameSystemFinancial},
			{Id: pkg.RoleMerchantOwner, Name: roleNameSystemSupport},
			{Id: pkg.RoleMerchantOwner, Name: roleNameSystemViewOnly},
		},
	}
)

func (s *Service) GetMerchantUsers(ctx context.Context, req *grpc.GetMerchantUsersRequest, res *grpc.GetMerchantUsersResponse) error {
	users, err := s.userRoleRepository.GetUsersForMerchant(req.MerchantId)

	if err != nil {
		res.Status = pkg.ResponseStatusSystemError
		res.Message = usersDbInternalError
		res.Message.Details = err.Error()

		return nil
	}

	res.Status = pkg.ResponseStatusOk
	res.Users = users

	return nil
}

func (s *Service) GetAdminUsers(ctx context.Context, _ *grpc.EmptyRequest, res *grpc.GetAdminUsersResponse) error {
	users, err := s.userRoleRepository.GetUsersForAdmin()

	if err != nil {
		res.Status = pkg.ResponseStatusSystemError
		res.Message = usersDbInternalError
		res.Message.Details = err.Error()

		return nil
	}
	res.Status = pkg.ResponseStatusOk
	res.Users = users

	return nil
}

func (s *Service) GetMerchantsForUser(ctx context.Context, req *grpc.GetMerchantsForUserRequest, res *grpc.GetMerchantsForUserResponse) error {
	users, err := s.userRoleRepository.GetMerchantsForUser(req.UserId)

	if err != nil {
		res.Status = pkg.ResponseStatusSystemError
		res.Message = usersDbInternalError
		res.Message.Details = err.Error()

		return nil
	}

	merchants := make([]*grpc.MerchantForUserInfo, len(users))

	for i, user := range users {
		merchant, err := s.merchant.GetById(user.MerchantId)
		if err != nil {
			zap.L().Error(
				"Can't get merchant by id",
				zap.Error(err),
				zap.String("merchant_id", user.MerchantId),
			)

			res.Status = pkg.ResponseStatusSystemError
			res.Message = usersDbInternalError
			res.Message.Details = err.Error()

			return nil
		}

		name := merchant.Id
		if merchant.Company != nil {
			name = merchant.Company.Name
		}

		merchants[i] = &grpc.MerchantForUserInfo{
			Id:   user.MerchantId,
			Name: name,
		}
	}

	res.Status = pkg.ResponseStatusOk
	res.Merchants = merchants
	return nil
}

func (s *Service) InviteUserMerchant(
	ctx context.Context,
	req *grpc.InviteUserMerchantRequest,
	res *grpc.InviteUserMerchantResponse,
) error {
	owner, err := s.userRoleRepository.GetMerchantUserByUserId(req.MerchantId, req.PerformerId)

	if err != nil || !owner.IsOwner() {
		zap.L().Error(errorUserNoHavePermission.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNoHavePermission

		return nil
	}

	merchant, err := s.merchant.GetById(owner.MerchantId)

	if err != nil {
		zap.L().Error(errorUserMerchantNotFound.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserMerchantNotFound

		return nil
	}

	user, err := s.userRoleRepository.GetMerchantUserByEmail(owner.MerchantId, req.Email)

	if (err != nil && err != mgo.ErrNotFound) || user != nil {
		zap.L().Error(errorUserAlreadyExist.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserAlreadyExist

		return nil
	}

	role := &billing.UserRole{
		Id:         bson.NewObjectId().Hex(),
		MerchantId: req.MerchantId,
		User: &billing.UserRoleProfile{
			Email:  req.Email,
			Status: pkg.UserRoleStatusInvited,
		},
	}

	if err = s.userRoleRepository.AddMerchantUser(role); err != nil {
		zap.L().Error(errorUserUnableToAdd.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToAdd

		return nil
	}

	expire := time.Now().Add(time.Hour * time.Duration(s.cfg.UserInviteTokenTimeout)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		claimType:   pkg.RoleTypeMerchant,
		claimEmail:  req.Email,
		claimRoleId: role.Id,
		claimExpire: expire,
	})
	tokenString, err := token.SignedString([]byte(s.cfg.UserInviteTokenSecret))

	if err != nil {
		zap.L().Error(errorUserUnableToCreateToken.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToCreateToken

		return nil
	}

	if err = s.sendInviteEmail(req.Email, owner.User.Email, owner.User.FirstName, owner.User.LastName, merchant.Company.Name, tokenString); err != nil {
		zap.L().Error(
			errorUserUnableToSendInvite.Message,
			zap.Error(err),
			zap.String("receiverEmail", req.Email),
			zap.String("senderEmail", owner.User.Email),
			zap.String("senderFirstName", owner.User.FirstName),
			zap.String("senderLastName", owner.User.LastName),
			zap.String("senderCompany", merchant.Company.Name),
			zap.String("token", tokenString),
		)
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToSendInvite

		return nil
	}

	res.Status = pkg.ResponseStatusOk
	res.Role = role

	return nil
}

func (s *Service) InviteUserAdmin(
	ctx context.Context,
	req *grpc.InviteUserAdminRequest,
	res *grpc.InviteUserAdminResponse,
) error {
	owner, err := s.userRoleRepository.GetAdminUserByUserId(req.PerformerId)

	if err != nil || !owner.IsAdmin() {
		zap.L().Error(errorUserNoHavePermission.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNoHavePermission

		return nil
	}

	user, err := s.userRoleRepository.GetAdminUserByEmail(req.Email)

	if (err != nil && err != mgo.ErrNotFound) || user != nil {
		zap.L().Error(errorUserAlreadyExist.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserAlreadyExist

		return nil
	}

	role := &billing.UserRole{
		Id:   bson.NewObjectId().Hex(),
		Role: req.Role,
		User: &billing.UserRoleProfile{
			Email:  req.Email,
			Status: pkg.UserRoleStatusInvited,
		},
	}

	if err = s.userRoleRepository.AddAdminUser(role); err != nil {
		zap.L().Error(errorUserUnableToAdd.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToAdd

		return nil
	}

	expire := time.Now().Add(time.Hour * time.Duration(s.cfg.UserInviteTokenTimeout)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		claimType:   pkg.RoleTypeSystem,
		claimEmail:  req.Email,
		claimRoleId: role.Id,
		claimExpire: expire,
	})
	tokenString, err := token.SignedString([]byte(s.cfg.UserInviteTokenSecret))

	if err != nil {
		zap.L().Error(errorUserUnableToCreateToken.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToCreateToken

		return nil
	}

	if err = s.sendInviteEmail(req.Email, owner.User.Email, owner.User.FirstName, owner.User.LastName, defaultCompanyName, tokenString); err != nil {
		zap.L().Error(
			errorUserUnableToSendInvite.Message,
			zap.Error(err),
			zap.String("receiverEmail", req.Email),
			zap.String("senderEmail", owner.User.Email),
			zap.String("senderFirstName", owner.User.FirstName),
			zap.String("senderLastName", owner.User.LastName),
			zap.String("token", tokenString),
		)
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToSendInvite

		return nil
	}

	res.Status = pkg.ResponseStatusOk
	res.Role = role

	return nil
}

func (s *Service) ResendInviteMerchant(
	ctx context.Context,
	req *grpc.ResendInviteMerchantRequest,
	res *grpc.EmptyResponseWithStatus,
) error {
	owner, err := s.userRoleRepository.GetMerchantUserByUserId(req.MerchantId, req.PerformerId)

	if err != nil || !owner.IsOwner() {
		zap.L().Error(errorUserNoHavePermission.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNoHavePermission

		return nil
	}

	merchant, err := s.merchant.GetById(owner.MerchantId)

	if err != nil {
		zap.L().Error(errorUserMerchantNotFound.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserMerchantNotFound

		return nil
	}

	role, err := s.userRoleRepository.GetMerchantUserById(req.RoleId)

	if err != nil {
		zap.L().Error(errorUserNotFound.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNotFound

		return nil
	}

	expire := time.Now().Add(time.Hour * time.Duration(s.cfg.UserInviteTokenTimeout)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		claimType:   pkg.RoleTypeMerchant,
		claimEmail:  role.User.Email,
		claimRoleId: role.Id,
		claimExpire: expire,
	})
	tokenString, err := token.SignedString([]byte(s.cfg.UserInviteTokenSecret))

	if err != nil {
		zap.L().Error(errorUserUnableToCreateToken.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToCreateToken

		return nil
	}

	if err = s.sendInviteEmail(role.User.Email, owner.User.Email, owner.User.FirstName, owner.User.LastName, merchant.Company.Name, tokenString); err != nil {
		zap.L().Error(
			errorUserUnableToSendInvite.Message,
			zap.Error(err),
			zap.String("receiverEmail", role.User.Email),
			zap.String("senderEmail", merchant.User.Email),
			zap.String("senderFirstName", merchant.User.FirstName),
			zap.String("senderLastName", merchant.User.LastName),
			zap.String("senderCompany", merchant.Company.Name),
			zap.String("tokenString", tokenString),
		)
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToSendInvite

		return nil
	}

	res.Status = pkg.ResponseStatusOk

	return nil
}

func (s *Service) ResendInviteAdmin(
	ctx context.Context,
	req *grpc.ResendInviteAdminRequest,
	res *grpc.EmptyResponseWithStatus,
) error {
	owner, err := s.userRoleRepository.GetAdminUserByUserId(req.PerformerId)

	if err != nil || owner.IsAdmin() {
		zap.L().Error(errorUserNoHavePermission.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNoHavePermission

		return nil
	}

	role, err := s.userRoleRepository.GetAdminUserById(req.RoleId)

	if (err != nil && err != mgo.ErrNotFound) || role != nil {
		zap.L().Error(errorUserAlreadyExist.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserAlreadyExist

		return nil
	}

	expire := time.Now().Add(time.Hour * time.Duration(s.cfg.UserInviteTokenTimeout)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		claimType:   pkg.RoleTypeSystem,
		claimEmail:  role.User.Email,
		claimRoleId: role.Id,
		claimExpire: expire,
	})
	tokenString, err := token.SignedString([]byte(s.cfg.UserInviteTokenSecret))

	if err != nil {
		zap.L().Error(errorUserUnableToCreateToken.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToCreateToken

		return nil
	}

	if err = s.sendInviteEmail(role.User.Email, owner.User.Email, owner.User.FirstName, owner.User.LastName, defaultCompanyName, tokenString); err != nil {
		zap.L().Error(
			errorUserUnableToSendInvite.Message,
			zap.Error(err),
			zap.String("receiverEmail", role.User.Email),
			zap.String("senderEmail", owner.User.Email),
			zap.String("senderFirstName", owner.User.FirstName),
			zap.String("senderLastName", owner.User.LastName),
			zap.String("tokenString", tokenString),
		)
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToSendInvite

		return nil
	}

	res.Status = pkg.ResponseStatusOk

	return nil
}

func (s *Service) AcceptInvite(
	ctx context.Context,
	req *grpc.AcceptInviteRequest,
	res *grpc.AcceptInviteResponse,
) error {
	claims, err := s.parseInviteToken(req.Token, req.Email)

	if err != nil {
		zap.L().Error("Error on parse invite token", zap.Error(err), zap.String("token", req.Token), zap.String("email", req.Email))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserInvalidToken

		return nil
	}

	var user *billing.UserRole

	switch claims[claimType] {
	case pkg.RoleTypeMerchant:
		user, err = s.userRoleRepository.GetMerchantUserById(claims[claimRoleId].(string))
		break
	case pkg.RoleTypeSystem:
		user, err = s.userRoleRepository.GetAdminUserById(claims[claimRoleId].(string))
		break
	default:
		err = errors.New(errorUserUnsupportedRoleType.Message)
	}

	if err != nil {
		zap.L().Error(errorUserNotFound.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNotFound
		return nil
	}

	if user.User.Status != pkg.UserRoleStatusInvited {
		zap.L().Error(errorUserInviteAlreadyAccepted.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserInviteAlreadyAccepted
		return nil
	}

	user.User.Status = pkg.UserRoleStatusAccepted
	user.User.UserId = req.UserId
	user.User.FirstName = req.FirstName
	user.User.LastName = req.LastName

	switch claims[claimType] {
	case pkg.RoleTypeMerchant:
		err = s.userRoleRepository.UpdateMerchantUser(user)
		break
	case pkg.RoleTypeSystem:
		err = s.userRoleRepository.UpdateAdminUser(user)
		break
	default:
		err = errors.New(errorUserUnsupportedRoleType.Message)
	}

	if err != nil {
		zap.L().Error(errorUserUnableToAdd.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToAdd
		return nil
	}

	casbinUserId := user.User.UserId

	if claims[claimType] == pkg.RoleTypeMerchant {
		casbinUserId = fmt.Sprintf(casbinMerchantUserMask, user.MerchantId, user.User.UserId)
	}

	_, err = s.casbinService.AddRoleForUser(ctx, &casbinProto.UserRoleRequest{
		User: casbinUserId,
		Role: user.Role,
	})

	if err != nil {
		zap.L().Error(errorUserUnableToAddToCasbin.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToAddToCasbin
		return nil
	}

	res.Status = pkg.ResponseStatusOk
	res.Role = user

	return nil
}

func (s *Service) GetMerchantUser(
	ctx context.Context,
	req *grpc.GetMerchantUserRequest,
	res *grpc.GetMerchantUserResponse,
) error {
	user, err := s.userRoleRepository.GetMerchantUserById(req.Id)

	if err != nil {
		zap.L().Error(errorUserNotFound.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNotFound

		return nil
	}

	res.Status = pkg.ResponseStatusOk
	res.Role = user

	return nil
}

func (s *Service) GetAdminUser(
	ctx context.Context,
	req *grpc.GetAdminUserRequest,
	res *grpc.GetAdminUserResponse,
) error {
	user, err := s.userRoleRepository.GetAdminUserById(req.Id)

	if err != nil {
		zap.L().Error(errorUserNotFound.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNotFound

		return nil
	}

	res.Status = pkg.ResponseStatusOk
	res.Role = user

	return nil
}

func (s *Service) CheckInviteToken(
	ctx context.Context,
	req *grpc.CheckInviteTokenRequest,
	res *grpc.CheckInviteTokenResponse,
) error {
	claims, err := s.parseInviteToken(req.Token, req.Email)

	if err != nil {
		zap.L().Error("Error on parse invite token", zap.Error(err), zap.String("token", req.Token), zap.String("email", req.Email))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserInvalidToken

		return nil
	}

	res.Status = pkg.ResponseStatusOk
	res.RoleId = claims[claimRoleId].(string)
	res.RoleType = claims[claimType].(string)

	return nil
}

func (s *Service) parseInviteToken(t string, email string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(errorUserInvalidToken.Message)
		}

		return []byte(s.cfg.UserInviteTokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token isn't valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("cannot read claims")
	}

	if claims[claimEmail] != email {
		return nil, errors.New(errorUserInvalidInviteEmail.Message)
	}

	return claims, nil
}

func (s *Service) sendInviteEmail(receiverEmail, senderEmail, senderFirstName, senderLastName, senderCompany, token string) error {
	// TODO: What will be the address of the links? How do we transfer the parameters (to the public or to the JWT key)?
	inviteLink := "/user/invite/member?key=" + token

	if senderCompany != defaultCompanyName {
		inviteLink = "/user/invite/merchant?key=" + token
	}

	payload := &postmarkSdrPkg.Payload{
		TemplateAlias: s.cfg.EmailInviteTemplate,
		TemplateModel: map[string]string{
			"sender_first_name": senderFirstName,
			"sender_last_name":  senderLastName,
			"sender_email":      senderEmail,
			"sender_company":    senderCompany,
			"invite_link":       inviteLink,
		},
		To: receiverEmail,
	}
	err := s.broker.Publish(postmarkSdrPkg.PostmarkSenderTopicName, payload, amqp.Table{})

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ChangeRoleForMerchantUser(ctx context.Context, req *grpc.ChangeRoleForMerchantUserRequest, res *grpc.EmptyResponseWithStatus) error {
	owner, err := s.userRoleRepository.GetMerchantUserByUserId(req.MerchantId, req.PerformerId)

	if err != nil || !owner.IsOwner() {
		zap.L().Error(errorUserNoHavePermission.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNoHavePermission

		return nil
	}

	user, err := s.userRoleRepository.GetMerchantUserById(req.RoleId)

	if err != nil {
		zap.L().Error(errorUserAlreadyExist.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNotFound

		return nil
	}

	if user.Role == req.Role {
		zap.L().Error(errorUserAlreadyHasRole.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserAlreadyHasRole

		return nil
	}

	user.Role = req.Role
	err = s.userRoleRepository.UpdateMerchantUser(user)

	if err != nil {
		zap.L().Error(errorUserNotFound.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusSystemError
		res.Message = errorUserUnableToSave

		return nil
	}

	casbinUserId := fmt.Sprintf(casbinMerchantUserMask, user.MerchantId, user.User.UserId)

	_, err = s.casbinService.DeleteUser(ctx, &casbinProto.UserRoleRequest{User: casbinUserId})

	if err != nil {
		zap.L().Error(errorUserUnableToDeleteFromCasbin.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusSystemError
		res.Message = errorUserUnableToDeleteFromCasbin
		return nil
	}

	_, err = s.casbinService.AddRoleForUser(ctx, &casbinProto.UserRoleRequest{
		User: casbinUserId,
		Role: req.Role,
	})

	if err != nil {
		zap.L().Error(errorUserUnableToDeleteFromCasbin.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusSystemError
		res.Message = errorUserUnableToDeleteFromCasbin
		return nil
	}

	res.Status = pkg.ResponseStatusOk

	return nil
}

func (s *Service) ChangeRoleForAdminUser(ctx context.Context, req *grpc.ChangeRoleForAdminUserRequest, res *grpc.EmptyResponseWithStatus) error {
	owner, err := s.userRoleRepository.GetAdminUserByUserId(req.PerformerId)

	if err != nil || !owner.IsAdmin() {
		zap.L().Error(errorUserNoHavePermission.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNoHavePermission

		return nil
	}

	user, err := s.userRoleRepository.GetAdminUserById(req.RoleId)

	if err != nil {
		zap.L().Error(errorUserAlreadyExist.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNotFound

		return nil
	}

	if user.Role == req.Role {
		zap.L().Error(errorUserAlreadyHasRole.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserAlreadyHasRole

		return nil
	}

	user.Role = req.Role
	err = s.userRoleRepository.UpdateAdminUser(user)

	if err != nil {
		zap.L().Error(errorUserUnableToSave.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusSystemError
		res.Message = errorUserUnableToSave
		res.Message.Details = err.Error()

		return nil
	}

	_, err = s.casbinService.DeleteUser(ctx, &casbinProto.UserRoleRequest{User: user.User.UserId})

	if err != nil {
		zap.L().Error(errorUserUnableToDeleteFromCasbin.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusSystemError
		res.Message = errorUserUnableToDeleteFromCasbin
		return nil
	}

	_, err = s.casbinService.AddRoleForUser(ctx, &casbinProto.UserRoleRequest{
		User: user.User.UserId,
		Role: req.Role,
	})

	if err != nil {
		zap.L().Error(errorUserUnableToDeleteFromCasbin.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusSystemError
		res.Message = errorUserUnableToDeleteFromCasbin
		return nil
	}

	res.Status = pkg.ResponseStatusOk

	return nil
}

func (s *Service) GetRoleList(ctx context.Context, req *grpc.GetRoleListRequest, res *grpc.GetRoleListResponse) error {
	list, ok := merchantUserRoles[req.Type]

	if !ok {
		zap.L().Error("unsupported user type", zap.Any("req", req))
		return nil
	}

	res.Items = list

	return nil
}

func (s *Service) DeleteMerchantUser(
	ctx context.Context,
	req *grpc.DeleteMerchantUserRequest,
	res *grpc.EmptyResponseWithStatus,
) error {
	owner, err := s.userRoleRepository.GetMerchantUserByUserId(req.MerchantId, req.PerformerId)

	if err != nil || !owner.IsOwner() {
		zap.L().Error(errorUserNoHavePermission.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNoHavePermission

		return nil
	}

	user, err := s.userRoleRepository.GetMerchantUserById(req.RoleId)

	if err != nil {
		zap.L().Error(errorUserAlreadyExist.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNotFound

		return nil
	}

	if err = s.userRoleRepository.DeleteMerchantUser(user); err != nil {
		zap.L().Error(errorUserUnableToDelete.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToDelete

		return nil
	}

	_, err = s.casbinService.DeleteUser(ctx, &casbinProto.UserRoleRequest{
		User: fmt.Sprintf(casbinMerchantUserMask, user.MerchantId, user.User.UserId),
	})

	if err != nil {
		zap.L().Error(errorUserUnableToDeleteFromCasbin.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToDeleteFromCasbin
		return nil
	}

	res.Status = pkg.ResponseStatusOk

	return nil
}

func (s *Service) DeleteAdminUser(
	ctx context.Context,
	req *grpc.DeleteAdminUserRequest,
	res *grpc.EmptyResponseWithStatus,
) error {
	owner, err := s.userRoleRepository.GetAdminUserByUserId(req.PerformerId)

	if err != nil || !owner.IsAdmin() {
		zap.L().Error(errorUserNoHavePermission.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNoHavePermission

		return nil
	}

	user, err := s.userRoleRepository.GetAdminUserById(req.RoleId)

	if err != nil {
		zap.L().Error(errorUserAlreadyExist.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserNotFound

		return nil
	}

	if err = s.userRoleRepository.DeleteAdminUser(user); err != nil {
		zap.L().Error(errorUserUnableToDelete.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToDelete

		return nil
	}

	_, err = s.casbinService.DeleteUser(ctx, &casbinProto.UserRoleRequest{User: user.User.UserId})

	if err != nil {
		zap.L().Error(errorUserUnableToDeleteFromCasbin.Message, zap.Error(err), zap.Any("req", req))
		res.Status = pkg.ResponseStatusBadData
		res.Message = errorUserUnableToDeleteFromCasbin
		return nil
	}

	res.Status = pkg.ResponseStatusOk

	return nil
}
