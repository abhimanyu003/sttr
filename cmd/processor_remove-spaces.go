// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/abhimanyu003/sttr/utils"
	"github.com/spf13/cobra"
)

var removeSpaces_flag_s string

func init() {
	removeSpacesCmd.Flags().StringVarP(&removeSpaces_flag_s, "separator", "s", "", "Separator to split spaces")
	rootCmd.AddCommand(removeSpacesCmd)
}

var removeSpacesCmd = &cobra.Command{
	Use:     "remove-spaces [string]",
	Short:   "Remove all spaces + new lines",
	Aliases: []string{"remove-space", "trim-spaces", "trim-space"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var in []byte
		var out string

		if len(args) == 0 {
			in = []byte(utils.ReadMultilineInput())
		} else {
			if fi, err := os.Stat(args[0]); err == nil && !fi.IsDir() {
				d, err := os.ReadFile(args[0])
				if err != nil {
					return err
				}
				in = d
			} else {
				in = []byte(args[0])
			}
		}

		flags := make([]processors.Flag, 0)
		p := processors.RemoveSpaces{}
		flags = append(flags, processors.Flag{Short: "s", Value: removeSpaces_flag_s})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
