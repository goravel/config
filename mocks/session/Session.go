// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	session "github.com/goravel/framework/contracts/session"
	mock "github.com/stretchr/testify/mock"
)

// Session is an autogenerated mock type for the Session type
type Session struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *Session) All() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// Exists provides a mock function with given fields: key
func (_m *Session) Exists(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Forget provides a mock function with given fields: keys
func (_m *Session) Forget(keys ...string) session.Session {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 session.Session
	if rf, ok := ret.Get(0).(func(...string) session.Session); ok {
		r0 = rf(keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(session.Session)
		}
	}

	return r0
}

// Get provides a mock function with given fields: key, defaultValue
func (_m *Session) Get(key string, defaultValue ...interface{}) interface{} {
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, defaultValue...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(key, defaultValue...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetId provides a mock function with given fields:
func (_m *Session) GetId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *Session) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Has provides a mock function with given fields: key
func (_m *Session) Has(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Missing provides a mock function with given fields: key
func (_m *Session) Missing(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Put provides a mock function with given fields: key, value
func (_m *Session) Put(key string, value interface{}) session.Session {
	ret := _m.Called(key, value)

	var r0 session.Session
	if rf, ok := ret.Get(0).(func(string, interface{}) session.Session); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(session.Session)
		}
	}

	return r0
}

// RegenerateToken provides a mock function with given fields:
func (_m *Session) RegenerateToken() session.Session {
	ret := _m.Called()

	var r0 session.Session
	if rf, ok := ret.Get(0).(func() session.Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(session.Session)
		}
	}

	return r0
}

// Save provides a mock function with given fields:
func (_m *Session) Save() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetId provides a mock function with given fields: id
func (_m *Session) SetId(id string) session.Session {
	ret := _m.Called(id)

	var r0 session.Session
	if rf, ok := ret.Get(0).(func(string) session.Session); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(session.Session)
		}
	}

	return r0
}

// SetName provides a mock function with given fields: name
func (_m *Session) SetName(name string) session.Session {
	ret := _m.Called(name)

	var r0 session.Session
	if rf, ok := ret.Get(0).(func(string) session.Session); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(session.Session)
		}
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Session) Start() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewSession creates a new instance of Session. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSession(t interface {
	mock.TestingT
	Cleanup(func())
}) *Session {
	mock := &Session{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
