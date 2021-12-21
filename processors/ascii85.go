package processors

import (
	"bytes"
	"encoding/ascii85"
	"io"
)

// Ascii85Encoding encode string to base64.
type Ascii85Encoding struct{}

func (p Ascii85Encoding) Name() string {
	return "ascii85-encode"
}

func (p Ascii85Encoding) Alias() []string {
	return []string{"ascii85-encoding", "base85-encode", "b85-encode"}
}

func (p Ascii85Encoding) Transform(data []byte, _ ...Flag) (string, error) {
	buf := &bytes.Buffer{}
	encoder := ascii85.NewEncoder(buf)
	_, err := encoder.Write(data)
	if err != nil {
		return "", err
	}
	if err := encoder.Close(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (p Ascii85Encoding) Flags() []Flag {
	return nil
}

func (p Ascii85Encoding) Title() string {
	return "Ascii85 / Base85 Encoding"
}

func (p Ascii85Encoding) Description() string {
	return "Encode your text to Ascii85 ( Base85 )"
}

func (p Ascii85Encoding) FilterValue() string {
	return p.Title()
}

// Ascii85Decoding encode string to Ascii aka base85.
type Ascii85Decoding struct{}

func (p Ascii85Decoding) Name() string {
	return "ascii85-decode"
}

func (p Ascii85Decoding) Alias() []string {
	return []string{"ascii85-decoding", "base85-decode", "b85-decode"}
}

func (p Ascii85Decoding) Transform(data []byte, _ ...Flag) (string, error) {
	decoder := ascii85.NewDecoder(bytes.NewReader(data))
	buf, err := io.ReadAll(decoder)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (p Ascii85Decoding) Flags() []Flag {
	return nil
}

func (p Ascii85Decoding) Title() string {
	return "Ascii85 / Base85 Decoding"
}

func (p Ascii85Decoding) Description() string {
	return "Decode your text to Ascii85 ( Base85 ) text"
}

func (p Ascii85Decoding) FilterValue() string {
	return p.Title()
}
