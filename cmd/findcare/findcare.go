package main

import (
	"os"

	"github.com/dwarukira/findcare/internal/commands"
	"github.com/urfave/cli"
)

var version = "development"

func main() {
	app := cli.NewApp()
	app.Name = "findcare"
	app.Usage = "Find a care provider"
	app.Version = version
	app.Copyright = "(c) 2021-2021 Duncan Warukira <dwarukira@gmail.com>"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		commands.VersionCommand,
		commands.ConfigCommand,
		commands.StartCommand,
		commands.MigrateCommand,
	}

	if err := app.Run(os.Args); err != nil {
		// log.Error(err)
	}

}
