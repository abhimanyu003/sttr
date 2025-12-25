package processors

import (
	"bytes"
	"fmt"

	"github.com/mdp/qrterminal/v3"
)

type QRCode struct{}

func (p QRCode) Name() string {
	return "qr"
}

func (p QRCode) Alias() []string {
	return []string{"qrcode", "qr-code"}
}

func (p QRCode) Transform(data []byte, f ...Flag) (string, error) {

	config := qrterminal.Config{
		Level:      qrterminal.M,
		HalfBlocks: true,
		QuietZone:  1,
	}

	for _, flag := range f {
		switch flag.Short {
		case "s":
			if size, ok := flag.Value.(uint); ok {
				if size >= 300 {
					config.QuietZone = 3
					config.Level = qrterminal.H
				} else if size >= 200 {
					config.QuietZone = 2
					config.Level = qrterminal.M
				} else {
					config.QuietZone = 1
					config.Level = qrterminal.L
				}
			}
		case "l":
			if level, ok := flag.Value.(string); ok {
				switch level {
				case "L", "low":
					config.Level = qrterminal.L
				case "M", "medium":
					config.Level = qrterminal.M
				case "H", "high":
					config.Level = qrterminal.H
				}
			}
		case "f":
			if full, ok := flag.Value.(bool); ok && full {
				config.HalfBlocks = false
			}
		}
	}

	var buf bytes.Buffer
	config.Writer = &buf

	qrterminal.GenerateWithConfig(string(data), config)

	return buf.String(), nil
}

func (p QRCode) Flags() []Flag {
	return []Flag{
		{
			Name:  "size",
			Short: "s",
			Desc:  "QR code size hint (affects error correction and quiet zone)",
			Value: uint(100),
			Type:  FlagUint,
		},
		{
			Name:  "level",
			Short: "l",
			Desc:  "Error correction level (L/low, M/medium, H/high)",
			Value: "H",
			Type:  FlagString,
		},
		{
			Name:  "full",
			Short: "f",
			Desc:  "Use full blocks instead of half blocks",
			Value: false,
			Type:  FlagBool,
		},
	}
}

func (p QRCode) Title() string {
	return fmt.Sprintf("QR Code (%s)", p.Name())
}

func (p QRCode) Description() string {
	return "Generate QR code in terminal"
}

func (p QRCode) FilterValue() string {
	return p.Title()
}

func (p QRCode) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{
		ChunkSize:    64 * 1024,
		BufferOutput: true,
		LineByLine:   false,
	}
}
