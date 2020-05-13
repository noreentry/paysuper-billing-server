// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import billingpb "github.com/paysuper/paysuper-proto/go/billingpb"
import context "context"
import mock "github.com/stretchr/testify/mock"

// PayoutRepositoryInterface is an autogenerated mock type for the PayoutRepositoryInterface type
type PayoutRepositoryInterface struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, merchantId, statuses, dateFrom, dateTo, limit, offset
func (_m *PayoutRepositoryInterface) Find(ctx context.Context, merchantId string, statuses []string, dateFrom string, dateTo string, limit int64, offset int64) ([]*billingpb.PayoutDocument, error) {
	ret := _m.Called(ctx, merchantId, statuses, dateFrom, dateTo, limit, offset)

	var r0 []*billingpb.PayoutDocument
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, string, string, int64, int64) []*billingpb.PayoutDocument); ok {
		r0 = rf(ctx, merchantId, statuses, dateFrom, dateTo, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*billingpb.PayoutDocument)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, string, string, int64, int64) error); ok {
		r1 = rf(ctx, merchantId, statuses, dateFrom, dateTo, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCount provides a mock function with given fields: ctx, merchantId, statuses, dateFrom, dateTo
func (_m *PayoutRepositoryInterface) FindCount(ctx context.Context, merchantId string, statuses []string, dateFrom string, dateTo string) (int64, error) {
	ret := _m.Called(ctx, merchantId, statuses, dateFrom, dateTo)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, string, string) int64); ok {
		r0 = rf(ctx, merchantId, statuses, dateFrom, dateTo)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, string, string) error); ok {
		r1 = rf(ctx, merchantId, statuses, dateFrom, dateTo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalanceAmount provides a mock function with given fields: _a0, _a1, _a2
func (_m *PayoutRepositoryInterface) GetBalanceAmount(_a0 context.Context, _a1 string, _a2 string) (float64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 float64
	if rf, ok := ret.Get(0).(func(context.Context, string, string) float64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: _a0, _a1
func (_m *PayoutRepositoryInterface) GetById(_a0 context.Context, _a1 string) (*billingpb.PayoutDocument, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billingpb.PayoutDocument
	if rf, ok := ret.Get(0).(func(context.Context, string) *billingpb.PayoutDocument); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.PayoutDocument)
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

// GetByIdMerchantId provides a mock function with given fields: _a0, _a1, _a2
func (_m *PayoutRepositoryInterface) GetByIdMerchantId(_a0 context.Context, _a1 string, _a2 string) (*billingpb.PayoutDocument, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *billingpb.PayoutDocument
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *billingpb.PayoutDocument); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.PayoutDocument)
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

// GetLast provides a mock function with given fields: _a0, _a1, _a2
func (_m *PayoutRepositoryInterface) GetLast(_a0 context.Context, _a1 string, _a2 string) (*billingpb.PayoutDocument, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *billingpb.PayoutDocument
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *billingpb.PayoutDocument); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billingpb.PayoutDocument)
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

// Insert provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *PayoutRepositoryInterface) Insert(_a0 context.Context, _a1 *billingpb.PayoutDocument, _a2 string, _a3 string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.PayoutDocument, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *PayoutRepositoryInterface) Update(_a0 context.Context, _a1 *billingpb.PayoutDocument, _a2 string, _a3 string) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billingpb.PayoutDocument, string, string) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
