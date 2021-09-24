package cmd

import (
	"github.com/abhimanyu003/sttr/ui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Use sttr in interactive mode",
	Long: `Launches a nice terminal UI where you
can explore the available processors`,
	Run: func(cmd *cobra.Command, args []string) {
		x := ui.Ui{}
		x.Render()
	},
}
