package processors

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

// HexToRGB convert hex color code to R, G, B codes
// here we are using data library helper.
type HexToRGB struct{}

func (p HexToRGB) Name() string {
	return "hex-rgb"
}

func (p HexToRGB) Alias() []string {
	return nil
}

func (p HexToRGB) Transform(data []byte, _ ...Flag) (string, error) {
	c, err := colorful.Hex(string(data))
	return fmt.Sprintf("%d, %d, %d", int(c.R*255), int(c.G*255), int(c.B*255)), err
}

func (p HexToRGB) Flags() []Flag {
	return nil
}

func (p HexToRGB) Title() string {
	return "Hex To RGB"
}

func (p HexToRGB) Description() string {
	return "Convert a #hex-color code to RGB"
}

func (p HexToRGB) FilterValue() string {
	return p.Title()
}
