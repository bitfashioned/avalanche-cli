// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// StatusChecker is an autogenerated mock type for the StatusChecker type
type StatusChecker struct {
	mock.Mock
}

// GetCurrentNetworkVersion provides a mock function with given fields:
func (_m *StatusChecker) GetCurrentNetworkVersion() (string, int, bool, error) {
	ret := _m.Called()

	var r0 string
	var r1 int
	var r2 bool
	var r3 error
	if rf, ok := ret.Get(0).(func() (string, int, bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func() bool); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(bool)
	}

	if rf, ok := ret.Get(3).(func() error); ok {
		r3 = rf()
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

type mockConstructorTestingTNewStatusChecker interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatusChecker creates a new instance of StatusChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatusChecker(t mockConstructorTestingTNewStatusChecker) *StatusChecker {
	mock := &StatusChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
