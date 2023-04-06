package processors

import (
	"encoding/base64"
	"fmt"
)

var base64RawFlag = Flag{
	Name:  "raw",
	Short: "r",
	Desc:  "unpadded base64 encoding",
	Value: false,
	Type:  FlagBool,
}

func checkBase64RawFlag(f []Flag) bool {
	raw := false
	for _, flag := range f {
		if flag.Name == "raw" || flag.Short == "r" {
			r, ok := flag.Value.(bool)
			if ok {
				raw = r
			}
		}
	}
	return raw
}

// Base64Encode encode string to base64.
type Base64Encode struct{}

func (p Base64Encode) Name() string {
	return "base64-encode"
}

func (p Base64Encode) Alias() []string {
	return []string{"b64-enc", "b64-encode"}
}

func (p Base64Encode) Transform(data []byte, f ...Flag) (string, error) {
	if checkBase64RawFlag(f) {
		return base64.RawStdEncoding.EncodeToString(data), nil
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func (p Base64Encode) Flags() []Flag {
	return []Flag{base64RawFlag}
}

func (p Base64Encode) Title() string {
	title := "Base64 Encoding"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Base64Encode) Description() string {
	return "Encode your text to Base64"
}

func (p Base64Encode) FilterValue() string {
	return p.Title()
}

// Base64Decode decode string from base64 to plain text.
type Base64Decode struct{}

func (p Base64Decode) Name() string {
	return "base64-decode"
}

func (p Base64Decode) Alias() []string {
	return []string{"b64-dec", "b64-decode"}
}

func (p Base64Decode) Transform(data []byte, f ...Flag) (string, error) {
	var decodedString []byte
	var err error
	if checkBase64RawFlag(f) {
		decodedString, err = base64.RawStdEncoding.DecodeString(string(data))
	} else {
		decodedString, err = base64.StdEncoding.DecodeString(string(data))
	}
	return string(decodedString), err
}

func (p Base64Decode) Flags() []Flag {
	return []Flag{base64RawFlag}
}

func (p Base64Decode) Title() string {
	title := "Base64 Decode"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Base64Decode) Description() string {
	return "Decode your base64 text"
}

func (p Base64Decode) FilterValue() string {
	return p.Title()
}
