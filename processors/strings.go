package processors

import (
	"fmt"
	"github.com/abhimanyu003/sttr/utils"
	"regexp"
	"strings"

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
	return "To Kebab case"
}

func (p Kebab) Description() string {
	return "Transform your text to kebab-case"
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

func (p Camel) Alias() []string {
	return nil
}

func (p Camel) Transform(data []byte, _ ...Flag) (string, error) {
	str := regexp.MustCompile(`\s+`).ReplaceAllString(string(data), " ")
	return strcase.ToCamel(str), nil
}

func (p Camel) Flags() []Flag {
	return nil
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
	return "Count Number of Words"
}

func (p CountWords) Description() string {
	return "Count the number of words in your text"
}

func (p CountWords) FilterValue() string {
	return p.Title()
}

// Reverse reverse a given string
// Example: "test" to "tset"
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
	return "Reverse text"
}

func (p Reverse) Description() string {
	return "Reverse Text ( txeT esreveR )"
}

func (p Reverse) FilterValue() string {
	return p.Title()
}

// Split a given string
// Example: "test,test"  to ["test", "test"]
type Split struct{}

func (p Split) Name() string {
	return "split"
}

func (p Split) Alias() []string {
	return nil
}

func (p Split) Transform(data []byte, s ...Flag) (string, error) {
	// default separator is ,
	separator := ","
	for _, flag := range s {
		if flag.Short == "s" {
			if s, ok := flag.Value.(string); ok {
				separator = s
			}

		}
	}

	// escape \n and \t
	if strings.Contains(separator, "\\n") {
		separator = strings.Replace(separator, "\\n", string([]byte{0xa}), -1)
	}
	if strings.Contains(separator, "\\t") {
		separator = strings.Replace(separator, "\\t", string([]byte{0x9}), -1)
	}
	strOut := strings.Join(strings.Split(string(data), separator), "\",\"")
	return fmt.Sprintf("[\"%s\"]", strOut), nil
}

func (p Split) Flags() []Flag {
	return []Flag{
		{
			Name:  "separator",
			Short: "s",
			Desc:  "Separator to split string",
			Value: ",",
			Type:  FlagString,
		},
	}
}

func (p Split) Title() string {
	return "Split text"
}

func (p Split) Description() string {
	return "Split Text to string list"
}

func (p Split) FilterValue() string {
	return p.Title()
}
