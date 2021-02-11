package commands

import (
	"context"
	"time"

	"github.com/dwarukira/findcare/internal/config"
	"github.com/urfave/cli"
)

// MigrateCommand registers the migrate cli command.
var MigrateCommand = cli.Command{
	Name:   "migrate",
	Usage:  "Initializes the index database if needed",
	Action: migrateAction,
}

func migrateAction(ctx *cli.Context) error {
	start := time.Now()

	conf := config.NewConfig(ctx)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := conf.Init(); err != nil {
		return err
	}

	log.Infoln("migrating database")

	conf.InitDb()

	elapsed := time.Since(start)

	log.Infof("database migration completed in %s", elapsed)

	return nil
}
