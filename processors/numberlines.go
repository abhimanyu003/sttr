package processors

import (
	"fmt"
	"strings"
)

// / Prepends consecutive number to each line of input
type LineNumberer struct{}

func (p LineNumberer) Name() string {
	return "number-lines"
}

func (p LineNumberer) Alias() []string {
	return []string{"nl"}
}

func (p LineNumberer) Transform(data []byte, _ ...Flag) (string, error) {
	var s = string(data)
	var counter = 1
	var result = ""
	for line := range strings.Lines(s) {
		if line == "\n" {
			result += line
		} else {
			result += fmt.Sprintf("%d. %s", counter, line)
			counter++
		}

	}
	return result, nil
}

func (p LineNumberer) Flags() []Flag {
	return nil
}

func (p LineNumberer) Title() string {
	return "Line numberer"
}

func (p LineNumberer) Description() string {
	return "Prepends consecutive number to each input line"
}

func (p LineNumberer) FilterValue() string {
	return p.Title()
}
