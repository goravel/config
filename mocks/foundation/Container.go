// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	auth "github.com/goravel/framework/contracts/auth"
	access "github.com/goravel/framework/contracts/auth/access"

	cache "github.com/goravel/framework/contracts/cache"

	config "github.com/goravel/framework/contracts/config"

	console "github.com/goravel/framework/contracts/console"

	context "context"

	crypt "github.com/goravel/framework/contracts/crypt"

	event "github.com/goravel/framework/contracts/event"

	filesystem "github.com/goravel/framework/contracts/filesystem"

	foundation "github.com/goravel/framework/contracts/foundation"

	grpc "github.com/goravel/framework/contracts/grpc"

	hash "github.com/goravel/framework/contracts/hash"

	http "github.com/goravel/framework/contracts/http"

	log "github.com/goravel/framework/contracts/log"

	mail "github.com/goravel/framework/contracts/mail"

	mock "github.com/stretchr/testify/mock"

	orm "github.com/goravel/framework/contracts/database/orm"

	queue "github.com/goravel/framework/contracts/queue"

	route "github.com/goravel/framework/contracts/route"

	schedule "github.com/goravel/framework/contracts/schedule"

	seeder "github.com/goravel/framework/contracts/database/seeder"

	session "github.com/goravel/framework/contracts/session"

	testing "github.com/goravel/framework/contracts/testing"

	translation "github.com/goravel/framework/contracts/translation"

	validation "github.com/goravel/framework/contracts/validation"
)

// Container is an autogenerated mock type for the Container type
type Container struct {
	mock.Mock
}

// Bind provides a mock function with given fields: key, callback
func (_m *Container) Bind(key interface{}, callback func(foundation.Application) (interface{}, error)) {
	_m.Called(key, callback)
}

// BindWith provides a mock function with given fields: key, callback
func (_m *Container) BindWith(key interface{}, callback func(foundation.Application, map[string]interface{}) (interface{}, error)) {
	_m.Called(key, callback)
}

// Instance provides a mock function with given fields: key, instance
func (_m *Container) Instance(key interface{}, instance interface{}) {
	_m.Called(key, instance)
}

// Make provides a mock function with given fields: key
func (_m *Container) Make(key interface{}) (interface{}, error) {
	ret := _m.Called(key)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (interface{}, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeArtisan provides a mock function with given fields:
func (_m *Container) MakeArtisan() console.Artisan {
	ret := _m.Called()

	var r0 console.Artisan
	if rf, ok := ret.Get(0).(func() console.Artisan); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(console.Artisan)
		}
	}

	return r0
}

// MakeAuth provides a mock function with given fields: ctx
func (_m *Container) MakeAuth(ctx http.Context) auth.Auth {
	ret := _m.Called(ctx)

	var r0 auth.Auth
	if rf, ok := ret.Get(0).(func(http.Context) auth.Auth); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(auth.Auth)
		}
	}

	return r0
}

// MakeCache provides a mock function with given fields:
func (_m *Container) MakeCache() cache.Cache {
	ret := _m.Called()

	var r0 cache.Cache
	if rf, ok := ret.Get(0).(func() cache.Cache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Cache)
		}
	}

	return r0
}

// MakeConfig provides a mock function with given fields:
func (_m *Container) MakeConfig() config.Config {
	ret := _m.Called()

	var r0 config.Config
	if rf, ok := ret.Get(0).(func() config.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.Config)
		}
	}

	return r0
}

// MakeCrypt provides a mock function with given fields:
func (_m *Container) MakeCrypt() crypt.Crypt {
	ret := _m.Called()

	var r0 crypt.Crypt
	if rf, ok := ret.Get(0).(func() crypt.Crypt); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypt.Crypt)
		}
	}

	return r0
}

// MakeEvent provides a mock function with given fields:
func (_m *Container) MakeEvent() event.Instance {
	ret := _m.Called()

	var r0 event.Instance
	if rf, ok := ret.Get(0).(func() event.Instance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Instance)
		}
	}

	return r0
}

// MakeGate provides a mock function with given fields:
func (_m *Container) MakeGate() access.Gate {
	ret := _m.Called()

	var r0 access.Gate
	if rf, ok := ret.Get(0).(func() access.Gate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(access.Gate)
		}
	}

	return r0
}

// MakeGrpc provides a mock function with given fields:
func (_m *Container) MakeGrpc() grpc.Grpc {
	ret := _m.Called()

	var r0 grpc.Grpc
	if rf, ok := ret.Get(0).(func() grpc.Grpc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.Grpc)
		}
	}

	return r0
}

// MakeHash provides a mock function with given fields:
func (_m *Container) MakeHash() hash.Hash {
	ret := _m.Called()

	var r0 hash.Hash
	if rf, ok := ret.Get(0).(func() hash.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(hash.Hash)
		}
	}

	return r0
}

// MakeLang provides a mock function with given fields: ctx
func (_m *Container) MakeLang(ctx context.Context) translation.Translator {
	ret := _m.Called(ctx)

	var r0 translation.Translator
	if rf, ok := ret.Get(0).(func(context.Context) translation.Translator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(translation.Translator)
		}
	}

	return r0
}

// MakeLog provides a mock function with given fields:
func (_m *Container) MakeLog() log.Log {
	ret := _m.Called()

	var r0 log.Log
	if rf, ok := ret.Get(0).(func() log.Log); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Log)
		}
	}

	return r0
}

// MakeMail provides a mock function with given fields:
func (_m *Container) MakeMail() mail.Mail {
	ret := _m.Called()

	var r0 mail.Mail
	if rf, ok := ret.Get(0).(func() mail.Mail); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mail.Mail)
		}
	}

	return r0
}

// MakeOrm provides a mock function with given fields:
func (_m *Container) MakeOrm() orm.Orm {
	ret := _m.Called()

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

// MakeQueue provides a mock function with given fields:
func (_m *Container) MakeQueue() queue.Queue {
	ret := _m.Called()

	var r0 queue.Queue
	if rf, ok := ret.Get(0).(func() queue.Queue); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(queue.Queue)
		}
	}

	return r0
}

// MakeRateLimiter provides a mock function with given fields:
func (_m *Container) MakeRateLimiter() http.RateLimiter {
	ret := _m.Called()

	var r0 http.RateLimiter
	if rf, ok := ret.Get(0).(func() http.RateLimiter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.RateLimiter)
		}
	}

	return r0
}

// MakeRoute provides a mock function with given fields:
func (_m *Container) MakeRoute() route.Route {
	ret := _m.Called()

	var r0 route.Route
	if rf, ok := ret.Get(0).(func() route.Route); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(route.Route)
		}
	}

	return r0
}

// MakeSchedule provides a mock function with given fields:
func (_m *Container) MakeSchedule() schedule.Schedule {
	ret := _m.Called()

	var r0 schedule.Schedule
	if rf, ok := ret.Get(0).(func() schedule.Schedule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(schedule.Schedule)
		}
	}

	return r0
}

// MakeSeeder provides a mock function with given fields:
func (_m *Container) MakeSeeder() seeder.Facade {
	ret := _m.Called()

	var r0 seeder.Facade
	if rf, ok := ret.Get(0).(func() seeder.Facade); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(seeder.Facade)
		}
	}

	return r0
}

// MakeSession provides a mock function with given fields:
func (_m *Container) MakeSession() session.Manager {
	ret := _m.Called()

	var r0 session.Manager
	if rf, ok := ret.Get(0).(func() session.Manager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(session.Manager)
		}
	}

	return r0
}

// MakeStorage provides a mock function with given fields:
func (_m *Container) MakeStorage() filesystem.Storage {
	ret := _m.Called()

	var r0 filesystem.Storage
	if rf, ok := ret.Get(0).(func() filesystem.Storage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(filesystem.Storage)
		}
	}

	return r0
}

// MakeTesting provides a mock function with given fields:
func (_m *Container) MakeTesting() testing.Testing {
	ret := _m.Called()

	var r0 testing.Testing
	if rf, ok := ret.Get(0).(func() testing.Testing); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(testing.Testing)
		}
	}

	return r0
}

// MakeValidation provides a mock function with given fields:
func (_m *Container) MakeValidation() validation.Validation {
	ret := _m.Called()

	var r0 validation.Validation
	if rf, ok := ret.Get(0).(func() validation.Validation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Validation)
		}
	}

	return r0
}

// MakeView provides a mock function with given fields:
func (_m *Container) MakeView() http.View {
	ret := _m.Called()

	var r0 http.View
	if rf, ok := ret.Get(0).(func() http.View); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.View)
		}
	}

	return r0
}

// MakeWith provides a mock function with given fields: key, parameters
func (_m *Container) MakeWith(key interface{}, parameters map[string]interface{}) (interface{}, error) {
	ret := _m.Called(key, parameters)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, map[string]interface{}) (interface{}, error)); ok {
		return rf(key, parameters)
	}
	if rf, ok := ret.Get(0).(func(interface{}, map[string]interface{}) interface{}); ok {
		r0 = rf(key, parameters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, map[string]interface{}) error); ok {
		r1 = rf(key, parameters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Singleton provides a mock function with given fields: key, callback
func (_m *Container) Singleton(key interface{}, callback func(foundation.Application) (interface{}, error)) {
	_m.Called(key, callback)
}

// NewContainer creates a new instance of Container. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContainer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Container {
	mock := &Container{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
