package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	rootCmd.AddCommand(docsCmd)
}

var docsCmd = &cobra.Command{
	Use:    "generate-docs",
	Short:  "Print the version of sttr",
	Hidden: true,
	Long:   `All software has a version (semantic at best). This is sttr's'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		linkHandler := func(name string) string {
			return fmt.Sprintf(`{{< relref "%s" >}}`, name)
		}

		filePrepender := func(filename string) string {
			name := filepath.Base(filename)
			base := strings.TrimSuffix(name, filepath.Ext(name))
			return fmt.Sprintf("---\ntitle: %s\n---\n", strings.Replace(base, "_", " ", -1))
		}

		dir := args[0]
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		rootCmd.DisableAutoGenTag = true
		return doc.GenMarkdownTreeCustom(rootCmd, dir, filePrepender, linkHandler)
	},
}
