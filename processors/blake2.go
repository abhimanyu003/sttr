package processors

import (
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

// BLAKE2b generates BLAKE2b hash
type BLAKE2b struct{}

// Implement StreamingProcessor interface
func (p BLAKE2b) CanStream() bool {
	return true
}

func (p BLAKE2b) PreferStream() bool {
	return true
}

func (p BLAKE2b) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	var size uint = 64 // Default BLAKE2b size
	for _, flag := range opts {
		if flag.Short == "s" {
			if s, ok := flag.Value.(uint); ok {
				size = s
			}
		}
	}

	if size < 1 || size > 64 {
		return fmt.Errorf("BLAKE2b size must be between 1 and 64 bytes")
	}

	hasher, err := blake2b.New(int(size), nil)
	if err != nil {
		return err
	}

	_, err = io.Copy(hasher, reader)
	if err != nil {
		return err
	}

	result := hex.EncodeToString(hasher.Sum(nil))
	_, err = writer.Write([]byte(result))
	return err
}

func (p BLAKE2b) Name() string {
	return "blake2b"
}

func (p BLAKE2b) Alias() []string {
	return []string{"blake2b-hash", "blake2b-sum"}
}

func (p BLAKE2b) Transform(data []byte, f ...Flag) (string, error) {
	var size uint = 64 // Default BLAKE2b size
	for _, flag := range f {
		if flag.Short == "s" {
			if s, ok := flag.Value.(uint); ok {
				size = s
			}
		}
	}

	if size < 1 || size > 64 {
		return "", fmt.Errorf("BLAKE2b size must be between 1 and 64 bytes")
	}

	hasher, err := blake2b.New(int(size), nil)
	if err != nil {
		return "", err
	}

	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (p BLAKE2b) Flags() []Flag {
	return []Flag{
		{
			Name:  "size",
			Short: "s",
			Desc:  "Hash size in bytes (1-64)",
			Value: uint(64),
			Type:  FlagUint,
		},
	}
}

func (p BLAKE2b) Title() string {
	return fmt.Sprintf("BLAKE2b Hash (%s)", p.Name())
}

func (p BLAKE2b) Description() string {
	return "Get the BLAKE2b hash of your text"
}

func (p BLAKE2b) FilterValue() string {
	return p.Title()
}

// BLAKE2s generates BLAKE2s hash
type BLAKE2s struct{}

// Implement StreamingProcessor interface
func (p BLAKE2s) CanStream() bool {
	return true
}

func (p BLAKE2s) PreferStream() bool {
	return true
}

func (p BLAKE2s) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	hasher, err := blake2s.New256(nil)
	if err != nil {
		return err
	}

	_, err = io.Copy(hasher, reader)
	if err != nil {
		return err
	}

	result := hex.EncodeToString(hasher.Sum(nil))
	_, err = writer.Write([]byte(result))
	return err
}

func (p BLAKE2s) Name() string {
	return "blake2s"
}

func (p BLAKE2s) Alias() []string {
	return []string{"blake2s-hash", "blake2s-sum"}
}

func (p BLAKE2s) Transform(data []byte, _ ...Flag) (string, error) {
	hasher, err := blake2s.New256(nil)
	if err != nil {
		return "", err
	}

	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (p BLAKE2s) Flags() []Flag {
	return nil
}

func (p BLAKE2s) Title() string {
	return fmt.Sprintf("BLAKE2s Hash (%s)", p.Name())
}

func (p BLAKE2s) Description() string {
	return "Get the BLAKE2s hash of your text"
}

func (p BLAKE2s) FilterValue() string {
	return p.Title()
}
