// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Errors is an autogenerated mock type for the Errors type
type Errors struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *Errors) All() map[string]map[string]string {
	ret := _m.Called()

	var r0 map[string]map[string]string
	if rf, ok := ret.Get(0).(func() map[string]map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]string)
		}
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *Errors) Get(key string) map[string]string {
	ret := _m.Called(key)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// Has provides a mock function with given fields: key
func (_m *Errors) Has(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// One provides a mock function with given fields: key
func (_m *Errors) One(key ...string) string {
	_va := make([]interface{}, len(key))
	for _i := range key {
		_va[_i] = key[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(key...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewErrors interface {
	mock.TestingT
	Cleanup(func())
}

// NewErrors creates a new instance of Errors. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewErrors(t mockConstructorTestingTNewErrors) *Errors {
	mock := &Errors{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
