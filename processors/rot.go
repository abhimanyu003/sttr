package processors

import "strings"

// ROT13Encode convert string to ROT13 encoding.
func ROT13Encode(input string) string {
	return strings.Map(rot13, input)
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
		} else {
			return r + 13
		}
	}
	// Do nothing.
	return r
}
