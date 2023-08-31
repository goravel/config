// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	http "github.com/goravel/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"

	nethttp "net/http"

	route "github.com/goravel/framework/contracts/route"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

// Any provides a mock function with given fields: relativePath, handler
func (_m *Router) Any(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Delete provides a mock function with given fields: relativePath, handler
func (_m *Router) Delete(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Get provides a mock function with given fields: relativePath, handler
func (_m *Router) Get(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Group provides a mock function with given fields: handler
func (_m *Router) Group(handler route.GroupFunc) {
	_m.Called(handler)
}

// Middleware provides a mock function with given fields: middlewares
func (_m *Router) Middleware(middlewares ...http.Middleware) route.Router {
	_va := make([]interface{}, len(middlewares))
	for _i := range middlewares {
		_va[_i] = middlewares[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 route.Router
	if rf, ok := ret.Get(0).(func(...http.Middleware) route.Router); ok {
		r0 = rf(middlewares...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Router)
		}
	}

	return r0
}

// Options provides a mock function with given fields: relativePath, handler
func (_m *Router) Options(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Patch provides a mock function with given fields: relativePath, handler
func (_m *Router) Patch(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Post provides a mock function with given fields: relativePath, handler
func (_m *Router) Post(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Prefix provides a mock function with given fields: addr
func (_m *Router) Prefix(addr string) route.Router {
	ret := _m.Called(addr)

	var r0 route.Router
	if rf, ok := ret.Get(0).(func(string) route.Router); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Router)
		}
	}

	return r0
}

// Put provides a mock function with given fields: relativePath, handler
func (_m *Router) Put(relativePath string, handler http.HandlerFunc) {
	_m.Called(relativePath, handler)
}

// Resource provides a mock function with given fields: relativePath, controller
func (_m *Router) Resource(relativePath string, controller http.ResourceController) {
	_m.Called(relativePath, controller)
}

// Static provides a mock function with given fields: relativePath, root
func (_m *Router) Static(relativePath string, root string) {
	_m.Called(relativePath, root)
}

// StaticFS provides a mock function with given fields: relativePath, fs
func (_m *Router) StaticFS(relativePath string, fs nethttp.FileSystem) {
	_m.Called(relativePath, fs)
}

// StaticFile provides a mock function with given fields: relativePath, filepath
func (_m *Router) StaticFile(relativePath string, filepath string) {
	_m.Called(relativePath, filepath)
}

// NewRouter creates a new instance of Router. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRouter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Router {
	mock := &Router{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
