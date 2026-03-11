package main

import (
	"context"
	"fmt"
	"os"

	"github.com/toto/cli-go-template/internal/app"
	"github.com/toto/cli-go-template/internal/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	runtime := app.DefaultRuntime(app.BuildInfo{
		Binary:  "gocli",
		Version: version,
		Commit:  commit,
		Date:    date,
	})

	if err := cli.Execute(context.Background(), runtime); err != nil {
		if cliErr, ok := err.(interface{ ExitCode() int }); ok {
			os.Exit(cliErr.ExitCode())
		}

		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
