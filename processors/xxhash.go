package processors

import (
	"fmt"

	"github.com/harsh16coder/xxhash"
)

// XXH64 encodes string to XXH64
type XXH64 struct{}

func (x XXH64) Name() string {
	return "xxh-64"
}

func (x XXH64) Alias() []string {
	return []string{"xxh64", "xxhash64", "xxhash-64"}
}

func (x XXH64) Transform(data []byte, _ ...Flag) (string, error) {
	h := xxhash.New64()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	s := h.Sum64()
	return fmt.Sprintf("%016x", s), nil
}

func (x XXH64) Flags() []Flag {
	return nil
}

func (x XXH64) Title() string {
	title := "xxHash - XXH64"
	return fmt.Sprintf("%s (%s)", title, x.Name())
}

func (x XXH64) Description() string {
	return "Get the XXH64 checksum of your text"
}

func (x XXH64) FilterValue() string {
	return x.Title()
}

// XX32 encodes string to XXH32
type XXH32 struct{}

func (x XXH32) Name() string {
	return "xxh-32"
}

func (x XXH32) Alias() []string {
	return []string{"xxh32", "xxhash32", "xxhash-32"}
}

func (x XXH32) Transform(data []byte, _ ...Flag) (string, error) {
	h := xxhash.New32()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	s := h.Sum32()
	return fmt.Sprintf("%08x", s), nil
}

func (x XXH32) Flags() []Flag {
	return nil
}

func (x XXH32) Title() string {
	title := "xxHash - XXH32"
	return fmt.Sprintf("%s (%s)", title, x.Name())
}

func (x XXH32) Description() string {
	return "Get the XXH32 checksum of your text"
}

func (x XXH32) FilterValue() string {
	return x.Title()
}

// XX128 encodes string to XXH32
type XXH128 struct{}

func (x XXH128) Name() string {
	return "xxh-128"
}

func (x XXH128) Alias() []string {
	return []string{"xxh128", "xxhash128", "xxhash-128"}
}

func (x XXH128) Transform(data []byte, _ ...Flag) (string, error) {
	h := xxhash.New128()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	s := h.Sum128()
	return fmt.Sprintf("%016x%016x", s.Hi, s.Lo), nil
}

func (x XXH128) Flags() []Flag {
	return nil
}

func (x XXH128) Title() string {
	title := "xxHash - XXH128"
	return fmt.Sprintf("%s (%s)", title, x.Name())
}

func (x XXH128) Description() string {
	return "Get the XXH128 checksum of your text"
}

func (x XXH128) FilterValue() string {
	return x.Title()
}
