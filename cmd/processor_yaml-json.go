// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

var yamlJson_flag_i bool

func init() {	
	yamlJsonCmd.Flags().BoolVarP(&yamlJson_flag_i, "indent", "i", false, "Indent the output (prettyprint)")
	rootCmd.AddCommand(yamlJsonCmd)
}

var yamlJsonCmd = &cobra.Command{
	Use:     "yaml-json",
	Short:   "Convert YAML to JSON text",
	Aliases: []string{"yml-json"},
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
		p := processors.YAMLToJSON{}
		flags = append(flags, processors.Flag{Short: "i", Value: yamlJson_flag_i})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
