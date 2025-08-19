package processors

import (
	"fmt"

	"github.com/cespare/xxhash/v2"
)

type XXHash64 struct{}

func (x XXHash64) Name() string {
	return "xxhash-64"
}

func (x XXHash64) Alias() []string {
	return []string{"xxhash-64"}
}

func (x XXHash64) Transform(data []byte, _ ...Flag) (string, error) {
	h := xxhash.New()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	s := h.Sum64()
	return fmt.Sprintf("%x", s), nil
}

func (x XXHash64) Flags() []Flag {
	return nil
}

func (x XXHash64) Title() string {
	title := "XXhash - 64"
	return fmt.Sprintf("%s (%s)", title, x.Name())
}

func (x XXHash64) Description() string {
	return "Get the XXHash64 checksum of your text"
}

func (x XXHash64) FilterValue() string {
	return x.Title()
}
