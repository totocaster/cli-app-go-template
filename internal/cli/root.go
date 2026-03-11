package cli

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/toto/cli-go-template/internal/app"
)

type rootOptions struct {
	text bool
}

// Execute builds the root command and runs it with the provided runtime.
func Execute(ctx context.Context, runtime app.Runtime) error {
	cmd := newRootCmd(runtime)
	cmd.SetContext(ctx)
	return cmd.ExecuteContext(ctx)
}

func newRootCmd(runtime app.Runtime) *cobra.Command {
	options := &rootOptions{}

	cmd := &cobra.Command{
		Use:           runtime.Build.Binary,
		Short:         "Starting point for a production-ready Go CLI",
		Long:          "A template-friendly Go CLI scaffold with modular Cobra commands, tests, CI, release automation, and documentation conventions.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.SetOut(runtime.Stdout)
	cmd.SetErr(runtime.Stderr)
	cmd.PersistentFlags().BoolVar(&options.text, "text", false, "Render human-readable output instead of JSON")

	cmd.AddCommand(
		newHelloCmd(runtime, options),
		newDoctorCmd(runtime, options),
		newConfigCmd(runtime, options),
		newCompletionCmd(runtime),
		newVersionCmd(runtime, options),
	)

	return cmd
}
