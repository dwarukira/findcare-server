package commands

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/dwarukira/findcare/internal/config"
	"github.com/dwarukira/findcare/internal/server"
	"github.com/dwarukira/findcare/pkg/fs"
	"github.com/dwarukira/findcare/pkg/txt"
	"github.com/sevlyar/go-daemon"
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
	cctx, cancel := context.WithCancel(context.Background())

	if err := conf.Init(); err != nil {
		log.Fatal(err)
	}

	// initialize the database
	conf.InitDb()

	// check if daemon is running, if not initialize the daemon
	dctx := new(daemon.Context)
	dctx.LogFileName = conf.LogFilename()
	dctx.PidFileName = conf.PIDFilename()
	dctx.Args = ctx.Args()

	if !daemon.WasReborn() && conf.DetachServer() {
		conf.Shutdown()
		cancel()

		if pid, ok := childAlreadyRunning(conf.PIDFilename()); ok {
			log.Infof("daemon already running with process id %v\n", pid)
			return nil
		}

		child, err := dctx.Reborn()
		if err != nil {
			log.Fatal(err)
		}

		if child != nil {
			if !fs.Overwrite(conf.PIDFilename(), []byte(strconv.Itoa(child.Pid))) {
				log.Fatalf("failed writing process id to %s", txt.Quote(conf.PIDFilename()))
			}

			log.Infof("daemon started with process id %v\n", child.Pid)
			return nil
		}
	}

	if conf.ReadOnly() {
		log.Infof("start: read-only mode enabled")
	}

	// start web server
	go server.Start(cctx, conf)

	// set up proper shutdown of daemon and web server
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Info("shutting down...")

	conf.Shutdown()
	cancel()
	err := dctx.Release()

	if err != nil {
		log.Error(err)
	}

	time.Sleep(3 * time.Second)

	return nil
}
