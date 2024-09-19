package migration

import (
	"fmt"

	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/database/migration"
	contractsorm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/log"
	"github.com/goravel/framework/database/migration/grammars"
)

var _ migration.Schema = (*Schema)(nil)

type Schema struct {
	blueprint  migration.Blueprint
	config     config.Config
	connection string
	grammar    migration.Grammar
	log        log.Log
	migrations []migration.Migration
	orm        contractsorm.Orm
}

func NewSchema(blueprint migration.Blueprint, config config.Config, connection string, log log.Log, orm contractsorm.Orm) (*Schema, error) {
	schema := &Schema{
		blueprint:  blueprint,
		config:     config,
		connection: connection,
		log:        log,
		orm:        orm,
	}

	if err := schema.initGrammar(); err != nil {
		return nil, err
	}

	return schema, nil
}

func (r *Schema) Connection(name string) migration.Schema {
	prefix := r.config.GetString(fmt.Sprintf("database.connections.%s.prefix", name))
	dbSchema := r.config.GetString(fmt.Sprintf("database.connections.%s.schema", name))
	blueprint := NewBlueprint(prefix, dbSchema)

	schema, err := NewSchema(blueprint, r.config, name, r.log, r.orm)
	if err != nil {
		r.log.Panic(err)
	}

	return schema
}

func (r *Schema) Create(table string, callback func(table migration.Blueprint)) {
	r.blueprint.SetTable(table)
	r.blueprint.Create()
	callback(r.blueprint)

	// TODO catch error and rollback
	_ = r.blueprint.Build(r.orm.Connection(r.connection).Query(), r.grammar)
}

func (r *Schema) DropIfExists(table string) {
	r.blueprint.SetTable(table)
	r.blueprint.DropIfExists()

	// TODO catch error
	_ = r.blueprint.Build(r.orm.Connection(r.connection).Query(), r.grammar)
}

func (r *Schema) Register(migrations []migration.Migration) {
	r.migrations = migrations
}

func (r *Schema) Sql(sql string) {
	// TODO catch error and rollback
	_, _ = r.orm.Connection(r.connection).Query().Exec(sql)
}

func (r *Schema) initGrammar() error {
	driver := r.config.GetString(fmt.Sprintf("database.connections.%s.driver", r.connection))

	switch driver {
	case contractsorm.DriverMysql.String():
		// TODO Optimize here when implementing Mysql driver
		return nil
	case contractsorm.DriverPostgres.String():
		r.grammar = grammars.NewPostgres()
		return nil
	//case ormcontract.DriverSqlserver:
	//	grammar = grammars.NewSqlserver()
	//case ormcontract.DriverSqlite:
	//	grammar = grammars.NewSqlite()
	default:
		return fmt.Errorf("unsupported database driver: %s", driver)
	}
}
