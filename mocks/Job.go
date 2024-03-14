// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Job is an autogenerated mock type for the Job type
type Job struct {
	mock.Mock
}

// Handle provides a mock function with given fields: args
func (_m *Job) Handle(args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Signature provides a mock function with given fields:
func (_m *Job) Signature() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Signature")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewJob creates a new instance of Job. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJob(t interface {
	mock.TestingT
	Cleanup(func())
}) *Job {
	mock := &Job{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
