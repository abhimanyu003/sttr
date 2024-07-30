package processors

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/abhimanyu003/sttr/utils"

	"github.com/iancoleman/strcase"
)

// Lower converts a string to lower case.
// Example: "THIS IS STRING" to "this is string".
type Lower struct{}

func (p Lower) Name() string {
	return "lower"
}

func (p Lower) Alias() []string {
	return nil
}

func (p Lower) Transform(data []byte, _ ...Flag) (string, error) {
	return strings.ToLower(string(data)), nil
}

func (p Lower) Flags() []Flag {
	return nil
}

func (p Lower) Title() string {
	title := "To Lower case"
	return fmt.Sprintf("%s (%s)", title, p.Name())
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

func (p Upper) Alias() []string {
	return nil
}

func (p Upper) Transform(data []byte, _ ...Flag) (string, error) {
	return strings.ToUpper(string(data)), nil
}

func (p Upper) Flags() []Flag {
	return nil
}

func (p Upper) Title() string {
	title := "To Upper case"
	return fmt.Sprintf("%s (%s)", title, p.Name())
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

func (p Title) Alias() []string {
	return nil
}

func (p Title) Transform(data []byte, _ ...Flag) (string, error) {
	return strings.Title(string(data)), nil
}

func (p Title) Flags() []Flag {
	return nil
}

func (p Title) Title() string {
	title := "To Title Case"
	return fmt.Sprintf("%s (%s)", title, p.Name())
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

func (p Snake) Alias() []string {
	return nil
}

func (p Snake) Transform(data []byte, _ ...Flag) (string, error) {
	str := regexp.MustCompile(`\s+`).ReplaceAllString(string(data), " ")
	return strcase.ToSnake(str), nil
}

func (p Snake) Flags() []Flag {
	return nil
}

func (p Snake) Title() string {
	title := "To Snake case"
	return fmt.Sprintf("%s (%s)", title, p.Name())
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
	return "kebab"
}

func (p Kebab) Alias() []string {
	return nil
}

func (p Kebab) Transform(data []byte, _ ...Flag) (string, error) {
	return utils.ToKebabCase(data), nil
}

func (p Kebab) Flags() []Flag {
	return nil
}

func (p Kebab) Title() string {
	title := "To Kebab case"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Kebab) Description() string {
	return "Transform your text to kebab-case"
}

func (p Kebab) FilterValue() string {
	return p.Title()
}

// Camel convert string to camelCase.
// Example: "this is string" to "thisIsString".
type Camel struct{}

func (p Camel) Name() string {
	return "camel"
}

func (p Camel) Alias() []string {
	return nil
}

func (p Camel) Transform(data []byte, _ ...Flag) (string, error) {
	str := regexp.MustCompile(`\s+`).ReplaceAllString(string(data), " ")
	return strcase.ToLowerCamel(str), nil
}

func (p Camel) Flags() []Flag {
	return nil
}

func (p Camel) Title() string {
	title := "To Camel case"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Camel) Description() string {
	return "Transform your text to camelCase"
}

func (p Camel) FilterValue() string {
	return p.Title()
}

// Pascal convert string to CamelCase.
// Example: "this is string" to "ThisIsString".
type Pascal struct{}

func (p Pascal) Name() string {
	return "pascal"
}

func (p Pascal) Alias() []string {
	return nil
}

func (p Pascal) Transform(data []byte, _ ...Flag) (string, error) {
	str := regexp.MustCompile(`\s+`).ReplaceAllString(string(data), " ")
	return strcase.ToCamel(str), nil
}

func (p Pascal) Flags() []Flag {
	return nil
}

func (p Pascal) Title() string {
	title := "To Pascal case"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Pascal) Description() string {
	return "Transform your text to PascalCase"
}

func (p Pascal) FilterValue() string {
	return p.Title()
}

// Slug convert string to StringToSlug. It's similar to Kebab case but URL Friendly.
// Example: "this is string" to "this-is-string".
type Slug struct{}

func (p Slug) Name() string {
	return "slug"
}

func (p Slug) Alias() []string {
	return nil
}

func (p Slug) Transform(data []byte, _ ...Flag) (string, error) {
	re := regexp.MustCompile("[^a-z0-9]+")
	return strings.Trim(re.ReplaceAllString(strings.ToLower(string(data)), "-"), "-"), nil
}

func (p Slug) Flags() []Flag {
	return nil
}

func (p Slug) Title() string {
	title := "To Slug case"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Slug) Description() string {
	return "Transform your text to slug-case"
}

func (p Slug) FilterValue() string {
	return p.Title()
}

// CountCharacters count number of characters including spaces.
type CountCharacters struct{}

func (p CountCharacters) Name() string {
	return "count-chars"
}

func (p CountCharacters) Alias() []string {
	return nil
}

func (p CountCharacters) Transform(data []byte, _ ...Flag) (string, error) {
	return fmt.Sprintf("%d", len([]rune(string(data)))), nil
}

func (p CountCharacters) Flags() []Flag {
	return nil
}

func (p CountCharacters) Title() string {
	title := "Count Number of Characters"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p CountCharacters) Description() string {
	return "Find the length of your text (including spaces)"
}

func (p CountCharacters) FilterValue() string {
	return p.Title()
}

// CountWords counts number of words in string.
// Example: "hello world" = 2.
type CountWords struct{}

func (p CountWords) Name() string {
	return "count-words"
}

func (p CountWords) Alias() []string {
	return nil
}

func (p CountWords) Transform(data []byte, _ ...Flag) (string, error) {
	return fmt.Sprintf("%d", len(strings.Fields(string(data)))), nil
}

func (p CountWords) Flags() []Flag {
	return nil
}

func (p CountWords) Title() string {
	title := "Count Number of Words"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p CountWords) Description() string {
	return "Count the number of words in your text"
}

func (p CountWords) FilterValue() string {
	return p.Title()
}

// Reverse reverses a given string
// Example: "test" to "tset".
type Reverse struct{}

func (p Reverse) Name() string {
	return "reverse"
}

func (p Reverse) Alias() []string {
	return nil
}

func (p Reverse) Transform(data []byte, _ ...Flag) (string, error) {
	result := ""
	for _, v := range data {
		result = string(v) + result
	}
	return result, nil
}

func (p Reverse) Flags() []Flag {
	return nil
}

func (p Reverse) Title() string {
	title := "Reverse text"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Reverse) Description() string {
	return "Reverse Text ( txeT esreveR )"
}

func (p Reverse) FilterValue() string {
	return p.Title()
}

// EscapeQuotes escapes quotes from a given string
// Example: "test" to \"test\".
type EscapeQuotes struct{}

func (p EscapeQuotes) Name() string {
	return "escape-quotes"
}

func (p EscapeQuotes) Alias() []string {
	return []string{"esc-quotes", "escape-quotes"}
}

func (p EscapeQuotes) Transform(data []byte, f ...Flag) (string, error) {
	result := ""
	for _, v := range data {
		for _, flag := range f {
			switch flag.Short {
			case "d":
				if v == '"' {
					result += "\\"
				}
			case "s":
				if v == '\'' {
					result += "\\"
				}
			}
		}
		if len(f) == 0 {
			if v == '"' {
				result += "\\"
			}
			if v == '\'' {
				result += "\\"
			}
		}
		result += string(v)
	}
	return result, nil
}

func (p EscapeQuotes) Flags() []Flag {
	return []Flag{
		{
			Name:  "double-quote",
			Short: "d",
			Desc:  "Escape double quote",
			Value: true,
			Type:  FlagBool,
		},
		{
			Name:  "single-quote",
			Short: "s",
			Desc:  "Escape single quote",
			Value: true,
			Type:  FlagBool,
		},
	}
}

func (p EscapeQuotes) Title() string {
	title := "Escape Quotes"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p EscapeQuotes) Description() string {
	return "Escapes single and double quotes by default"
}

func (p EscapeQuotes) FilterValue() string {
	return p.Title()
}
