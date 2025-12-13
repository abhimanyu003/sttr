package processors

import (
	"reflect"
	"testing"
)

func TestCrockfordBase32Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"crockford-b32-enc", "cb32-encode"},
		description: "Encode your text to Crockford Base32",
		filterValue: "Crockford Base32 Encode (crockford-base32-encode)",
		flags: []Flag{
			{
				Name:  "checksum",
				Short: "c",
				Desc:  "Add Crockford checksum",
				Value: false,
				Type:  FlagBool,
			},
		},
		name:  "crockford-base32-encode",
		title: "Crockford Base32 Encode (crockford-base32-encode)",
	}
	p := CrockfordBase32Encode{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
	}
	if got := p.Flags(); !reflect.DeepEqual(got, test.flags) {
		t.Errorf("Flags() = %v, want %v", got, test.flags)
	}
	if got := p.Name(); got != test.name {
		t.Errorf("Name() = %v, want %v", got, test.name)
	}
	if got := p.Title(); got != test.title {
		t.Errorf("Title() = %v, want %v", got, test.title)
	}
}

func TestCrockfordBase32Encode_Transform(t *testing.T) {
	type args struct {
		data []byte
		in1  []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Hello world",
			args: args{data: []byte("hello world")},
			want: "38CNP6RVS0EXQQ4V34",
		},
		{
			name: "Empty string",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "Single character",
			args: args{data: []byte("a")},
			want: "31",
		},
		{
			name: "With checksum",
			args: args{
				data: []byte("hello"),
				in1:  []Flag{{Short: "c", Value: true}},
			},
			want: "D1JPRV3FM", // Last character is checksum
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CrockfordBase32Encode{}
			got, err := p.Transform(tt.args.data, tt.args.in1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transform() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrockfordBase32Decode_Transform(t *testing.T) {
	type args struct {
		data []byte
		in1  []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Hello world",
			args: args{data: []byte("38CNP6RVS0EXQQ4V34")},
			want: "hello world",
		},
		{
			name: "Empty string",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "Single character",
			args: args{data: []byte("31")},
			want: "a",
		},
		{
			name: "With ambiguous characters (O->0, I->1)",
			args: args{data: []byte("31")}, // Should handle O and I normalization
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CrockfordBase32Decode{}
			got, err := p.Transform(tt.args.data, tt.args.in1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transform() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase58Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"b58-enc", "b58-encode"},
		description: "Encode your text to Base58",
		filterValue: "Base58 Encode (base58-encode)",
		flags: []Flag{
			{
				Name:  "check",
				Short: "c",
				Desc:  "Use Base58Check encoding (with checksum)",
				Value: false,
				Type:  FlagBool,
			},
		},
		name:  "base58-encode",
		title: "Base58 Encode (base58-encode)",
	}
	p := Base58Encode{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
	}
	if got := p.Flags(); !reflect.DeepEqual(got, test.flags) {
		t.Errorf("Flags() = %v, want %v", got, test.flags)
	}
	if got := p.Name(); got != test.name {
		t.Errorf("Name() = %v, want %v", got, test.name)
	}
	if got := p.Title(); got != test.title {
		t.Errorf("Title() = %v, want %v", got, test.title)
	}
}

func TestBase58Encode_Transform(t *testing.T) {
	type args struct {
		data []byte
		in1  []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Hello world",
			args: args{data: []byte("hello world")},
			want: "StV1DL6CwTryKyV",
		},
		{
			name: "Empty string",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "Single character",
			args: args{data: []byte("a")},
			want: "2g",
		},
		{
			name: "Leading zeros",
			args: args{data: []byte("\x00\x00hello")},
			want: "11Cn8eVZg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base58Encode{}
			got, err := p.Transform(tt.args.data, tt.args.in1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transform() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase58Decode_Transform(t *testing.T) {
	type args struct {
		data []byte
		in1  []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Hello world",
			args: args{data: []byte("StV1DL6CwTryKyV")},
			want: "hello world",
		},
		{
			name: "Empty string",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "Single character",
			args: args{data: []byte("2g")},
			want: "a",
		},
		{
			name: "Leading zeros",
			args: args{data: []byte("11Cn8eVZg")},
			want: "\x00\x00hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base58Decode{}
			got, err := p.Transform(tt.args.data, tt.args.in1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transform() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase62Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"b62-enc", "b62-encode"},
		description: "Encode your text to Base62",
		filterValue: "Base62 Encode (base62-encode)",
		flags: []Flag{
			{
				Name:  "prefix",
				Short: "p",
				Desc:  "Add prefix to encoded string",
				Value: "",
				Type:  FlagString,
			},
		},
		name:  "base62-encode",
		title: "Base62 Encode (base62-encode)",
	}
	p := Base62Encode{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
	}
	if got := p.Flags(); !reflect.DeepEqual(got, test.flags) {
		t.Errorf("Flags() = %v, want %v", got, test.flags)
	}
	if got := p.Name(); got != test.name {
		t.Errorf("Name() = %v, want %v", got, test.name)
	}
	if got := p.Title(); got != test.title {
		t.Errorf("Title() = %v, want %v", got, test.title)
	}
}

func TestBase62Encode_Transform(t *testing.T) {
	type args struct {
		data []byte
		in1  []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Hello world",
			args: args{data: []byte("hello world")},
			want: "AAwf93rvy4aWQVw",
		},
		{
			name: "Empty string",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "Single character",
			args: args{data: []byte("a")},
			want: "1Z",
		},
		{
			name: "With prefix",
			args: args{
				data: []byte("hello"),
				in1:  []Flag{{Short: "p", Value: "gh"}},
			},
			want: "gh_7tQLFHz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base62Encode{}
			got, err := p.Transform(tt.args.data, tt.args.in1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transform() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase62Decode_Transform(t *testing.T) {
	type args struct {
		data []byte
		in1  []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Hello world",
			args: args{data: []byte("AAwf93rvy4aWQVw")},
			want: "hello world",
		},
		{
			name: "Empty string",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "Single character",
			args: args{data: []byte("1Z")},
			want: "a",
		},
		{
			name: "With prefix",
			args: args{data: []byte("gh_7tQLFHz")},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base62Decode{}
			got, err := p.Transform(tt.args.data, tt.args.in1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transform() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test round-trip encoding/decoding
func TestCrockfordBase32_RoundTrip(t *testing.T) {
	testCases := []string{
		"hello world",
		"The quick brown fox jumps over the lazy dog",
		"",
		"a",
		"123456789",
		"Special chars: !@#$%^&*()",
	}

	encoder := CrockfordBase32Encode{}
	decoder := CrockfordBase32Decode{}

	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			encoded, err := encoder.Transform([]byte(testCase))
			if err != nil {
				t.Errorf("Encode error: %v", err)
				return
			}

			decoded, err := decoder.Transform([]byte(encoded))
			if err != nil {
				t.Errorf("Decode error: %v", err)
				return
			}

			if decoded != testCase {
				t.Errorf("Round-trip failed: got %v, want %v", decoded, testCase)
			}
		})
	}
}

func TestBase58_RoundTrip(t *testing.T) {
	testCases := []string{
		"hello world",
		"The quick brown fox jumps over the lazy dog",
		"",
		"a",
		"123456789",
		"\x00\x00hello", // Test leading zeros
	}

	encoder := Base58Encode{}
	decoder := Base58Decode{}

	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			encoded, err := encoder.Transform([]byte(testCase))
			if err != nil {
				t.Errorf("Encode error: %v", err)
				return
			}

			decoded, err := decoder.Transform([]byte(encoded))
			if err != nil {
				t.Errorf("Decode error: %v", err)
				return
			}

			if decoded != testCase {
				t.Errorf("Round-trip failed: got %v, want %v", decoded, testCase)
			}
		})
	}
}

func TestBase62_RoundTrip(t *testing.T) {
	testCases := []string{
		"hello world",
		"The quick brown fox jumps over the lazy dog",
		"",
		"a",
		"123456789",
		"Special chars: !@#$%^&*()",
	}

	encoder := Base62Encode{}
	decoder := Base62Decode{}

	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			encoded, err := encoder.Transform([]byte(testCase))
			if err != nil {
				t.Errorf("Encode error: %v", err)
				return
			}

			decoded, err := decoder.Transform([]byte(encoded))
			if err != nil {
				t.Errorf("Decode error: %v", err)
				return
			}

			if decoded != testCase {
				t.Errorf("Round-trip failed: got %v, want %v", decoded, testCase)
			}
		})
	}
}
