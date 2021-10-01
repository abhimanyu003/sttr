package processors

import (
	"reflect"
	"testing"
)

func TestROT13Encode_Transform(t *testing.T) {
	type args struct {
		data []byte
		in1   []Flag
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
			want: "gur dhvpx oebja sbk wh`cf bire n ynml qbt",
		}, {
			name: "String Uppercase",
			args: args{data: []byte("THE QUICK BROWN FOX JUMPS OVER A LAZY DOG")},
			want: "GUR DHVPX OEBJA SBK WH@CF BIRE N YNML QBT",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "123345\nnopq\n456\n123\nnop\n567\n7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ROT13Encode{}
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

func TestROT13Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"rot13", "rot13-enc"},
		description: "Encode your text to ROT13",
		filterValue: "ROT13 Encode",
		flags:       nil,
		name:        "rot13-encode",
		title:       "ROT13 Encode",
	}
	p := ROT13Encode{}
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
