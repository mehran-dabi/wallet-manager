// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// IWalletController is an autogenerated mock type for the IWalletController type
type IWalletController struct {
	mock.Mock
}

// AddFund provides a mock function with given fields: c
func (_m *IWalletController) AddFund(c *gin.Context) {
	_m.Called(c)
}

// Create provides a mock function with given fields: c
func (_m *IWalletController) Create(c *gin.Context) {
	_m.Called(c)
}

// GetByID provides a mock function with given fields: c
func (_m *IWalletController) GetByID(c *gin.Context) {
	_m.Called(c)
}

// GetByUserID provides a mock function with given fields: c
func (_m *IWalletController) GetByUserID(c *gin.Context) {
	_m.Called(c)
}

// SubtractFund provides a mock function with given fields: c
func (_m *IWalletController) SubtractFund(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewIWalletController interface {
	mock.TestingT
	Cleanup(func())
}

// NewIWalletController creates a new instance of IWalletController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIWalletController(t mockConstructorTestingTNewIWalletController) *IWalletController {
	mock := &IWalletController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}