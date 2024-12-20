// Code generated by mockery. DO NOT EDIT.

package orm

import (
	orm "github.com/goravel/framework/contracts/database/orm"
	mock "github.com/stretchr/testify/mock"
)

// ObserverWithRestored is an autogenerated mock type for the ObserverWithRestored type
type ObserverWithRestored struct {
	mock.Mock
}

type ObserverWithRestored_Expecter struct {
	mock *mock.Mock
}

func (_m *ObserverWithRestored) EXPECT() *ObserverWithRestored_Expecter {
	return &ObserverWithRestored_Expecter{mock: &_m.Mock}
}

// Restored provides a mock function with given fields: _a0
func (_m *ObserverWithRestored) Restored(_a0 orm.Event) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Restored")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(orm.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ObserverWithRestored_Restored_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Restored'
type ObserverWithRestored_Restored_Call struct {
	*mock.Call
}

// Restored is a helper method to define mock.On call
//   - _a0 orm.Event
func (_e *ObserverWithRestored_Expecter) Restored(_a0 interface{}) *ObserverWithRestored_Restored_Call {
	return &ObserverWithRestored_Restored_Call{Call: _e.mock.On("Restored", _a0)}
}

func (_c *ObserverWithRestored_Restored_Call) Run(run func(_a0 orm.Event)) *ObserverWithRestored_Restored_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(orm.Event))
	})
	return _c
}

func (_c *ObserverWithRestored_Restored_Call) Return(_a0 error) *ObserverWithRestored_Restored_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ObserverWithRestored_Restored_Call) RunAndReturn(run func(orm.Event) error) *ObserverWithRestored_Restored_Call {
	_c.Call.Return(run)
	return _c
}

// NewObserverWithRestored creates a new instance of ObserverWithRestored. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewObserverWithRestored(t interface {
	mock.TestingT
	Cleanup(func())
}) *ObserverWithRestored {
	mock := &ObserverWithRestored{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
