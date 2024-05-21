package console

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/support/color"
	"github.com/goravel/framework/support/file"
	"github.com/goravel/framework/support/str"
)

type TestMakeCommand struct {
}

func NewTestMakeCommand() *TestMakeCommand {
	return &TestMakeCommand{}
}

// Signature The name and signature of the console command.
func (receiver *TestMakeCommand) Signature() string {
	return "make:test"
}

// Description The console command description.
func (receiver *TestMakeCommand) Description() string {
	return "Create a new test class"
}

// Extend The console command extend.
func (receiver *TestMakeCommand) Extend() command.Extend {
	return command.Extend{
		Category: "make",
		Flags: []command.Flag{
			&command.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Create the test even if it already exists",
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *TestMakeCommand) Handle(ctx console.Context) error {
	name := ctx.Argument(0)
	if name == "" {
		var err error
		name, err = ctx.Ask("Enter the test name", console.AskOption{
			Validate: func(s string) error {
				if s == "" {
					return errors.New("the test name cannot be empty")
				}

				return nil
			},
		})
		if err != nil {
			return err
		}
	}

	force := ctx.OptionBool("force")
	path := receiver.getPath(name)
	if !force && file.Exists(path) {
		color.Red().Println("The test already exists. Use the --force flag to overwrite")
		return nil
	}

	stub := receiver.getStub()

	if err := file.Create(path, receiver.populateStub(stub, name)); err != nil {
		return err
	}

	color.Green().Println("Test created successfully")

	return nil
}

func (receiver *TestMakeCommand) getStub() string {
	return Stubs{}.Test()
}

// populateStub Populate the place-holders in the command stub.
func (receiver *TestMakeCommand) populateStub(stub string, name string) string {
	controllerName, packageName, _ := receiver.parseName(name)

	stub = strings.ReplaceAll(stub, "DummyTest", str.Case2Camel(controllerName))
	stub = strings.ReplaceAll(stub, "DummyPackage", packageName)

	return stub
}

// getPath Get the full path to the command.
func (receiver *TestMakeCommand) getPath(name string) string {
	pwd, _ := os.Getwd()

	controllerName, _, folderPath := receiver.parseName(name)

	return filepath.Join(pwd, "tests", folderPath, str.Camel2Case(controllerName)+".go")
}

// parseName Parse the name to get the controller name, package name and folder path.
func (receiver *TestMakeCommand) parseName(name string) (string, string, string) {
	name = strings.TrimSuffix(name, ".go")

	segments := strings.Split(name, "/")

	controllerName := segments[len(segments)-1]

	packageName := "tests"
	folderPath := ""

	if len(segments) > 1 {
		folderPath = filepath.Join(segments[:len(segments)-1]...)
		packageName = segments[len(segments)-2]
	}

	return controllerName, packageName, folderPath
}
