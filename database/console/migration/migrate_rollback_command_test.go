package migration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goravel/framework/contracts/database/migration"
	"github.com/goravel/framework/database/gorm"
	mocksconsole "github.com/goravel/framework/mocks/console"
	mocksmigration "github.com/goravel/framework/mocks/database/migration"
	"github.com/goravel/framework/support/env"
	"github.com/goravel/framework/support/file"
)

func TestMigrateRollbackCommand(t *testing.T) {
	if env.IsWindows() {
		t.Skip("Skipping tests of using docker")
	}

	testQueries := gorm.NewTestQueries().Queries()
	for driver, testQuery := range testQueries {
		query := testQuery.Query()

		mockConfig := testQuery.MockConfig()
		mockConfig.EXPECT().GetString("database.migrations.table").Return("migrations").Once()
		mockConfig.EXPECT().GetString("database.migrations.driver").Return(migration.DriverSql).Once()
		mockConfig.EXPECT().GetString(fmt.Sprintf("database.connections.%s.charset", testQuery.Docker().Driver().String())).Return("utf8bm4").Once()

		createMigrations(driver)

		mockContext := mocksconsole.NewContext(t)
		mockContext.On("Option", "step").Return("1").Once()

		mockSchema := mocksmigration.NewSchema(t)

		migrateCommand := NewMigrateCommand(mockConfig, mockSchema)
		assert.Nil(t, migrateCommand.Handle(mockContext))

		var agent Agent
		err := query.Where("name", "goravel").FirstOrFail(&agent)
		assert.Nil(t, err)
		assert.True(t, agent.ID > 0)

		migrateRollbackCommand := NewMigrateRollbackCommand(mockConfig)
		assert.Nil(t, migrateRollbackCommand.Handle(mockContext))

		var agent1 Agent
		err = query.Where("name", "goravel").FirstOrFail(&agent1)
		assert.Error(t, err)
	}

	defer assert.Nil(t, file.Remove("database"))
}
