package processors

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

var List = []list.Item{
	Zeropad{},
	Lower{},
	Upper{},
	Title{},
	Snake{},
	Kebab{},
	Slug{},
	Camel{},
	Reverse{},
	CountCharacters{},
	CountWords{},
	CountLines{},
}

type Processor interface {

	// Name is the name of a processor used as the CLI command, must be one lowercase word,
	// hyphens are allowed
	Name() string

	// Transform is the text transformation function, implemented by the processor
	Transform(input string, opts ...Flag) (string, error)

	// Flags are flags that could be used to transform the text
	Flags() []Flag
}

type FlagType string

const (
	FlagInt    = FlagType("Int")
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
	// Type - required the type of the flag, supported are: string, int, uint, bool
	Type FlagType

	// Value - optional default value of the flag
	Value interface{}
}

// Zeropad is an Example processor to show how to add text processors,
// it checks if the input is a number and pads it with zeros
// Example implements 'Item' and 'DefaultItem' from package 'github.com/charmbracelet/bubbles/list'
// to work with the ui, and `Processor` from this package to do the text transformation and generation
// of the cli commands
// After implementing add the struct to List
type Zeropad struct{}

func (p Zeropad) Name() string {
	return "zeropad"
}

func (p Zeropad) Transform(input string, f ...Flag) (string, error) {
	input = strings.TrimSpace(input)
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return "", fmt.Errorf("number expected: '%s'", input)
	}

	n := 1
	for _, flag := range f {
		if flag.Short == "n" {
			x, ok := flag.Value.(int)
			if ok {
				n = x
			}
		}
	}
	return fmt.Sprintf("%s%s", strings.Repeat("0", n), input), nil
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
