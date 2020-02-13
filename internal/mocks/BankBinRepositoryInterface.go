// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import pkg "github.com/paysuper/paysuper-billing-server/internal/pkg"

// BankBinRepositoryInterface is an autogenerated mock type for the BankBinRepositoryInterface type
type BankBinRepositoryInterface struct {
	mock.Mock
}

// GetByBin provides a mock function with given fields: _a0, _a1
func (_m *BankBinRepositoryInterface) GetByBin(_a0 context.Context, _a1 int32) (*pkg.BinData, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pkg.BinData
	if rf, ok := ret.Get(0).(func(context.Context, int32) *pkg.BinData); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pkg.BinData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0, _a1
func (_m *BankBinRepositoryInterface) Insert(_a0 context.Context, _a1 *pkg.BinData) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pkg.BinData) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MultipleInsert provides a mock function with given fields: _a0, _a1
func (_m *BankBinRepositoryInterface) MultipleInsert(_a0 context.Context, _a1 []*pkg.BinData) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*pkg.BinData) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}