package cli

import (
	"github.com/spf13/cobra"
	"github.com/toto/cli-go-template/internal/app"
	"github.com/toto/cli-go-template/internal/output"
)

type versionResponse struct {
	Binary  string `json:"binary"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func newVersionCmd(runtime app.Runtime, options *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print build version information",
		Aliases: []string{"ver"},
		RunE: func(cmd *cobra.Command, args []string) error {
			response := versionResponse{
				Binary:  runtime.Build.Binary,
				Version: runtime.Build.Version,
				Commit:  runtime.Build.Commit,
				Date:    runtime.Build.Date,
			}

			if options.text {
				if err := output.Line(runtime.Stdout, "%s %s", response.Binary, response.Version); err != nil {
					return err
				}
				if err := output.Line(runtime.Stdout, "commit: %s", response.Commit); err != nil {
					return err
				}
				return output.Line(runtime.Stdout, "built: %s", response.Date)
			}

			return output.JSON(runtime.Stdout, response)
		},
	}
}
