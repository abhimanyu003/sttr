//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/abhimanyu003/sttr/utils"
	list2 "github.com/charmbracelet/bubbles/list"

	"github.com/abhimanyu003/sttr/processors"
)

type data struct {
	Name  string
	Camel string
	Desc  string
	SName string
	Alias []string
	Flags []processors.Flag
}

func main() {
	list := processors.List
	for _, item := range list {
		p, ok := item.(processors.Processor)
		if !ok {
			log.Printf("item is not a processor: %v", item)
			continue
		}
		i, ok := item.(list2.DefaultItem)
		if !ok {
			log.Printf("item is not a list.DefaultItem: %v", item)
			continue
		}

		d := data{
			Name:  p.Name(),
			Alias: p.Alias(),
			Camel: utils.ToLowerCamelCase([]byte(p.Name())),
			SName: fmt.Sprintf("%T", p),
			Desc:  i.Description(),
			Flags: p.Flags(),
		}
		if d.Name == "" {
			log.Print("processor has no name")
			continue
		}
		generate(d)
	}
}

func generate(d data) {
	file, err := os.Create("./cmd/processor_" + d.Name + ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	funcMap := template.FuncMap{
		"Lower":     strings.ToLower,
		"ListAlias": ListAlias,
	}

	tmpl, err := template.New("test").Funcs(funcMap).Parse(t)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(file, d)
	if err != nil {
		log.Fatal(err)
	}
}

func ListAlias(l []string) string {
	if len(l) == 0 {
		return ""
	}
	sb := strings.Builder{}
	for i, s := range l {
		sb.WriteString(fmt.Sprintf(`"%s"`, s))
		if i < len(l)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

const t = `// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)
{{- $camel := .Camel -}}

{{ with .Flags }}
{{- $len := len . -}}
{{- if eq $len 1 }}{{ range . }}

var {{ $camel }}_flag_{{ .Short }} {{ .Type.String | Lower }}{{ end }}{{ end }}
{{- if gt $len 1 }}

var (
{{- range . }}		
	{{ $camel }}_flag_{{ .Short }} {{ .Type.String | Lower }}{{ end }}
){{ end -}}
{{ end }}

func init() {
{{- range .Flags }}{{ if .Type.IsString }}
	{{ $camel }}Cmd.Flags().{{ .Type }}VarP(&{{ $camel }}_flag_{{ .Short }}, "{{ .Name }}", "{{ .Short }}", "{{ .Value }}", "{{ .Desc }}")
{{- else }}	
	{{ $camel }}Cmd.Flags().{{ .Type }}VarP(&{{ $camel }}_flag_{{ .Short }}, "{{ .Name }}", "{{ .Short }}", {{ .Value }}, "{{ .Desc }}")
{{- end }}	
{{- end }}
	rootCmd.AddCommand({{ .Camel }}Cmd)
}

var {{ .Camel }}Cmd = &cobra.Command{
	Use:     "{{ .Name }}",
	Short:   "{{ .Desc }}",
	Aliases: []string{ {{- .Alias | ListAlias -}} },
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
		p := {{ .SName }}{}
		{{- range .Flags }}
		flags = append(flags, processors.Flag{Short: "{{.Short}}", Value: {{ $camel }}_flag_{{ .Short }}})
		{{- end }}

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(os.Stdout, "%s\n", out)
		return err
	},
}
`
