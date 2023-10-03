// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	filesystem "github.com/goravel/framework/contracts/filesystem"
	http "github.com/goravel/framework/contracts/http"

	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"

	validation "github.com/goravel/framework/contracts/validation"
)

// ContextRequest is an autogenerated mock type for the ContextRequest type
type ContextRequest struct {
	mock.Mock
}

// AbortWithStatus provides a mock function with given fields: code
func (_m *ContextRequest) AbortWithStatus(code int) {
	_m.Called(code)
}

// AbortWithStatusJson provides a mock function with given fields: code, jsonObj
func (_m *ContextRequest) AbortWithStatusJson(code int, jsonObj interface{}) {
	_m.Called(code, jsonObj)
}

// All provides a mock function with given fields:
func (_m *ContextRequest) All() map[string]interface{} {
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

// Bind provides a mock function with given fields: obj
func (_m *ContextRequest) Bind(obj interface{}) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// File provides a mock function with given fields: name
func (_m *ContextRequest) File(name string) (filesystem.File, error) {
	ret := _m.Called(name)

	var r0 filesystem.File
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (filesystem.File, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) filesystem.File); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.File)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FullUrl provides a mock function with given fields:
func (_m *ContextRequest) FullUrl() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Header provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) Header(key string, defaultValue ...string) string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Headers provides a mock function with given fields:
func (_m *ContextRequest) Headers() nethttp.Header {
	ret := _m.Called()

	var r0 nethttp.Header
	if rf, ok := ret.Get(0).(func() nethttp.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nethttp.Header)
		}
	}

	return r0
}

// Host provides a mock function with given fields:
func (_m *ContextRequest) Host() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Input provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) Input(key string, defaultValue ...string) string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// InputArray provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) InputArray(key string, defaultValue ...[]string) []string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, ...[]string) []string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// InputBool provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) InputBool(key string, defaultValue ...bool) bool {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...bool) bool); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InputInt provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) InputInt(key string, defaultValue ...int) int {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// InputInt64 provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) InputInt64(key string, defaultValue ...int64) int64 {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...int64) int64); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// InputMap provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) InputMap(key string, defaultValue ...map[string]string) map[string]string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, ...map[string]string) map[string]string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// Ip provides a mock function with given fields:
func (_m *ContextRequest) Ip() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Method provides a mock function with given fields:
func (_m *ContextRequest) Method() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Next provides a mock function with given fields:
func (_m *ContextRequest) Next() {
	_m.Called()
}

// Origin provides a mock function with given fields:
func (_m *ContextRequest) Origin() *nethttp.Request {
	ret := _m.Called()

	var r0 *nethttp.Request
	if rf, ok := ret.Get(0).(func() *nethttp.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nethttp.Request)
		}
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *ContextRequest) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Queries provides a mock function with given fields:
func (_m *ContextRequest) Queries() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// Query provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) Query(key string, defaultValue ...string) string {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// QueryArray provides a mock function with given fields: key
func (_m *ContextRequest) QueryArray(key string) []string {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// QueryBool provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) QueryBool(key string, defaultValue ...bool) bool {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...bool) bool); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// QueryInt provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) QueryInt(key string, defaultValue ...int) int {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, ...int) int); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// QueryInt64 provides a mock function with given fields: key, defaultValue
func (_m *ContextRequest) QueryInt64(key string, defaultValue ...int64) int64 {
	_va := make([]interface{}, len(defaultValue))
	for _i := range defaultValue {
		_va[_i] = defaultValue[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...int64) int64); ok {
		r0 = rf(key, defaultValue...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// QueryMap provides a mock function with given fields: key
func (_m *ContextRequest) QueryMap(key string) map[string]string {
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

// Route provides a mock function with given fields: key
func (_m *ContextRequest) Route(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RouteInt provides a mock function with given fields: key
func (_m *ContextRequest) RouteInt(key string) int {
	ret := _m.Called(key)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// RouteInt64 provides a mock function with given fields: key
func (_m *ContextRequest) RouteInt64(key string) int64 {
	ret := _m.Called(key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Url provides a mock function with given fields:
func (_m *ContextRequest) Url() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Validate provides a mock function with given fields: rules, options
func (_m *ContextRequest) Validate(rules map[string]string, options ...validation.Option) (validation.Validator, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, rules)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 validation.Validator
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]string, ...validation.Option) (validation.Validator, error)); ok {
		return rf(rules, options...)
	}
	if rf, ok := ret.Get(0).(func(map[string]string, ...validation.Option) validation.Validator); ok {
		r0 = rf(rules, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Validator)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]string, ...validation.Option) error); ok {
		r1 = rf(rules, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateRequest provides a mock function with given fields: request
func (_m *ContextRequest) ValidateRequest(request http.FormRequest) (validation.Errors, error) {
	ret := _m.Called(request)

	var r0 validation.Errors
	var r1 error
	if rf, ok := ret.Get(0).(func(http.FormRequest) (validation.Errors, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(http.FormRequest) validation.Errors); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Errors)
		}
	}

	if rf, ok := ret.Get(1).(func(http.FormRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewContextRequest creates a new instance of ContextRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContextRequest(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContextRequest {
	mock := &ContextRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
