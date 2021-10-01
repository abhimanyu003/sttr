package processors

import "encoding/hex"

type HexEncode struct{}

func (p HexEncode) Name() string {
	return "hex-encode"
}

func (p HexEncode) Alias() []string {
	return []string{"hex-enc", "hexadecimal-encode"}
}

func (p HexEncode) Transform(data []byte, _ ...Flag) (string, error) {
	return hex.EncodeToString(data), nil
}

func (p HexEncode) Flags() []Flag {
	return nil
}

func (p HexEncode) Title() string {
	return "Hex Encode"
}

func (p HexEncode) Description() string {
	return "Encode your text Hex"
}

func (p HexEncode) FilterValue() string {
	return p.Title()
}

type HexDecode struct{}

func (p HexDecode) Name() string {
	return "hex-decode"
}

func (p HexDecode) Alias() []string {
	return []string{"hex-dec", "hexadecimal-decode"}
}

func (p HexDecode) Transform(data []byte, _ ...Flag) (string, error) {
	output, err := hex.DecodeString(string(data))

	if err != nil {
		return "", err
	}
	return string(output), nil
}

func (p HexDecode) Flags() []Flag {
	return nil
}

func (p HexDecode) Title() string {
	return "Hex Decode"
}

func (p HexDecode) Description() string {
	return "Convert Hexadecimal to String"
}

func (p HexDecode) FilterValue() string {
	return p.Title()
}
