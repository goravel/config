// Code generated by mockery. DO NOT EDIT.

package migration

import (
	migration "github.com/goravel/framework/contracts/database/migration"
	mock "github.com/stretchr/testify/mock"
)

// Migrator is an autogenerated mock type for the Migrator type
type Migrator struct {
	mock.Mock
}

type Migrator_Expecter struct {
	mock *mock.Mock
}

func (_m *Migrator) EXPECT() *Migrator_Expecter {
	return &Migrator_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: name
func (_m *Migrator) Create(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Migrator_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - name string
func (_e *Migrator_Expecter) Create(name interface{}) *Migrator_Create_Call {
	return &Migrator_Create_Call{Call: _e.mock.On("Create", name)}
}

func (_c *Migrator_Create_Call) Run(run func(name string)) *Migrator_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Migrator_Create_Call) Return(_a0 error) *Migrator_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Migrator_Create_Call) RunAndReturn(run func(string) error) *Migrator_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Fresh provides a mock function with no fields
func (_m *Migrator) Fresh() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Fresh")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Fresh_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fresh'
type Migrator_Fresh_Call struct {
	*mock.Call
}

// Fresh is a helper method to define mock.On call
func (_e *Migrator_Expecter) Fresh() *Migrator_Fresh_Call {
	return &Migrator_Fresh_Call{Call: _e.mock.On("Fresh")}
}

func (_c *Migrator_Fresh_Call) Run(run func()) *Migrator_Fresh_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Migrator_Fresh_Call) Return(_a0 error) *Migrator_Fresh_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Migrator_Fresh_Call) RunAndReturn(run func() error) *Migrator_Fresh_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with no fields
func (_m *Migrator) Reset() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Reset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type Migrator_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *Migrator_Expecter) Reset() *Migrator_Reset_Call {
	return &Migrator_Reset_Call{Call: _e.mock.On("Reset")}
}

func (_c *Migrator_Reset_Call) Run(run func()) *Migrator_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Migrator_Reset_Call) Return(_a0 error) *Migrator_Reset_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Migrator_Reset_Call) RunAndReturn(run func() error) *Migrator_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with given fields: step, batch
func (_m *Migrator) Rollback(step int, batch int) error {
	ret := _m.Called(step, batch)

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(step, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type Migrator_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
//   - step int
//   - batch int
func (_e *Migrator_Expecter) Rollback(step interface{}, batch interface{}) *Migrator_Rollback_Call {
	return &Migrator_Rollback_Call{Call: _e.mock.On("Rollback", step, batch)}
}

func (_c *Migrator_Rollback_Call) Run(run func(step int, batch int)) *Migrator_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *Migrator_Rollback_Call) Return(_a0 error) *Migrator_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Migrator_Rollback_Call) RunAndReturn(run func(int, int) error) *Migrator_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with no fields
func (_m *Migrator) Run() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrator_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type Migrator_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
func (_e *Migrator_Expecter) Run() *Migrator_Run_Call {
	return &Migrator_Run_Call{Call: _e.mock.On("Run")}
}

func (_c *Migrator_Run_Call) Run(run func()) *Migrator_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Migrator_Run_Call) Return(_a0 error) *Migrator_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Migrator_Run_Call) RunAndReturn(run func() error) *Migrator_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with no fields
func (_m *Migrator) Status() ([]migration.Status, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 []migration.Status
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]migration.Status, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []migration.Status); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]migration.Status)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrator_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type Migrator_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
func (_e *Migrator_Expecter) Status() *Migrator_Status_Call {
	return &Migrator_Status_Call{Call: _e.mock.On("Status")}
}

func (_c *Migrator_Status_Call) Run(run func()) *Migrator_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Migrator_Status_Call) Return(_a0 []migration.Status, _a1 error) *Migrator_Status_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Migrator_Status_Call) RunAndReturn(run func() ([]migration.Status, error)) *Migrator_Status_Call {
	_c.Call.Return(run)
	return _c
}

// NewMigrator creates a new instance of Migrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMigrator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Migrator {
	mock := &Migrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
