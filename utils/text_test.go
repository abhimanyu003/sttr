package utils

import (
	"testing"
)

func TestTrimTrailingLinebreaks(t *testing.T) {
	inputs := []string{
		"a", "aa", "Hello World", "a ", "Hello World ",
		" Hello World", " a ", "", " ", "  ", "Hello\nWorld", "Hello\rWorld", "Hello\r\nWord",
	}
	type args struct {
		input     []string
		linebreak string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "no linebreak",
			args: args{input: inputs, linebreak: ""},
		},
		{
			name: "lf",
			args: args{input: inputs, linebreak: "\n"},
		},
		{
			name: "cr",
			args: args{input: inputs, linebreak: "\r"},
		},
		{
			name: "crlf",
			args: args{input: inputs, linebreak: "\r\n"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, in := range tt.args.input {
				if got := TrimTrailingLinebreaks(in + tt.args.linebreak); got != in {
					t.Errorf("TrimTrailingLinebreaks() = %v, want %v", []byte(got), []byte(in))
				}
			}
		})
	}
}
