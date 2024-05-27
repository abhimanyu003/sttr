package processors

import (
	"fmt"
	"strings"
)

// ROT13Encode converts string to ROT13 encoding.
type ROT13Encode struct{}

func (p ROT13Encode) Name() string {
	return "rot13-encode"
}

func (p ROT13Encode) Alias() []string {
	return []string{"rot13", "rot13-enc"}
}

func (p ROT13Encode) Transform(data []byte, _ ...Flag) (string, error) {
	return strings.Map(rot13, string(data)), nil
}

func (p ROT13Encode) Flags() []Flag {
	return nil
}

func (p ROT13Encode) Title() string {
	title := "ROT13 Encode"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p ROT13Encode) Description() string {
	return "Encode your text to ROT13"
}

func (p ROT13Encode) FilterValue() string {
	return p.Title()
}

// rot13 private helper function for converting rune into rot13.
func rot13(r rune) rune {
	if r >= 'a' && r <= 'z' {
		// Rotate lowercase letters 13 places.
		if r >= 'm' {
			return r - 13
		}

		return r + 13
	} else if r >= 'A' && r <= 'Z' {
		// Rotate uppercase letters 13 places.
		if r >= 'M' {
			return r - 13
		}

		return r + 13
	}
	// Do nothing.
	return r
}
