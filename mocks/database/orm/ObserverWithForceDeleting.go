// Code generated by mockery. DO NOT EDIT.

package orm

import (
	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// ObserverWithForceDeleting is an autogenerated mock type for the ObserverWithForceDeleting type
type ObserverWithForceDeleting struct {
	mock.Mock
}

type ObserverWithForceDeleting_Expecter struct {
	mock *mock.Mock
}

func (_m *ObserverWithForceDeleting) EXPECT() *ObserverWithForceDeleting_Expecter {
	return &ObserverWithForceDeleting_Expecter{mock: &_m.Mock}
}

// ForceDeleting provides a mock function with given fields: _a0
func (_m *ObserverWithForceDeleting) ForceDeleting(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ForceDeleting")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ObserverWithForceDeleting_ForceDeleting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForceDeleting'
type ObserverWithForceDeleting_ForceDeleting_Call struct {
	*mock.Call
}

// ForceDeleting is a helper method to define mock.On call
//   - _a0 orm.Event
func (_e *ObserverWithForceDeleting_Expecter) ForceDeleting(_a0 interface{}) *ObserverWithForceDeleting_ForceDeleting_Call {
	return &ObserverWithForceDeleting_ForceDeleting_Call{Call: _e.mock.On("ForceDeleting", _a0)}
}

func (_c *ObserverWithForceDeleting_ForceDeleting_Call) Run(run func(_a0 orm.Event)) *ObserverWithForceDeleting_ForceDeleting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(orm.Event))
	})
	return _c
}

func (_c *ObserverWithForceDeleting_ForceDeleting_Call) Return(_a0 error) *ObserverWithForceDeleting_ForceDeleting_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ObserverWithForceDeleting_ForceDeleting_Call) RunAndReturn(run func(orm.Event) error) *ObserverWithForceDeleting_ForceDeleting_Call {
	_c.Call.Return(run)
	return _c
}

// NewObserverWithForceDeleting creates a new instance of ObserverWithForceDeleting. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewObserverWithForceDeleting(t interface {
	mock.TestingT
	Cleanup(func())
}) *ObserverWithForceDeleting {
	mock := &ObserverWithForceDeleting{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
