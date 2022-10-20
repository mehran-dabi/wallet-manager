// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "wallet-manager/domain/wallet/dto"

	mock "github.com/stretchr/testify/mock"
)

// IWalletService is an autogenerated mock type for the IWalletService type
type IWalletService struct {
	mock.Mock
}

// AddFund provides a mock function with given fields: ctx, ID, fund
func (_m *IWalletService) AddFund(ctx context.Context, ID int64, fund uint64) (*dto.Wallet, error) {
	ret := _m.Called(ctx, ID, fund)

	var r0 *dto.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint64) *dto.Wallet); ok {
		r0 = rf(ctx, ID, fund)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, uint64) error); ok {
		r1 = rf(ctx, ID, fund)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, userID
func (_m *IWalletService) Create(ctx context.Context, userID int64) (*dto.Wallet, error) {
	ret := _m.Called(ctx, userID)

	var r0 *dto.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, int64) *dto.Wallet); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *IWalletService) GetByID(ctx context.Context, ID int64) (*dto.Wallet, error) {
	ret := _m.Called(ctx, ID)

	var r0 *dto.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, int64) *dto.Wallet); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: ctx, userID
func (_m *IWalletService) GetByUserID(ctx context.Context, userID int64) (*dto.Wallet, error) {
	ret := _m.Called(ctx, userID)

	var r0 *dto.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, int64) *dto.Wallet); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubtractFund provides a mock function with given fields: ctx, ID, fund
func (_m *IWalletService) SubtractFund(ctx context.Context, ID int64, fund uint64) (*dto.Wallet, error) {
	ret := _m.Called(ctx, ID, fund)

	var r0 *dto.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint64) *dto.Wallet); ok {
		r0 = rf(ctx, ID, fund)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, uint64) error); ok {
		r1 = rf(ctx, ID, fund)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIWalletService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIWalletService creates a new instance of IWalletService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIWalletService(t mockConstructorTestingTNewIWalletService) *IWalletService {
	mock := &IWalletService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}