// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/grpc.proto

package grpc // import "github.com/ProtocolONE/payone-billing-service/pkg/proto/grpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import billing "github.com/ProtocolONE/payone-billing-service/pkg/proto/billing"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{0}
}
func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (dst *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(dst, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{1}
}
func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (dst *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(dst, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

type PaymentCreateRequest struct {
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-" structure:"-"`
}

func (m *PaymentCreateRequest) Reset()         { *m = PaymentCreateRequest{} }
func (m *PaymentCreateRequest) String() string { return proto.CompactTextString(m) }
func (*PaymentCreateRequest) ProtoMessage()    {}
func (*PaymentCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{2}
}
func (m *PaymentCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentCreateRequest.Unmarshal(m, b)
}
func (m *PaymentCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentCreateRequest.Marshal(b, m, deterministic)
}
func (dst *PaymentCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentCreateRequest.Merge(dst, src)
}
func (m *PaymentCreateRequest) XXX_Size() int {
	return xxx_messageInfo_PaymentCreateRequest.Size(m)
}
func (m *PaymentCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentCreateRequest proto.InternalMessageInfo

func (m *PaymentCreateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type PaymentCreateResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	RedirectUrl          string   `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *PaymentCreateResponse) Reset()         { *m = PaymentCreateResponse{} }
func (m *PaymentCreateResponse) String() string { return proto.CompactTextString(m) }
func (*PaymentCreateResponse) ProtoMessage()    {}
func (*PaymentCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{3}
}
func (m *PaymentCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentCreateResponse.Unmarshal(m, b)
}
func (m *PaymentCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentCreateResponse.Marshal(b, m, deterministic)
}
func (dst *PaymentCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentCreateResponse.Merge(dst, src)
}
func (m *PaymentCreateResponse) XXX_Size() int {
	return xxx_messageInfo_PaymentCreateResponse.Size(m)
}
func (m *PaymentCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentCreateResponse proto.InternalMessageInfo

func (m *PaymentCreateResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PaymentCreateResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func (m *PaymentCreateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type PaymentFormJsonDataRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Scheme               string   `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Host                 string   `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *PaymentFormJsonDataRequest) Reset()         { *m = PaymentFormJsonDataRequest{} }
func (m *PaymentFormJsonDataRequest) String() string { return proto.CompactTextString(m) }
func (*PaymentFormJsonDataRequest) ProtoMessage()    {}
func (*PaymentFormJsonDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{4}
}
func (m *PaymentFormJsonDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentFormJsonDataRequest.Unmarshal(m, b)
}
func (m *PaymentFormJsonDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentFormJsonDataRequest.Marshal(b, m, deterministic)
}
func (dst *PaymentFormJsonDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentFormJsonDataRequest.Merge(dst, src)
}
func (m *PaymentFormJsonDataRequest) XXX_Size() int {
	return xxx_messageInfo_PaymentFormJsonDataRequest.Size(m)
}
func (m *PaymentFormJsonDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentFormJsonDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentFormJsonDataRequest proto.InternalMessageInfo

func (m *PaymentFormJsonDataRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *PaymentFormJsonDataRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *PaymentFormJsonDataRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type PaymentFormJsonDataProject struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// @inject_tag: json:"url_success,omitempty"
	UrlSuccess string `protobuf:"bytes,2,opt,name=url_success,json=urlSuccess,proto3" json:"url_success,omitempty"`
	// @inject_tag: json:"url_fail,omitempty"
	UrlFail              string   `protobuf:"bytes,3,opt,name=url_fail,json=urlFail,proto3" json:"url_fail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *PaymentFormJsonDataProject) Reset()         { *m = PaymentFormJsonDataProject{} }
func (m *PaymentFormJsonDataProject) String() string { return proto.CompactTextString(m) }
func (*PaymentFormJsonDataProject) ProtoMessage()    {}
func (*PaymentFormJsonDataProject) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{5}
}
func (m *PaymentFormJsonDataProject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentFormJsonDataProject.Unmarshal(m, b)
}
func (m *PaymentFormJsonDataProject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentFormJsonDataProject.Marshal(b, m, deterministic)
}
func (dst *PaymentFormJsonDataProject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentFormJsonDataProject.Merge(dst, src)
}
func (m *PaymentFormJsonDataProject) XXX_Size() int {
	return xxx_messageInfo_PaymentFormJsonDataProject.Size(m)
}
func (m *PaymentFormJsonDataProject) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentFormJsonDataProject.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentFormJsonDataProject proto.InternalMessageInfo

func (m *PaymentFormJsonDataProject) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PaymentFormJsonDataProject) GetUrlSuccess() string {
	if m != nil {
		return m.UrlSuccess
	}
	return ""
}

func (m *PaymentFormJsonDataProject) GetUrlFail() string {
	if m != nil {
		return m.UrlFail
	}
	return ""
}

type PaymentFormJsonDataResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"account,omitempty"
	Account               string                              `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	HasVat                bool                                `protobuf:"varint,3,opt,name=has_vat,json=hasVat,proto3" json:"has_vat,omitempty"`
	HasUserCommission     bool                                `protobuf:"varint,4,opt,name=has_user_commission,json=hasUserCommission,proto3" json:"has_user_commission,omitempty"`
	Project               *PaymentFormJsonDataProject         `protobuf:"bytes,5,opt,name=project,proto3" json:"project,omitempty"`
	PaymentMethods        []*billing.PaymentFormPaymentMethod `protobuf:"bytes,6,rep,name=payment_methods,json=paymentMethods,proto3" json:"payment_methods,omitempty"`
	InlineFormRedirectUrl string                              `protobuf:"bytes,7,opt,name=inline_form_redirect_url,json=inlineFormRedirectUrl,proto3" json:"inline_form_redirect_url,omitempty"`
	Token                 string                              `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                            `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized      []byte                              `json:"-" bson:"-" structure:"-"`
	XXX_sizecache         int32                               `json:"-" bson:"-" structure:"-"`
}

func (m *PaymentFormJsonDataResponse) Reset()         { *m = PaymentFormJsonDataResponse{} }
func (m *PaymentFormJsonDataResponse) String() string { return proto.CompactTextString(m) }
func (*PaymentFormJsonDataResponse) ProtoMessage()    {}
func (*PaymentFormJsonDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{6}
}
func (m *PaymentFormJsonDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentFormJsonDataResponse.Unmarshal(m, b)
}
func (m *PaymentFormJsonDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentFormJsonDataResponse.Marshal(b, m, deterministic)
}
func (dst *PaymentFormJsonDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentFormJsonDataResponse.Merge(dst, src)
}
func (m *PaymentFormJsonDataResponse) XXX_Size() int {
	return xxx_messageInfo_PaymentFormJsonDataResponse.Size(m)
}
func (m *PaymentFormJsonDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentFormJsonDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentFormJsonDataResponse proto.InternalMessageInfo

func (m *PaymentFormJsonDataResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaymentFormJsonDataResponse) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *PaymentFormJsonDataResponse) GetHasVat() bool {
	if m != nil {
		return m.HasVat
	}
	return false
}

func (m *PaymentFormJsonDataResponse) GetHasUserCommission() bool {
	if m != nil {
		return m.HasUserCommission
	}
	return false
}

func (m *PaymentFormJsonDataResponse) GetProject() *PaymentFormJsonDataProject {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *PaymentFormJsonDataResponse) GetPaymentMethods() []*billing.PaymentFormPaymentMethod {
	if m != nil {
		return m.PaymentMethods
	}
	return nil
}

func (m *PaymentFormJsonDataResponse) GetInlineFormRedirectUrl() string {
	if m != nil {
		return m.InlineFormRedirectUrl
	}
	return ""
}

func (m *PaymentFormJsonDataResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type PaymentNotifyRequest struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Request              *any.Any `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	RawRequest           string   `protobuf:"bytes,3,opt,name=raw_request,json=rawRequest,proto3" json:"raw_request,omitempty"`
	Signature            string   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *PaymentNotifyRequest) Reset()         { *m = PaymentNotifyRequest{} }
func (m *PaymentNotifyRequest) String() string { return proto.CompactTextString(m) }
func (*PaymentNotifyRequest) ProtoMessage()    {}
func (*PaymentNotifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{7}
}
func (m *PaymentNotifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentNotifyRequest.Unmarshal(m, b)
}
func (m *PaymentNotifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentNotifyRequest.Marshal(b, m, deterministic)
}
func (dst *PaymentNotifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentNotifyRequest.Merge(dst, src)
}
func (m *PaymentNotifyRequest) XXX_Size() int {
	return xxx_messageInfo_PaymentNotifyRequest.Size(m)
}
func (m *PaymentNotifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentNotifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentNotifyRequest proto.InternalMessageInfo

func (m *PaymentNotifyRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *PaymentNotifyRequest) GetRequest() *any.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *PaymentNotifyRequest) GetRawRequest() string {
	if m != nil {
		return m.RawRequest
	}
	return ""
}

func (m *PaymentNotifyRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type PaymentNotifyResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *PaymentNotifyResponse) Reset()         { *m = PaymentNotifyResponse{} }
func (m *PaymentNotifyResponse) String() string { return proto.CompactTextString(m) }
func (*PaymentNotifyResponse) ProtoMessage()    {}
func (*PaymentNotifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_400e69b0b3b24cb9, []int{8}
}
func (m *PaymentNotifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentNotifyResponse.Unmarshal(m, b)
}
func (m *PaymentNotifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentNotifyResponse.Marshal(b, m, deterministic)
}
func (dst *PaymentNotifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentNotifyResponse.Merge(dst, src)
}
func (m *PaymentNotifyResponse) XXX_Size() int {
	return xxx_messageInfo_PaymentNotifyResponse.Size(m)
}
func (m *PaymentNotifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentNotifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentNotifyResponse proto.InternalMessageInfo

func (m *PaymentNotifyResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PaymentNotifyResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "grpc.EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "grpc.EmptyResponse")
	proto.RegisterType((*PaymentCreateRequest)(nil), "grpc.PaymentCreateRequest")
	proto.RegisterMapType((map[string]string)(nil), "grpc.PaymentCreateRequest.DataEntry")
	proto.RegisterType((*PaymentCreateResponse)(nil), "grpc.PaymentCreateResponse")
	proto.RegisterType((*PaymentFormJsonDataRequest)(nil), "grpc.PaymentFormJsonDataRequest")
	proto.RegisterType((*PaymentFormJsonDataProject)(nil), "grpc.PaymentFormJsonDataProject")
	proto.RegisterType((*PaymentFormJsonDataResponse)(nil), "grpc.PaymentFormJsonDataResponse")
	proto.RegisterType((*PaymentNotifyRequest)(nil), "grpc.PaymentNotifyRequest")
	proto.RegisterType((*PaymentNotifyResponse)(nil), "grpc.PaymentNotifyResponse")
}

func init() { proto.RegisterFile("grpc/grpc.proto", fileDescriptor_grpc_400e69b0b3b24cb9) }

var fileDescriptor_grpc_400e69b0b3b24cb9 = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0x8e, 0x7f, 0x62, 0xc7, 0x74, 0xea, 0xb4, 0xcc, 0x4f, 0x15, 0xa5, 0x40, 0x1d, 0xa1, 0x87,
	0x5c, 0x22, 0x03, 0xee, 0x21, 0x69, 0x50, 0x14, 0x68, 0x5c, 0x07, 0x68, 0x80, 0x26, 0xae, 0x82,
	0xf4, 0xd0, 0x8b, 0x40, 0x53, 0xb4, 0xc5, 0x9a, 0x12, 0x55, 0x92, 0x4a, 0xa0, 0x47, 0xd8, 0x87,
	0x58, 0x60, 0xdf, 0x6e, 0x5f, 0x63, 0x21, 0x92, 0x72, 0xec, 0xc0, 0xd9, 0xdd, 0x8b, 0x3d, 0x7f,
	0x9c, 0x6f, 0x86, 0xf3, 0x0d, 0x05, 0xf6, 0xe6, 0x22, 0xc3, 0x83, 0xf2, 0xc7, 0xcf, 0x04, 0x57,
	0x1c, 0x36, 0x4b, 0xd9, 0x3d, 0x9c, 0x52, 0xc6, 0x68, 0x3a, 0x1f, 0xd8, 0x7f, 0xe3, 0x74, 0x8f,
	0xe7, 0x9c, 0xcf, 0x19, 0x19, 0x68, 0x6d, 0x9a, 0xcf, 0x06, 0x28, 0x2d, 0x8c, 0xcb, 0xeb, 0x81,
	0xdd, 0x71, 0x92, 0xa9, 0x22, 0x20, 0xff, 0xe7, 0x44, 0x2a, 0x6f, 0x0f, 0x7c, 0x63, 0x75, 0x99,
	0xf1, 0x54, 0x12, 0xef, 0x5d, 0x0d, 0x1c, 0x4c, 0x50, 0x91, 0x90, 0x54, 0x8d, 0x04, 0x41, 0x8a,
	0xd8, 0x48, 0x78, 0x09, 0x9a, 0x11, 0x52, 0xc8, 0xa9, 0xf5, 0x1b, 0x67, 0xdd, 0xe1, 0x4f, 0xbe,
	0x2e, 0x66, 0x53, 0xa4, 0xff, 0x07, 0x52, 0x68, 0x9c, 0x2a, 0x51, 0x04, 0xfa, 0x84, 0x7b, 0x01,
	0x3a, 0x4b, 0x13, 0xfc, 0x16, 0x34, 0x16, 0xa4, 0x70, 0x6a, 0xfd, 0xda, 0x59, 0x27, 0x28, 0x45,
	0x78, 0x00, 0xb6, 0x9f, 0x10, 0xcb, 0x89, 0x53, 0xd7, 0x36, 0xa3, 0x5c, 0xd5, 0x2f, 0x6b, 0x5e,
	0x0c, 0x0e, 0x5f, 0x01, 0x98, 0x22, 0xe1, 0x11, 0x68, 0x49, 0x85, 0x54, 0x2e, 0x75, 0x9e, 0xed,
	0xc0, 0x6a, 0xf0, 0x14, 0xec, 0x0a, 0x12, 0x51, 0x41, 0xb0, 0x0a, 0x73, 0xc1, 0x6c, 0xc6, 0x6e,
	0x65, 0x7b, 0x14, 0xac, 0x44, 0x23, 0x42, 0x70, 0xe1, 0x34, 0x0c, 0x9a, 0x56, 0x3c, 0x0c, 0x5c,
	0x8b, 0x74, 0xc3, 0x45, 0x72, 0x2b, 0x79, 0x5a, 0x56, 0x5c, 0xb5, 0x7e, 0x0c, 0x76, 0xb8, 0x88,
	0x88, 0x08, 0x69, 0x64, 0x0b, 0x6f, 0x6b, 0xfd, 0xcf, 0x48, 0x57, 0x82, 0x63, 0x92, 0x54, 0xd5,
	0x5b, 0x0d, 0x42, 0xd0, 0x8c, 0xb9, 0x54, 0x16, 0x45, 0xcb, 0x1e, 0xdb, 0x08, 0x32, 0x11, 0xfc,
	0x3f, 0x82, 0x55, 0x79, 0x22, 0x45, 0x09, 0xb1, 0x00, 0x5a, 0x86, 0x3f, 0x82, 0x6e, 0x2e, 0x58,
	0x28, 0x73, 0x8c, 0x89, 0x94, 0x16, 0x02, 0xe4, 0x82, 0x3d, 0x18, 0x4b, 0x59, 0x59, 0x19, 0x30,
	0x43, 0x94, 0x59, 0xa8, 0x76, 0x2e, 0xd8, 0x0d, 0xa2, 0xcc, 0xfb, 0x58, 0x07, 0x27, 0x1b, 0x7b,
	0xb2, 0x77, 0xd8, 0x03, 0xf5, 0x65, 0x3b, 0x75, 0x1a, 0x41, 0x07, 0xb4, 0x11, 0xc6, 0x3c, 0x4f,
	0x95, 0xc5, 0xa9, 0x54, 0xf8, 0x3d, 0x68, 0xc7, 0x48, 0x86, 0x4f, 0xc8, 0xb4, 0xb3, 0x13, 0xb4,
	0x62, 0x24, 0xff, 0x41, 0x0a, 0xfa, 0x60, 0xbf, 0x74, 0xe4, 0x92, 0x88, 0x10, 0xf3, 0x24, 0xa1,
	0x52, 0x52, 0x9e, 0x3a, 0x4d, 0x1d, 0xf4, 0x5d, 0x8c, 0xe4, 0xa3, 0x24, 0x62, 0xb4, 0x74, 0xc0,
	0x2b, 0xd0, 0xce, 0x4c, 0xb7, 0xce, 0x76, 0xbf, 0x76, 0xd6, 0x1d, 0xf6, 0xd7, 0x58, 0xb4, 0xe1,
	0x56, 0x82, 0xea, 0x00, 0xbc, 0x05, 0x7b, 0x99, 0x09, 0x0b, 0x13, 0xa2, 0x62, 0x1e, 0x49, 0xa7,
	0xa5, 0x99, 0x78, 0xea, 0x57, 0xe4, 0x5f, 0x49, 0x63, 0xc5, 0xbf, 0x74, 0x64, 0xd0, 0xcb, 0x56,
	0x55, 0x09, 0x2f, 0x80, 0x43, 0x53, 0x46, 0x53, 0x12, 0xce, 0xb8, 0x48, 0xc2, 0x35, 0xca, 0xb4,
	0x75, 0xef, 0x87, 0xc6, 0x5f, 0xa6, 0x0a, 0xd6, 0xc9, 0xa3, 0xf8, 0x82, 0xa4, 0xce, 0x8e, 0x21,
	0x8f, 0x56, 0xbc, 0x0f, 0x2f, 0x2b, 0x73, 0xc7, 0x15, 0x9d, 0x15, 0x5f, 0xc1, 0x1b, 0x1f, 0xb4,
	0x85, 0x89, 0xd2, 0xb7, 0xdd, 0x1d, 0x1e, 0xf8, 0x66, 0x69, 0xfd, 0x6a, 0x69, 0xfd, 0xdf, 0xd3,
	0x22, 0xa8, 0x82, 0x4a, 0x26, 0x08, 0xf4, 0x1c, 0x56, 0x67, 0xcc, 0xac, 0x81, 0x40, 0xcf, 0x15,
	0xd6, 0x0f, 0xa0, 0x23, 0xe9, 0x3c, 0x45, 0x2a, 0x17, 0x44, 0x4f, 0xa0, 0x13, 0xbc, 0x18, 0xbc,
	0xf1, 0x72, 0x93, 0xaa, 0x0a, 0xbf, 0xb0, 0x49, 0xcb, 0x35, 0xa9, 0xaf, 0xac, 0xc9, 0xf0, 0x7d,
	0x03, 0xf4, 0xae, 0xcd, 0x6d, 0x3f, 0x10, 0xf1, 0x44, 0x31, 0x81, 0x23, 0x00, 0xef, 0xcb, 0x9e,
	0xcc, 0x86, 0x4e, 0x04, 0xd7, 0xbc, 0x3c, 0x59, 0x0e, 0x65, 0xc5, 0x69, 0x4b, 0x75, 0x7b, 0xeb,
	0x4e, 0x6f, 0x0b, 0xe2, 0xb7, 0x36, 0x43, 0x27, 0x7b, 0x9b, 0x25, 0x55, 0xc6, 0xd3, 0xcf, 0x44,
	0xd8, 0x77, 0x6d, 0x0b, 0xfe, 0xfd, 0xea, 0x61, 0xab, 0xd2, 0xbb, 0x6f, 0x3f, 0x65, 0xee, 0xc9,
	0x46, 0xdf, 0x32, 0xe5, 0x03, 0x38, 0xaa, 0x5c, 0x88, 0xb1, 0x29, 0xc2, 0x8b, 0xcd, 0x49, 0xd7,
	0x68, 0xf1, 0x2a, 0xe9, 0xfa, 0x40, 0xbc, 0x2d, 0xf8, 0x0b, 0xd8, 0x0d, 0xc8, 0x34, 0xa7, 0x2c,
	0x1a, 0x21, 0x1c, 0x13, 0x08, 0x4d, 0xf8, 0xea, 0xb3, 0xed, 0xee, 0xaf, 0xd9, 0xaa, 0xa3, 0xd7,
	0xbf, 0xfd, 0xfb, 0xeb, 0x9c, 0xaa, 0x38, 0x9f, 0xfa, 0x98, 0x27, 0x83, 0x49, 0xc9, 0x24, 0xcc,
	0xd9, 0xfd, 0xdd, 0x78, 0x90, 0xa1, 0x82, 0xa7, 0xe4, 0xdc, 0x5e, 0xfc, 0xb9, 0x34, 0xd3, 0x1b,
	0x64, 0x8b, 0xb9, 0xf9, 0x4a, 0xe8, 0x6f, 0xcb, 0xb4, 0xa5, 0xe5, 0x9f, 0x3f, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x83, 0xbd, 0x0d, 0xc8, 0x6f, 0x06, 0x00, 0x00,
}
