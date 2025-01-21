package console

import (
	"fmt"

	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/support/str"
	"github.com/spf13/cast"
)

type ShowCommand struct {
	config config.Config
	schema schema.Schema
}

type databaseInfo struct {
	Name            string
	Version         string
	Database        string
	Host            string
	Port            int
	Username        string
	OpenConnections int
	Tables          []schema.Table `gorm:"-"`
	Views           []schema.View  `gorm:"-"`
}

type queryResult struct{ Value string }

func NewShowCommand(config config.Config, schema schema.Schema) *ShowCommand {
	return &ShowCommand{
		config: config,
		schema: schema,
	}
}

// Signature The name and signature of the console command.
func (r *ShowCommand) Signature() string {
	return "db:show"
}

// Description The console command description.
func (r *ShowCommand) Description() string {
	return "Display information about the given database"
}

// Extend The console command extend.
func (r *ShowCommand) Extend() command.Extend {
	return command.Extend{
		Category: "db",
		Flags: []command.Flag{
			&command.StringFlag{
				Name:    "database",
				Aliases: []string{"d"},
				Usage:   "The database connection",
			},
			&command.BoolFlag{
				Name:    "views",
				Aliases: []string{"v"},
				Usage:   "Show the database views",
			},
		},
	}
}

// Handle Execute the console command.
func (r *ShowCommand) Handle(ctx console.Context) error {
	if got := ctx.Argument(0); len(got) > 0 {
		ctx.Error(fmt.Sprintf("No arguments expected for '%s' command, got '%s'.", r.Signature(), got))
		return nil
	}

	r.schema = r.schema.Connection(ctx.Option("database"))
	dbConfig := r.schema.Orm().Config()
	info := databaseInfo{
		Database: dbConfig.Database,
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
		Username: dbConfig.Username,
	}

	var err error
	info.Name, info.Version, info.OpenConnections, err = r.getDataBaseInfo()
	if err != nil {
		ctx.Error(err.Error())
		return nil
	}

	if info.Tables, err = r.schema.GetTables(); err != nil {
		ctx.Error(err.Error())
		return nil
	}

	if ctx.OptionBool("views") {
		if info.Views, err = r.schema.GetViews(); err != nil {
			ctx.Error(err.Error())
			return nil
		}
	}
	r.display(ctx, info)

	return nil
}

func (r *ShowCommand) getDataBaseInfo() (name, version string, openConnections int, err error) {
	name = r.schema.Orm().Name()
	version = r.schema.Orm().Version()

	db, err := r.schema.Orm().DB()
	if err != nil {
		return
	}

	openConnections = db.Stats().OpenConnections

	return
}

func (r *ShowCommand) display(ctx console.Context, info databaseInfo) {
	ctx.NewLine()
	ctx.TwoColumnDetail(fmt.Sprintf("<fg=green;op=bold>%s</>", info.Name), info.Version)
	ctx.TwoColumnDetail("Database", info.Database)
	ctx.TwoColumnDetail("Host", info.Host)
	ctx.TwoColumnDetail("Port", cast.ToString(info.Port))
	ctx.TwoColumnDetail("Username", info.Username)
	ctx.TwoColumnDetail("Open Connections", cast.ToString(info.OpenConnections))
	ctx.TwoColumnDetail("Tables", cast.ToString(len(info.Tables)))
	if size := func() (size int) {
		for i := range info.Tables {
			size += info.Tables[i].Size
		}
		return
	}(); size > 0 {
		ctx.TwoColumnDetail("Total Size", fmt.Sprintf("%.3f MB", float64(size)/1024/1024))
	}
	ctx.NewLine()
	if len(info.Tables) > 0 {
		ctx.TwoColumnDetail("<fg=green;op=bold>Tables</>", "<fg=yellow;op=bold>Size (MB)</>")
		for i := range info.Tables {
			ctx.TwoColumnDetail(info.Tables[i].Name, fmt.Sprintf("%.3f", float64(info.Tables[i].Size)/1024/1024))
		}
		ctx.NewLine()
	}
	if len(info.Views) > 0 {
		ctx.TwoColumnDetail("<fg=green;op=bold>Views</>", "<fg=yellow;op=bold>Rows</>")
		for i := range info.Views {
			if !str.Of(info.Views[i].Name).StartsWith("pg_catalog", "information_schema", "spt_") {
				var rows int64
				if err := r.schema.Orm().Query().Table(info.Views[i].Name).Count(&rows); err != nil {
					ctx.Error(err.Error())
					return
				}
				ctx.TwoColumnDetail(info.Views[i].Name, fmt.Sprintf("%d", rows))
			}
		}
		ctx.NewLine()
	}
}
