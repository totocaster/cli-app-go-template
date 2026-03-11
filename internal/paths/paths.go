package paths

import (
	"os"
	"path/filepath"
)

// ConfigDir returns the per-user config directory for the given application name.
func ConfigDir(appName string) (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(base, appName), nil
}
