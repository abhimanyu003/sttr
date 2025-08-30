package processors

import (
	"fmt"
	"strconv"
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

func nonEmptyCount(strs []string) int {
	var count = 0
	for _, s := range strs {
		if s != "" {
			count++
		}
	}
	return count
}
func (p LineNumberer) Transform(data []byte, _ ...Flag) (string, error) {
	var s = string(data)
	var counter = 1
	var result = ""
	var lines = strings.Split(s, "\n")
	var nec = nonEmptyCount(lines)
	var maxDigits = len(strconv.Itoa(nec))
	for _, line := range lines {
		if line != "" {
			line = fmt.Sprintf("%*d. %s", maxDigits, counter, line)
			counter++
		}
		result += line + "\n"

	}
	result = strings.TrimSuffix(result, "\n")
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
