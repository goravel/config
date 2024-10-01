package db

import (
	"fmt"

	contractsconfig "github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/database"
)

type Configs struct {
	config     contractsconfig.Config
	connection string
}

func NewConfigs(config contractsconfig.Config, connection string) *Configs {
	return &Configs{
		config:     config,
		connection: connection,
	}
}

func (c *Configs) Reads() []database.FullConfig {
	configs := c.config.Get(fmt.Sprintf("database.connections.%s.read", c.connection))
	if configs, ok := configs.([]database.Config); ok {
		return c.fillDefault(configs)
	}

	return nil
}

func (c *Configs) Writes() []database.FullConfig {
	configs := c.config.Get(fmt.Sprintf("database.connections.%s.write", c.connection))
	if configs, ok := configs.([]database.Config); ok {
		return c.fillDefault(configs)
	}

	// Use default db configuration when write is empty
	return c.fillDefault([]database.Config{{}})
}

func (c *Configs) fillDefault(configs []database.Config) []database.FullConfig {
	if len(configs) == 0 {
		return nil
	}

	var fullConfigs []database.FullConfig
	driver := database.Driver(c.config.GetString(fmt.Sprintf("database.connections.%s.driver", c.connection)))

	for _, config := range configs {
		fullConfig := database.FullConfig{
			Config:     config,
			Connection: c.connection,
			Driver:     driver,
			Prefix:     c.config.GetString(fmt.Sprintf("database.connections.%s.prefix", c.connection)),
			Singular:   c.config.GetBool(fmt.Sprintf("database.connections.%s.singular", c.connection)),
		}
		if driver != database.DriverSqlite {
			if fullConfig.Host == "" {
				fullConfig.Host = c.config.GetString(fmt.Sprintf("database.connections.%s.host", c.connection))
			}
			if fullConfig.Port == 0 {
				fullConfig.Port = c.config.GetInt(fmt.Sprintf("database.connections.%s.port", c.connection))
			}
			if fullConfig.Username == "" {
				fullConfig.Username = c.config.GetString(fmt.Sprintf("database.connections.%s.username", c.connection))
			}
			if fullConfig.Password == "" {
				fullConfig.Password = c.config.GetString(fmt.Sprintf("database.connections.%s.password", c.connection))
			}
			if driver == database.DriverMysql {
				fullConfig.Charset = c.config.GetString(fmt.Sprintf("database.connections.%s.charset", c.connection))
			}
			if driver == database.DriverMysql || driver == database.DriverSqlserver {
				fullConfig.Loc = c.config.GetString(fmt.Sprintf("database.connections.%s.loc", c.connection))
			}
			if driver == database.DriverPostgres {
				fullConfig.Sslmode = c.config.GetString(fmt.Sprintf("database.connections.%s.sslmode", c.connection))
				fullConfig.Timezone = c.config.GetString(fmt.Sprintf("database.connections.%s.timezone", c.connection))
			}
		}
		if config.Database == "" {
			fullConfig.Database = c.config.GetString(fmt.Sprintf("database.connections.%s.database", c.connection))
		}
		fullConfigs = append(fullConfigs, fullConfig)
	}

	return fullConfigs
}
