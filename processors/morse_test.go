package processors

import (
	"reflect"
	"testing"
)

func TestMorseCodeEncode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"morse-enc", "morse-encode", "morse-code-encode", "morse-code-enc"},
		description: "Encode your text to Morse Code",
		filterValue: "Morse Code Encoding (morse-encode)",
		flags:       nil,
		name:        "morse-encode",
		title:       "Morse Code Encoding (morse-encode)",
	}
	p := MorseCodeEncode{}
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

func TestMorseCodeEncode_Transform(t *testing.T) {
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
			name: "String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
			want: "- .... . / --.- ..- .. -.-. -.- / -... .-. --- .-- -. / ..-. --- -..- / .--- ..- -- .--. ... / --- ...- . .-. / .- / .-.. .- --.. -.-- / -.. --- --.",
		},
		{
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "",
		},
		{
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: ".---- ..--- ...-- ...-- ....- ..... / .- -... -.-. -.. / ....- ..... -.... / .---- ..--- ...-- / .- -... -.-. / ..... -.... --... / --... ---.. ----. -----",
		},
		{
			name: "Test For unicode standard Encoding",
			args: args{data: []byte("�")},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MorseCodeEncode{}
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

func TestMorseCodeDecode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"morse-dec", "morse-decode", "morse-code-decode", "morse-code-dec"},
		description: "Decode Morse Code to text",
		filterValue: "Morse Code Decode (morse-decode)",
		flags:       []Flag{morseDecodeLangFlag},
		name:        "morse-decode",
		title:       "Morse Code Decode (morse-decode)",
	}
	p := MorseCodeDecode{}
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

func TestMorseCodeDecode_Transform(t *testing.T) {
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
			name: "String",
			args: args{data: []byte("- .... . / --.- ..- .. -.-. -.- / -... .-. --- .-- -. / ..-. --- -..- / .--- ..- -- .--. ... / --- ...- . .-. / .- / .-.. .- --.. -.-- / -.. --- --.")},
			want: "THE QUICK BROWN FOX JUMPS OVER A LAZY DOG",
		},
		{
			name: "Multi line string",
			args: args{data: []byte(".---- ..--- ...-- ...-- ....- ..... / .- -... -.-. -.. / ....- ..... -.... / .---- ..--- ...-- / .- -... -.-. / ..... -.... --... / --... ---.. ----. -----")},
			want: "123345 ABCD 456 123 ABC 567 7890",
		},
		{
			name: "Test For specific language (he)",
			args: args{data: []byte(".- -... --. / -.."), in1: []Flag{{Short: "l", Value: "he"}}},
			want: "אבג ד",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MorseCodeDecode{}
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
