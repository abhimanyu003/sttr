// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

var base64UrlDecode_flag_r bool

func init() {
	base64UrlDecodeCmd.Flags().BoolVarP(&base64UrlDecode_flag_r, "raw", "r", false, "unpadded base64 encoding")
	rootCmd.AddCommand(base64UrlDecodeCmd)
}

var base64UrlDecodeCmd = &cobra.Command{
	Use:     "base64url-decode",
	Short:   "Decode your base64 text with URL Safe",
	Aliases: []string{"b64url-dec", "b64url-decode"},
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
		p := processors.Base64URLDecode{}
		flags = append(flags, processors.Flag{Short: "r", Value: base64UrlDecode_flag_r})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(os.Stdout, "%s\n", out)
		return err
	},
}
