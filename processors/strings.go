package processors

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"regexp"
	"sort"
	"strings"
)

// StringToTitle convert sting to title case.
// Example: "this is string" to "This is String".
func StringToTitle(input string) string {
	return strings.Title(input)
}

// StringToLower convert sting to lower case.
// Example: "THIS IS STRING" to "this is string".
func StringToLower(input string) string {
	return strings.ToLower(input)
}

// StringToUpper convert sting to lower case.
// Example: "this is string" to "THIS IS STRING".
func StringToUpper(input string) string {
	return strings.ToUpper(input)
}

// StringToSnakeCase convert sting to snake_case.
// Example: "this is string" to "this_is_string".
func StringToSnakeCase(input string) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")

	return strcase.ToSnake(input)
}

// StringToKebab convert sting to kebab-case.
// Example: "this is string" to "this-is-string".
func StringToKebab(input string) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")

	return strcase.ToKebab(input)
}

// StringToCamel convert sting to CamelCase.
// Example: "this is string" to "ThisIsString".
func StringToCamel(input string) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")

	return strcase.ToCamel(input)
}

// StringToSlug convert sting to StringToSlug. It's similar to Kebab case but URL Friendly.
// Example: "this is string" to "this-is-string".
func StringToSlug(input string) string {
	re := regexp.MustCompile("[^a-z0-9]+")

	return strings.Trim(re.ReplaceAllString(strings.ToLower(input), "-"), "-")
}

// CountNumberCharacters count number of Characters including spaces.
func CountNumberCharacters(input string) string {
	return fmt.Sprintf("%d", len(input))
}

// SortLines sort given list, it's not a natural sort.
func SortLines(input string) string {
	sorted := strings.Split(input, "\n")
	sort.Strings(sorted)

	return strings.Join(sorted, "\n")
}

// StringReverse revers a given string
// Example: "test" to "tset"
func StringReverse(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return result
}
