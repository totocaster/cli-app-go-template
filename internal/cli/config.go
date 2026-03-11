package cli

import (
	"errors"
	"io/fs"

	"github.com/spf13/cobra"
	"github.com/toto/cli-go-template/internal/app"
	"github.com/toto/cli-go-template/internal/config"
	"github.com/toto/cli-go-template/internal/output"
)

type configPathResponse struct {
	Dir  string `json:"dir"`
	File string `json:"file"`
}

type configInitResponse struct {
	Dir     string `json:"dir"`
	File    string `json:"file"`
	Created bool   `json:"created"`
}

func newConfigCmd(runtime app.Runtime, options *rootOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Inspect or initialize the config scaffold",
	}

	cmd.AddCommand(
		newConfigPathCmd(runtime, options),
		newConfigTemplateCmd(runtime),
		newConfigInitCmd(runtime, options),
	)

	return cmd
}

func newConfigPathCmd(runtime app.Runtime, options *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use:     "path",
		Short:   "Print the default config directory and file path",
		Example: "  gocli config path\n  gocli --text config path",
		RunE: func(cmd *cobra.Command, args []string) error {
			location, err := configLocation(runtime)
			if err != nil {
				return err
			}

			response := configPathResponse{
				Dir:  location.Dir,
				File: location.File,
			}

			if options.text {
				return output.Line(runtime.Stdout, "%s", response.File)
			}

			return output.JSON(runtime.Stdout, response)
		},
	}
}

func newConfigTemplateCmd(runtime app.Runtime) *cobra.Command {
	return &cobra.Command{
		Use:     "template",
		Short:   "Print the starter config template",
		Example: "  gocli config template",
		RunE: func(cmd *cobra.Command, args []string) error {
			return output.Raw(runtime.Stdout, config.Template(runtime.Build.Binary))
		},
	}
}

func newConfigInitCmd(runtime app.Runtime, options *rootOptions) *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Create the default config file if it does not exist",
		Example: "  gocli config init\n  gocli --text config init --force",
		RunE: func(cmd *cobra.Command, args []string) error {
			location, err := configLocation(runtime)
			if err != nil {
				return err
			}

			err = config.Init(location, runtime.Build.Binary, force)
			if err != nil {
				if errors.Is(err, fs.ErrExist) {
					return newExitError(2, err)
				}
				return err
			}

			response := configInitResponse{
				Dir:     location.Dir,
				File:    location.File,
				Created: true,
			}

			if options.text {
				return output.Line(runtime.Stdout, "created %s", response.File)
			}

			return output.JSON(runtime.Stdout, response)
		},
	}

	cmd.Flags().BoolVar(&force, "force", false, "Overwrite the config file if it already exists")

	return cmd
}

func configLocation(runtime app.Runtime) (config.Location, error) {
	configDir, err := runtime.ConfigDir()
	if err != nil {
		return config.Location{}, err
	}

	return config.NewLocation(configDir), nil
}
