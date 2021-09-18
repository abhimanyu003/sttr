package processors

import (
	"net/url"
)

func URLEncode(input string) string {
	return url.QueryEscape(input)
}

func URLDecode(input string) string {
	res, _ := url.QueryUnescape(input)
	return res
}
