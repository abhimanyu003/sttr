package processors

import (
	"fmt"
	"strconv"
	"strings"
)

// NumberLines Prepends consecutive number to each line of input
type NumberLines struct{}

// Implement ConfigurableStreamingProcessor interface for line-by-line processing
func (p NumberLines) GetStreamingConfig() StreamingConfig {
	return StreamingConfig{
		ChunkSize:    64 * 1024, // 64KB chunks
		BufferOutput: true,      // Need to count lines first
		LineByLine:   false,     // Need full input to calculate max digits
	}
}

func (p NumberLines) Name() string {
	return "number-lines"
}

func (p NumberLines) Alias() []string {
	return []string{"nl", "line-numbers", "line-number", "number-line", "numberlines", "numberline"}
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
func (p NumberLines) Transform(data []byte, _ ...Flag) (string, error) {
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

func (p NumberLines) Flags() []Flag {
	return nil
}

func (p NumberLines) Title() string {
	return "Line numberer"
}

func (p NumberLines) Description() string {
	return "Prepends consecutive number to each input line"
}

func (p NumberLines) FilterValue() string {
	return p.Title()
}
