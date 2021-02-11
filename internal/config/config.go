package config

import (
	"sync"

	"github.com/dwarukira/findcare/internal/event"
	"github.com/dwarukira/findcare/internal/hub"
	"github.com/dwarukira/findcare/pkg/fs"
	"github.com/dwarukira/findcare/pkg/txt"
	"github.com/jinzhu/gorm"
	"github.com/klauspost/cpuid"
	"github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var log = event.Log
var once sync.Once

// Config holds database, cache and all parameters of project
type Config struct {
	once    sync.Once
	db      *gorm.DB
	options *Options
	hub     *hub.Config
}

func init() {

}

func NewConfig(ctx *cli.Context) *Config {
	c := &Config{
		options: NewOptions(ctx),
	}

	if configFile := c.ConfigFile(); c.options.ConfigFile == "" && fs.FileExists(configFile) {
		if err := c.options.Load(configFile); err != nil {
			log.Warnf("config: %s", err)
		} else {
			log.Debugf("config: options loaded from %s", txt.Quote(configFile))
		}
	}

	return c
}

// Debug tests if debug mode is enabled.
func (c *Config) Debug() bool {
	return c.options.Debug
}

// Version returns the application version.
func (c *Config) Version() string {
	return c.options.Version
}

// ReadOnly tests if photo directories are write protected.
func (c *Config) ReadOnly() bool {
	return c.options.ReadOnly
}

// LogLevel returns the logrus log level.
func (c *Config) LogLevel() logrus.Level {
	if c.Debug() {
		c.options.LogLevel = "debug"
	}

	if logLevel, err := logrus.ParseLevel(c.options.LogLevel); err == nil {
		return logLevel
	} else {
		return logrus.InfoLevel
	}
}

// Init creates directories, parses additional config files, opens a database connection and initializes dependencies.
func (c *Config) Init() error {

	if cpuName := cpuid.CPU.BrandName; cpuName != "" {
		log.Debugf("config: running on %s", txt.Quote(cpuid.CPU.BrandName))
	}

	// c.Propagate()

	return c.connectDb()
}

// Propagate updates config options in other packages as needed.
func (c *Config) Propagate() {
	log.SetLevel(c.LogLevel())

}

// Shutdown services and workers.
func (c *Config) Shutdown() {

	if err := c.CloseDb(); err != nil {
		log.Errorf("could not close database connection: %s", err)
	} else {
		log.Info("closed database connection")
	}
}
