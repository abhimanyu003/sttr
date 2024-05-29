package processors

import (
	"fmt"
	"html"
)

// HTMLEncode escapes string to HTML
type HTMLEncode struct{}

func (p HTMLEncode) Name() string {
	return "html-encode"
}

func (p HTMLEncode) Alias() []string {
	return []string{"html-enc", "html-escape"}
}

func (p HTMLEncode) Transform(data []byte, _ ...Flag) (string, error) {
	return html.EscapeString(string(data)), nil
}

func (p HTMLEncode) Flags() []Flag {
	return nil
}

func (p HTMLEncode) Title() string {
	title := "HTML Encode"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p HTMLEncode) Description() string {
	return "Escape your HTML"
}

func (p HTMLEncode) FilterValue() string {
	return p.Title()
}

// HTMLDecode unescapes HTML to string
type HTMLDecode struct{}

func (p HTMLDecode) Name() string {
	return "html-decode"
}

func (p HTMLDecode) Alias() []string {
	return []string{"html-dec", "html-unescape"}
}

func (p HTMLDecode) Transform(data []byte, _ ...Flag) (string, error) {
	return html.UnescapeString(string(data)), nil
}

func (p HTMLDecode) Flags() []Flag {
	return nil
}

func (p HTMLDecode) Title() string {
	title := "HTML Decode"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p HTMLDecode) Description() string {
	return "Unescape your HTML"
}

func (p HTMLDecode) FilterValue() string {
	return p.Title()
}
