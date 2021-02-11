package hub

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dwarukira/findcare/pkg/fs"
	"github.com/dwarukira/findcare/pkg/txt"
	"gopkg.in/yaml.v2"
)

// Config represents backend api credentials for maps & geodata.
type Config struct {
	Key      string `json:"key" yaml:"Key"`
	Secret   string `json:"secret" yaml:"Secret"`
	Session  string `json:"session" yaml:"Session"`
	Status   string `json:"status" yaml:"Status"`
	Version  string `json:"version" yaml:"Version"`
	Serial   string `json:"serial" yaml:"Serial"`
	FileName string `json:"-" yaml:"-"`
}

// NewConfig creates a new backend api credentials instance.
func NewConfig(version, fileName, serial string) *Config {
	return &Config{
		Key:      "",
		Secret:   "",
		Session:  "",
		Status:   "",
		Version:  version,
		Serial:   serial,
		FileName: fileName,
	}
}

// Propagate updates backend api credentials in other packages.
func (c *Config) Propagate() {

}

// Sanitize verifies and sanitizes backend api credentials.
func (c *Config) Sanitize() {
	c.Key = strings.ToLower(c.Key)

	if c.Secret != "" {
		if c.Key != fmt.Sprintf("%x", sha1.Sum([]byte(c.Secret))) {
			c.Key = ""
			c.Secret = ""
			c.Session = ""
			c.Status = ""
		}
	}
}

func (c *Config) Load() error {
	if !fs.FileExists(c.FileName) {
		return fmt.Errorf("settings file not found: %s", txt.Quote(c.FileName))
	}

	mutex.Lock()
	defer mutex.Unlock()

	yamlConfig, err := ioutil.ReadFile(c.FileName)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlConfig, c); err != nil {
		return err
	}

	c.Sanitize()
	c.Propagate()

	return nil
}
