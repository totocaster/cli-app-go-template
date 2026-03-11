package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// FileName is the default configuration file name for the CLI.
const FileName = "config.toml"

// Location describes where the default config directory and file live.
type Location struct {
	Dir  string `json:"dir"`
	File string `json:"file"`
}

// NewLocation derives the default config file path from a config directory.
func NewLocation(configDir string) Location {
	return Location{
		Dir:  configDir,
		File: filepath.Join(configDir, FileName),
	}
}

// Exists reports whether the config file already exists.
func Exists(location Location) (bool, error) {
	_, err := os.Stat(location.File)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}

// Template returns the starter config file contents for the CLI.
func Template(appName string) string {
	return fmt.Sprintf(`# %s configuration
# Replace this scaffold with your real application settings.

log_level = "info"
output = "json"
`, appName)
}

// Init creates or overwrites the default config file.
func Init(location Location, appName string, force bool) error {
	if err := os.MkdirAll(location.Dir, 0o755); err != nil {
		return err
	}

	exists, err := Exists(location)
	if err != nil {
		return err
	}

	if exists && !force {
		return fs.ErrExist
	}

	return os.WriteFile(location.File, []byte(Template(appName)), 0o600)
}
