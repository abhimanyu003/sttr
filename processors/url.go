package processors

import (
	"net/url"
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
	return "URL Encode"
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
	return "URL Decode"
}

func (p URLDecode) Description() string {
	return "Decode URL entities"
}

func (p URLDecode) FilterValue() string {
	return p.Title()
}
