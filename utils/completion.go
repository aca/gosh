package utils

import (
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

const completionMsg = `To load completions:

Bash:

$ source <({{.Command}} completion bash)

# To load completions for each session, execute once:
Linux:
  $ {{.Command}} completion bash > /etc/bash_completion.d/{{.Command}}
MacOS:
  $ {{.Command}} completion bash > /usr/local/etc/bash_completion.d/{{.Command}}

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ {{.Command}} completion zsh > "${fpath[1]}/_{{.Command}}"

# You will need to start a new shell for this setup to take effect.

Fish:

$ {{.Command}} completion fish | source

# To load completions for each session, execute once:
$ {{.Command}} completion fish > ~/.config/fish/completions/{{.Command}}.fish
`

func NewCompletionCommand(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "completion [bash|zsh|fish|powershell]",
		Short:                 "Generate completion script",
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletion(os.Stdout)
			}
		},
	}

	t := template.Must(template.New("completion_msg").Parse(completionMsg))
	cmdName := struct {
		Command string
	}{
		Command: name,
	}

	sb := &strings.Builder{}
	t.Execute(sb, cmdName)
	cmd.Long = sb.String()

	return cmd
}
