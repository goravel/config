package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/goravel/framework/contracts/config"
	contractsorm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/database/gorm"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/color"
)

type Orm struct {
	ctx        context.Context
	config     config.Config
	connection string
	query      contractsorm.Query
	queries    map[string]contractsorm.Query
}

func NewOrm(ctx context.Context, config config.Config, connection string, query contractsorm.Query) (*Orm, error) {
	return &Orm{
		ctx:        ctx,
		config:     config,
		connection: connection,
		query:      query,
		queries: map[string]contractsorm.Query{
			connection: query,
		},
	}, nil
}

func BuildOrm(ctx context.Context, config config.Config, connection string) (*Orm, error) {
	query, err := gorm.BuildQuery(ctx, config, connection)
	if err != nil {
		return nil, fmt.Errorf("[Orm] Build query for %s connection error: %v", connection, err)
	}

	return NewOrm(ctx, config, connection, query)
}

func (r *Orm) Connection(name string) contractsorm.Orm {
	if name == "" {
		name = r.config.GetString("database.default")
	}
	if instance, exist := r.queries[name]; exist {
		return &Orm{
			ctx:        r.ctx,
			config:     r.config,
			connection: name,
			query:      instance,
			queries:    r.queries,
		}
	}

	queue, err := gorm.BuildQuery(r.ctx, r.config, name)
	if err != nil || queue == nil {
		color.Red().Println(fmt.Sprintf("[Orm] Init %s connection error: %v", name, err))

		return nil
	}

	r.queries[name] = queue

	return &Orm{
		ctx:        r.ctx,
		config:     r.config,
		connection: name,
		query:      queue,
		queries:    r.queries,
	}
}

func (r *Orm) DB() (*sql.DB, error) {
	query := r.Query().(*gorm.Query)

	return query.Instance().DB()
}

func (r *Orm) Query() contractsorm.Query {
	return r.query
}

func (r *Orm) Factory() contractsorm.Factory {
	return NewFactoryImpl(r.Query())
}

func (r *Orm) Observe(model any, observer contractsorm.Observer) {
	orm.Observers = append(orm.Observers, orm.Observer{
		Model:    model,
		Observer: observer,
	})
}

func (r *Orm) Refresh() {
	appFacade.Refresh(BindingOrm)
}

func (r *Orm) Transaction(txFunc func(tx contractsorm.Query) error) error {
	tx, err := r.Query().Begin()
	if err != nil {
		return err
	}

	if err := txFunc(tx); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("rollback error: %v", err)
		}

		return err
	} else {
		return tx.Commit()
	}
}

func (r *Orm) WithContext(ctx context.Context) contractsorm.Orm {
	for _, query := range r.queries {
		query := query.(*gorm.Query)
		query.SetContext(ctx)
	}

	query := r.query.(*gorm.Query)
	query.SetContext(ctx)

	return &Orm{
		ctx:        ctx,
		config:     r.config,
		connection: r.connection,
		query:      query,
		queries:    r.queries,
	}
}
