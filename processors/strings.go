package processors

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"regexp"
	"sort"
	"strings"
)

func StringToTitle(input string) string {
	return strings.Title(input)
}

func StringToLower(input string) string {
	return strings.ToLower(input)
}

func StringToUpper(input string) string {
	return strings.ToUpper(input)
}

func StringToSnakeCase(input string) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")

	return strcase.ToSnake(input)
}

func StringToKebab(input string) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")

	return strcase.ToKebab(input)
}

func StringToCamel(input string) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")

	return strcase.ToCamel(input)
}

func StringToSlug(input string) string {
	re := regexp.MustCompile("[^a-z0-9]+")

	return strings.Trim(re.ReplaceAllString(strings.ToLower(input), "-"), "-")
}

func CountNumberCharacters(input string) string {
	return fmt.Sprintf("%d", len(input))
}

func SortLines(input string) string {
	sorted := strings.Split(input, "\n")
	sort.Strings(sorted)

	return strings.Join(sorted, "\n")
}

func StringReverse(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return result
}
