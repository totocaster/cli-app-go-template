package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toto/cli-go-template/internal/app"
)

func newCompletionCmd(runtime app.Runtime) *cobra.Command {
	cmd := &cobra.Command{
		Use:       "completion [bash|zsh|fish|powershell]",
		Short:     "Generate shell completion scripts",
		Long:      "Generate shell completion scripts so the installed binary behaves like a polished end-user CLI.",
		ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Example: "  gocli completion zsh > \"${fpath[1]}/_gocli\"\n" +
			"  gocli completion bash > ~/.local/share/bash-completion/completions/gocli\n" +
			"  gocli completion fish > ~/.config/fish/completions/gocli.fish",
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return cmd.Root().GenBashCompletionV2(runtime.Stdout, true)
			case "zsh":
				return cmd.Root().GenZshCompletion(runtime.Stdout)
			case "fish":
				return cmd.Root().GenFishCompletion(runtime.Stdout, true)
			case "powershell":
				return cmd.Root().GenPowerShellCompletionWithDesc(runtime.Stdout)
			default:
				return fmt.Errorf("unsupported shell %q", args[0])
			}
		},
	}

	return cmd
}
