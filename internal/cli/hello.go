package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/toto/cli-go-template/internal/app"
	"github.com/toto/cli-go-template/internal/output"
)

type helloResponse struct {
	App       string `json:"app"`
	Name      string `json:"name"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func newHelloCmd(runtime app.Runtime, options *rootOptions) *cobra.Command {
	var uppercase bool

	cmd := &cobra.Command{
		Use:     "hello [name]",
		Short:   "Disposable example command to replace with your first real feature",
		Args:    cobra.MaximumNArgs(1),
		Example: "  gocli hello\n  gocli hello Toto\n  gocli --text hello Toto --uppercase",
		RunE: func(cmd *cobra.Command, args []string) error {
			name := "world"
			if len(args) == 1 {
				name = args[0]
			}

			message := fmt.Sprintf("hello, %s", name)
			if uppercase {
				message = strings.ToUpper(message)
			}

			response := helloResponse{
				App:       runtime.Build.Binary,
				Name:      name,
				Message:   message,
				Timestamp: runtime.Now().UTC().Format("2006-01-02T15:04:05Z07:00"),
			}

			if options.text {
				return output.Line(runtime.Stdout, response.Message)
			}

			return output.JSON(runtime.Stdout, response)
		},
	}

	cmd.Flags().BoolVar(&uppercase, "uppercase", false, "Render the greeting in uppercase")

	return cmd
}
