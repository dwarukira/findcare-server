package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

// VersionCommand registers the version cli command.
var VersionCommand = cli.Command{
	Name:   "version",
	Usage:  "Shows version information",
	Action: versionAction,
}

// versionAction prints the current version
func versionAction(ctx *cli.Context) error {
	// conf := config.NewConfig(ctx)

	fmt.Println("1")

	return nil
}
