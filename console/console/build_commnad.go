package console

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"os"
	"os/exec"
	"runtime"
)

type BuildCommand struct {
	config config.Config
}

func NewBuildCommand(config config.Config) *BuildCommand {
	return &BuildCommand{
		config: config,
	}
}

// Signature The name and signature of the console command.
func (receiver *BuildCommand) Signature() string {
	return "build"
}

// Description The console command description.
func (receiver *BuildCommand) Description() string {
	return "Build the application"
}

// Extend The console command extend.
func (receiver *BuildCommand) Extend() command.Extend {
	return command.Extend{
		Category: "build",
		Flags: []command.Flag{
			&command.StringFlag{
				Name:  "system",
				Value: runtime.GOOS,
				Usage: "target system os",
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *BuildCommand) Handle(ctx console.Context) error {
	if receiver.config.GetString("app.env") == "production" {
		color.Yellowln("**************************************")
		color.Yellowln("*     Application In Production!     *")
		color.Yellowln("**************************************")
		color.Println(color.New(color.Green).Sprintf("Do you really wish to run this command? (yes/no) ") + "[" + color.New(color.Yellow).Sprintf("no") + "]" + ":")

		var result string
		_, err := fmt.Scanln(&result)
		if err != nil {
			color.Redln(err.Error())

			return nil
		}

		if result != "yes" {
			color.Yellowln("Command Canceled")

			return nil
		}
	}

	systemOption := ctx.Option("system")

	if err := receiver.buildTheApplication(systemOption); err != nil {
		color.Redln(err.Error())

		return nil
	}

	color.Greenln("Built successfully.")

	return nil
}

// buildTheApplication Build the application executable.
func (receiver *BuildCommand) buildTheApplication(system string) error {
	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOOS", system)
	os.Setenv("GOARCH", "amd64")
	cmd := exec.Command(
		"go",
		"build",
		".",
	)
	_, err := cmd.Output()
	return err
}
