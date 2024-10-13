// Code generated by mockery. DO NOT EDIT.

package migration

import (
	migration "github.com/goravel/framework/contracts/database/migration"
	mock "github.com/stretchr/testify/mock"

	orm "github.com/goravel/framework/contracts/database/orm"
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

// Connection provides a mock function with given fields: name
func (_m *Schema) Connection(name string) migration.Schema {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Connection")
	}

	var r0 migration.Schema
	if rf, ok := ret.Get(0).(func(string) migration.Schema); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(migration.Schema)
		}
	}

	return r0
}

// Schema_Connection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connection'
type Schema_Connection_Call struct {
	*mock.Call
}

// Connection is a helper method to define mock.On call
//   - name string
func (_e *Schema_Expecter) Connection(name interface{}) *Schema_Connection_Call {
	return &Schema_Connection_Call{Call: _e.mock.On("Connection", name)}
}

func (_c *Schema_Connection_Call) Run(run func(name string)) *Schema_Connection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Schema_Connection_Call) Return(_a0 migration.Schema) *Schema_Connection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Schema_Connection_Call) RunAndReturn(run func(string) migration.Schema) *Schema_Connection_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: table, callback
func (_m *Schema) Create(table string, callback func(migration.Blueprint)) {
	_m.Called(table, callback)
}

// Schema_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Schema_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - table string
//   - callback func(migration.Blueprint)
func (_e *Schema_Expecter) Create(table interface{}, callback interface{}) *Schema_Create_Call {
	return &Schema_Create_Call{Call: _e.mock.On("Create", table, callback)}
}

func (_c *Schema_Create_Call) Run(run func(table string, callback func(migration.Blueprint))) *Schema_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func(migration.Blueprint)))
	})
	return _c
}

func (_c *Schema_Create_Call) Return() *Schema_Create_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schema_Create_Call) RunAndReturn(run func(string, func(migration.Blueprint))) *Schema_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DropIfExists provides a mock function with given fields: table
func (_m *Schema) DropIfExists(table string) {
	_m.Called(table)
}

// Schema_DropIfExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropIfExists'
type Schema_DropIfExists_Call struct {
	*mock.Call
}

// DropIfExists is a helper method to define mock.On call
//   - table string
func (_e *Schema_Expecter) DropIfExists(table interface{}) *Schema_DropIfExists_Call {
	return &Schema_DropIfExists_Call{Call: _e.mock.On("DropIfExists", table)}
}

func (_c *Schema_DropIfExists_Call) Run(run func(table string)) *Schema_DropIfExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Schema_DropIfExists_Call) Return() *Schema_DropIfExists_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schema_DropIfExists_Call) RunAndReturn(run func(string)) *Schema_DropIfExists_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnection provides a mock function with given fields:
func (_m *Schema) GetConnection() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetConnection")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Schema_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type Schema_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
func (_e *Schema_Expecter) GetConnection() *Schema_GetConnection_Call {
	return &Schema_GetConnection_Call{Call: _e.mock.On("GetConnection")}
}

func (_c *Schema_GetConnection_Call) Run(run func()) *Schema_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Schema_GetConnection_Call) Return(_a0 string) *Schema_GetConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Schema_GetConnection_Call) RunAndReturn(run func() string) *Schema_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetTables provides a mock function with given fields:
func (_m *Schema) GetTables() ([]migration.Table, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTables")
	}

	var r0 []migration.Table
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]migration.Table, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []migration.Table); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]migration.Table)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Schema_GetTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTables'
type Schema_GetTables_Call struct {
	*mock.Call
}

// GetTables is a helper method to define mock.On call
func (_e *Schema_Expecter) GetTables() *Schema_GetTables_Call {
	return &Schema_GetTables_Call{Call: _e.mock.On("GetTables")}
}

func (_c *Schema_GetTables_Call) Run(run func()) *Schema_GetTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Schema_GetTables_Call) Return(_a0 []migration.Table, _a1 error) *Schema_GetTables_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Schema_GetTables_Call) RunAndReturn(run func() ([]migration.Table, error)) *Schema_GetTables_Call {
	_c.Call.Return(run)
	return _c
}

// HasTable provides a mock function with given fields: table
func (_m *Schema) HasTable(table string) bool {
	ret := _m.Called(table)

	if len(ret) == 0 {
		panic("no return value specified for HasTable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(table)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Schema_HasTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasTable'
type Schema_HasTable_Call struct {
	*mock.Call
}

// HasTable is a helper method to define mock.On call
//   - table string
func (_e *Schema_Expecter) HasTable(table interface{}) *Schema_HasTable_Call {
	return &Schema_HasTable_Call{Call: _e.mock.On("HasTable", table)}
}

func (_c *Schema_HasTable_Call) Run(run func(table string)) *Schema_HasTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Schema_HasTable_Call) Return(_a0 bool) *Schema_HasTable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Schema_HasTable_Call) RunAndReturn(run func(string) bool) *Schema_HasTable_Call {
	_c.Call.Return(run)
	return _c
}

// Migrations provides a mock function with given fields:
func (_m *Schema) Migrations() []migration.Migration {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Migrations")
	}

	var r0 []migration.Migration
	if rf, ok := ret.Get(0).(func() []migration.Migration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]migration.Migration)
		}
	}

	return r0
}

// Schema_Migrations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Migrations'
type Schema_Migrations_Call struct {
	*mock.Call
}

// Migrations is a helper method to define mock.On call
func (_e *Schema_Expecter) Migrations() *Schema_Migrations_Call {
	return &Schema_Migrations_Call{Call: _e.mock.On("Migrations")}
}

func (_c *Schema_Migrations_Call) Run(run func()) *Schema_Migrations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Schema_Migrations_Call) Return(_a0 []migration.Migration) *Schema_Migrations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Schema_Migrations_Call) RunAndReturn(run func() []migration.Migration) *Schema_Migrations_Call {
	_c.Call.Return(run)
	return _c
}

// Orm provides a mock function with given fields:
func (_m *Schema) Orm() orm.Orm {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Orm")
	}

	var r0 orm.Orm
	if rf, ok := ret.Get(0).(func() orm.Orm); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Orm)
		}
	}

	return r0
}

// Schema_Orm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Orm'
type Schema_Orm_Call struct {
	*mock.Call
}

// Orm is a helper method to define mock.On call
func (_e *Schema_Expecter) Orm() *Schema_Orm_Call {
	return &Schema_Orm_Call{Call: _e.mock.On("Orm")}
}

func (_c *Schema_Orm_Call) Run(run func()) *Schema_Orm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Schema_Orm_Call) Return(_a0 orm.Orm) *Schema_Orm_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Schema_Orm_Call) RunAndReturn(run func() orm.Orm) *Schema_Orm_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: _a0
func (_m *Schema) Register(_a0 []migration.Migration) {
	_m.Called(_a0)
}

// Schema_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Schema_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 []migration.Migration
func (_e *Schema_Expecter) Register(_a0 interface{}) *Schema_Register_Call {
	return &Schema_Register_Call{Call: _e.mock.On("Register", _a0)}
}

func (_c *Schema_Register_Call) Run(run func(_a0 []migration.Migration)) *Schema_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]migration.Migration))
	})
	return _c
}

func (_c *Schema_Register_Call) Return() *Schema_Register_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schema_Register_Call) RunAndReturn(run func([]migration.Migration)) *Schema_Register_Call {
	_c.Call.Return(run)
	return _c
}

// SetConnection provides a mock function with given fields: name
func (_m *Schema) SetConnection(name string) {
	_m.Called(name)
}

// Schema_SetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConnection'
type Schema_SetConnection_Call struct {
	*mock.Call
}

// SetConnection is a helper method to define mock.On call
//   - name string
func (_e *Schema_Expecter) SetConnection(name interface{}) *Schema_SetConnection_Call {
	return &Schema_SetConnection_Call{Call: _e.mock.On("SetConnection", name)}
}

func (_c *Schema_SetConnection_Call) Run(run func(name string)) *Schema_SetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Schema_SetConnection_Call) Return() *Schema_SetConnection_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schema_SetConnection_Call) RunAndReturn(run func(string)) *Schema_SetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// Sql provides a mock function with given fields: sql
func (_m *Schema) Sql(sql string) {
	_m.Called(sql)
}

// Schema_Sql_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sql'
type Schema_Sql_Call struct {
	*mock.Call
}

// Sql is a helper method to define mock.On call
//   - sql string
func (_e *Schema_Expecter) Sql(sql interface{}) *Schema_Sql_Call {
	return &Schema_Sql_Call{Call: _e.mock.On("Sql", sql)}
}

func (_c *Schema_Sql_Call) Run(run func(sql string)) *Schema_Sql_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Schema_Sql_Call) Return() *Schema_Sql_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schema_Sql_Call) RunAndReturn(run func(string)) *Schema_Sql_Call {
	_c.Call.Return(run)
	return _c
}

// Table provides a mock function with given fields: table, callback
func (_m *Schema) Table(table string, callback func(migration.Blueprint)) {
	_m.Called(table, callback)
}

// Schema_Table_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Table'
type Schema_Table_Call struct {
	*mock.Call
}

// Table is a helper method to define mock.On call
//   - table string
//   - callback func(migration.Blueprint)
func (_e *Schema_Expecter) Table(table interface{}, callback interface{}) *Schema_Table_Call {
	return &Schema_Table_Call{Call: _e.mock.On("Table", table, callback)}
}

func (_c *Schema_Table_Call) Run(run func(table string, callback func(migration.Blueprint))) *Schema_Table_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func(migration.Blueprint)))
	})
	return _c
}

func (_c *Schema_Table_Call) Return() *Schema_Table_Call {
	_c.Call.Return()
	return _c
}

func (_c *Schema_Table_Call) RunAndReturn(run func(string, func(migration.Blueprint))) *Schema_Table_Call {
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
