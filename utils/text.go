package utils

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

// ReadMultilineInput read multiple lines from stdin,
// to return the input, two empty line are expected
func ReadMultilineInput() string {
	str := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	empty := 0
	for {
		scanner.Scan()
		text := scanner.Text()
		str = append(str, text)
		if len(text) != 0 {
			empty = 0
		} else {
			empty++
			if empty == 2 {
				break
			}
		}
	}
	// Use collected inputs
	return strings.Join(str[:len(str)-1], "\n")
}

// TrimTrailingLinebreaks removes trailing linebreaks ('r', '\n', '\r\n\') from
// the input
func TrimTrailingLinebreaks(input string) string {
	buf := []byte(input)
	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i] != '\n' && buf[i] != '\r' {
			return string(buf[:i+1])
		}
	}
	return ""
}

func ToKebabCase(input string) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")
	return strcase.ToKebab(input)
}

func ToLowerCamelCase(input string) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")
	return strcase.ToLowerCamel(input)
}
