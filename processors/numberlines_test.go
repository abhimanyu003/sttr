package processors

import (
	"reflect"
	"testing"
)

func TestLineNumberer_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"nl"},
		description: "Prepends consecutive number to each input line",
		filterValue: "Line numberer",
		flags:       nil,
		name:        "number-lines",
		title:       "Line numberer",
	}
	p := LineNumberer{}
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

func TestLineNumberer_Transform(t *testing.T) {
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
			want: "1. this is string",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "1. 😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***")},
			want: "1. Hello\n2. foo\n3. bar\n4. foo\n5. f00\n6. %*&^*&^&\n7. ***",
		}, {
			name: "Test multi lingual character",
			args: args{data: []byte("नमस्ते")},
			want: "1. नमस्ते",
		},
		{
			name: "Lines with empty lines between",
			args: args{data: []byte("this is string\n\nfollowed by empty line and another")},
			want: "1. this is string\n\n2. followed by empty line and another",
		},
		{
			name: "Lines with first empty line",
			args: args{data: []byte("\nthis is string after newline\n\nfollowed by empty line and another")},
			want: "\n1. this is string after newline\n\n2. followed by empty line and another",
		},
		{
			name: "Lines with space-only lines between",
			args: args{data: []byte("this is string\n   \nfollowed by three-space line and another")},
			want: "1. this is string\n2.    \n3. followed by three-space line and another",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := LineNumberer{}
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
