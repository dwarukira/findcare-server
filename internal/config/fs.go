package config

import (
	"os/user"
	"path/filepath"

	"github.com/dwarukira/findcare/pkg/fs"
)

// ConfigFile returns the config file name.
func (c *Config) ConfigFile() string {
	if c.options.ConfigFile == "" || !fs.FileExists(c.options.ConfigFile) {
		return filepath.Join(c.ConfigPath(), "options.yml")
	}

	return c.options.ConfigFile
}

// ConfigPath returns the config path.
func (c *Config) ConfigPath() string {
	if c.options.ConfigPath == "" {
		if fs.PathExists(filepath.Join(c.StoragePath(), "settings")) {
			return filepath.Join(c.StoragePath(), "settings")
		}

		return filepath.Join(c.StoragePath(), "config")
	}

	return fs.Abs(c.options.ConfigPath)
}

// StoragePath returns the path for generated files like cache and index.
func (c *Config) StoragePath() string {
	if c.options.StoragePath == "" {
		const dirName = "storage"

		// Default directories.
		originalsDir := fs.Abs(filepath.Join(c.OriginalsPath(), fs.HiddenPath, dirName))
		storageDir := fs.Abs(dirName)

		// Find existing directories.
		if fs.PathWritable(originalsDir) && !c.ReadOnly() {
			return originalsDir
		} else if fs.PathWritable(storageDir) && c.ReadOnly() {
			return storageDir
		}

		// Use .hr in home directory?
		if usr, _ := user.Current(); usr.HomeDir != "" {
			p := fs.Abs(filepath.Join(usr.HomeDir, fs.HiddenPath, dirName))

			if fs.PathWritable(p) || c.ReadOnly() {
				return p
			}
		}

		// Store cache and index in "originals/.photoprism/storage".
		return originalsDir
	}

	return fs.Abs(c.options.StoragePath)
}

// OriginalsPath returns the originals.
func (c *Config) OriginalsPath() string {
	if c.options.OriginalsPath == "" {
		// Try to find the right directory by iterating through a list.
		c.options.OriginalsPath = fs.FindDir(fs.OriginalPaths)
	}

	return fs.Abs(c.options.OriginalsPath)
}
