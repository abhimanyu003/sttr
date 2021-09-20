package processors

import (
	"net/url"
)

// URLEncode encode url string
func URLEncode(input string) string {
	return url.QueryEscape(input)
}

// URLDecode decode url string
func URLDecode(input string) string {
	res, _ := url.QueryUnescape(input)
	return res
}
