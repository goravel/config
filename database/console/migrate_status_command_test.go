package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ormcontract "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/database/gorm"
	configmock "github.com/goravel/framework/mocks/config"
	consolemocks "github.com/goravel/framework/mocks/console"
	"github.com/goravel/framework/support/docker"
	"github.com/goravel/framework/support/env"
)

func TestMigrateStatusCommand(t *testing.T) {
	if env.IsWindows() {
		t.Skip("Skipping tests of using docker")
	}

	var (
		mockConfig *configmock.Config
		query      ormcontract.Query
	)

	beforeEach := func() {
		mockConfig = &configmock.Config{}
	}

	tests := []struct {
		name  string
		setup func()
	}{
		{
			name: "mysql",
			setup: func() {
				mysqlQuery := gorm.NewTestQuery(docker.Mysql())
				query = mysqlQuery.Query()
				mockConfig = mysqlQuery.MockConfig()
				createMysqlMigrations()

			},
		},
		{
			name: "postgres",
			setup: func() {
				postgresQuery := gorm.NewTestQuery(docker.Postgres())
				query = postgresQuery.Query()
				mockConfig = postgresQuery.MockConfig()
				createPostgresMigrations()
			},
		},
		{
			name: "sqlserver",
			setup: func() {
				sqlserverQuery := gorm.NewTestQuery(docker.Sqlserver())
				query = sqlserverQuery.Query()
				mockConfig = sqlserverQuery.MockConfig()
				createSqlserverMigrations()
			},
		},
		{
			name: "sqlite",
			setup: func() {
				sqliteQuery := gorm.NewTestQuery(docker.Sqlite())
				query = sqliteQuery.Query()
				mockConfig = sqliteQuery.MockConfig()
				createSqliteMigrations()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beforeEach()
			test.setup()

			mockContext := &consolemocks.Context{}

			migrateCommand := NewMigrateCommand(mockConfig)
			assert.Nil(t, migrateCommand.Handle(mockContext))

			migrateStatusCommand := NewMigrateStatusCommand(mockConfig)
			assert.Nil(t, migrateStatusCommand.Handle(mockContext))

			res, err := query.Table("migrations").Where("dirty", false).Update("dirty", true)
			assert.Nil(t, err)
			assert.Equal(t, int64(1), res.RowsAffected)

			assert.Nil(t, migrateStatusCommand.Handle(mockContext))

			removeMigrations()
		})
	}
}
