package commands

import (
	"fmt"

	"github.com/dwarukira/findcare/internal/config"
	"github.com/urfave/cli"
)

// ConfigCommand registers the display config cli command.
var ConfigCommand = cli.Command{
	Name:   "config",
	Usage:  "Displays global configuration values",
	Action: configAction,
}

// configAction lists configuration options and their values.
func configAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)

	fmt.Printf("%-25s VALUE\n", "NAME")

	// Feature flags.
	fmt.Printf("%-25s %t\n", "debug", conf.Debug())

	// Config path and main file.
	fmt.Printf("%-25s %s\n", "config-file", conf.ConfigFile())
	fmt.Printf("%-25s %s\n", "config-path", conf.ConfigPath())

	// Logging.
	// fmt.Printf("%-25s %s\n", "log-level", conf.LogLevel())
	// fmt.Printf("%-25s %s\n", "log-filename", conf.LogFilename())

	// HTTP server configuration.
	fmt.Printf("%-25s %s\n", "http-host", conf.HttpHost())
	fmt.Printf("%-25s %d\n", "http-port", conf.HttpPort())
	fmt.Printf("%-25s %s\n", "http-mode", conf.HttpMode())

	return nil
}
