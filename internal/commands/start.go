package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/dwarukira/findcare/internal/config"
	"github.com/urfave/cli"
)

// StartCommand registers the start cli command.
var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"up"},
	Usage:   "Starts web server",
	Flags:   startFlags,
	Action:  startAction,
}

var startFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "detach-server, d",
		Usage:  "detach from the console (daemon mode)",
		EnvVar: "HR_DETACH_SERVER",
	},
	cli.BoolFlag{
		Name:  "config, c",
		Usage: "show config",
	},
}

// startAction start the web server and initializes the daemon
func startAction(ctx *cli.Context) error {

	conf := config.NewConfig(ctx)

	if ctx.IsSet("config") {
		fmt.Printf("NAME                  VALUE\n")
		fmt.Printf("detach-server         %t\n", conf.DetachServer())

		fmt.Printf("http-host             %s\n", conf.HttpHost())
		fmt.Printf("http-port             %d\n", conf.HttpPort())
		fmt.Printf("http-mode             %s\n", conf.HttpMode())

		return nil
	}

	if conf.HttpPort() < 1 || conf.HttpPort() > 65535 {
		log.Fatal("server port must be a number between 1 and 65535")
	}

	// pass this context down the chain
	_, cancel := context.WithCancel(context.Background())

	log.Info("shutting down...")

	cancel()

	time.Sleep(3 * time.Second)

	return nil
}
