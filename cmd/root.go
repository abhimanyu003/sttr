package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sttr",
	Short: "sttr is a fast and flexible string/text converter",
	Long: ` 
                Complete documentation is available at https://github.com/abhimanyu003/sttr`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var processor string

func init() {
	rootCmd.Flags().StringVarP(&processor, "processor", "p", "", "processor to use")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
