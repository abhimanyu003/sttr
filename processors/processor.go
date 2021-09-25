package processors

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

var List = []list.Item{
	Base64Decode{},
	Base64Encode{},
	Camel{},
	CountCharacters{},
	CountLines{},
	CountWords{},
	FormatJSON{},
	HexToRGB{},
	HTMLEncode{},
	HTMLDecode{},
	JSONToYAML{},
	Kebab{},
	Lower{},
	MD5Encode{},
	Reverse{},
	ROT13Encode{},
	SHA1Encode{},
	SHA256Encode{},
	SHA512Encode{},
	Slug{},
	Snake{},
	Title{},
	Upper{},
	URLDecode{},
	URLEncode{},
	YAMLToJSON{},
	Zeropad{},
}

type Processor interface {

	// Name is the name of a processor used as the CLI command, must be one lowercase word,
	// hyphens are allowed
	Name() string

	// Alias is an optional array of alias names for the processor
	Alias() []string

	// Transform is the text transformation function, implemented by the processor
	Transform(input string, opts ...Flag) (string, error)

	// Flags are flags that could be used to transform the text
	Flags() []Flag
}

type FlagType string

func (f FlagType) String() string {
	return string(f)
}

func (f FlagType) IsString() bool {
	return f == FlagString
}

const (
	FlagInt    = FlagType("Int")
	FlagUint   = FlagType("Uint")
	FlagBool   = FlagType("Bool")
	FlagString = FlagType("String")
)

type Flag struct {
	// Name - required (long version) of the flag, lowercase (with hyphens)
	Name string

	// Short - required (single character, lowercase) of the flag
	Short string

	// Desc - required, a short description of the flag
	Desc string
	// Type - required the type of the flag
	Type FlagType

	// Value - optional default value of the flag
	Value interface{}
}

// Zeropad is an Example processor to show how to add text processors,
// it checks if the data is a number and pads it with zeros
// Example implements 'Item' and 'DefaultItem' from package 'github.com/charmbracelet/bubbles/list'
// to work with the ui, and `Processor` from this package to do the text transformation and generation
// of the cli commands
// After implementing add the struct to List
type Zeropad struct{}

func (p Zeropad) Name() string {
	return "zeropad"
}

func (p Zeropad) Alias() []string {
	return nil
}

func (p Zeropad) Transform(input string, f ...Flag) (string, error) {
	input = strings.TrimSpace(input)
	neg := ""
	i, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return "", fmt.Errorf("number expected: '%s'", input)
	}
	if i < 0 {
		neg = "-"
		input = input[1:]
	}

	n := 1
	pre := ""
	for _, flag := range f {
		if flag.Short == "n" {
			x, ok := flag.Value.(int)
			if ok {
				n = x
			}
		} else if flag.Short == "p" {
			x, ok := flag.Value.(string)
			if ok {
				pre = x
			}
		}
	}
	return fmt.Sprintf("%s%s%s%s", pre, neg, strings.Repeat("0", n), input), nil
}

func (p Zeropad) Flags() []Flag {
	return []Flag{
		{
			Name:  "number-of-zeros",
			Short: "n",
			Desc:  "Number of zeros to be padded",
			Value: 5,
			Type:  FlagInt,
		},
		{
			Name:  "prefix",
			Short: "p",
			Desc:  "The number get prefixed with this",
			Value: "",
			Type:  FlagString,
		},
	}
}

func (p Zeropad) Title() string {
	return strings.Title(p.Name())
}

func (p Zeropad) Description() string {
	return "Pad a number with zeros"
}

func (p Zeropad) FilterValue() string {
	return p.Title()
}
