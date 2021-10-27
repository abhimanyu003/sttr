package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(sttr completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ sttr completion bash > /etc/bash_completion.d/sttr
  # macOS:
  $ sttr completion bash > /usr/local/etc/bash_completion.d/sttr

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  # Generate a _sttr completion script and put it somewhere in your $fpath
  $ sttr completion zsh > /usr/local/share/zsh/site-functions/_sttr

  # You will need to start a new shell for this setup to take effect.

fish:

  $ sttr completion fish | source

  # To load completions for each session, execute once:
  $ sttr completion fish > ~/.config/fish/completions/sttr.fish

PowerShell:

  PS> sttr completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> sttr completion powershell > sttr.ps1
  # and source this file from your PowerShell profile.
`,
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
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}
