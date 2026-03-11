package paths

import (
	"path/filepath"
	"testing"
)

func TestConfigDir(t *testing.T) {
	t.Parallel()

	dir, err := ConfigDir("gocli")
	if err != nil {
		t.Fatalf("ConfigDir returned error: %v", err)
	}

	if dir == "" {
		t.Fatal("ConfigDir returned an empty path")
	}

	if filepath.Base(dir) != "gocli" {
		t.Fatalf("expected config dir to end with gocli, got %q", dir)
	}
}
