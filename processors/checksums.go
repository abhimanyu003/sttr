package processors

import (
	"fmt"
	"hash/adler32"
	"hash/crc32"
)

// CRC32 generates CRC32 checksum
type CRC32 struct{}

func (p CRC32) Name() string {
	return "crc32"
}

func (p CRC32) Alias() []string {
	return []string{"crc32-sum", "crc32-checksum"}
}

func (p CRC32) Transform(data []byte, f ...Flag) (string, error) {
	var polynomial string = "ieee"
	for _, flag := range f {
		if flag.Short == "p" {
			if pol, ok := flag.Value.(string); ok {
				polynomial = pol
			}
		}
	}

	var table *crc32.Table
	switch polynomial {
	case "ieee":
		table = crc32.IEEETable
	case "castagnoli":
		table = crc32.MakeTable(crc32.Castagnoli)
	case "koopman":
		table = crc32.MakeTable(crc32.Koopman)
	default:
		return "", fmt.Errorf("unsupported polynomial: %s (supported: ieee, castagnoli, koopman)", polynomial)
	}

	checksum := crc32.Checksum(data, table)
	return fmt.Sprintf("%08x", checksum), nil
}

func (p CRC32) Flags() []Flag {
	return []Flag{
		{
			Name:  "polynomial",
			Short: "p",
			Desc:  "CRC32 polynomial (ieee, castagnoli, koopman)",
			Value: "ieee",
			Type:  FlagString,
		},
	}
}

func (p CRC32) Title() string {
	return fmt.Sprintf("CRC32 Checksum (%s)", p.Name())
}

func (p CRC32) Description() string {
	return "Get the CRC32 checksum of your text"
}

func (p CRC32) FilterValue() string {
	return p.Title()
}

// Adler32 generates Adler32 checksum
type Adler32 struct{}

func (p Adler32) Name() string {
	return "adler32"
}

func (p Adler32) Alias() []string {
	return []string{"adler32-sum", "adler32-checksum"}
}

func (p Adler32) Transform(data []byte, _ ...Flag) (string, error) {
	checksum := adler32.Checksum(data)
	return fmt.Sprintf("%08x", checksum), nil
}

func (p Adler32) Flags() []Flag {
	return nil
}

func (p Adler32) Title() string {
	return fmt.Sprintf("Adler32 Checksum (%s)", p.Name())
}

func (p Adler32) Description() string {
	return "Get the Adler32 checksum of your text"
}

func (p Adler32) FilterValue() string {
	return p.Title()
}
