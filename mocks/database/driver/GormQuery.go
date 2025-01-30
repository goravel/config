// Code generated by mockery. DO NOT EDIT.

package driver

import (
	clause "gorm.io/gorm/clause"

	mock "github.com/stretchr/testify/mock"
)

// GormQuery is an autogenerated mock type for the GormQuery type
type GormQuery struct {
	mock.Mock
}

type GormQuery_Expecter struct {
	mock *mock.Mock
}

func (_m *GormQuery) EXPECT() *GormQuery_Expecter {
	return &GormQuery_Expecter{mock: &_m.Mock}
}

// LockForUpdate provides a mock function with no fields
func (_m *GormQuery) LockForUpdate() clause.Expression {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LockForUpdate")
	}

	var r0 clause.Expression
	if rf, ok := ret.Get(0).(func() clause.Expression); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clause.Expression)
		}
	}

	return r0
}

// GormQuery_LockForUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LockForUpdate'
type GormQuery_LockForUpdate_Call struct {
	*mock.Call
}

// LockForUpdate is a helper method to define mock.On call
func (_e *GormQuery_Expecter) LockForUpdate() *GormQuery_LockForUpdate_Call {
	return &GormQuery_LockForUpdate_Call{Call: _e.mock.On("LockForUpdate")}
}

func (_c *GormQuery_LockForUpdate_Call) Run(run func()) *GormQuery_LockForUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GormQuery_LockForUpdate_Call) Return(_a0 clause.Expression) *GormQuery_LockForUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GormQuery_LockForUpdate_Call) RunAndReturn(run func() clause.Expression) *GormQuery_LockForUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// RandomOrder provides a mock function with no fields
func (_m *GormQuery) RandomOrder() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RandomOrder")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GormQuery_RandomOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RandomOrder'
type GormQuery_RandomOrder_Call struct {
	*mock.Call
}

// RandomOrder is a helper method to define mock.On call
func (_e *GormQuery_Expecter) RandomOrder() *GormQuery_RandomOrder_Call {
	return &GormQuery_RandomOrder_Call{Call: _e.mock.On("RandomOrder")}
}

func (_c *GormQuery_RandomOrder_Call) Run(run func()) *GormQuery_RandomOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GormQuery_RandomOrder_Call) Return(_a0 string) *GormQuery_RandomOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GormQuery_RandomOrder_Call) RunAndReturn(run func() string) *GormQuery_RandomOrder_Call {
	_c.Call.Return(run)
	return _c
}

// SharedLock provides a mock function with no fields
func (_m *GormQuery) SharedLock() clause.Expression {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SharedLock")
	}

	var r0 clause.Expression
	if rf, ok := ret.Get(0).(func() clause.Expression); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clause.Expression)
		}
	}

	return r0
}

// GormQuery_SharedLock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SharedLock'
type GormQuery_SharedLock_Call struct {
	*mock.Call
}

// SharedLock is a helper method to define mock.On call
func (_e *GormQuery_Expecter) SharedLock() *GormQuery_SharedLock_Call {
	return &GormQuery_SharedLock_Call{Call: _e.mock.On("SharedLock")}
}

func (_c *GormQuery_SharedLock_Call) Run(run func()) *GormQuery_SharedLock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GormQuery_SharedLock_Call) Return(_a0 clause.Expression) *GormQuery_SharedLock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GormQuery_SharedLock_Call) RunAndReturn(run func() clause.Expression) *GormQuery_SharedLock_Call {
	_c.Call.Return(run)
	return _c
}

// NewGormQuery creates a new instance of GormQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGormQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *GormQuery {
	mock := &GormQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
