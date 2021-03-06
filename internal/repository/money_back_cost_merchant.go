package repository

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/paysuper/paysuper-billing-server/internal/database"
	internalPkg "github.com/paysuper/paysuper-billing-server/internal/pkg"
	"github.com/paysuper/paysuper-billing-server/internal/repository/models"
	"github.com/paysuper/paysuper-billing-server/pkg"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	tools "github.com/paysuper/paysuper-tools/number"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	mongodb "gopkg.in/paysuper/paysuper-database-mongo.v2"
)

const (
	collectionMoneyBackCostMerchant = "money_back_cost_merchant"

	cacheMoneyBackCostMerchantKey   = "pucm:m:%s:n:%s:pc:%s:ur:%s:r:%s:c:%s:ps:%d:mcc:%s"
	cacheMoneyBackCostMerchantKeyId = "pucm:id:%s"
	cacheMoneyBackCostMerchantAll   = "pucm:all:m:%s"
)

type moneyBackCostMerchantRepository repository

// NewMoneyBackCostMerchantRepository create and return an object for working with the cost of merchant for money back.
// The returned object implements the MoneyBackCostMerchantRepositoryInterface interface.
func NewMoneyBackCostMerchantRepository(db mongodb.SourceInterface, cache database.CacheInterface) MoneyBackCostMerchantRepositoryInterface {
	s := &moneyBackCostMerchantRepository{db: db, cache: cache, mapper: models.NewMoneyBackCostMerchantMapper()}
	return s
}

func (r *moneyBackCostMerchantRepository) Insert(ctx context.Context, obj *billingpb.MoneyBackCostMerchant) (err error) {
	obj.FixAmount = tools.FormatAmount(obj.FixAmount)
	obj.Percent = tools.ToPrecise(obj.Percent)
	obj.IsActive = true

	mgo, err := r.mapper.MapObjectToMgo(obj)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseMapModelFailed,
			zap.Error(err),
			zap.Any(pkg.ErrorDatabaseFieldQuery, obj),
		)
		return err
	}

	_, err = r.db.Collection(collectionMoneyBackCostMerchant).InsertOne(ctx, mgo)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldOperation, pkg.ErrorDatabaseFieldOperationInsert),
			zap.Any(pkg.ErrorDatabaseFieldDocument, obj),
		)
		return err
	}

	_ = r.updateCaches(obj)

	return nil
}

func (r *moneyBackCostMerchantRepository) MultipleInsert(ctx context.Context, objs []*billingpb.MoneyBackCostMerchant) error {
	c := make([]interface{}, len(objs))

	for i, v := range objs {
		if v.Id == "" {
			v.Id = primitive.NewObjectID().Hex()
		}

		v.FixAmount = tools.FormatAmount(v.FixAmount)
		v.Percent = tools.ToPrecise(v.Percent)
		v.IsActive = true

		mgo, err := r.mapper.MapObjectToMgo(v)

		if err != nil {
			zap.L().Error(
				pkg.ErrorDatabaseMapModelFailed,
				zap.Error(err),
				zap.Any(pkg.ErrorDatabaseFieldQuery, v),
			)
		}

		c[i] = mgo
	}

	_, err := r.db.Collection(collectionMoneyBackCostMerchant).InsertMany(ctx, c)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldOperation, pkg.ErrorDatabaseFieldOperationInsert),
			zap.Any(pkg.ErrorDatabaseFieldQuery, c),
		)
		return err
	}

	for _, v := range objs {
		_ = r.updateCaches(v)
	}

	return nil
}

func (r *moneyBackCostMerchantRepository) Update(ctx context.Context, obj *billingpb.MoneyBackCostMerchant) error {
	oid, err := primitive.ObjectIDFromHex(obj.Id)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseInvalidObjectId,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldQuery, obj.Id),
		)
		return err
	}

	obj.FixAmount = tools.FormatAmount(obj.FixAmount)
	obj.Percent = tools.ToPrecise(obj.Percent)
	obj.IsActive = true

	mgo, err := r.mapper.MapObjectToMgo(obj)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseMapModelFailed,
			zap.Error(err),
			zap.Any(pkg.ErrorDatabaseFieldQuery, obj),
		)
		return err
	}

	filter := bson.M{"_id": oid}
	_, err = r.db.Collection(collectionMoneyBackCostMerchant).ReplaceOne(ctx, filter, mgo)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldOperation, pkg.ErrorDatabaseFieldOperationUpdate),
			zap.Any(pkg.ErrorDatabaseFieldQuery, obj),
		)
		return err
	}

	return r.updateCaches(obj)
}

func (r *moneyBackCostMerchantRepository) GetById(ctx context.Context, id string) (*billingpb.MoneyBackCostMerchant, error) {
	obj := &billingpb.MoneyBackCostMerchant{}
	key := fmt.Sprintf(cacheMoneyBackCostMerchantKeyId, id)

	if err := r.cache.Get(key, obj); err == nil {
		return obj, nil
	}

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseInvalidObjectId,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldQuery, id),
		)
		return nil, err
	}

	var mgo = models.MgoMoneyBackCostMerchant{}
	query := bson.M{"_id": oid, "is_active": true}
	err = r.db.Collection(collectionMoneyBackCostMerchant).FindOne(ctx, query).Decode(&mgo)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.Any(pkg.ErrorDatabaseFieldQuery, query),
		)
		return nil, err
	}

	o, err := r.mapper.MapMgoToObject(&mgo)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseMapModelFailed,
			zap.Error(err),
			zap.Any(pkg.ErrorDatabaseFieldQuery, mgo),
		)
		return nil, err
	}

	obj = o.(*billingpb.MoneyBackCostMerchant)

	_ = r.cache.Set(key, obj, 0)

	return obj, nil
}

func (r *moneyBackCostMerchantRepository) GetAllForMerchant(ctx context.Context, merchantId string) (*billingpb.MoneyBackCostMerchantList, error) {
	item := billingpb.MoneyBackCostMerchantList{}
	key := fmt.Sprintf(cacheMoneyBackCostMerchantAll, merchantId)

	if err := r.cache.Get(key, &item); err == nil {
		return &item, nil
	}

	oid, err := primitive.ObjectIDFromHex(merchantId)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseInvalidObjectId,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldQuery, merchantId),
		)
		return nil, err
	}

	query := bson.M{"merchant_id": oid, "is_active": true}
	opts := options.Find().
		SetSort(bson.M{"name": 1, "payout_currency": 1, "region": 1, "country": 1, "mcc_code": 1})
	cursor, err := r.db.Collection(collectionMoneyBackCostMerchant).Find(ctx, query, opts)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.Any(pkg.ErrorDatabaseFieldQuery, query),
		)
		return nil, err
	}

	var list []*models.MgoMoneyBackCostMerchant
	err = cursor.All(ctx, &list)

	if err != nil {
		zap.L().Error(
			pkg.ErrorQueryCursorExecutionFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.Any(pkg.ErrorDatabaseFieldQuery, query),
		)
		return nil, err
	}

	objs := make([]*billingpb.MoneyBackCostMerchant, len(list))

	for i, obj := range list {
		v, err := r.mapper.MapMgoToObject(obj)
		if err != nil {
			zap.L().Error(
				pkg.ErrorDatabaseMapModelFailed,
				zap.Error(err),
				zap.Any(pkg.ErrorDatabaseFieldQuery, obj),
			)
			return nil, err
		}
		objs[i] = v.(*billingpb.MoneyBackCostMerchant)
	}

	item.Items = objs

	err = r.cache.Set(key, &item, 0)

	if err != nil {
		zap.L().Error(
			pkg.ErrorCacheQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorCacheFieldCmd, "SET"),
			zap.String(pkg.ErrorCacheFieldKey, key),
			zap.Any(pkg.ErrorCacheFieldData, item),
		)
	}

	return &item, nil
}

func (r *moneyBackCostMerchantRepository) Find(
	ctx context.Context,
	merchantId string,
	name string,
	payoutCurrency string,
	undoReason string,
	region string,
	country string,
	mccCode string,
	paymentStage int32,
) ([]*internalPkg.MoneyBackCostMerchantSet, error) {
	c := make([]*internalPkg.MoneyBackCostMerchantSet, 0)
	key := fmt.Sprintf(cacheMoneyBackCostMerchantKey, merchantId, name, payoutCurrency, undoReason, region, country, paymentStage, mccCode)

	if err := r.cache.Get(key, &c); err == nil {
		return c, nil
	}

	merchantOid, err := primitive.ObjectIDFromHex(merchantId)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseInvalidObjectId,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldQuery, merchantId),
		)
		return nil, err
	}

	matchQuery := bson.M{
		"merchant_id":     merchantOid,
		"name":            primitive.Regex{Pattern: "^" + name + "$", Options: "i"},
		"payout_currency": payoutCurrency,
		"undo_reason":     undoReason,
		"payment_stage":   paymentStage,
		"is_active":       true,
		"mcc_code":        mccCode,
		"$or": []bson.M{
			{
				"country": country,
				"region":  region,
			},
			{
				"$or": []bson.M{
					{"country": ""},
					{"country": bson.M{"$exists": false}},
				},
				"region": region,
			},
		},
	}

	query := []bson.M{
		{
			"$match": matchQuery,
		},
		{
			"$group": bson.M{
				"_id": "$country",
				"set": bson.M{"$push": "$$ROOT"},
			},
		},
		{
			"$sort": bson.M{"_id": -1},
		},
	}

	cursor, err := r.db.Collection(collectionMoneyBackCostMerchant).Aggregate(ctx, query)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.Any(pkg.ErrorDatabaseFieldQuery, query),
		)
		return nil, err
	}

	mgo := make([]*internalPkg.MgoMoneyBackCostMerchantSet, 0)
	err = cursor.All(ctx, &mgo)

	if err != nil {
		zap.L().Error(
			pkg.ErrorQueryCursorExecutionFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.Any(pkg.ErrorDatabaseFieldQuery, query),
		)
		return nil, err
	}

	list := make([]*billingpb.MoneyBackCostMerchant, 0)

	for _, objs := range mgo {
		list = nil

		for _, obj := range objs.Set {
			v, err := r.mapper.MapMgoToObject(obj)

			if err != nil {
				zap.L().Error(
					pkg.ErrorDatabaseMapModelFailed,
					zap.Error(err),
					zap.Any(pkg.ErrorDatabaseFieldQuery, obj),
				)
				return nil, err
			}

			list = append(list, v.(*billingpb.MoneyBackCostMerchant))
		}

		c = append(c, &internalPkg.MoneyBackCostMerchantSet{Id: objs.Id, Set: list})
	}

	if err = r.cache.Set(key, c, 0); err != nil {
		zap.L().Error(
			pkg.ErrorCacheQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorCacheFieldCmd, "SET"),
			zap.String(pkg.ErrorCacheFieldKey, key),
			zap.Any(pkg.ErrorCacheFieldData, c),
		)
	}

	return c, nil
}

func (r *moneyBackCostMerchantRepository) Delete(ctx context.Context, obj *billingpb.MoneyBackCostMerchant) error {
	oid, err := primitive.ObjectIDFromHex(obj.Id)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseInvalidObjectId,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldQuery, obj.Id),
		)
		return err
	}

	obj.UpdatedAt = ptypes.TimestampNow()
	obj.IsActive = false

	mgo, err := r.mapper.MapObjectToMgo(obj)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseMapModelFailed,
			zap.Error(err),
			zap.Any(pkg.ErrorDatabaseFieldQuery, obj),
		)
		return err
	}

	filter := bson.M{"_id": oid}
	_, err = r.db.Collection(collectionMoneyBackCostMerchant).ReplaceOne(ctx, filter, mgo)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.Any(pkg.ErrorDatabaseFieldQuery, obj),
		)
		return err
	}

	return r.updateCaches(obj)
}

func (r *moneyBackCostMerchantRepository) updateCaches(obj *billingpb.MoneyBackCostMerchant) error {
	groupKeys := []string{
		fmt.Sprintf(cacheMoneyBackCostMerchantKey, obj.MerchantId, obj.Name, obj.PayoutCurrency, obj.UndoReason, obj.Region, obj.Country, obj.PaymentStage, obj.MccCode),
		fmt.Sprintf(cacheMoneyBackCostMerchantKey, obj.MerchantId, obj.Name, obj.PayoutCurrency, obj.UndoReason, obj.Region, "", obj.PaymentStage, obj.MccCode),
		fmt.Sprintf(cacheMoneyBackCostMerchantAll, obj.MerchantId),
	}

	for _, key := range groupKeys {
		if err := r.cache.Delete(key); err != nil {
			zap.L().Error(
				pkg.ErrorCacheQueryFailed,
				zap.Error(err),
				zap.String(pkg.ErrorCacheFieldCmd, "DELETE"),
				zap.String(pkg.ErrorCacheFieldKey, key),
			)
			return err
		}
	}

	key := fmt.Sprintf(cacheMoneyBackCostMerchantKeyId, obj.Id)

	if err := r.cache.Delete(key); err != nil {
		zap.L().Error(
			pkg.ErrorCacheQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorCacheFieldCmd, "DELETE"),
			zap.String(pkg.ErrorCacheFieldKey, key),
		)
		return err
	}

	if obj.IsActive {
		if err := r.cache.Set(key, obj, 0); err != nil {
			zap.L().Error(
				pkg.ErrorCacheQueryFailed,
				zap.Error(err),
				zap.String(pkg.ErrorCacheFieldCmd, "SET"),
				zap.String(pkg.ErrorCacheFieldKey, key),
				zap.Any(pkg.ErrorCacheFieldData, obj),
			)
			return err
		}
	}

	return nil
}

func (r *moneyBackCostMerchantRepository) DeleteAndInsertMany(
	ctx context.Context,
	merchantId string,
	tariffs []*billingpb.MoneyBackCostMerchant,
) error {
	merchantOid, err := primitive.ObjectIDFromHex(merchantId)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseInvalidObjectId,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.String(pkg.ErrorDatabaseFieldQuery, merchantId),
		)
		return err
	}

	filter := bson.M{"merchant_id": merchantOid, "is_active": true}
	_, err = r.db.Collection(collectionMoneyBackCostMerchant).DeleteMany(ctx, filter)

	if err != nil {
		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, collectionMoneyBackCostMerchant),
			zap.Any(pkg.ErrorDatabaseFieldQuery, filter),
		)
		return err
	}

	return r.MultipleInsert(ctx, tariffs)
}
