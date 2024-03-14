// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// DispatchesEvents is an autogenerated mock type for the DispatchesEvents type
type DispatchesEvents struct {
	mock.Mock
}

// DispatchesEvents provides a mock function with given fields:
func (_m *DispatchesEvents) DispatchesEvents() map[orm.EventType]func(orm.Event) error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DispatchesEvents")
	}

	var r0 map[orm.EventType]func(orm.Event) error
	if rf, ok := ret.Get(0).(func() map[orm.EventType]func(orm.Event) error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[orm.EventType]func(orm.Event) error)
		}
	}

	return r0
}

// NewDispatchesEvents creates a new instance of DispatchesEvents. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDispatchesEvents(t interface {
	mock.TestingT
	Cleanup(func())
}) *DispatchesEvents {
	mock := &DispatchesEvents{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
