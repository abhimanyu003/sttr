package processors

import (
	"reflect"
	"testing"
)

func TestAscii85Encoding_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"ascii85-encoding", "base85-encode", "b85-encode"},
		description: "Encode your text to Ascii85 ( Base85 )",
		filterValue: "Ascii85 / Base85 Encoding (ascii85-encode)",
		flags:       nil,
		name:        "ascii85-encode",
		title:       "Ascii85 / Base85 Encoding (ascii85-encode)",
	}
	p := ASCII85Encoding{}
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

func TestAscii85Encoding_Transform(t *testing.T) {
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
			args: args{data: []byte("abhimanyu")},
			want: "@:EnaD..=-FT",
		}, {
			name: "Emoji",
			args: args{data: []byte("😇🙃🙂😉😌😙😗🇮🇳")},
			want: "n=Q)'n=Q,$n=Q,#n=Q))n=Q),n=Q)9n=Q)7n=PK=n=PKB",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "0etO@1c5VK@UipU1c7/u0etNl@:E^R2)[B#2`NfO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ASCII85Encoding{}
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

func TestAscii85Decoding_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"ascii85-decoding", "base85-decode", "b85-decode"},
		description: "Decode your text to Ascii85 ( Base85 ) text",
		filterValue: "Ascii85 / Base85 Decoding (ascii85-decode)",
		flags:       nil,
		name:        "ascii85-decode",
		title:       "Ascii85 / Base85 Decoding (ascii85-decode)",
	}
	p := ASCII85Decoding{}
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

func TestAscii85Decoding_Transform(t *testing.T) {
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
			args: args{data: []byte("@:EnaD..=-FT")},
			want: "abhimanyu",
		}, {
			name: "Emoji",
			args: args{data: []byte("n=Q)'n=Q,$n=Q,#n=Q))n=Q),n=Q)9n=Q)7n=PK=n=PKB")},
			want: "😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("0etO@1c5VK@UipU1c7/u0etNl@:E^R2)[B#2`NfO")},
			want: "123345\nabcd\n456\n123\nabc\n567\n7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ASCII85Decoding{}
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
