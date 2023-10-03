// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	access "github.com/goravel/framework/contracts/auth/access"

	mock "github.com/stretchr/testify/mock"
)

// Gate is an autogenerated mock type for the Gate type
type Gate struct {
	mock.Mock
}

// After provides a mock function with given fields: callback
func (_m *Gate) After(callback func(context.Context, string, map[string]interface{}, access.Response) access.Response) {
	_m.Called(callback)
}

// Allows provides a mock function with given fields: ability, arguments
func (_m *Gate) Allows(ability string, arguments map[string]interface{}) bool {
	ret := _m.Called(ability, arguments)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) bool); ok {
		r0 = rf(ability, arguments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Any provides a mock function with given fields: abilities, arguments
func (_m *Gate) Any(abilities []string, arguments map[string]interface{}) bool {
	ret := _m.Called(abilities, arguments)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]string, map[string]interface{}) bool); ok {
		r0 = rf(abilities, arguments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Before provides a mock function with given fields: callback
func (_m *Gate) Before(callback func(context.Context, string, map[string]interface{}) access.Response) {
	_m.Called(callback)
}

// Define provides a mock function with given fields: ability, callback
func (_m *Gate) Define(ability string, callback func(context.Context, map[string]interface{}) access.Response) {
	_m.Called(ability, callback)
}

// Denies provides a mock function with given fields: ability, arguments
func (_m *Gate) Denies(ability string, arguments map[string]interface{}) bool {
	ret := _m.Called(ability, arguments)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) bool); ok {
		r0 = rf(ability, arguments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Inspect provides a mock function with given fields: ability, arguments
func (_m *Gate) Inspect(ability string, arguments map[string]interface{}) access.Response {
	ret := _m.Called(ability, arguments)

	var r0 access.Response
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) access.Response); ok {
		r0 = rf(ability, arguments)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(access.Response)
		}
	}

	return r0
}

// None provides a mock function with given fields: abilities, arguments
func (_m *Gate) None(abilities []string, arguments map[string]interface{}) bool {
	ret := _m.Called(abilities, arguments)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]string, map[string]interface{}) bool); ok {
		r0 = rf(abilities, arguments)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// WithContext provides a mock function with given fields: ctx
func (_m *Gate) WithContext(ctx context.Context) access.Gate {
	ret := _m.Called(ctx)

	var r0 access.Gate
	if rf, ok := ret.Get(0).(func(context.Context) access.Gate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(access.Gate)
		}
	}

	return r0
}

// NewGate creates a new instance of Gate. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGate(t interface {
	mock.TestingT
	Cleanup(func())
}) *Gate {
	mock := &Gate{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
