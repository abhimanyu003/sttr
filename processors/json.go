package processors

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/vmihailenco/msgpack/v5"
)

// FormatJSON format given string to a JSON with Indent.
type FormatJSON struct{}

func (p FormatJSON) Name() string {
	return "json"
}

func (p FormatJSON) Alias() []string {
	return nil
}

func (p FormatJSON) Transform(data []byte, f ...Flag) (string, error) {
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objmap)
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
	return "JSON To YAML"
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

	var rawData interface{}

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
	return "JSON To MSGPACK"
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

	var rawData interface{}

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
	return "MSGPACK To JSON"
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
	return "YAML To JSON"
}

func (p YAMLToJSON) Description() string {
	return "Convert YAML to JSON text"
}

func (p YAMLToJSON) FilterValue() string {
	return p.Title()
}
