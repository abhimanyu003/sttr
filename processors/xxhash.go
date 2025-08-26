package processors

import (
	"fmt"

	"github.com/cespare/xxhash/v2"
)

type XXH64 struct{}

func (x XXH64) Name() string {
	return "xxh-64"
}

func (x XXH64) Alias() []string {
	return []string{""}
}

func (x XXH64) Transform(data []byte, _ ...Flag) (string, error) {
	h := xxhash.New()
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
