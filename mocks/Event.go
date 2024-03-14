// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// Event is an autogenerated mock type for the Event type
type Event struct {
	mock.Mock
}

// Context provides a mock function with given fields:
func (_m *Event) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// GetAttribute provides a mock function with given fields: key
func (_m *Event) GetAttribute(key string) interface{} {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for GetAttribute")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetOriginal provides a mock function with given fields: key, def
func (_m *Event) GetOriginal(key string, def ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, def...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetOriginal")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(key, def...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// IsClean provides a mock function with given fields: columns
func (_m *Event) IsClean(columns ...string) bool {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IsClean")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = rf(columns...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsDirty provides a mock function with given fields: columns
func (_m *Event) IsDirty(columns ...string) bool {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for IsDirty")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = rf(columns...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Query provides a mock function with given fields:
func (_m *Event) Query() orm.Query {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// SetAttribute provides a mock function with given fields: key, value
func (_m *Event) SetAttribute(key string, value interface{}) {
	_m.Called(key, value)
}

// NewEvent creates a new instance of Event. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEvent(t interface {
	mock.TestingT
	Cleanup(func())
}) *Event {
	mock := &Event{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
