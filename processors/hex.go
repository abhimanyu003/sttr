package processors

import (
	"encoding/hex"
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

// HexToRGB convert hex color code to R, G, B codes
// here we are using input library helper.
func HexToRGB(input string) string {
	c, _ := colorful.Hex(input)

	return fmt.Sprintf("%d, %d, %d", int(c.R*255), int(c.G*255), int(c.B*255))
}

// HexToString convert hex color code to R, G, B codes
// here we are using input library helper.
func HexToString(input string) string {
	output, _ := hex.DecodeString(input)

	return string(output)
}

// StringToHex convert hex color code to R, G, B codes
// here we are using input library helper.
func StringToHex(input string) string {
	return hex.EncodeToString([]byte(input))
}
