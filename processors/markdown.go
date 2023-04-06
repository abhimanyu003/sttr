package processors

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
)

// Markdown convert markdown to HTML.
type Markdown struct{}

func (p Markdown) Name() string {
	return "markdown-html"
}

func (p Markdown) Alias() []string {
	return []string{"md-html"}
}

func (p Markdown) Transform(data []byte, _ ...Flag) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(data, &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (p Markdown) Flags() []Flag {
	return nil
}

func (p Markdown) Title() string {
	title := "Markdown to HTML"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Markdown) Description() string {
	return "Convert Markdown to HTML"
}

func (p Markdown) FilterValue() string {
	return p.Title()
}
