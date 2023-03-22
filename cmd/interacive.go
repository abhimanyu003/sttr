package cmd

import (
	"github.com/abhimanyu003/sttr/ui"
	"github.com/spf13/cobra"
	"io"
)

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Use sttr in interactive mode",
	Long: `Launches a nice terminal UI where you
can explore the available processors interactively`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		in := ""

		if len(args) == 0 {
			all, err := io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
			in = string(all)
		} else {
			in = args[0]
		}

		x := ui.New(in)
		x.Render()
		return err
	},
}
