package processors

import (
	"encoding/base64"
)

// Base64Encode encode string to base64.
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode decode string from base64 to plain text.
func Base64Decode(data string) string {
	decodedString, _ := base64.StdEncoding.DecodeString(data)

	return string(decodedString)
}
