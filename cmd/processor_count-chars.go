// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(countCharsCmd)
}

var countCharsCmd = &cobra.Command{
	Use:     "count-chars",
	Short:   "Find the length of your text (including spaces)",
	Aliases: []string{},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var in []byte
		var out string

		if len(args) == 0 {
			in, err = ioutil.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
		} else {
			if fi, err := os.Stat(args[0]); err == nil && !fi.IsDir() {
				d, err := ioutil.ReadFile(args[0])
				if err != nil {
					return err
				}
				in = d
			} else {
				in = []byte(args[0])
			}
		}

		flags := make([]processors.Flag, 0)
		p := processors.CountCharacters{}

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
