package processors

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/vmihailenco/msgpack/v5"
	"gitlab.com/abhimanyusharma003/go-ordered-json"
)

// FormatJSON format given string to a JSON with Indent.
type FormatJSON struct{}

func (p FormatJSON) Name() string {
	return "json"
}

func (p FormatJSON) Alias() []string {
	return nil
}

// unmarshalJson converts given bytes to json.RawMessage
// it checks if input is of type array or non array.
func unmarshalJSON(data []byte) (any, error) {
	nonArray := ordered.NewOrderedMap()
	arrayBased := make([]ordered.OrderedMap, 0)
	err := json.Unmarshal(data, &nonArray)
	if err == nil {
		return nonArray, nil
	}

	err = json.Unmarshal(data, &arrayBased)
	if err == nil {
		return arrayBased, nil
	}

	return nil, err
}

func (p FormatJSON) Transform(data []byte, f ...Flag) (string, error) {
	objmap, err := unmarshalJSON(data)
	if err != nil {
		return "", err
	}
	var indent bool
	for _, flag := range f {
		if flag.Short == "i" {
			if b, ok := flag.Value.(bool); ok {
				indent = b
			}
		}
	}
	var newJSON []byte
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
	title := "Format JSON"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p FormatJSON) Description() string {
	return "Format your text as JSON ( json decode )"
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

func (p JSONToYAML) Transform(data []byte, _ ...Flag) (string, error) {
	y, err := yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(y), nil
}

func (p JSONToYAML) Flags() []Flag {
	return nil
}

func (p JSONToYAML) Title() string {
	title := "JSON To YAML"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p JSONToYAML) Description() string {
	return "Convert JSON to YAML text"
}

func (p JSONToYAML) FilterValue() string {
	return p.Title()
}

// JSONToMSGPACK convert JSON to MSGPACK string.
type JSONToMSGPACK struct{}

func (p JSONToMSGPACK) Name() string {
	return "json-msgpack"
}

func (p JSONToMSGPACK) Alias() []string {
	return []string{}
}

func (p JSONToMSGPACK) Transform(data []byte, _ ...Flag) (string, error) {
	var rawData any

	err := json.Unmarshal(data, &rawData)

	if err != nil {
		return "", err
	}

	m, err := msgpack.Marshal(rawData)
	if err != nil {
		return "", err
	}
	return string(m), nil
}

func (p JSONToMSGPACK) Flags() []Flag {
	return nil
}

func (p JSONToMSGPACK) Title() string {
	title := "JSON To MSGPACK"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p JSONToMSGPACK) Description() string {
	return "Convert JSON to MSGPACK text"
}

func (p JSONToMSGPACK) FilterValue() string {
	return p.Title()
}

// MSGPACKToJSON convert MSGPACK to JSON string.
type MSGPACKToJSON struct{}

func (p MSGPACKToJSON) Name() string {
	return "msgpack-json"
}

func (p MSGPACKToJSON) Alias() []string {
	return []string{}
}

func (p MSGPACKToJSON) Transform(data []byte, _ ...Flag) (string, error) {
	var rawData any

	err := msgpack.Unmarshal(data, &rawData)

	if err != nil {
		return "", err
	}

	m, err := json.Marshal(rawData)
	if err != nil {
		return "", err
	}
	return string(m), nil
}

func (p MSGPACKToJSON) Flags() []Flag {
	return nil
}

func (p MSGPACKToJSON) Title() string {
	title := "MSGPACK To JSON"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p MSGPACKToJSON) Description() string {
	return "Convert MSGPACK to JSON text"
}

func (p MSGPACKToJSON) FilterValue() string {
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

func (p YAMLToJSON) Transform(data []byte, f ...Flag) (string, error) {
	y, err := yaml.YAMLToJSON(data)
	if err != nil {
		return "", err
	}
	j := FormatJSON{}
	return j.Transform(y, f...)
}

func (p YAMLToJSON) Flags() []Flag {
	return []Flag{
		{Name: "indent", Short: "i", Desc: "Indent the output (prettyprint)", Type: FlagBool, Value: false},
	}
}

func (p YAMLToJSON) Title() string {
	title := "YAML To JSON"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p YAMLToJSON) Description() string {
	return "Convert YAML to JSON text"
}

func (p YAMLToJSON) FilterValue() string {
	return p.Title()
}

// JSONUnescape unescape given string to a JSON with Indent.
type JSONUnescape struct{}

func (p JSONUnescape) Name() string {
	return "json-unescape"
}

func (p JSONUnescape) Alias() []string {
	return []string{"json-unesc"}
}

func (p JSONUnescape) Transform(data []byte, f ...Flag) (string, error) {
	s, err := strconv.Unquote(`"` + strings.Trim(string(data), " ") + `"`)
	if err != nil {
		return "", err
	}
	objmap, err := unmarshalJSON([]byte(s))
	if err != nil {
		return "", err
	}

	var indent bool
	for _, flag := range f {
		if flag.Short == "i" {
			if b, ok := flag.Value.(bool); ok {
				indent = b
			}
		}
	}
	var newJSON []byte
	if indent {
		newJSON, err = json.MarshalIndent(objmap, "", "  ")
	} else {
		newJSON, err = json.Marshal(objmap)
	}

	return string(newJSON), err
}

func (p JSONUnescape) Flags() []Flag {
	return []Flag{
		{Name: "indent", Short: "i", Desc: "Indent the output (prettyprint)", Type: FlagBool, Value: false},
	}
}

func (p JSONUnescape) Title() string {
	title := "JSON Unescape"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p JSONUnescape) Description() string {
	return "JSON Unescape"
}

func (p JSONUnescape) FilterValue() string {
	return p.Title()
}

// JSONEscape unescape given string to a JSON with Indent.
type JSONEscape struct{}

func (p JSONEscape) Name() string {
	return "json-escape"
}

func (p JSONEscape) Alias() []string {
	return []string{"json-esc"}
}

func (p JSONEscape) Transform(data []byte, f ...Flag) (string, error) {
	objmap, err := unmarshalJSON(data)
	if err != nil {
		return "", err
	}

	newJSON, err := json.Marshal(objmap)
	if err != nil {
		return "", err
	}

	output := strconv.Quote(string(newJSON))
	output = strings.TrimLeft(output, `"`)
	output = strings.TrimRight(output, `"`)

	return output, err
}

func (p JSONEscape) Flags() []Flag {
	return nil
}

func (p JSONEscape) Title() string {
	title := "JSON Escape"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p JSONEscape) Description() string {
	return "JSON Escape"
}

func (p JSONEscape) FilterValue() string {
	return p.Title()
}
