// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// CartHandler is an autogenerated mock type for the CartHandler type
type CartHandler struct {
	mock.Mock
}

// AddCart provides a mock function with given fields:
func (_m *CartHandler) AddCart() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// DeleteCart provides a mock function with given fields:
func (_m *CartHandler) DeleteCart() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// MyCart provides a mock function with given fields:
func (_m *CartHandler) MyCart() echo.HandlerFunc {
	ret := _m.Called()

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

type mockConstructorTestingTNewCartHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartHandler creates a new instance of CartHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartHandler(t mockConstructorTestingTNewCartHandler) *CartHandler {
	mock := &CartHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
