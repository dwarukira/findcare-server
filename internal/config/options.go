package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/dwarukira/findcare/pkg/fs"
	"github.com/urfave/cli"

	yaml "gopkg.in/yaml.v2"
)

// Database drivers (sql dialects).
const (
	MySQL    = "mysql"
	MariaDB  = "mariadb"
	SQLite   = "sqlite3"
	Postgres = "postgres" // TODO: Requires GORM 2.0 for generic column data types
)

type Options struct {
	Name              string `json:"-"`
	Version           string `json:"-"`
	Copyright         string `json:"-"`
	ReadOnly          bool   `yaml:"ReadOnly" json:"ReadOnly" flag:"read-only"`
	ConfigPath        string `yaml:"ConfigPath" json:"-" flag:"config-path"`
	ConfigFile        string `json:"-"`
	OriginalsPath     string `yaml:"OriginalsPath" json:"-" flag:"originals-path"`
	OriginalsLimit    int64  `yaml:"OriginalsLimit" json:"OriginalsLimit" flag:"originals-limit"`
	ImportPath        string `yaml:"ImportPath" json:"-" flag:"import-path"`
	StoragePath       string `yaml:"StoragePath" json:"-" flag:"storage-path"`
	Debug             bool   `yaml:"Debug" json:"Debug" flag:"debug"`
	Public            bool   `yaml:"Public" json:"-" flag:"public"`
	LogLevel          string `yaml:"LogLevel" json:"-" flag:"log-level"`
	LogFilename       string `yaml:"LogFilename" json:"-" flag:"log-filename"`
	DatabaseDriver    string `yaml:"DatabaseDriver" json:"-" flag:"database-driver"`
	DatabaseDsn       string `yaml:"DatabaseDsn" json:"-" flag:"database-dsn"`
	DatabaseServer    string `yaml:"DatabaseServer" json:"-" flag:"database-server"`
	DatabaseName      string `yaml:"DatabaseName" json:"-" flag:"database-name"`
	DatabaseUser      string `yaml:"DatabaseUser" json:"-" flag:"database-user"`
	DatabasePassword  string `yaml:"DatabasePassword" json:"-" flag:"database-password"`
	DatabaseConns     int    `yaml:"DatabaseConns" json:"-" flag:"database-conns"`
	DatabaseConnsIdle int    `yaml:"DatabaseConnsIdle" json:"-" flag:"database-conns-idle"`
	HttpHost          string `yaml:"HttpHost" json:"-" flag:"http-host"`
	HttpPort          int    `yaml:"HttpPort" json:"-" flag:"http-port"`
	HttpMode          string `yaml:"HttpMode" json:"-" flag:"http-mode"`
	HttpCompression   string `yaml:"HttpCompression" json:"-" flag:"http-compression"`
	DetachServer      bool   `yaml:"DetachServer" json:"-" flag:"detach-server"`
}

func NewOptions(ctx *cli.Context) *Options {
	c := &Options{}

	if ctx == nil {
		return c
	}

	c.Name = ctx.App.Name
	c.Copyright = ctx.App.Copyright
	c.Version = ctx.App.Version
	c.ConfigFile = fs.Abs(ctx.GlobalString("config-file"))

	if err := c.Load(c.ConfigFile); err != nil {
		log.Debug(err)
	}

	if err := c.SetContext(ctx); err != nil {
		log.Error(err)
	}

	return c
}

// Load uses a yaml config file to initiate the configuration entity.
func (c *Options) Load(fileName string) error {
	if fileName == "" {
		return nil
	}

	if !fs.FileExists(fileName) {
		return errors.New(fmt.Sprintf("config: %s not found", fileName))
	}

	yamlConfig, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlConfig, c)
}

// SetContext uses options from the CLI to setup configuration overrides
// for the entity.
func (c *Options) SetContext(ctx *cli.Context) error {
	v := reflect.ValueOf(c).Elem()

	// Iterate through all config fields.
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)

		tagValue := v.Type().Field(i).Tag.Get("flag")

		// Automatically assign options to fields with "flag" tag.
		if tagValue != "" {
			switch t := fieldValue.Interface().(type) {
			case int, int64:
				// Only if explicitly set or current value is empty (use default).
				if ctx.IsSet(tagValue) {
					f := ctx.Int64(tagValue)
					fieldValue.SetInt(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.Int() == 0 {
					f := ctx.GlobalInt64(tagValue)
					fieldValue.SetInt(f)
				}
			case uint, uint64:
				// Only if explicitly set or current value is empty (use default).
				if ctx.IsSet(tagValue) {
					f := ctx.Uint64(tagValue)
					fieldValue.SetUint(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.Uint() == 0 {
					f := ctx.GlobalUint64(tagValue)
					fieldValue.SetUint(f)
				}
			case string:
				// Only if explicitly set or current value is empty (use default)
				if ctx.IsSet(tagValue) {
					f := ctx.String(tagValue)
					fieldValue.SetString(f)
				} else if ctx.GlobalIsSet(tagValue) || fieldValue.String() == "" {
					f := ctx.GlobalString(tagValue)
					fieldValue.SetString(f)
				}
			case bool:
				if ctx.IsSet(tagValue) {
					f := ctx.Bool(tagValue)
					fieldValue.SetBool(f)
				} else if ctx.GlobalIsSet(tagValue) {
					f := ctx.GlobalBool(tagValue)
					fieldValue.SetBool(f)
				}
			default:
				log.Warnf("can't assign value of type %s from cli flag %s", t, tagValue)
			}
		}
	}

	return nil
}
