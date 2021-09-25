package processors

import (
	"encoding/json"

	"github.com/ghodss/yaml"
)

// FormatJSON format given string to a JSON with Indent.
type FormatJSON struct{}

func (p FormatJSON) Name() string {
	return "json"
}

func (p FormatJSON) Alias() []string {
	return nil
}

func (p FormatJSON) Transform(input string, f ...Flag) (string, error) {
	var objmap map[string]*json.RawMessage
	_ = json.Unmarshal([]byte(input), &objmap)

	var indent bool
	for _, flag := range f {
		if flag.Short == "i" {
			if b, ok := flag.Value.(bool); ok {
				indent = b
			}
		}
	}
	var newJSON []byte
	var err error
	if indent {
		newJSON, err = json.MarshalIndent(objmap, "", "  ")
	} else {
		newJSON, err = json.Marshal(objmap)
	}

	return string(newJSON), err
}

func (p FormatJSON) Flags() []Flag {
	return []Flag{
		{Name: "indent", Short: "i", Desc: "Indent the output (prettyprint)", Type: FlagBool, Value: false},
	}
}

func (p FormatJSON) Title() string {
	return "Format JSON"
}

func (p FormatJSON) Description() string {
	return "Format your text as JSON"
}

func (p FormatJSON) FilterValue() string {
	return p.Title()
}

// JSONToYAML convert JSON to YAML string.
type JSONToYAML struct{}

func (p JSONToYAML) Name() string {
	return "json-yaml"
}

func (p JSONToYAML) Alias() []string {
	return []string{"json-yml"}
}

func (p JSONToYAML) Transform(input string, _ ...Flag) (string, error) {
	y, _ := yaml.JSONToYAML([]byte(input))
	return string(y), nil
}

func (p JSONToYAML) Flags() []Flag {
	return nil
}

func (p JSONToYAML) Title() string {
	return "JSON To YAML"
}

func (p JSONToYAML) Description() string {
	return "Convert JSON to YAML text"
}

func (p JSONToYAML) FilterValue() string {
	return p.Title()
}

// YAMLToJSON convert YAML to JSON string with formatted output.
type YAMLToJSON struct{}

func (p YAMLToJSON) Name() string {
	return "yaml-json"
}

func (p YAMLToJSON) Alias() []string {
	return []string{"yml-json"}
}

func (p YAMLToJSON) Transform(input string, f ...Flag) (string, error) {
	y, _ := yaml.YAMLToJSON([]byte(input))
	j := FormatJSON{}
	return j.Transform(string(y), f...)
}

func (p YAMLToJSON) Flags() []Flag {
	return []Flag{
		{Name: "indent", Short: "i", Desc: "Indent the output (prettyprint)", Type: FlagBool, Value: false},
	}
}

func (p YAMLToJSON) Title() string {
	return "YAML To JSON"
}

func (p YAMLToJSON) Description() string {
	return "Convert YAML to JSON text"
}

func (p YAMLToJSON) FilterValue() string {
	return p.Title()
}
