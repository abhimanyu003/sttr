package main

import (
	"github.com/abhimanyu003/sttr/cmd"

	"github.com/abhimanyu003/sttr/processors"

	"github.com/charmbracelet/bubbles/list"
)

// version specify version of application using ldflags
var version = "dev"

var items = []list.Item{
	item{title: "Base64 Encoding", args: []string{"b64-enc", "b64-encode", "base64-encode"}, desc: "Encode your text to Base64", processor: processors.Base64Encode},
	item{title: "Base64 Decode", args: []string{"b64-dec", "b64-decode", "base64-decode"}, desc: "Decode your Base64 text", processor: processors.Base64Decode},

	item{title: "URL Encode", args: []string{"url-enc", "url-encode"}, desc: "Encode URL entities", processor: processors.URLEncode},
	item{title: "URL Decode", args: []string{"url-dec", "url-encode"}, desc: "Decode URL entities", processor: processors.URLDecode},

	item{title: "ROT13 Encode", args: []string{"rot13", "rot13-encode"}, desc: "Encode your text to ROT13", processor: processors.ROT13Encode},

	item{title: "To Title Case", args: []string{"title-case"}, desc: "Convert your text to Title Case", processor: processors.StringToTitle},
	item{title: "To Lower Case", args: []string{"lower-case"}, desc: "Convert your text to lower case", processor: processors.StringToLower},
	item{title: "To Upper Case", args: []string{"upper-case"}, desc: "Convert your text to UPPER CASE", processor: processors.StringToUpper},
	item{title: "To Snake Case", args: []string{"snake-case"}, desc: "Convert your text to snake_case", processor: processors.StringToSnakeCase},
	item{title: "To Kebab Case", args: []string{"kebab-case"}, desc: "Convert your text to keybab-case", processor: processors.StringToKebab},
	item{title: "To Slug Case", args: []string{"slug-case"}, desc: "Convert your text to slug-case", processor: processors.StringToSlug},
	item{title: "To Camel Case", args: []string{"camel-case"}, desc: "Convert your text to CamelCase", processor: processors.StringToCamel},
	item{title: "Reverse String", args: []string{"reverse"}, desc: "Reverse String ( gnirtS esreveR )", processor: processors.StringReverse},

	item{
		title: "Count Number of Characters", args: []string{"count-chars"}, desc: "Find the length of your text (including spaces)",
		processor: processors.CountNumberCharacters,
	},
	item{
		title: "Count Number of Words", args: []string{"count-words"}, desc: "Count the number of words in your text",
		processor: processors.CountWords,
	},
	item{
		title: "Count Number of Lines", args: []string{"count-lines"}, desc: "Count the number of lines in your text",
		processor: processors.CountLines,
	},

	item{title: "MD5 Sum", args: []string{"md5"}, desc: "Get the MD5 checksum of your text", processor: processors.MD5Encode},
	item{title: "SHA1 Sum", args: []string{"sha1"}, desc: "Get the SHA1 checksum of your text", processor: processors.SHA1Encode},
	item{title: "SHA256 Sum", args: []string{"sha256"}, desc: "Get the SHA256 checksum of your text", processor: processors.SHA256Encode},
	item{title: "SHA512 Sum", args: []string{"sha512"}, desc: "Get the SHA512 checksum of your text", processor: processors.SHA512Encode},

	item{title: "Format JSON", args: []string{"json"}, desc: "Format your text as JSON", processor: processors.FormatJSON},
	item{title: "JSON To YAML", args: []string{"json-yaml"}, desc: "Convert JSON to YAML text", processor: processors.JSONToYAML},
	item{title: "YAML To JSON", args: []string{"yaml-json"}, desc: "Convert YAML to JSON text", processor: processors.YAMLToJSON},

	item{title: "Hex To RGB", args: []string{"hex-rgb"}, desc: "Convert a #Hex code to RGB", processor: processors.HexToRGB},

	item{title: "Sort Lines", args: []string{"sort-lines"}, desc: "Sort lines alphabetically", processor: processors.SortLines},
}

type item struct {
	title     string
	args      []string
	desc      string
	processor func(string) string
}

func main() {
	cmd.Version = version
	cmd.Execute()
}
