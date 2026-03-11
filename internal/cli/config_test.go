package cli

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestConfigInitCreatesFile(t *testing.T) {
	t.Parallel()

	rt := newTestRuntime(t)
	cmd := newRootCmd(rt.runtime)
	cmd.SetArgs([]string{"config", "init"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute returned error: %v", err)
	}

	configPath := filepath.Join(rt.dir, "config.toml")
	if _, err := os.Stat(configPath); err != nil {
		t.Fatalf("expected config file to exist: %v", err)
	}
}

func TestConfigInitReturnsExitErrorWhenFileExists(t *testing.T) {
	t.Parallel()

	rt := newTestRuntime(t)
	cmd := newRootCmd(rt.runtime)
	cmd.SetArgs([]string{"config", "init"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("first Execute returned error: %v", err)
	}

	cmd = newRootCmd(rt.runtime)
	cmd.SetArgs([]string{"config", "init"})

	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected second Execute to fail")
	}

	if !errors.Is(err, fs.ErrExist) {
		t.Fatalf("expected fs.ErrExist, got %v", err)
	}

	exitErr, ok := err.(interface{ ExitCode() int })
	if !ok {
		t.Fatalf("expected exit error, got %T", err)
	}

	if exitErr.ExitCode() != 2 {
		t.Fatalf("unexpected exit code %d", exitErr.ExitCode())
	}
}
