// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

var extractEmails_flag_s string

func init() {
	extractEmailsCmd.Flags().StringVarP(&extractEmails_flag_s, "separator", "s", "", "Separator to split multiple emails")
	rootCmd.AddCommand(extractEmailsCmd)
}

var extractEmailsCmd = &cobra.Command{
	Use:     "extract-emails",
	Short:   "Extract emails from given text",
	Aliases: []string{"find-emails", "find-email", "extract-email"},
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
		p := processors.ExtractEmails{}
		flags = append(flags, processors.Flag{Short: "s", Value: extractEmails_flag_s})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(os.Stdout, "%s\n", out)
		return err
	},
}
