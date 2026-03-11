package app

import (
	"io"
	"os"
	"time"

	"github.com/toto/cli-go-template/internal/paths"
)

// Runtime carries the process dependencies a command needs to execute.
type Runtime struct {
	Stdout    io.Writer
	Stderr    io.Writer
	Now       func() time.Time
	ConfigDir func() (string, error)
	Build     BuildInfo
}

// DefaultRuntime wires the real process dependencies for the CLI.
func DefaultRuntime(build BuildInfo) Runtime {
	return Runtime{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Now:    time.Now,
		ConfigDir: func() (string, error) {
			return paths.ConfigDir(build.Binary)
		},
		Build: build,
	}
}
