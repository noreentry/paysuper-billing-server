// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/grpc.proto

package grpc // import "github.com/ProtocolONE/payone-billing-service/pkg/proto/grpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/ProtocolONE/payone-billing-service/pkg/proto/billing"

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
	return fileDescriptor_grpc_21bce49ce84bbd73, []int{0}
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
	return fileDescriptor_grpc_21bce49ce84bbd73, []int{1}
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
	return fileDescriptor_grpc_21bce49ce84bbd73, []int{2}
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
	return fileDescriptor_grpc_21bce49ce84bbd73, []int{3}
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

type FindByStringValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *FindByStringValue) Reset()         { *m = FindByStringValue{} }
func (m *FindByStringValue) String() string { return proto.CompactTextString(m) }
func (*FindByStringValue) ProtoMessage()    {}
func (*FindByStringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_21bce49ce84bbd73, []int{4}
}
func (m *FindByStringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByStringValue.Unmarshal(m, b)
}
func (m *FindByStringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByStringValue.Marshal(b, m, deterministic)
}
func (dst *FindByStringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByStringValue.Merge(dst, src)
}
func (m *FindByStringValue) XXX_Size() int {
	return xxx_messageInfo_FindByStringValue.Size(m)
}
func (m *FindByStringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByStringValue.DiscardUnknown(m)
}

var xxx_messageInfo_FindByStringValue proto.InternalMessageInfo

func (m *FindByStringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "grpc.EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "grpc.EmptyResponse")
	proto.RegisterType((*PaymentCreateRequest)(nil), "grpc.PaymentCreateRequest")
	proto.RegisterMapType((map[string]string)(nil), "grpc.PaymentCreateRequest.DataEntry")
	proto.RegisterType((*PaymentCreateResponse)(nil), "grpc.PaymentCreateResponse")
	proto.RegisterType((*FindByStringValue)(nil), "grpc.FindByStringValue")
}

func init() { proto.RegisterFile("grpc/grpc.proto", fileDescriptor_grpc_21bce49ce84bbd73) }

var fileDescriptor_grpc_21bce49ce84bbd73 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0xb6, 0x52, 0xa7, 0x21, 0x85, 0xa5, 0x85, 0xc8, 0xbd, 0x14, 0x8b, 0x43, 0x39,
	0xd4, 0x96, 0xca, 0x81, 0x82, 0x10, 0x87, 0x84, 0xf4, 0x80, 0x04, 0x0d, 0xae, 0xe0, 0x00, 0x07,
	0xb4, 0xb1, 0x47, 0xce, 0xaa, 0xce, 0xae, 0x99, 0x5d, 0x57, 0xf2, 0x4f, 0xe0, 0x5f, 0xf1, 0xd3,
	0x90, 0x77, 0xd7, 0x25, 0x2d, 0xe1, 0x62, 0xcf, 0xc7, 0xce, 0x9b, 0xb7, 0x6f, 0x1f, 0xec, 0x17,
	0x54, 0x65, 0x49, 0xfb, 0x89, 0x2b, 0x52, 0x46, 0xb1, 0xad, 0x36, 0x0e, 0x0f, 0x17, 0xa2, 0x2c,
	0x85, 0x2c, 0x12, 0xff, 0x77, 0xcd, 0x68, 0x04, 0xc3, 0xd9, 0xaa, 0x32, 0x4d, 0x8a, 0x3f, 0x6b,
	0xd4, 0x26, 0xda, 0x87, 0x07, 0x3e, 0xd7, 0x95, 0x92, 0x1a, 0xa3, 0x5f, 0x01, 0x1c, 0xcc, 0x79,
	0xb3, 0x42, 0x69, 0xa6, 0x84, 0xdc, 0xa0, 0x3f, 0xc9, 0xce, 0x61, 0x2b, 0xe7, 0x86, 0x8f, 0x83,
	0xe3, 0xc1, 0xc9, 0xde, 0xd9, 0xf3, 0xd8, 0x6e, 0xdc, 0x74, 0x32, 0x7e, 0xcf, 0x0d, 0x9f, 0x49,
	0x43, 0x4d, 0x6a, 0x27, 0xc2, 0x57, 0xb0, 0x7b, 0x5b, 0x62, 0x0f, 0x61, 0x70, 0x8d, 0xcd, 0x38,
	0x38, 0x0e, 0x4e, 0x76, 0xd3, 0x36, 0x64, 0x07, 0xb0, 0x7d, 0xc3, 0xcb, 0x1a, 0xc7, 0x7d, 0x5b,
	0x73, 0xc9, 0x9b, 0xfe, 0x79, 0x10, 0x2d, 0xe1, 0xf0, 0xde, 0x02, 0x47, 0x92, 0x3d, 0x81, 0x1d,
	0x6d, 0xb8, 0xa9, 0xb5, 0xc5, 0xd9, 0x4e, 0x7d, 0xc6, 0x9e, 0xc1, 0x90, 0x30, 0x17, 0x84, 0x99,
	0xf9, 0x51, 0x53, 0xe9, 0x11, 0xf7, 0xba, 0xda, 0x17, 0x2a, 0xdb, 0x6d, 0x48, 0xa4, 0x68, 0x3c,
	0x70, 0xdb, 0x6c, 0x12, 0xbd, 0x80, 0x47, 0x17, 0x42, 0xe6, 0x93, 0xe6, 0xca, 0x90, 0x90, 0xc5,
	0xd7, 0x96, 0xc2, 0x5f, 0x62, 0xc1, 0x1a, 0xb1, 0xb3, 0xdf, 0x7d, 0x18, 0x4d, 0x9c, 0xa6, 0x57,
	0x48, 0x37, 0x22, 0x43, 0x36, 0x05, 0x76, 0x49, 0x39, 0x92, 0x63, 0x39, 0x27, 0x95, 0xa1, 0xd6,
	0xec, 0x28, 0xee, 0xa4, 0x5f, 0x6b, 0x7a, 0x8d, 0xc2, 0xd1, 0xdd, 0x66, 0xd4, 0x63, 0xdf, 0x21,
	0xf4, 0x97, 0xbd, 0x50, 0xb4, 0xfa, 0xa0, 0x95, 0x6c, 0x45, 0xeb, 0xc0, 0x9e, 0x3a, 0xbd, 0xff,
	0x21, 0x19, 0x46, 0xb7, 0x40, 0x6b, 0xd3, 0x3e, 0xfc, 0x88, 0x66, 0xa9, 0x72, 0x1d, 0xf5, 0xd8,
	0xe7, 0x7b, 0x8f, 0xda, 0xc1, 0x86, 0xff, 0x7f, 0xc6, 0xf0, 0x68, 0x63, 0xcf, 0xdb, 0xa4, 0xc7,
	0x5e, 0xc3, 0x30, 0xc5, 0x45, 0x2d, 0xca, 0x7c, 0xca, 0xb3, 0x25, 0x32, 0xe6, 0x8e, 0xaf, 0xbb,
	0x2b, 0x7c, 0x7c, 0xa7, 0xd6, 0x8d, 0x4e, 0xde, 0x7d, 0x7b, 0x5b, 0x08, 0xb3, 0xac, 0x17, 0x71,
	0xa6, 0x56, 0xc9, 0xbc, 0x35, 0x66, 0xa6, 0xca, 0xcb, 0x4f, 0xb3, 0xa4, 0xe2, 0x8d, 0x92, 0x78,
	0xea, 0xaf, 0x74, 0xaa, 0x9d, 0xc0, 0x49, 0x75, 0x5d, 0x24, 0xd6, 0xbf, 0xd6, 0xe7, 0x8b, 0x1d,
	0x1b, 0xbf, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x38, 0xea, 0x99, 0xfb, 0x02, 0x00, 0x00,
}
