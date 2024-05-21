package console

import (
	"path/filepath"
	"strings"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/support/color"
	"github.com/goravel/framework/support/file"
)

type PackageMakeCommand struct{}

func NewPackageMakeCommand() *PackageMakeCommand {
	return &PackageMakeCommand{}
}

// Signature The name and signature of the console command.
func (receiver *PackageMakeCommand) Signature() string {
	return "make:package"
}

// Description The console command description.
func (receiver *PackageMakeCommand) Description() string {
	return "Create a package template"
}

// Extend The console command extend.
func (receiver *PackageMakeCommand) Extend() command.Extend {
	return command.Extend{
		Category: "make",
		Flags: []command.Flag{
			&command.StringFlag{
				Name:    "root",
				Aliases: []string{"r"},
				Usage:   "The root path of package, default: packages",
				Value:   "packages",
			},
			&command.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Create the package even if it already exists",
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *PackageMakeCommand) Handle(ctx console.Context) error {
	pkg := ctx.Argument(0)
	if pkg == "" {
		color.Red().Println("Not enough arguments (missing: name)")

		return nil
	}

	force := ctx.OptionBool("force")
	pkg = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(pkg, "/", "_"), "-", "_"), ".", "_")
	root := ctx.Option("root") + "/" + pkg

	if !force && file.Exists(root) {
		color.Redf("Package %s already exists\n", pkg)
		return nil
	}

	packageName := packageName(pkg)
	packageMakeCommandStubs := NewPackageMakeCommandStubs(pkg, root)
	files := map[string]func() string{
		"README.md":                        packageMakeCommandStubs.Readme,
		"service_provider.go":              packageMakeCommandStubs.ServiceProvider,
		packageName + ".go":                packageMakeCommandStubs.Main,
		"config/" + packageName + ".go":    packageMakeCommandStubs.Config,
		"contracts/" + packageName + ".go": packageMakeCommandStubs.Contracts,
		"facades/" + packageName + ".go":   packageMakeCommandStubs.Facades,
	}

	for path, content := range files {
		if err := file.Create(filepath.Join(root, path), content()); err != nil {
			return err
		}
	}

	color.Green().Printf("Package created successfully: %s\n", root)

	return nil
}

func packageName(name string) string {
	nameSlice := strings.Split(name, "/")
	lastName := nameSlice[len(nameSlice)-1]

	return strings.ReplaceAll(strings.ReplaceAll(lastName, "-", "_"), ".", "_")
}
