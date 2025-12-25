package processors

import (
	"strings"
	"testing"
)

func TestQRCode_Transform(t *testing.T) {
	p := QRCode{}

	tests := []struct {
		name  string
		input string
		flags []Flag
		want  string
	}{
		{
			name:  "simple text",
			input: "Hello",
			flags: []Flag{},
			want:  "█",
		},
		{
			name:  "with size flag",
			input: "Test",
			flags: []Flag{{Short: "s", Value: uint(300)}},
			want:  "█",
		},
		{
			name:  "with level flag",
			input: "Test",
			flags: []Flag{{Short: "l", Value: "H"}},
			want:  "█",
		},
		{
			name:  "with full blocks",
			input: "Test",
			flags: []Flag{{Short: "f", Value: true}},
			want:  "█",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Transform([]byte(tt.input), tt.flags...)
			if err != nil {
				t.Errorf("QRCode.Transform() error = %v", err)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("QRCode.Transform() = %v, want to contain %v", got, tt.want)
			}

			if len(strings.TrimSpace(got)) == 0 {
				t.Errorf("QRCode.Transform() returned empty output")
			}
		})
	}
}

func TestQRCode_Name(t *testing.T) {
	p := QRCode{}
	if got := p.Name(); got != "qr" {
		t.Errorf("QRCode.Name() = %v, want %v", got, "qr")
	}
}

func TestQRCode_Alias(t *testing.T) {
	p := QRCode{}
	aliases := p.Alias()
	expected := []string{"qrcode", "qr-code"}

	if len(aliases) != len(expected) {
		t.Errorf("QRCode.Alias() length = %v, want %v", len(aliases), len(expected))
		return
	}

	for i, alias := range aliases {
		if alias != expected[i] {
			t.Errorf("QRCode.Alias()[%d] = %v, want %v", i, alias, expected[i])
		}
	}
}

func TestQRCode_Flags(t *testing.T) {
	p := QRCode{}
	flags := p.Flags()

	if len(flags) != 3 {
		t.Errorf("QRCode.Flags() length = %v, want %v", len(flags), 3)
	}

	flagNames := make(map[string]bool)
	for _, flag := range flags {
		flagNames[flag.Short] = true
	}

	expectedFlags := []string{"s", "l", "f"}
	for _, expected := range expectedFlags {
		if !flagNames[expected] {
			t.Errorf("QRCode.Flags() missing flag with short name %v", expected)
		}
	}
}
