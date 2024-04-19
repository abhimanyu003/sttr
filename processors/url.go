package processors

import (
	"fmt"
	"mvdan.cc/xurls/v2"
	"net/url"
	"strings"
)

// URLEncode encode url string.
type URLEncode struct{}

func (p URLEncode) Name() string {
	return "url-encode"
}

func (p URLEncode) Alias() []string {
	return []string{"url-enc"}
}

func (p URLEncode) Transform(data []byte, _ ...Flag) (string, error) {
	return url.QueryEscape(string(data)), nil
}

func (p URLEncode) Flags() []Flag {
	return nil
}

func (p URLEncode) Title() string {
	title := "URL Encode"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p URLEncode) Description() string {
	return "Encode URL entities"
}

func (p URLEncode) FilterValue() string {
	return p.Title()
}

// URLDecode decode url string.
type URLDecode struct{}

func (p URLDecode) Name() string {
	return "url-decode"
}

func (p URLDecode) Alias() []string {
	return []string{"url-dec"}
}

func (p URLDecode) Transform(data []byte, _ ...Flag) (string, error) {
	res, _ := url.QueryUnescape(string(data))
	return res, nil
}

func (p URLDecode) Flags() []Flag {
	return nil
}

func (p URLDecode) Title() string {
	title := "URL Decode"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p URLDecode) Description() string {
	return "Decode URL entities"
}

func (p URLDecode) FilterValue() string {
	return p.Title()
}

// ExtractURLs decode url string.
type ExtractURLs struct{}

func (p ExtractURLs) Name() string {
	return "extract-url"
}

func (p ExtractURLs) Alias() []string {
	return []string{"url-ext", "extract-urls", "ext-url"}
}

func (p ExtractURLs) Transform(data []byte, _ ...Flag) (string, error) {
	rxRelaxed := xurls.Relaxed()
	urls := rxRelaxed.FindAllString(string(data), -1)

	var output string

	for _, u := range urls {
		output = output + u + "\n"
	}

	output = strings.TrimSuffix(output, "\n")

	return output, nil
}

func (p ExtractURLs) Flags() []Flag {
	return nil
}

func (p ExtractURLs) Title() string {
	title := "Extract URLs"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p ExtractURLs) Description() string {
	return "Extract URLs from text"
}

func (p ExtractURLs) FilterValue() string {
	return p.Title()
}
