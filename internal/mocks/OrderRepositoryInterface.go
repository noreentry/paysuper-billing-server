// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import billingpb "github.com/paysuper/paysuper-proto/go/billingpb"
import context "context"
import mock "github.com/stretchr/testify/mock"

// OrderRepositoryInterface is an autogenerated mock type for the OrderRepositoryInterface type
type OrderRepositoryInterface struct {
	mock.Mock
}

// GetById provides a mock function with given fields: _a0, _a1
func (_m *OrderRepositoryInterface) GetById(_a0 context.Context, _a1 string) (*billingpb.Order, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billingpb.Order
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.Order); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByProjectOrderId provides a mock function with given fields: _a0, _a1, _a2
func (_m *OrderRepositoryInterface) GetByProjectOrderId(_a0 context.Context, _a1 string, _a2 string) (*billingpb.Order, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *billingpb.Order
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *billingpb.Order); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByRefundReceiptNumber provides a mock function with given fields: _a0, _a1
func (_m *OrderRepositoryInterface) GetByRefundReceiptNumber(_a0 context.Context, _a1 string) (*billingpb.Order, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billingpb.Order
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.Order); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUuid provides a mock function with given fields: _a0, _a1
func (_m *OrderRepositoryInterface) GetByUuid(_a0 context.Context, _a1 string) (*billingpb.Order, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billingpb.Order
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.Order); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0, _a1
func (_m *OrderRepositoryInterface) Insert(_a0 context.Context, _a1 *billingpb.Order) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.Order) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *OrderRepositoryInterface) Update(_a0 context.Context, _a1 *billingpb.Order) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.Order) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrderView provides a mock function with given fields: _a0, _a1
func (_m *OrderRepositoryInterface) UpdateOrderView(_a0 context.Context, _a1 []string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
