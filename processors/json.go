package processors

import (
	"encoding/json"
	"github.com/ghodss/yaml"
)

// FormatJSON format given string to a JSON with Indent.
func FormatJSON(input string) string {
	var jsonIndent []byte

	var objmap map[string]*json.RawMessage

	_ = json.Unmarshal([]byte(input), &objmap)
	jsonIndent, _ = json.MarshalIndent(objmap, "", "  ")

	return string(jsonIndent)
}

// JSONToYAML convert JSON to YAML string.
func JSONToYAML(input string) string {
	y, _ := yaml.JSONToYAML([]byte(input))

	return string(y)
}

// YAMLToJSON convert YAML to JSON string with formatted output.
func YAMLToJSON(input string) string {
	y, _ := yaml.YAMLToJSON([]byte(input))

	return FormatJSON(string(y))
}
