package processors

import (
	"encoding/base64"
)

// Base64URLEncode encode string to base64.
type Base64URLEncode struct{}

func (p Base64URLEncode) Name() string {
	return "base64url-encode"
}

func (p Base64URLEncode) Alias() []string {
	return []string{"b64url-enc", "b64url-encode"}
}

func (p Base64URLEncode) Transform(data []byte, f ...Flag) (string, error) {
	raw := checkBase64RawFlag(f)

	if raw {
		return base64.RawURLEncoding.EncodeToString(data), nil
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

func (p Base64URLEncode) Flags() []Flag {
	return []Flag{base64RawFlag}
}

func (p Base64URLEncode) Title() string {
	return "Base64URL Encoding"
}

func (p Base64URLEncode) Description() string {
	return "Encode your text to Base64 with URL Safe"
}

func (p Base64URLEncode) FilterValue() string {
	return p.Title()
}

// Base64URLDecode decode string from base64 to plain text.
type Base64URLDecode struct{}

func (p Base64URLDecode) Name() string {
	return "base64url-decode"
}

func (p Base64URLDecode) Alias() []string {
	return []string{"b64url-dec", "b64url-decode"}
}

func (p Base64URLDecode) Transform(data []byte, f ...Flag) (string, error) {
	raw := checkBase64RawFlag(f)
	var decodedString []byte
	var err error
	if raw {
		decodedString, err = base64.RawURLEncoding.DecodeString(string(data))
	} else {
		decodedString, err = base64.URLEncoding.DecodeString(string(data))
	}
	return string(decodedString), err
}

func (p Base64URLDecode) Flags() []Flag {
	return []Flag{base64RawFlag}
}

func (p Base64URLDecode) Title() string {
	return "Base64URL Decode"
}

func (p Base64URLDecode) Description() string {
	return "Decode your base64 text with URL Safe"
}

func (p Base64URLDecode) FilterValue() string {
	return p.Title()
}
