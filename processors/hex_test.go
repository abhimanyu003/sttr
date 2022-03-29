package processors

import (
	"reflect"
	"testing"
)

func TestHexEncode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"hex-enc", "hexadecimal-encode"},
		description: "Encode your text Hex",
		filterValue: "Hex Encode",
		flags:       nil,
		name:        "hex-encode",
		title:       "Hex Encode",
	}
	p := HexEncode{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("Flags() = %v, want %v", got, test.filterValue)
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

func TestHexEncode_Transform(t *testing.T) {
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
			name: "Test empty string",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "String",
			args: args{data: []byte("this is string")},
			want: "7468697320697320737472696e67",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "f09f9883f09f9887f09f9983f09f9982f09f9889f09f988cf09f9899f09f9897f09f87aef09f87b3",
		}, {
			name: "Multi line string",
			args: args{data: []byte("Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***")},
			want: "48656c6c6f0a666f6f0a6261720a666f6f0a6630300a252a265e2a265e260a2a2a2a",
		}, {
			name: "Test multi lingual character",
			args: args{data: []byte("नमस्ते")},
			want: "e0a4a8e0a4aee0a4b8e0a58de0a4a4e0a587",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := HexEncode{}
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

func TestHexDecode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"hex-dec", "hexadecimal-decode"},
		description: "Convert Hexadecimal to String",
		filterValue: "Hex Decode",
		flags:       nil,
		name:        "hex-decode",
		title:       "Hex Decode",
	}
	p := HexDecode{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("Flags() = %v, want %v", got, test.filterValue)
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

func TestHexDecode_Transform(t *testing.T) {
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
			name:    "Test empty hex",
			args:    args{data: []byte("")},
			want:    "",
			wantErr: false,
		},
		{
			name:    "Test invalid hex",
			args:    args{data: []byte("this is invalid hex")},
			want:    "",
			wantErr: true,
		},
		{
			name: "String",
			args: args{data: []byte("7468697320697320737472696e67")},
			want: "this is string",
		}, {
			name: "Emoji",
			args: args{data: []byte("f09f9883f09f9887f09f9983f09f9982f09f9889f09f988cf09f9899f09f9897f09f87aef09f87b3")},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("48656c6c6f0a666f6f0a6261720a666f6f0a6630300a252a265e2a265e260a2a2a2a")},
			want: "Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := HexDecode{}
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
