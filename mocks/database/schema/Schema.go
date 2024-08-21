// Code generated by mockery. DO NOT EDIT.

package schema

import (
	schema "github.com/goravel/framework/contracts/database/schema"
	mock "github.com/stretchr/testify/mock"
)

// Schema is an autogenerated mock type for the Schema type
type Schema struct {
	mock.Mock
}

type Schema_Expecter struct {
	mock *mock.Mock
}

func (_m *Schema) EXPECT() *Schema_Expecter {
	return &Schema_Expecter{mock: &_m.Mock}
}

// Connection provides a mock function with given fields:
func (_m *Schema) Connection() schema.Schema {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Connection")
	}

	var r0 schema.Schema
	if rf, ok := ret.Get(0).(func() schema.Schema); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schema.Schema)
		}
	}

	return r0
}

// Schema_Connection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connection'
type Schema_Connection_Call struct {
	*mock.Call
}

// Connection is a helper method to define mock.On call
func (_e *Schema_Expecter) Connection() *Schema_Connection_Call {
	return &Schema_Connection_Call{Call: _e.mock.On("Connection")}
}

func (_c *Schema_Connection_Call) Run(run func()) *Schema_Connection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Schema_Connection_Call) Return(_a0 schema.Schema) *Schema_Connection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Schema_Connection_Call) RunAndReturn(run func() schema.Schema) *Schema_Connection_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: _a0
func (_m *Schema) Register(_a0 []schema.Migration) {
	_m.Called(_a0)
}

// Schema_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Schema_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 []schema.Migration
func (_e *Schema_Expecter) Register(_a0 interface{}) *Schema_Register_Call {
	return &Schema_Register_Call{Call: _e.mock.On("Register", _a0)}
}

func (_c *Schema_Register_Call) Run(run func(_a0 []schema.Migration)) *Schema_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]schema.Migration))
	})
	return _c
}

func (_c *Schema_Register_Call) Return() *Schema_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schema_Register_Call) RunAndReturn(run func([]schema.Migration)) *Schema_Register_Call {
	_c.Call.Return(run)
	return _c
}

// Sql provides a mock function with given fields: callback
func (_m *Schema) Sql(callback func(schema.Blueprint)) {
	_m.Called(callback)
}

// Schema_Sql_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sql'
type Schema_Sql_Call struct {
	*mock.Call
}

// Sql is a helper method to define mock.On call
//   - callback func(schema.Blueprint)
func (_e *Schema_Expecter) Sql(callback interface{}) *Schema_Sql_Call {
	return &Schema_Sql_Call{Call: _e.mock.On("Sql", callback)}
}

func (_c *Schema_Sql_Call) Run(run func(callback func(schema.Blueprint))) *Schema_Sql_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(schema.Blueprint)))
	})
	return _c
}

func (_c *Schema_Sql_Call) Return() *Schema_Sql_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schema_Sql_Call) RunAndReturn(run func(func(schema.Blueprint))) *Schema_Sql_Call {
	_c.Call.Return(run)
	return _c
}

// NewSchema creates a new instance of Schema. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSchema(t interface {
	mock.TestingT
	Cleanup(func())
}) *Schema {
	mock := &Schema{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
