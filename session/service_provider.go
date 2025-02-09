package session

import (
	"github.com/goravel/framework/contracts"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/contracts/session"
	"github.com/goravel/framework/errors"
)

var (
	SessionFacade session.Manager
	ConfigFacade  config.Config
)

type ServiceProvider struct {
}

func (session *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(contracts.BindingSession, func(app foundation.Application) (any, error) {
		c := app.MakeConfig()
		if c == nil {
			return nil, errors.ConfigFacadeNotSet.SetModule(errors.ModuleSession)
		}

		j := app.GetJson()
		if j == nil {
			return nil, errors.JSONParserNotSet.SetModule(errors.ModuleSession)
		}

		return NewManager(c, j), nil
	})
}

func (session *ServiceProvider) Boot(app foundation.Application) {
	SessionFacade = app.MakeSession()
	ConfigFacade = app.MakeConfig()
}
