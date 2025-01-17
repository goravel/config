package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/database/schema/constants"
	mocksorm "github.com/goravel/framework/mocks/database/orm"
	mocksschema "github.com/goravel/framework/mocks/database/schema"
	"github.com/goravel/framework/support/convert"
)

type BlueprintTestSuite struct {
	suite.Suite
	blueprint *Blueprint
}

func TestBlueprintTestSuite(t *testing.T) {
	suite.Run(t, new(BlueprintTestSuite))
}

func (s *BlueprintTestSuite) SetupTest() {
	s.blueprint = NewBlueprint(nil, "goravel_", "users")
}

func (s *BlueprintTestSuite) TestAddAttributeCommands() {
	var (
		mockGrammar      *mocksschema.Grammar
		columnDefinition = &ColumnDefinition{
			comment: convert.Pointer("comment"),
		}
	)

	tests := []struct {
		name           string
		columns        []*ColumnDefinition
		setup          func()
		expectCommands []*schema.Command
	}{
		{
			name: "Should not add command when columns is empty",
			setup: func() {
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{"test"}).Once()
			},
		},
		{
			name:    "Should not add command when columns is not empty but GetAttributeCommands does not contain a valid command",
			columns: []*ColumnDefinition{columnDefinition},
			setup: func() {
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{"test"}).Once()
			},
		},
		{
			name:    "Should add comment command when columns is not empty and GetAttributeCommands contains a comment command",
			columns: []*ColumnDefinition{columnDefinition},
			setup: func() {
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{"comment"}).Once()
			},
			expectCommands: []*schema.Command{
				{
					Column: columnDefinition,
					Name:   "comment",
				},
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			mockGrammar = mocksschema.NewGrammar(s.T())
			s.blueprint.columns = test.columns
			test.setup()

			s.blueprint.addAttributeCommands(mockGrammar)
			s.Equal(test.expectCommands, s.blueprint.commands)
		})
	}
}

func (s *BlueprintTestSuite) TestBigIncrements() {
	name := "name"
	s.blueprint.BigIncrements(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		autoIncrement: convert.Pointer(true),
		name:          &name,
		unsigned:      convert.Pointer(true),
		ttype:         convert.Pointer("bigInteger"),
	})
}

func (s *BlueprintTestSuite) TestBigInteger() {
	name := "name"
	s.blueprint.BigInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:  &name,
		ttype: convert.Pointer("bigInteger"),
	})

	s.blueprint.BigInteger(name).AutoIncrement().Unsigned()
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		autoIncrement: convert.Pointer(true),
		name:          &name,
		unsigned:      convert.Pointer(true),
		ttype:         convert.Pointer("bigInteger"),
	})
}

func (s *BlueprintTestSuite) TestBuild() {
	mockQuery := mocksorm.NewQuery(s.T())
	mockGrammar := mocksschema.NewGrammar(s.T())

	tests := []struct {
		name        string
		setup       func()
		expectError string
	}{
		{
			name: "Successfully execute single SQL statement",
			setup: func() {
				// Create a table with a string column
				s.blueprint.Create()
				s.blueprint.String("name")

				// Mock the grammar to return SQL statements
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileCreate(s.blueprint).
					Return("CREATE TABLE users (name VARCHAR(255))").Once()

				// Mock successful query execution
				mockQuery.EXPECT().Exec("CREATE TABLE users (name VARCHAR(255))").
					Return(nil, nil).Once()
			},
		},
		{
			name: "Successfully execute multiple SQL statements",
			setup: func() {
				// Create a table with multiple columns
				s.blueprint.Create()
				s.blueprint.String("name")
				s.blueprint.Integer("age")

				// Mock the grammar to return multiple SQL statements
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileCreate(s.blueprint).
					Return("CREATE TABLE users (name VARCHAR(255), age INTEGER)").Once()

				// Mock successful query execution
				mockQuery.EXPECT().Exec("CREATE TABLE users (name VARCHAR(255), age INTEGER)").
					Return(nil, nil).Once()
			},
		},
		{
			name: "Return error when query execution fails",
			setup: func() {
				s.blueprint.Create()
				s.blueprint.String("name")

				// Mock the grammar to return SQL statements
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileCreate(s.blueprint).
					Return("CREATE TABLE users (name VARCHAR(255))").Once()

				// Mock failed query execution
				mockQuery.EXPECT().Exec("CREATE TABLE users (name VARCHAR(255))").
					Return(nil, assert.AnError).Once()
			},
			expectError: assert.AnError.Error(),
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// Reset blueprint for each test
			s.SetupTest()

			// Setup test case
			test.setup()

			// Execute Build
			err := s.blueprint.Build(mockQuery, mockGrammar)

			// Verify results
			if test.expectError == "" {
				s.Nil(err)
			} else {
				s.EqualError(err, test.expectError)
			}
		})
	}
}

func (s *BlueprintTestSuite) TestChar() {
	column := "name"
	customLength := 100
	length := constants.DefaultStringLength
	ttype := "char"
	s.blueprint.Char(column)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		length: &length,
		name:   &column,
		ttype:  &ttype,
	})

	s.blueprint.Char(column, customLength)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		length: &customLength,
		name:   &column,
		ttype:  &ttype,
	})
}

func (s *BlueprintTestSuite) TestCreateIndexName() {
	name := s.blueprint.createIndexName("index", []string{"id", "name-1", "name.2"})
	s.Equal("goravel_users_id_name_1_name_2_index", name)

	s.blueprint.table = "public.users"
	name = s.blueprint.createIndexName("index", []string{"id", "name-1", "name.2"})
	s.Equal("public_goravel_users_id_name_1_name_2_index", name)
}

func (s *BlueprintTestSuite) TestDecimal() {
	name := "name"
	s.blueprint.Decimal(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:  &name,
		ttype: convert.Pointer("decimal"),
	})
}

func (s *BlueprintTestSuite) TestDouble() {
	name := "name"
	s.blueprint.Double(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:  &name,
		ttype: convert.Pointer("double"),
	})
}

func (s *BlueprintTestSuite) TestFloat() {
	name := "name"
	s.blueprint.Float(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:      &name,
		precision: convert.Pointer(53),
		ttype:     convert.Pointer("float"),
	})

	s.blueprint.Float(name, 10)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:      &name,
		precision: convert.Pointer(10),
		ttype:     convert.Pointer("float"),
	})
}

func (s *BlueprintTestSuite) TestGetAddedColumns() {
	name := "name"
	addedColumn := &ColumnDefinition{
		name: &name,
	}

	s.blueprint.columns = []*ColumnDefinition{addedColumn}

	s.Len(s.blueprint.GetAddedColumns(), 1)
	s.Equal(addedColumn, s.blueprint.GetAddedColumns()[0])
}

func (s *BlueprintTestSuite) TestHasCommand() {
	s.False(s.blueprint.HasCommand(constants.CommandCreate))
	s.blueprint.Create()
	s.True(s.blueprint.HasCommand(constants.CommandCreate))
}

func (s *BlueprintTestSuite) TestIntegerIncrements() {
	name := "name"
	s.blueprint.IntegerIncrements(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		autoIncrement: convert.Pointer(true),
		name:          &name,
		unsigned:      convert.Pointer(true),
		ttype:         convert.Pointer("integer"),
	})
}

func (s *BlueprintTestSuite) TestIndexCommand() {
	s.blueprint.indexCommand("index", []string{"id", "name"})
	s.Contains(s.blueprint.commands, &schema.Command{
		Columns: []string{"id", "name"},
		Name:    "index",
		Index:   "goravel_users_id_name_index",
	})

	s.blueprint.indexCommand("index", []string{"id", "name"}, schema.IndexConfig{
		Algorithm: "custom_algorithm",
		Name:      "custom_name",
	})
	s.Contains(s.blueprint.commands, &schema.Command{
		Algorithm: "custom_algorithm",
		Columns:   []string{"id", "name"},
		Name:      "index",
		Index:     "custom_name",
	})
}

func (s *BlueprintTestSuite) TestIsCreate() {
	s.False(s.blueprint.isCreate())
	s.blueprint.commands = []*schema.Command{
		{
			Name: constants.CommandCreate,
		},
	}
	s.True(s.blueprint.isCreate())
}

func (s *BlueprintTestSuite) TestID() {
	s.blueprint.ID()
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		autoIncrement: convert.Pointer(true),
		name:          convert.Pointer("id"),
		unsigned:      convert.Pointer(true),
		ttype:         convert.Pointer("bigInteger"),
	})

	s.blueprint.ID("name")
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		autoIncrement: convert.Pointer(true),
		name:          convert.Pointer("name"),
		unsigned:      convert.Pointer(true),
		ttype:         convert.Pointer("bigInteger"),
	})
}

func (s *BlueprintTestSuite) TestInteger() {
	name := "name"
	s.blueprint.Integer(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:  &name,
		ttype: convert.Pointer("integer"),
	})
}

func (s *BlueprintTestSuite) TestMediumIncrements() {
	name := "name"
	s.blueprint.MediumIncrements(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		autoIncrement: convert.Pointer(true),
		name:          &name,
		unsigned:      convert.Pointer(true),
		ttype:         convert.Pointer("mediumInteger"),
	})
}

func (s *BlueprintTestSuite) TestMediumInteger() {
	name := "name"
	s.blueprint.MediumInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:  &name,
		ttype: convert.Pointer("mediumInteger"),
	})
}

func (s *BlueprintTestSuite) TestSmallIncrements() {
	name := "name"
	s.blueprint.SmallIncrements(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		autoIncrement: convert.Pointer(true),
		name:          &name,
		unsigned:      convert.Pointer(true),
		ttype:         convert.Pointer("smallInteger"),
	})
}

func (s *BlueprintTestSuite) TestSmallInteger() {
	name := "name"
	s.blueprint.SmallInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:  &name,
		ttype: convert.Pointer("smallInteger"),
	})
}

func (s *BlueprintTestSuite) TestString() {
	column := "name"
	customLength := 100
	length := constants.DefaultStringLength
	ttype := "string"
	s.blueprint.String(column)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		length: &length,
		name:   &column,
		ttype:  &ttype,
	})

	s.blueprint.String(column, customLength)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		length: &customLength,
		name:   &column,
		ttype:  &ttype,
	})
}

func (s *BlueprintTestSuite) TestTinyIncrements() {
	name := "name"
	s.blueprint.TinyIncrements(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		autoIncrement: convert.Pointer(true),
		name:          &name,
		unsigned:      convert.Pointer(true),
		ttype:         convert.Pointer("tinyInteger"),
	})
}

func (s *BlueprintTestSuite) TestTinyInteger() {
	name := "name"
	s.blueprint.TinyInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:  &name,
		ttype: convert.Pointer("tinyInteger"),
	})
}

func (s *BlueprintTestSuite) TestToSql() {
	mockGrammar := mocksschema.NewGrammar(s.T())

	tests := []struct {
		name        string
		setup       func()
		expectedSQL []string
	}{
		{
			name: "Add command with change",
			setup: func() {
				s.blueprint.String("name").Change()
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileChange(s.blueprint, s.blueprint.commands[0]).
					Return([]string{"ALTER TABLE users MODIFY name VARCHAR(255)"}).Once()
			},
			expectedSQL: []string{"ALTER TABLE users MODIFY name VARCHAR(255)"},
		},
		{
			name: "Add command without change",
			setup: func() {
				s.blueprint.String("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileAdd(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users ADD name VARCHAR(255)").Once()
			},
			expectedSQL: []string{"ALTER TABLE users ADD name VARCHAR(255)"},
		},
		{
			name: "Comment command",
			setup: func() {
				s.blueprint.String("name").Comment("test comment")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{"comment"}).Once()
				mockGrammar.EXPECT().CompileAdd(s.blueprint, s.blueprint.commands[0]).
					Return("sql1").Once()
				mockGrammar.EXPECT().CompileComment(s.blueprint, &schema.Command{
					Column: s.blueprint.columns[0],
					Name:   constants.CommandComment,
				}).Return("sql2").Once()
			},
			expectedSQL: []string{"sql1", "sql2"},
		},
		{
			name: "Default command",
			setup: func() {
				s.blueprint.String("name").Default("test")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{"default"}).Once()
				mockGrammar.EXPECT().CompileAdd(s.blueprint, s.blueprint.commands[0]).
					Return("sql1").Once()
				mockGrammar.EXPECT().CompileDefault(s.blueprint, &schema.Command{
					Column: s.blueprint.columns[0],
					Name:   constants.CommandDefault,
				}).Return("sql2").Once()
			},
			expectedSQL: []string{"sql1", "sql2"},
		},
		{
			name: "Drop command",
			setup: func() {
				s.blueprint.Drop()
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileDrop(s.blueprint).
					Return("DROP TABLE users").Once()
			},
			expectedSQL: []string{"DROP TABLE users"},
		},
		{
			name: "Drop column command",
			setup: func() {
				s.blueprint.DropColumn("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileDropColumn(s.blueprint, s.blueprint.commands[0]).
					Return([]string{"ALTER TABLE users DROP COLUMN name"}).Once()
			},
			expectedSQL: []string{"ALTER TABLE users DROP COLUMN name"},
		},
		{
			name: "Drop foreign command",
			setup: func() {
				s.blueprint.DropForeign("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileDropForeign(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users DROP FOREIGN KEY name").Once()
			},
			expectedSQL: []string{"ALTER TABLE users DROP FOREIGN KEY name"},
		},
		{
			name: "Drop full text command",
			setup: func() {
				s.blueprint.DropFullText("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileDropFullText(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users DROP FULLTEXT INDEX name").Once()
			},
			expectedSQL: []string{"ALTER TABLE users DROP FULLTEXT INDEX name"},
		},
		{
			name: "Drop if exists command",
			setup: func() {
				s.blueprint.DropIfExists()
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileDropIfExists(s.blueprint).
					Return("DROP TABLE IF EXISTS users").Once()
			},
			expectedSQL: []string{"DROP TABLE IF EXISTS users"},
		},
		{
			name: "Drop index command",
			setup: func() {
				s.blueprint.DropIndex("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileDropIndex(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users DROP INDEX name").Once()
			},
			expectedSQL: []string{"ALTER TABLE users DROP INDEX name"},
		},
		{
			name: "Drop primary command",
			setup: func() {
				s.blueprint.DropPrimary("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileDropPrimary(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users DROP PRIMARY KEY name").Once()
			},
			expectedSQL: []string{"ALTER TABLE users DROP PRIMARY KEY name"},
		},
		{
			name: "Drop unique command",
			setup: func() {
				s.blueprint.DropUnique("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileDropUnique(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users DROP UNIQUE name").Once()
			},
			expectedSQL: []string{"ALTER TABLE users DROP UNIQUE name"},
		},
		{
			name: "Foreign key command",
			setup: func() {
				s.blueprint.Foreign("user_id").References("id").On("users")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileForeign(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)").Once()
			},
			expectedSQL: []string{"ALTER TABLE users ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)"},
		},
		{
			name: "Index command",
			setup: func() {
				s.blueprint.Index("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileIndex(s.blueprint, s.blueprint.commands[0]).
					Return("CREATE INDEX goravel_users_name_index ON users (name)").Once()
			},
			expectedSQL: []string{"CREATE INDEX goravel_users_name_index ON users (name)"},
		},
		{
			name: "Primary key command",
			setup: func() {
				s.blueprint.Primary("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompilePrimary(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users ADD PRIMARY KEY (name)").Once()
			},
			expectedSQL: []string{"ALTER TABLE users ADD PRIMARY KEY (name)"},
		},
		{
			name: "Rename command",
			setup: func() {
				s.blueprint.Rename("new_users")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileRename(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users RENAME TO new_users").Once()
			},
			expectedSQL: []string{"ALTER TABLE users RENAME TO new_users"},
		},
		{
			name: "Rename index command",
			setup: func() {
				s.blueprint.RenameIndex("old_index", "new_index")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileRenameIndex(s.blueprint.schema, s.blueprint, s.blueprint.commands[0]).
					Return([]string{"ALTER INDEX old_index RENAME TO new_index"}).Once()
			},
			expectedSQL: []string{"ALTER INDEX old_index RENAME TO new_index"},
		},
		{
			name: "Unique command",
			setup: func() {
				s.blueprint.Unique("name")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileUnique(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users ADD CONSTRAINT unique_name UNIQUE (name)").Once()
			},
			expectedSQL: []string{"ALTER TABLE users ADD CONSTRAINT unique_name UNIQUE (name)"},
		},
		{
			name: "Multiple commands",
			setup: func() {
				s.blueprint.String("name")
				s.blueprint.Primary("id")
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
				mockGrammar.EXPECT().CompileAdd(s.blueprint, s.blueprint.commands[0]).
					Return("ALTER TABLE users ADD name VARCHAR(255)").Once()
				mockGrammar.EXPECT().CompilePrimary(s.blueprint, s.blueprint.commands[1]).
					Return("ALTER TABLE users ADD PRIMARY KEY (id)").Once()
			},
			expectedSQL: []string{
				"ALTER TABLE users ADD name VARCHAR(255)",
				"ALTER TABLE users ADD PRIMARY KEY (id)",
			},
		},
		{
			name: "Skip command",
			setup: func() {
				s.blueprint.Create()
				s.blueprint.commands[0].ShouldBeSkipped = true
				mockGrammar.EXPECT().GetAttributeCommands().Return([]string{}).Once()
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// Reset blueprint for each test
			s.SetupTest()

			// Setup test case
			test.setup()

			// Execute ToSql
			statements := s.blueprint.ToSql(mockGrammar)

			// Verify results
			s.Equal(test.expectedSQL, statements)
		})
	}
}

func (s *BlueprintTestSuite) TestUnsignedBigInteger() {
	name := "name"
	s.blueprint.UnsignedBigInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:     &name,
		ttype:    convert.Pointer("bigInteger"),
		unsigned: convert.Pointer(true),
	})
}

func (s *BlueprintTestSuite) TestUnsignedInteger() {
	name := "name"
	s.blueprint.UnsignedInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:     &name,
		ttype:    convert.Pointer("integer"),
		unsigned: convert.Pointer(true),
	})
}

func (s *BlueprintTestSuite) TestUnsignedMediumInteger() {
	name := "name"
	s.blueprint.UnsignedMediumInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:     &name,
		ttype:    convert.Pointer("mediumInteger"),
		unsigned: convert.Pointer(true),
	})
}

func (s *BlueprintTestSuite) TestUnsignedSmallInteger() {
	name := "name"
	s.blueprint.UnsignedSmallInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:     &name,
		ttype:    convert.Pointer("smallInteger"),
		unsigned: convert.Pointer(true),
	})
}

func (s *BlueprintTestSuite) TestUnsignedTinyInteger() {
	name := "name"
	s.blueprint.UnsignedTinyInteger(name)
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		name:     &name,
		ttype:    convert.Pointer("tinyInteger"),
		unsigned: convert.Pointer(true),
	})
}

func (s *BlueprintTestSuite) TestChange() {
	column := "name"
	customLength := 100
	s.blueprint.String(column, customLength).Change()
	s.Contains(s.blueprint.GetAddedColumns(), &ColumnDefinition{
		length: &customLength,
		name:   &column,
		change: true,
		ttype:  convert.Pointer("string"),
	})
}
