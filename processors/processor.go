package processors

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
)

var List = []list.Item{
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
	Example{},
}

type Processor interface {

	// Name is the name of a processor used as the CLI command, must be one lowercase word
	Name() string

	// Transform is the text transformation function, implemented by the processor
	Transform(input string, opts ...Option) string
}

type Option struct {
	Name  string
	Value interface{}
}

// Example is an Example processor to show how to add text processors
// Example implements 'Item' and 'DefaultItem' from package 'github.com/charmbracelet/bubbles/list'
// to work with the ui, and `Processor` from this package to do the text transformation and generation
// of the cli commands
// After implementing add the struct to List
type Example struct{}

func (p Example) Name() string {
	return "name"
}

func (p Example) Transform(input string, _ ...Option) string {
	return fmt.Sprintf("Example: %s", input)
}

func (p Example) Title() string {
	return "Example"
}

func (p Example) Description() string {
	return "Prefix your input with 'Example: '"
}

func (p Example) FilterValue() string {
	return p.Title()
}
