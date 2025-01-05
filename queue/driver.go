package queue

import (
	"fmt"

	"github.com/goravel/framework/contracts/queue"
)

const DriverSync string = "sync"
const DriverASync string = "async"
const DriverMachinery string = "machinery" // TODO: Will be removed in v1.17
const DriverCustom string = "custom"

type Driver interface {
	New(store string) (queue.Driver, error)
}

type DriverImpl struct {
	connection string
	config     *Config
}

func NewDriverImpl(connection string, config *Config) *DriverImpl {
	return &DriverImpl{
		connection: connection,
		config:     config,
	}
}

func (d *DriverImpl) New() (queue.Driver, error) {
	switch d.config.Driver(d.connection) {
	case DriverSync:
		return NewSync(d.connection), nil
	case DriverASync:
		return NewASync(d.connection, d.config.Size(d.connection)), nil
	case DriverMachinery:
		return NewMachinery(d.connection, d.config, LogFacade), nil // TODO: Will be removed in v1.17
	case DriverCustom:
		return d.custom(d.connection)
	default:
		return nil, fmt.Errorf("invalid driver: %s, only support sync, async, custom\n", d.connection)
	}
}

func (d *DriverImpl) custom(connection string) (queue.Driver, error) {
	custom := d.config.Via(connection)
	if driver, ok := custom.(queue.Driver); ok {
		return driver, nil
	}
	if driver, ok := custom.(func() (queue.Driver, error)); ok {
		return driver()
	}

	return nil, fmt.Errorf("%s doesn't implement contracts/queue/driver\n", connection)
}
