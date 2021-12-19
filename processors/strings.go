package processors

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/mcnijman/go-emailaddress"

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

// CountLines count number of words in string.
// Example: "line 1\n line 2" = 2
type CountLines struct{}

func (p CountLines) Name() string {
	return "count-lines"
}

func (p CountLines) Alias() []string {
	return nil
}

func (p CountLines) Transform(data []byte, _ ...Flag) (string, error) {
	var lines int
	if len(data) > 0 {
		lines = strings.Count(string(data), "\n") + 1
	}
	return fmt.Sprintf("%d", lines), nil
}

func (p CountLines) Flags() []Flag {
	return nil
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

func (p SortLines) Alias() []string {
	return nil
}

func (p SortLines) Transform(data []byte, _ ...Flag) (string, error) {
	sorted := strings.Split(string(data), "\n")
	sort.Strings(sorted)
	return strings.Join(sorted, "\n"), nil
}

func (p SortLines) Flags() []Flag {
	return nil
}

func (p SortLines) Title() string {
	return "Sort Lines"
}

func (p SortLines) Description() string {
	return "Sort lines alphabetically"
}

func (p SortLines) FilterValue() string {
	return p.Title()
}

// ShuffleLines sort given lines, in random order.
type ShuffleLines struct{}

func (p ShuffleLines) Name() string {
	return "shuffle-lines"
}

func (p ShuffleLines) Alias() []string {
	return nil
}

func (p ShuffleLines) Transform(data []byte, _ ...Flag) (string, error) {
	seed, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(time.Now().Nanosecond())))
	if err != nil {
		return "", err
	}
	rand.Seed(seed.Int64())

	shuffle := strings.Split(string(data), "\n")
	rand.Shuffle(len(shuffle), func(i, j int) {
		shuffle[i], shuffle[j] = shuffle[j], shuffle[i]
	})
	return strings.Join(shuffle, "\n"), nil
}

func (p ShuffleLines) Flags() []Flag {
	return nil
}

func (p ShuffleLines) Title() string {
	return "Shuffle Lines"
}

func (p ShuffleLines) Description() string {
	return "Shuffle lines randomly"
}

func (p ShuffleLines) FilterValue() string {
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

// ExtractEmails will pluck all the valid emails from a given text.
type ExtractEmails struct{}

func (p ExtractEmails) Name() string {
	return "extract-emails"
}

func (p ExtractEmails) Alias() []string {
	return []string{"find-emails", "find-email", "extract-email"}
}

func (p ExtractEmails) Transform(data []byte, f ...Flag) (string, error) {
	var emails []string
	extracted := emailaddress.FindWithIcannSuffix(data, false)
	for _, e := range extracted {
		emails = append(emails, e.String())
	}

	separator := "\n"
	for _, flag := range f {
		if flag.Short == "s" {
			x, ok := flag.Value.(string)
			if ok {
				separator = x
			}
		}
	}
	return strings.Join(emails, separator), nil
}

func (p ExtractEmails) Flags() []Flag {
	return []Flag{
		{
			Name:  "separator",
			Short: "s",
			Desc:  "Separator to split multiple emails",
			Value: "",
			Type:  FlagString,
		},
	}
}

func (p ExtractEmails) Title() string {
	return "Extract Emails"
}

func (p ExtractEmails) Description() string {
	return "Extract emails from given text"
}

func (p ExtractEmails) FilterValue() string {
	return p.Title()
}
