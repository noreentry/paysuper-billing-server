// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: grpc/grpc.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc/grpc.proto

It has these top-level messages:
	EmptyRequest
	EmptyResponse
	PaymentCreateRequest
	PaymentCreateResponse
	FindByStringValue
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import billing "github.com/ProtocolONE/payone-billing-service/pkg/proto/billing"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = billing.PaymentFormPaymentMethods{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for BillingService service

type BillingService interface {
	OrderCreateProcess(ctx context.Context, in *billing.OrderCreateRequest, opts ...client.CallOption) (*billing.Order, error)
	PaymentFormJsonDataProcess(ctx context.Context, in *FindByStringValue, opts ...client.CallOption) (*billing.PaymentFormPaymentMethods, error)
	PaymentCreateProcess(ctx context.Context, in *PaymentCreateRequest, opts ...client.CallOption) (*PaymentCreateResponse, error)
	RebuildCache(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyResponse, error)
}

type billingService struct {
	c    client.Client
	name string
}

func NewBillingService(name string, c client.Client) BillingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "grpc"
	}
	return &billingService{
		c:    c,
		name: name,
	}
}

func (c *billingService) OrderCreateProcess(ctx context.Context, in *billing.OrderCreateRequest, opts ...client.CallOption) (*billing.Order, error) {
	req := c.c.NewRequest(c.name, "BillingService.OrderCreateProcess", in)
	out := new(billing.Order)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) PaymentFormJsonDataProcess(ctx context.Context, in *FindByStringValue, opts ...client.CallOption) (*billing.PaymentFormPaymentMethods, error) {
	req := c.c.NewRequest(c.name, "BillingService.PaymentFormJsonDataProcess", in)
	out := new(billing.PaymentFormPaymentMethods)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) PaymentCreateProcess(ctx context.Context, in *PaymentCreateRequest, opts ...client.CallOption) (*PaymentCreateResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.PaymentCreateProcess", in)
	out := new(PaymentCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingService) RebuildCache(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "BillingService.RebuildCache", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BillingService service

type BillingServiceHandler interface {
	OrderCreateProcess(context.Context, *billing.OrderCreateRequest, *billing.Order) error
	PaymentFormJsonDataProcess(context.Context, *FindByStringValue, *billing.PaymentFormPaymentMethods) error
	PaymentCreateProcess(context.Context, *PaymentCreateRequest, *PaymentCreateResponse) error
	RebuildCache(context.Context, *EmptyRequest, *EmptyResponse) error
}

func RegisterBillingServiceHandler(s server.Server, hdlr BillingServiceHandler, opts ...server.HandlerOption) error {
	type billingService interface {
		OrderCreateProcess(ctx context.Context, in *billing.OrderCreateRequest, out *billing.Order) error
		PaymentFormJsonDataProcess(ctx context.Context, in *FindByStringValue, out *billing.PaymentFormPaymentMethods) error
		PaymentCreateProcess(ctx context.Context, in *PaymentCreateRequest, out *PaymentCreateResponse) error
		RebuildCache(ctx context.Context, in *EmptyRequest, out *EmptyResponse) error
	}
	type BillingService struct {
		billingService
	}
	h := &billingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BillingService{h}, opts...))
}

type billingServiceHandler struct {
	BillingServiceHandler
}

func (h *billingServiceHandler) OrderCreateProcess(ctx context.Context, in *billing.OrderCreateRequest, out *billing.Order) error {
	return h.BillingServiceHandler.OrderCreateProcess(ctx, in, out)
}

func (h *billingServiceHandler) PaymentFormJsonDataProcess(ctx context.Context, in *FindByStringValue, out *billing.PaymentFormPaymentMethods) error {
	return h.BillingServiceHandler.PaymentFormJsonDataProcess(ctx, in, out)
}

func (h *billingServiceHandler) PaymentCreateProcess(ctx context.Context, in *PaymentCreateRequest, out *PaymentCreateResponse) error {
	return h.BillingServiceHandler.PaymentCreateProcess(ctx, in, out)
}

func (h *billingServiceHandler) RebuildCache(ctx context.Context, in *EmptyRequest, out *EmptyResponse) error {
	return h.BillingServiceHandler.RebuildCache(ctx, in, out)
}
