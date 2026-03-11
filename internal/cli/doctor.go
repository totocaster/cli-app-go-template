package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/toto/cli-go-template/internal/app"
	"github.com/toto/cli-go-template/internal/config"
	"github.com/toto/cli-go-template/internal/output"
)

type doctorResponse struct {
	App          string `json:"app"`
	Version      string `json:"version"`
	ConfigDir    string `json:"config_dir"`
	ConfigFile   string `json:"config_file"`
	ConfigExists bool   `json:"config_exists"`
	CurrentTime  string `json:"current_time"`
}

func newDoctorCmd(runtime app.Runtime, options *rootOptions) *cobra.Command {
	return &cobra.Command{
		Use:     "doctor",
		Short:   "Sanity-check the local runtime and config path conventions",
		Example: "  gocli doctor\n  gocli --text doctor",
		RunE: func(cmd *cobra.Command, args []string) error {
			configDir, err := runtime.ConfigDir()
			if err != nil {
				return err
			}

			location := config.NewLocation(configDir)
			_, err = os.Stat(location.File)
			configExists := err == nil
			if err != nil && !os.IsNotExist(err) {
				return err
			}

			response := doctorResponse{
				App:          runtime.Build.Binary,
				Version:      runtime.Build.Version,
				ConfigDir:    location.Dir,
				ConfigFile:   location.File,
				ConfigExists: configExists,
				CurrentTime:  runtime.Now().UTC().Format("2006-01-02T15:04:05Z07:00"),
			}

			if options.text {
				if err := output.Line(runtime.Stdout, "binary: %s", response.App); err != nil {
					return err
				}
				if err := output.Line(runtime.Stdout, "version: %s", response.Version); err != nil {
					return err
				}
				if err := output.Line(runtime.Stdout, "config_dir: %s", response.ConfigDir); err != nil {
					return err
				}
				if err := output.Line(runtime.Stdout, "config_file: %s", response.ConfigFile); err != nil {
					return err
				}
				if err := output.Line(runtime.Stdout, "config_exists: %t", response.ConfigExists); err != nil {
					return err
				}
				return output.Line(runtime.Stdout, "current_time: %s", response.CurrentTime)
			}

			return output.JSON(runtime.Stdout, response)
		},
	}
}
