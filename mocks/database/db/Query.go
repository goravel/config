// Code generated by mockery. DO NOT EDIT.

package db

import (
	db "github.com/goravel/framework/contracts/database/db"
	mock "github.com/stretchr/testify/mock"
)

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

type Query_Expecter struct {
	mock *mock.Mock
}

func (_m *Query) EXPECT() *Query_Expecter {
	return &Query_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with no fields
func (_m *Query) Delete() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Query_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
func (_e *Query_Expecter) Delete() *Query_Delete_Call {
	return &Query_Delete_Call{Call: _e.mock.On("Delete")}
}

func (_c *Query_Delete_Call) Run(run func()) *Query_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Query_Delete_Call) Return(_a0 error) *Query_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Query_Delete_Call) RunAndReturn(run func() error) *Query_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: dest
func (_m *Query) Get(dest interface{}) error {
	ret := _m.Called(dest)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Query_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - dest interface{}
func (_e *Query_Expecter) Get(dest interface{}) *Query_Get_Call {
	return &Query_Get_Call{Call: _e.mock.On("Get", dest)}
}

func (_c *Query_Get_Call) Run(run func(dest interface{})) *Query_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *Query_Get_Call) Return(_a0 error) *Query_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Query_Get_Call) RunAndReturn(run func(interface{}) error) *Query_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with no fields
func (_m *Query) Insert() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type Query_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
func (_e *Query_Expecter) Insert() *Query_Insert_Call {
	return &Query_Insert_Call{Call: _e.mock.On("Insert")}
}

func (_c *Query_Insert_Call) Run(run func()) *Query_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Query_Insert_Call) Return(_a0 error) *Query_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Query_Insert_Call) RunAndReturn(run func() error) *Query_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with no fields
func (_m *Query) Update() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Query_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
func (_e *Query_Expecter) Update() *Query_Update_Call {
	return &Query_Update_Call{Call: _e.mock.On("Update")}
}

func (_c *Query_Update_Call) Run(run func()) *Query_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Query_Update_Call) Return(_a0 error) *Query_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Query_Update_Call) RunAndReturn(run func() error) *Query_Update_Call {
	_c.Call.Return(run)
	return _c
}

// Where provides a mock function with given fields: query, args
func (_m *Query) Where(query interface{}, args ...interface{}) db.Query {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Where")
	}

	var r0 db.Query
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) db.Query); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Query)
		}
	}

	return r0
}

// Query_Where_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Where'
type Query_Where_Call struct {
	*mock.Call
}

// Where is a helper method to define mock.On call
//   - query interface{}
//   - args ...interface{}
func (_e *Query_Expecter) Where(query interface{}, args ...interface{}) *Query_Where_Call {
	return &Query_Where_Call{Call: _e.mock.On("Where",
		append([]interface{}{query}, args...)...)}
}

func (_c *Query_Where_Call) Run(run func(query interface{}, args ...interface{})) *Query_Where_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Query_Where_Call) Return(_a0 db.Query) *Query_Where_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Query_Where_Call) RunAndReturn(run func(interface{}, ...interface{}) db.Query) *Query_Where_Call {
	_c.Call.Return(run)
	return _c
}

// NewQuery creates a new instance of Query. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *Query {
	mock := &Query{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
