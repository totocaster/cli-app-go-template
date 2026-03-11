package cli

import (
	"bytes"
	"testing"
	"time"

	"github.com/toto/cli-go-template/internal/app"
)

type testRuntime struct {
	runtime app.Runtime
	stdout  *bytes.Buffer
	stderr  *bytes.Buffer
	dir     string
}

func newTestRuntime(t *testing.T) testRuntime {
	t.Helper()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	dir := t.TempDir()

	runtime := app.Runtime{
		Stdout: stdout,
		Stderr: stderr,
		Now: func() time.Time {
			return time.Date(2026, time.March, 11, 9, 30, 0, 0, time.UTC)
		},
		ConfigDir: func() (string, error) {
			return dir, nil
		},
		Build: app.BuildInfo{
			Binary:  "gocli",
			Version: "test",
			Commit:  "abc123",
			Date:    "2026-03-11T09:30:00Z",
		},
	}

	return testRuntime{
		runtime: runtime,
		stdout:  stdout,
		stderr:  stderr,
		dir:     dir,
	}
}
