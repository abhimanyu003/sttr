package processors

import (
	"fmt"
	"regexp"
	"strings"
)

// RemoveNewLines removes newlines from string.
type RemoveNewLines struct{}

func (p RemoveNewLines) Name() string {
	return "remove-newlines"
}

func (p RemoveNewLines) Alias() []string {
	return []string{"remove-new-lines", "trim-newlines", "trim-new-lines"}
}

func (p RemoveNewLines) Transform(data []byte, f ...Flag) (string, error) {
	separator := " "
	for _, flag := range f {
		if flag.Short == "s" {
			x, ok := flag.Value.(string)
			if ok {
				separator = x
			}
		}
	}

	str := regexp.MustCompile(`[\r\n]+`).
		ReplaceAllString(strings.TrimSpace(string(data)), separator)
	return str, nil
}

func (p RemoveNewLines) Flags() []Flag {
	return []Flag{
		{
			Name:  "separator",
			Short: "s",
			Desc:  "Separator to split multiple lines",
			Value: "",
			Type:  FlagString,
		},
	}
}

func (p RemoveNewLines) Title() string {
	title := "Remove all new lines"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p RemoveNewLines) Description() string {
	return "Remove all new lines"
}

func (p RemoveNewLines) FilterValue() string {
	return p.Title()
}

// RemoveSpaces removes all the spaces from string.
type RemoveSpaces struct{}

func (p RemoveSpaces) Name() string {
	return "remove-spaces"
}

func (p RemoveSpaces) Alias() []string {
	return []string{"remove-space", "trim-spaces", "trim-space"}
}

func (p RemoveSpaces) Transform(data []byte, f ...Flag) (string, error) {
	separator := ""
	for _, flag := range f {
		if flag.Short == "s" {
			x, ok := flag.Value.(string)
			if ok {
				separator = x
			}
		}
	}

	str := regexp.MustCompile(`[\s\r\n]+`).
		ReplaceAllString(strings.TrimSpace(string(data)), separator)
	return str, nil
}

func (p RemoveSpaces) Flags() []Flag {
	return []Flag{
		{
			Name:  "separator",
			Short: "s",
			Desc:  "Separator to split spaces",
			Value: "",
			Type:  FlagString,
		},
	}
}

func (p RemoveSpaces) Title() string {
	title := "Remove all spaces + new lines"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p RemoveSpaces) Description() string {
	return "Remove all spaces + new lines"
}

func (p RemoveSpaces) FilterValue() string {
	return p.Title()
}
