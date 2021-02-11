package commands

import "github.com/urfave/cli"

// MigrateCommand registers the migrate cli command.
var MigrateCommand = cli.Command{
	Name:   "migrate",
	Usage:  "Initializes the index database if needed",
	Action: migrateAction,
}

func migrateAction(ctx *cli.Context) error {

	return nil
}
