// Code generated by mockery. DO NOT EDIT.

package http

import (
	http "github.com/goravel/framework/contracts/http"
	mock "github.com/stretchr/testify/mock"
)

// ResponseStatus is an autogenerated mock type for the ResponseStatus type
type ResponseStatus struct {
	mock.Mock
}

type ResponseStatus_Expecter struct {
	mock *mock.Mock
}

func (_m *ResponseStatus) EXPECT() *ResponseStatus_Expecter {
	return &ResponseStatus_Expecter{mock: &_m.Mock}
}

// Data provides a mock function with given fields: contentType, data
func (_m *ResponseStatus) Data(contentType string, data []byte) http.ResponseWithAbort {
	ret := _m.Called(contentType, data)

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 http.ResponseWithAbort
	if rf, ok := ret.Get(0).(func(string, []byte) http.ResponseWithAbort); ok {
		r0 = rf(contentType, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseWithAbort)
		}
	}

	return r0
}

// ResponseStatus_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type ResponseStatus_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
//   - contentType string
//   - data []byte
func (_e *ResponseStatus_Expecter) Data(contentType interface{}, data interface{}) *ResponseStatus_Data_Call {
	return &ResponseStatus_Data_Call{Call: _e.mock.On("Data", contentType, data)}
}

func (_c *ResponseStatus_Data_Call) Run(run func(contentType string, data []byte)) *ResponseStatus_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]byte))
	})
	return _c
}

func (_c *ResponseStatus_Data_Call) Return(_a0 http.ResponseWithAbort) *ResponseStatus_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseStatus_Data_Call) RunAndReturn(run func(string, []byte) http.ResponseWithAbort) *ResponseStatus_Data_Call {
	_c.Call.Return(run)
	return _c
}

// Json provides a mock function with given fields: obj
func (_m *ResponseStatus) Json(obj interface{}) http.ResponseWithAbort {
	ret := _m.Called(obj)

	if len(ret) == 0 {
		panic("no return value specified for Json")
	}

	var r0 http.ResponseWithAbort
	if rf, ok := ret.Get(0).(func(interface{}) http.ResponseWithAbort); ok {
		r0 = rf(obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseWithAbort)
		}
	}

	return r0
}

// ResponseStatus_Json_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Json'
type ResponseStatus_Json_Call struct {
	*mock.Call
}

// Json is a helper method to define mock.On call
//   - obj interface{}
func (_e *ResponseStatus_Expecter) Json(obj interface{}) *ResponseStatus_Json_Call {
	return &ResponseStatus_Json_Call{Call: _e.mock.On("Json", obj)}
}

func (_c *ResponseStatus_Json_Call) Run(run func(obj interface{})) *ResponseStatus_Json_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ResponseStatus_Json_Call) Return(_a0 http.ResponseWithAbort) *ResponseStatus_Json_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseStatus_Json_Call) RunAndReturn(run func(interface{}) http.ResponseWithAbort) *ResponseStatus_Json_Call {
	_c.Call.Return(run)
	return _c
}

// Stream provides a mock function with given fields: step
func (_m *ResponseStatus) Stream(step func(http.StreamWriter) error) http.Response {
	ret := _m.Called(step)

	if len(ret) == 0 {
		panic("no return value specified for Stream")
	}

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(func(http.StreamWriter) error) http.Response); ok {
		r0 = rf(step)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	return r0
}

// ResponseStatus_Stream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stream'
type ResponseStatus_Stream_Call struct {
	*mock.Call
}

// Stream is a helper method to define mock.On call
//   - step func(http.StreamWriter) error
func (_e *ResponseStatus_Expecter) Stream(step interface{}) *ResponseStatus_Stream_Call {
	return &ResponseStatus_Stream_Call{Call: _e.mock.On("Stream", step)}
}

func (_c *ResponseStatus_Stream_Call) Run(run func(step func(http.StreamWriter) error)) *ResponseStatus_Stream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(http.StreamWriter) error))
	})
	return _c
}

func (_c *ResponseStatus_Stream_Call) Return(_a0 http.Response) *ResponseStatus_Stream_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseStatus_Stream_Call) RunAndReturn(run func(func(http.StreamWriter) error) http.Response) *ResponseStatus_Stream_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields: format, values
func (_m *ResponseStatus) String(format string, values ...interface{}) http.ResponseWithAbort {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 http.ResponseWithAbort
	if rf, ok := ret.Get(0).(func(string, ...interface{}) http.ResponseWithAbort); ok {
		r0 = rf(format, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.ResponseWithAbort)
		}
	}

	return r0
}

// ResponseStatus_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type ResponseStatus_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
//   - format string
//   - values ...interface{}
func (_e *ResponseStatus_Expecter) String(format interface{}, values ...interface{}) *ResponseStatus_String_Call {
	return &ResponseStatus_String_Call{Call: _e.mock.On("String",
		append([]interface{}{format}, values...)...)}
}

func (_c *ResponseStatus_String_Call) Run(run func(format string, values ...interface{})) *ResponseStatus_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *ResponseStatus_String_Call) Return(_a0 http.ResponseWithAbort) *ResponseStatus_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResponseStatus_String_Call) RunAndReturn(run func(string, ...interface{}) http.ResponseWithAbort) *ResponseStatus_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewResponseStatus creates a new instance of ResponseStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResponseStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResponseStatus {
	mock := &ResponseStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
