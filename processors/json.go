package processors

import (
	"encoding/json"
	"github.com/ghodss/yaml"
)

func FormatJSON(input string) string {
	var jsonIndent []byte

	var objmap map[string]*json.RawMessage

	_ = json.Unmarshal([]byte(input), &objmap)
	jsonIndent, _ = json.MarshalIndent(objmap, "", "  ")

	return string(jsonIndent)
}

func JSONToYAML(input string) string {
	y, _ := yaml.JSONToYAML([]byte(input))

	return string(y)
}

func YAMLToJSON(input string) string {
	y, _ := yaml.YAMLToJSON([]byte(input))

	return FormatJSON(string(y))
}
