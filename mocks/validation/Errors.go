// Code generated by mockery. DO NOT EDIT.

package validation

import mock "github.com/stretchr/testify/mock"

// Errors is an autogenerated mock type for the Errors type
type Errors struct {
	mock.Mock
}

type Errors_Expecter struct {
	mock *mock.Mock
}

func (_m *Errors) EXPECT() *Errors_Expecter {
	return &Errors_Expecter{mock: &_m.Mock}
}

// All provides a mock function with no fields
func (_m *Errors) All() map[string]map[string]string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 map[string]map[string]string
	if rf, ok := ret.Get(0).(func() map[string]map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]string)
		}
	}

	return r0
}

// Errors_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type Errors_All_Call struct {
	*mock.Call
}

// All is a helper method to define mock.On call
func (_e *Errors_Expecter) All() *Errors_All_Call {
	return &Errors_All_Call{Call: _e.mock.On("All")}
}

func (_c *Errors_All_Call) Run(run func()) *Errors_All_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Errors_All_Call) Return(_a0 map[string]map[string]string) *Errors_All_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Errors_All_Call) RunAndReturn(run func() map[string]map[string]string) *Errors_All_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *Errors) Get(key string) map[string]string {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// Errors_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Errors_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *Errors_Expecter) Get(key interface{}) *Errors_Get_Call {
	return &Errors_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *Errors_Get_Call) Run(run func(key string)) *Errors_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Errors_Get_Call) Return(_a0 map[string]string) *Errors_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Errors_Get_Call) RunAndReturn(run func(string) map[string]string) *Errors_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Has provides a mock function with given fields: key
func (_m *Errors) Has(key string) bool {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Has")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Errors_Has_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Has'
type Errors_Has_Call struct {
	*mock.Call
}

// Has is a helper method to define mock.On call
//   - key string
func (_e *Errors_Expecter) Has(key interface{}) *Errors_Has_Call {
	return &Errors_Has_Call{Call: _e.mock.On("Has", key)}
}

func (_c *Errors_Has_Call) Run(run func(key string)) *Errors_Has_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Errors_Has_Call) Return(_a0 bool) *Errors_Has_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Errors_Has_Call) RunAndReturn(run func(string) bool) *Errors_Has_Call {
	_c.Call.Return(run)
	return _c
}

// One provides a mock function with given fields: key
func (_m *Errors) One(key ...string) string {
	_va := make([]interface{}, len(key))
	for _i := range key {
		_va[_i] = key[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for One")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(key...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Errors_One_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'One'
type Errors_One_Call struct {
	*mock.Call
}

// One is a helper method to define mock.On call
//   - key ...string
func (_e *Errors_Expecter) One(key ...interface{}) *Errors_One_Call {
	return &Errors_One_Call{Call: _e.mock.On("One",
		append([]interface{}{}, key...)...)}
}

func (_c *Errors_One_Call) Run(run func(key ...string)) *Errors_One_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Errors_One_Call) Return(_a0 string) *Errors_One_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Errors_One_Call) RunAndReturn(run func(...string) string) *Errors_One_Call {
	_c.Call.Return(run)
	return _c
}

// NewErrors creates a new instance of Errors. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewErrors(t interface {
	mock.TestingT
	Cleanup(func())
}) *Errors {
	mock := &Errors{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
