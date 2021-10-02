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
	return strings.Join(str[:len(str)-2], "\n")
}

func ToKebabCase(input []byte) string {
	str := regexp.MustCompile(`\s+`).ReplaceAllString(string(input), " ")
	return strcase.ToKebab(str)
}

func ToLowerCamelCase(input []byte) string {
	str := regexp.MustCompile(`\s+`).ReplaceAllString(string(input), " ")
	return strcase.ToLowerCamel(str)
}
