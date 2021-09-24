package processors

import (
	"encoding/base64"
)

// Base64Encode encode string to base64.
type Base64Encode struct{}

func (p Base64Encode) Name() string {
	return "base64-encode"
}

func (p Base64Encode) Alias() []string {
	return []string{"b64-enc", "base64-encode"}
}

func (p Base64Encode) Transform(data string, _ ...Flag) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(data)), nil
}

func (p Base64Encode) Flags() []Flag {
	return nil
}

func (p Base64Encode) Title() string {
	return "Base64 Encoding"
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

func (p Base64Decode) Transform(data string, _ ...Flag) (string, error) {
	decodedString, err := base64.StdEncoding.DecodeString(data)
	return string(decodedString), err
}

func (p Base64Decode) Flags() []Flag {
	return nil
}

func (p Base64Decode) Title() string {
	return "Base64 Decode"
}

func (p Base64Decode) Description() string {
	return "Decode your base64 text"
}

func (p Base64Decode) FilterValue() string {
	return p.Title()
}
