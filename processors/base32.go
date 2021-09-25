package processors

import (
	"encoding/base32"
)

// Base32Encoding encode string to base64.
type Base32Encoding struct{}

func (p Base32Encoding) Name() string {
	return "base32-encode"
}

func (p Base32Encoding) Alias() []string {
	return []string{"b32-enc", "b32-encode"}
}

func (p Base32Encoding) Transform(data string, _ ...Flag) (string, error) {
	return base32.StdEncoding.EncodeToString([]byte(data)), nil
}

func (p Base32Encoding) Flags() []Flag {
	return nil
}

func (p Base32Encoding) Title() string {
	return "Base32 Encoding"
}

func (p Base32Encoding) Description() string {
	return "Encode your text to Base32"
}

func (p Base32Encoding) FilterValue() string {
	return p.Title()
}

// Base32Decode decode string from base64 to plain text.
type Base32Decode struct{}

func (p Base32Decode) Name() string {
	return "base32-decode"
}

func (p Base32Decode) Alias() []string {
	return []string{"b32-dec", "b32-decode"}
}

func (p Base32Decode) Transform(data string, _ ...Flag) (string, error) {
	decodedString, err := base32.StdEncoding.DecodeString(data)
	return string(decodedString), err
}

func (p Base32Decode) Flags() []Flag {
	return nil
}

func (p Base32Decode) Title() string {
	return "Base32 Decode"
}

func (p Base32Decode) Description() string {
	return "Decode your base32 text"
}

func (p Base32Decode) FilterValue() string {
	return p.Title()
}
