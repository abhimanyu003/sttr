package processors

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

func HexToRGB(input string) string {
	c, _ := colorful.Hex(input)
	return fmt.Sprintf("%d, %d, %d", int(c.R*255), int(c.G*255), int(c.B*255))
}
