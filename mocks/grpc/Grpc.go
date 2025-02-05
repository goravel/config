// Code generated by mockery. DO NOT EDIT.

package grpc

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	net "net"
)

// Grpc is an autogenerated mock type for the Grpc type
type Grpc struct {
	mock.Mock
}

type Grpc_Expecter struct {
	mock *mock.Mock
}

func (_m *Grpc) EXPECT() *Grpc_Expecter {
	return &Grpc_Expecter{mock: &_m.Mock}
}

// Client provides a mock function with given fields: ctx, name
func (_m *Grpc) Client(ctx context.Context, name string) (*grpc.ClientConn, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for Client")
	}

	var r0 *grpc.ClientConn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*grpc.ClientConn, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *grpc.ClientConn); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.ClientConn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Grpc_Client_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Client'
type Grpc_Client_Call struct {
	*mock.Call
}

// Client is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *Grpc_Expecter) Client(ctx interface{}, name interface{}) *Grpc_Client_Call {
	return &Grpc_Client_Call{Call: _e.mock.On("Client", ctx, name)}
}

func (_c *Grpc_Client_Call) Run(run func(ctx context.Context, name string)) *Grpc_Client_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Grpc_Client_Call) Return(_a0 *grpc.ClientConn, _a1 error) *Grpc_Client_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Grpc_Client_Call) RunAndReturn(run func(context.Context, string) (*grpc.ClientConn, error)) *Grpc_Client_Call {
	_c.Call.Return(run)
	return _c
}

// Listen provides a mock function with given fields: l
func (_m *Grpc) Listen(l net.Listener) error {
	ret := _m.Called(l)

	if len(ret) == 0 {
		panic("no return value specified for Listen")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(net.Listener) error); ok {
		r0 = rf(l)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Grpc_Listen_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Listen'
type Grpc_Listen_Call struct {
	*mock.Call
}

// Listen is a helper method to define mock.On call
//   - l net.Listener
func (_e *Grpc_Expecter) Listen(l interface{}) *Grpc_Listen_Call {
	return &Grpc_Listen_Call{Call: _e.mock.On("Listen", l)}
}

func (_c *Grpc_Listen_Call) Run(run func(l net.Listener)) *Grpc_Listen_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.Listener))
	})
	return _c
}

func (_c *Grpc_Listen_Call) Return(_a0 error) *Grpc_Listen_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Grpc_Listen_Call) RunAndReturn(run func(net.Listener) error) *Grpc_Listen_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: host
func (_m *Grpc) Run(host ...string) error {
	_va := make([]interface{}, len(host))
	for _i := range host {
		_va[_i] = host[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(host...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Grpc_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type Grpc_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - host ...string
func (_e *Grpc_Expecter) Run(host ...interface{}) *Grpc_Run_Call {
	return &Grpc_Run_Call{Call: _e.mock.On("Run",
		append([]interface{}{}, host...)...)}
}

func (_c *Grpc_Run_Call) Run(run func(host ...string)) *Grpc_Run_Call {
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

func (_c *Grpc_Run_Call) Return(_a0 error) *Grpc_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Grpc_Run_Call) RunAndReturn(run func(...string) error) *Grpc_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Server provides a mock function with no fields
func (_m *Grpc) Server() *grpc.Server {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Server")
	}

	var r0 *grpc.Server
	if rf, ok := ret.Get(0).(func() *grpc.Server); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.Server)
		}
	}

	return r0
}

// Grpc_Server_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Server'
type Grpc_Server_Call struct {
	*mock.Call
}

// Server is a helper method to define mock.On call
func (_e *Grpc_Expecter) Server() *Grpc_Server_Call {
	return &Grpc_Server_Call{Call: _e.mock.On("Server")}
}

func (_c *Grpc_Server_Call) Run(run func()) *Grpc_Server_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Grpc_Server_Call) Return(_a0 *grpc.Server) *Grpc_Server_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Grpc_Server_Call) RunAndReturn(run func() *grpc.Server) *Grpc_Server_Call {
	_c.Call.Return(run)
	return _c
}

// Shutdown provides a mock function with given fields: force
func (_m *Grpc) Shutdown(force ...bool) {
	_va := make([]interface{}, len(force))
	for _i := range force {
		_va[_i] = force[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Grpc_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type Grpc_Shutdown_Call struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
//   - force ...bool
func (_e *Grpc_Expecter) Shutdown(force ...interface{}) *Grpc_Shutdown_Call {
	return &Grpc_Shutdown_Call{Call: _e.mock.On("Shutdown",
		append([]interface{}{}, force...)...)}
}

func (_c *Grpc_Shutdown_Call) Run(run func(force ...bool)) *Grpc_Shutdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]bool, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(bool)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Grpc_Shutdown_Call) Return() *Grpc_Shutdown_Call {
	_c.Call.Return()
	return _c
}

func (_c *Grpc_Shutdown_Call) RunAndReturn(run func(...bool)) *Grpc_Shutdown_Call {
	_c.Run(run)
	return _c
}

// UnaryClientInterceptorGroups provides a mock function with given fields: _a0
func (_m *Grpc) UnaryClientInterceptorGroups(_a0 map[string][]grpc.UnaryClientInterceptor) {
	_m.Called(_a0)
}

// Grpc_UnaryClientInterceptorGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnaryClientInterceptorGroups'
type Grpc_UnaryClientInterceptorGroups_Call struct {
	*mock.Call
}

// UnaryClientInterceptorGroups is a helper method to define mock.On call
//   - _a0 map[string][]grpc.UnaryClientInterceptor
func (_e *Grpc_Expecter) UnaryClientInterceptorGroups(_a0 interface{}) *Grpc_UnaryClientInterceptorGroups_Call {
	return &Grpc_UnaryClientInterceptorGroups_Call{Call: _e.mock.On("UnaryClientInterceptorGroups", _a0)}
}

func (_c *Grpc_UnaryClientInterceptorGroups_Call) Run(run func(_a0 map[string][]grpc.UnaryClientInterceptor)) *Grpc_UnaryClientInterceptorGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string][]grpc.UnaryClientInterceptor))
	})
	return _c
}

func (_c *Grpc_UnaryClientInterceptorGroups_Call) Return() *Grpc_UnaryClientInterceptorGroups_Call {
	_c.Call.Return()
	return _c
}

func (_c *Grpc_UnaryClientInterceptorGroups_Call) RunAndReturn(run func(map[string][]grpc.UnaryClientInterceptor)) *Grpc_UnaryClientInterceptorGroups_Call {
	_c.Run(run)
	return _c
}

// UnaryServerInterceptors provides a mock function with given fields: _a0
func (_m *Grpc) UnaryServerInterceptors(_a0 []grpc.UnaryServerInterceptor) {
	_m.Called(_a0)
}

// Grpc_UnaryServerInterceptors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnaryServerInterceptors'
type Grpc_UnaryServerInterceptors_Call struct {
	*mock.Call
}

// UnaryServerInterceptors is a helper method to define mock.On call
//   - _a0 []grpc.UnaryServerInterceptor
func (_e *Grpc_Expecter) UnaryServerInterceptors(_a0 interface{}) *Grpc_UnaryServerInterceptors_Call {
	return &Grpc_UnaryServerInterceptors_Call{Call: _e.mock.On("UnaryServerInterceptors", _a0)}
}

func (_c *Grpc_UnaryServerInterceptors_Call) Run(run func(_a0 []grpc.UnaryServerInterceptor)) *Grpc_UnaryServerInterceptors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]grpc.UnaryServerInterceptor))
	})
	return _c
}

func (_c *Grpc_UnaryServerInterceptors_Call) Return() *Grpc_UnaryServerInterceptors_Call {
	_c.Call.Return()
	return _c
}

func (_c *Grpc_UnaryServerInterceptors_Call) RunAndReturn(run func([]grpc.UnaryServerInterceptor)) *Grpc_UnaryServerInterceptors_Call {
	_c.Run(run)
	return _c
}

// NewGrpc creates a new instance of Grpc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGrpc(t interface {
	mock.TestingT
	Cleanup(func())
}) *Grpc {
	mock := &Grpc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
