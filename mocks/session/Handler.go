// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Handler) Close() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Destroy provides a mock function with given fields: id
func (_m *Handler) Destroy(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Gc provides a mock function with given fields: maxLifetime
func (_m *Handler) Gc(maxLifetime int) (int, bool) {
	ret := _m.Called(maxLifetime)

	var r0 int
	var r1 bool
	if rf, ok := ret.Get(0).(func(int) (int, bool)); ok {
		return rf(maxLifetime)
	}
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(maxLifetime)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int) bool); ok {
		r1 = rf(maxLifetime)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Open provides a mock function with given fields: path, name
func (_m *Handler) Open(path string, name string) bool {
	ret := _m.Called(path, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(path, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Read provides a mock function with given fields: id
func (_m *Handler) Read(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Write provides a mock function with given fields: id, data
func (_m *Handler) Write(id string, data string) error {
	ret := _m.Called(id, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
