package processors

import (
	"fmt"
	"strings"
)

// ROT13 converts string with ROT13 cypher.
// https://en.wikipedia.org/wiki/ROT13
type ROT13 struct{}

func (p ROT13) Name() string {
	return "rot13"
}

func (p ROT13) Alias() []string {
	return []string{"rot13-encode", "rot13-enc"}
}

func (p ROT13) Transform(data []byte, _ ...Flag) (string, error) {
	return strings.Map(rot13, string(data)), nil
}

func (p ROT13) Flags() []Flag {
	return nil
}

func (p ROT13) Title() string {
	title := "ROT13 Letter Substitution"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p ROT13) Description() string {
	return "Cipher/Decipher your text with ROT13 letter substitution"
}

func (p ROT13) FilterValue() string {
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
