package processors

import (
	"html"
)

type HTMLEncode struct{}

func (p HTMLEncode) Name() string {
	return "html-encode"
}

func (p HTMLEncode) Alias() []string {
	return []string{"html-enc", "html-escape"}
}

func (p HTMLEncode) Transform(input string, _ ...Flag) (string, error) {
	return html.EscapeString(input), nil
}

func (p HTMLEncode) Flags() []Flag {
	return nil
}

func (p HTMLEncode) Title() string {
	return "HTML Encode"
}

func (p HTMLEncode) Description() string {
	return "Escape your HTML"
}

func (p HTMLEncode) FilterValue() string {
	return p.Title()
}

type HTMLDecode struct{}

func (p HTMLDecode) Name() string {
	return "html-decode"
}

func (p HTMLDecode) Alias() []string {
	return []string{"html-dec", "html-unescape"}
}

func (p HTMLDecode) Transform(input string, _ ...Flag) (string, error) {
	return html.UnescapeString(input), nil
}

func (p HTMLDecode) Flags() []Flag {
	return nil
}

func (p HTMLDecode) Title() string {
	return "HTML Decode"
}

func (p HTMLDecode) Description() string {
	return "Unescape your HTML"
}

func (p HTMLDecode) FilterValue() string {
	return p.Title()
}
