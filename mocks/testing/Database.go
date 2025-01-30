// Code generated by mockery. DO NOT EDIT.

package testing

import (
	seeder "github.com/goravel/framework/contracts/database/seeder"
	testing "github.com/goravel/framework/contracts/testing"
	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

type Database_Expecter struct {
	mock *mock.Mock
}

func (_m *Database) EXPECT() *Database_Expecter {
	return &Database_Expecter{mock: &_m.Mock}
}

// Build provides a mock function with no fields
func (_m *Database) Build() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type Database_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
func (_e *Database_Expecter) Build() *Database_Build_Call {
	return &Database_Build_Call{Call: _e.mock.On("Build")}
}

func (_c *Database_Build_Call) Run(run func()) *Database_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Database_Build_Call) Return(_a0 error) *Database_Build_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Build_Call) RunAndReturn(run func() error) *Database_Build_Call {
	_c.Call.Return(run)
	return _c
}

// Config provides a mock function with no fields
func (_m *Database) Config() testing.DatabaseConfig {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Config")
	}

	var r0 testing.DatabaseConfig
	if rf, ok := ret.Get(0).(func() testing.DatabaseConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(testing.DatabaseConfig)
	}

	return r0
}

// Database_Config_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Config'
type Database_Config_Call struct {
	*mock.Call
}

// Config is a helper method to define mock.On call
func (_e *Database_Expecter) Config() *Database_Config_Call {
	return &Database_Config_Call{Call: _e.mock.On("Config")}
}

func (_c *Database_Config_Call) Run(run func()) *Database_Config_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Database_Config_Call) Return(_a0 testing.DatabaseConfig) *Database_Config_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Config_Call) RunAndReturn(run func() testing.DatabaseConfig) *Database_Config_Call {
	_c.Call.Return(run)
	return _c
}

// Database provides a mock function with given fields: name
func (_m *Database) Database(name string) (testing.DatabaseDriver, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Database")
	}

	var r0 testing.DatabaseDriver
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (testing.DatabaseDriver, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) testing.DatabaseDriver); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(testing.DatabaseDriver)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Database_Database_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Database'
type Database_Database_Call struct {
	*mock.Call
}

// Database is a helper method to define mock.On call
//   - name string
func (_e *Database_Expecter) Database(name interface{}) *Database_Database_Call {
	return &Database_Database_Call{Call: _e.mock.On("Database", name)}
}

func (_c *Database_Database_Call) Run(run func(name string)) *Database_Database_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Database_Database_Call) Return(_a0 testing.DatabaseDriver, _a1 error) *Database_Database_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Database_Database_Call) RunAndReturn(run func(string) (testing.DatabaseDriver, error)) *Database_Database_Call {
	_c.Call.Return(run)
	return _c
}

// Driver provides a mock function with no fields
func (_m *Database) Driver() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Driver")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Database_Driver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Driver'
type Database_Driver_Call struct {
	*mock.Call
}

// Driver is a helper method to define mock.On call
func (_e *Database_Expecter) Driver() *Database_Driver_Call {
	return &Database_Driver_Call{Call: _e.mock.On("Driver")}
}

func (_c *Database_Driver_Call) Run(run func()) *Database_Driver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Database_Driver_Call) Return(_a0 string) *Database_Driver_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Driver_Call) RunAndReturn(run func() string) *Database_Driver_Call {
	_c.Call.Return(run)
	return _c
}

// Fresh provides a mock function with no fields
func (_m *Database) Fresh() error {
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

// Database_Fresh_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fresh'
type Database_Fresh_Call struct {
	*mock.Call
}

// Fresh is a helper method to define mock.On call
func (_e *Database_Expecter) Fresh() *Database_Fresh_Call {
	return &Database_Fresh_Call{Call: _e.mock.On("Fresh")}
}

func (_c *Database_Fresh_Call) Run(run func()) *Database_Fresh_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Database_Fresh_Call) Return(_a0 error) *Database_Fresh_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Fresh_Call) RunAndReturn(run func() error) *Database_Fresh_Call {
	_c.Call.Return(run)
	return _c
}

// Image provides a mock function with given fields: image
func (_m *Database) Image(image testing.Image) {
	_m.Called(image)
}

// Database_Image_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Image'
type Database_Image_Call struct {
	*mock.Call
}

// Image is a helper method to define mock.On call
//   - image testing.Image
func (_e *Database_Expecter) Image(image interface{}) *Database_Image_Call {
	return &Database_Image_Call{Call: _e.mock.On("Image", image)}
}

func (_c *Database_Image_Call) Run(run func(image testing.Image)) *Database_Image_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(testing.Image))
	})
	return _c
}

func (_c *Database_Image_Call) Return() *Database_Image_Call {
	_c.Call.Return()
	return _c
}

func (_c *Database_Image_Call) RunAndReturn(run func(testing.Image)) *Database_Image_Call {
	_c.Run(run)
	return _c
}

// Migrate provides a mock function with no fields
func (_m *Database) Migrate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Migrate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database_Migrate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Migrate'
type Database_Migrate_Call struct {
	*mock.Call
}

// Migrate is a helper method to define mock.On call
func (_e *Database_Expecter) Migrate() *Database_Migrate_Call {
	return &Database_Migrate_Call{Call: _e.mock.On("Migrate")}
}

func (_c *Database_Migrate_Call) Run(run func()) *Database_Migrate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Database_Migrate_Call) Return(_a0 error) *Database_Migrate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Migrate_Call) RunAndReturn(run func() error) *Database_Migrate_Call {
	_c.Call.Return(run)
	return _c
}

// Ready provides a mock function with no fields
func (_m *Database) Ready() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ready")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database_Ready_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ready'
type Database_Ready_Call struct {
	*mock.Call
}

// Ready is a helper method to define mock.On call
func (_e *Database_Expecter) Ready() *Database_Ready_Call {
	return &Database_Ready_Call{Call: _e.mock.On("Ready")}
}

func (_c *Database_Ready_Call) Run(run func()) *Database_Ready_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Database_Ready_Call) Return(_a0 error) *Database_Ready_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Ready_Call) RunAndReturn(run func() error) *Database_Ready_Call {
	_c.Call.Return(run)
	return _c
}

// Reuse provides a mock function with given fields: containerID, port
func (_m *Database) Reuse(containerID string, port int) error {
	ret := _m.Called(containerID, port)

	if len(ret) == 0 {
		panic("no return value specified for Reuse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(containerID, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database_Reuse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reuse'
type Database_Reuse_Call struct {
	*mock.Call
}

// Reuse is a helper method to define mock.On call
//   - containerID string
//   - port int
func (_e *Database_Expecter) Reuse(containerID interface{}, port interface{}) *Database_Reuse_Call {
	return &Database_Reuse_Call{Call: _e.mock.On("Reuse", containerID, port)}
}

func (_c *Database_Reuse_Call) Run(run func(containerID string, port int)) *Database_Reuse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *Database_Reuse_Call) Return(_a0 error) *Database_Reuse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Reuse_Call) RunAndReturn(run func(string, int) error) *Database_Reuse_Call {
	_c.Call.Return(run)
	return _c
}

// Seed provides a mock function with given fields: seeders
func (_m *Database) Seed(seeders ...seeder.Seeder) error {
	_va := make([]interface{}, len(seeders))
	for _i := range seeders {
		_va[_i] = seeders[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Seed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...seeder.Seeder) error); ok {
		r0 = rf(seeders...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database_Seed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Seed'
type Database_Seed_Call struct {
	*mock.Call
}

// Seed is a helper method to define mock.On call
//   - seeders ...seeder.Seeder
func (_e *Database_Expecter) Seed(seeders ...interface{}) *Database_Seed_Call {
	return &Database_Seed_Call{Call: _e.mock.On("Seed",
		append([]interface{}{}, seeders...)...)}
}

func (_c *Database_Seed_Call) Run(run func(seeders ...seeder.Seeder)) *Database_Seed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]seeder.Seeder, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(seeder.Seeder)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Database_Seed_Call) Return(_a0 error) *Database_Seed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Seed_Call) RunAndReturn(run func(...seeder.Seeder) error) *Database_Seed_Call {
	_c.Call.Return(run)
	return _c
}

// Shutdown provides a mock function with no fields
func (_m *Database) Shutdown() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Shutdown")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type Database_Shutdown_Call struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
func (_e *Database_Expecter) Shutdown() *Database_Shutdown_Call {
	return &Database_Shutdown_Call{Call: _e.mock.On("Shutdown")}
}

func (_c *Database_Shutdown_Call) Run(run func()) *Database_Shutdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Database_Shutdown_Call) Return(_a0 error) *Database_Shutdown_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Database_Shutdown_Call) RunAndReturn(run func() error) *Database_Shutdown_Call {
	_c.Call.Return(run)
	return _c
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
