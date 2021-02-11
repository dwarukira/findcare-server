package commands

import (
	"fmt"
	"strings"
	"unicode/utf8"

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

	dbDriver := conf.DatabaseDriver()

	fmt.Printf("%-25s VALUE\n", "NAME")

	// Feature flags.
	fmt.Printf("%-25s %t\n", "debug", conf.Debug())

	// Config path and main file.
	fmt.Printf("%-25s %s\n", "config-file", conf.ConfigFile())
	fmt.Printf("%-25s %s\n", "config-path", conf.ConfigPath())

	// Logging.
	fmt.Printf("%-25s %s\n", "log-level", conf.LogLevel())
	fmt.Printf("%-25s %s\n", "log-filename", conf.LogFilename())

	// Asset path and file names.
	fmt.Printf("%-25s %s\n", "static-path", conf.StaticPath())
	// fmt.Printf("%-25s %s\n", "build-path", conf.BuildPath())
	// fmt.Printf("%-25s %s\n", "img-path", conf.ImgPath())
	fmt.Printf("%-25s %s\n", "templates-path", conf.TemplatesPath())

	// HTTP server configuration.
	fmt.Printf("%-25s %s\n", "http-host", conf.HttpHost())
	fmt.Printf("%-25s %d\n", "http-port", conf.HttpPort())
	fmt.Printf("%-25s %s\n", "http-mode", conf.HttpMode())

	// Database configuration.
	fmt.Printf("%-25s %s\n", "database-driver", dbDriver)
	fmt.Printf("%-25s %s\n", "database-server", conf.DatabaseServer())
	fmt.Printf("%-25s %s\n", "database-host", conf.DatabaseHost())
	fmt.Printf("%-25s %s\n", "database-port", conf.DatabasePortString())
	fmt.Printf("%-25s %s\n", "database-name", conf.DatabaseName())
	fmt.Printf("%-25s %s\n", "database-user", conf.DatabaseUser())
	fmt.Printf("%-25s %s\n", "database-password", strings.Repeat("*", utf8.RuneCountInString(conf.DatabasePassword())))
	fmt.Printf("%-25s %d\n", "database-conns", conf.DatabaseConns())
	fmt.Printf("%-25s %d\n", "database-conns-idle", conf.DatabaseConnsIdle())

	return nil
}
