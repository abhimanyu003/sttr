package processors

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
)

// Lower converts a string to lower case.
// Example: "THIS IS STRING" to "this is string".
type Lower struct{}

func (p Lower) Name() string {
	return "lower"
}

func (p Lower) Transform(input string, _ ...Option) string {
	return strings.ToLower(input)
}

func (p Lower) Title() string {
	return "To Lower case"
}

func (p Lower) Description() string {
	return "Transform your text to lower case"
}

func (p Lower) FilterValue() string {
	return p.Title()
}

// Upper convert string to upper case.
// Example: "this is string" to "THIS IS STRING".
type Upper struct{}

func (p Upper) Name() string {
	return "upper"
}

func (p Upper) Transform(input string, _ ...Option) string {
	return strings.ToUpper(input)
}

func (p Upper) Title() string {
	return "To Upper case"
}

func (p Upper) Description() string {
	return "Transform your text to UPPER CASE"
}

func (p Upper) FilterValue() string {
	return p.Title()
}

// Title convert string to title case.
// Example: "this is string" to "This Is String".
type Title struct{}

func (p Title) Name() string {
	return "title"
}

func (p Title) Transform(input string, _ ...Option) string {
	return strings.Title(input)
}

func (p Title) Title() string {
	return "To Title Case"
}

func (p Title) Description() string {
	return "Transform your text to Title Case"
}

func (p Title) FilterValue() string {
	return p.Title()
}

// Snake convert string to snake_case.
// Example: "this is string" to "this_is_string".
type Snake struct{}

func (p Snake) Name() string {
	return "snake"
}

func (p Snake) Transform(input string, _ ...Option) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")
	return strcase.ToSnake(input)
}

func (p Snake) Title() string {
	return "To Snake case"
}

func (p Snake) Description() string {
	return "Transform your text to snake_case"
}

func (p Snake) FilterValue() string {
	return p.Title()
}

// Kebab convert string to kebab-case.
// Example: "this is string" to "this-is-string".
type Kebab struct{}

func (p Kebab) Name() string {
	return "kebap"
}

func (p Kebab) Transform(input string, _ ...Option) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")
	return strcase.ToKebab(input)
}

func (p Kebab) Title() string {
	return "To Kebab case"
}

func (p Kebab) Description() string {
	return "Transform your text to kebap-case"
}

func (p Kebab) FilterValue() string {
	return p.Title()
}

// Camel convert string to CamelCase.
// Example: "this is string" to "ThisIsString".
type Camel struct{}

func (p Camel) Name() string {
	return "camel"
}

func (p Camel) Transform(input string, _ ...Option) string {
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ")
	return strcase.ToCamel(input)
}

func (p Camel) Title() string {
	return "To Camel case"
}

func (p Camel) Description() string {
	return "Transform your text to CamelCase"
}

func (p Camel) FilterValue() string {
	return p.Title()
}

// Slug convert string to StringToSlug. It's similar to Kebab case but URL Friendly.
// Example: "this is string" to "this-is-string".
type Slug struct{}

func (p Slug) Name() string {
	return "slug"
}

func (p Slug) Transform(input string, _ ...Option) string {
	re := regexp.MustCompile("[^a-z0-9]+")
	return strings.Trim(re.ReplaceAllString(strings.ToLower(input), "-"), "-")
}

func (p Slug) Title() string {
	return "To Slug case"
}

func (p Slug) Description() string {
	return "Transform your text to slug-case"
}

func (p Slug) FilterValue() string {
	return p.Title()
}

// CountCharacters count number of Characters including spaces.
type CountCharacters struct{}

func (p CountCharacters) Name() string {
	return "count-chars"
}

func (p CountCharacters) Transform(input string, _ ...Option) string {
	return fmt.Sprintf("%d", len(input))
}

func (p CountCharacters) Title() string {
	return "Count Number of Characters"
}

func (p CountCharacters) Description() string {
	return "Find the length of your text (including spaces)"
}

func (p CountCharacters) FilterValue() string {
	return p.Title()
}

// CountWords count number of words in string.
// Example: "hello world" = 2
type CountWords struct{}

func (p CountWords) Name() string {
	return "count-words"
}

func (p CountWords) Transform(input string, _ ...Option) string {
	return fmt.Sprintf("%d", len(strings.Fields(input)))
}

func (p CountWords) Title() string {
	return "Count Number of Words"
}

func (p CountWords) Description() string {
	return "Count the number of words in your text"
}

func (p CountWords) FilterValue() string {
	return p.Title()
}

// CountLines count number of words in string.
// Example: "line 1\n line 2" = 2
type CountLines struct{}

func (p CountLines) Name() string {
	return "count-lines"
}

func (p CountLines) Transform(input string, _ ...Option) string {
	lines := strings.Count(input, "\n")
	if len(input) > 0 && !strings.HasSuffix(input, "\n") {
		lines++
	}
	return fmt.Sprintf("%d", lines)
}

func (p CountLines) Title() string {
	return "Count Number of Lines"
}

func (p CountLines) Description() string {
	return "Count the number of lines in your text"
}

func (p CountLines) FilterValue() string {
	return p.Title()
}

// SortLines sort given lines, it's not a natural sort.
// Example: 2\n 1\n -> 1\n 2\n
type SortLines struct{}

func (p SortLines) Name() string {
	return "sort-lines"
}

func (p SortLines) Transform(input string, _ ...Option) string {
	sorted := strings.Split(input, "\n")
	sort.Strings(sorted)
	return strings.Join(sorted, "\n")
}

func (p SortLines) Title() string {
	return "SortLines"
}

func (p SortLines) Description() string {
	return "Sort lines alphabetically"
}

func (p SortLines) FilterValue() string {
	return p.Title()
}

// Reverse reverse a given string
// Example: "test" to "tset"
type Reverse struct{}

func (p Reverse) Name() string {
	return "reverse"
}

func (p Reverse) Transform(input string, _ ...Option) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return result
}

func (p Reverse) Title() string {
	return "Reverse text"
}

func (p Reverse) Description() string {
	return "Reverse Text ( txeT esreveR )"
}

func (p Reverse) FilterValue() string {
	return p.Title()
}
