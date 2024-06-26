// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/abhimanyu003/sttr/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sha224Cmd)
}

var sha224Cmd = &cobra.Command{
	Use:     "sha224 [string]",
	Short:   "Get the SHA-224 checksum of your text",
	Aliases: []string{"sha224-sum"},
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
		p := processors.SHA224{}

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
