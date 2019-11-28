// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import billing "github.com/paysuper/paysuper-billing-server/pkg/proto/billing"
import context "context"
import mock "github.com/stretchr/testify/mock"
import pkg "github.com/paysuper/paysuper-billing-server/internal/pkg"

// CountryServiceInterface is an autogenerated mock type for the CountryServiceInterface type
type CountryServiceInterface struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: _a0
func (_m *CountryServiceInterface) GetAll(_a0 context.Context) (*billing.CountriesList, error) {
	ret := _m.Called(_a0)

	var r0 *billing.CountriesList
	if rf, ok := ret.Get(0).(func(context.Context) *billing.CountriesList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.CountriesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIsoCodeA2 provides a mock function with given fields: _a0, _a1
func (_m *CountryServiceInterface) GetByIsoCodeA2(_a0 context.Context, _a1 string) (*billing.Country, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *billing.Country
	if rf, ok := ret.Get(0).(func(context.Context, string) *billing.Country); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.Country)
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

// GetCountriesAndRegionsByTariffRegion provides a mock function with given fields: ctx, tariffRegion
func (_m *CountryServiceInterface) GetCountriesAndRegionsByTariffRegion(ctx context.Context, tariffRegion string) ([]*pkg.CountryAndRegionItem, error) {
	ret := _m.Called(ctx, tariffRegion)

	var r0 []*pkg.CountryAndRegionItem
	if rf, ok := ret.Get(0).(func(context.Context, string) []*pkg.CountryAndRegionItem); ok {
		r0 = rf(ctx, tariffRegion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pkg.CountryAndRegionItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tariffRegion)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCountriesWithVatEnabled provides a mock function with given fields: _a0
func (_m *CountryServiceInterface) GetCountriesWithVatEnabled(_a0 context.Context) (*billing.CountriesList, error) {
	ret := _m.Called(_a0)

	var r0 *billing.CountriesList
	if rf, ok := ret.Get(0).(func(context.Context) *billing.CountriesList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*billing.CountriesList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0, _a1
func (_m *CountryServiceInterface) Insert(_a0 context.Context, _a1 *billing.Country) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billing.Country) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsRegionExists provides a mock function with given fields: _a0, _a1
func (_m *CountryServiceInterface) IsRegionExists(_a0 context.Context, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTariffRegionExists provides a mock function with given fields: _a0
func (_m *CountryServiceInterface) IsTariffRegionExists(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MultipleInsert provides a mock function with given fields: _a0, _a1
func (_m *CountryServiceInterface) MultipleInsert(_a0 context.Context, _a1 []*billing.Country) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*billing.Country) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *CountryServiceInterface) Update(_a0 context.Context, _a1 *billing.Country) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *billing.Country) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
