package processors

import (
	"reflect"
	"testing"
)

func TestBase32Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"b32-enc", "b32-encode"},
		description: "Encode your text to Base32",
		filterValue: "Base32 Encoding (base32-encode)",
		flags:       nil,
		name:        "base32-encode",
		title:       "Base32 Encoding (base32-encode)",
	}
	p := Base32Encoding{}
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

func TestBase32Encode_Transform(t *testing.T) {
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
			want: "ORUGKIDROVUWG2ZAMJZG653OEBTG66BANJ2W24DTEBXXMZLSEBQSA3DBPJ4SAZDPM4======",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "6CPZRA7QT6MIP4E7TGB7BH4ZQLYJ7GEJ6CPZRDHQT6MJT4E7TCL7BH4HV3YJ7B5T",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "GEZDGMZUGUFGCYTDMQFDINJWBIYTEMYKMFRGGCRVGY3QUNZYHEYA====",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base32Encoding{}
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

func TestBase32Decode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"b32-dec", "b32-decode"},
		description: "Decode your base32 text",
		filterValue: "Base32 Decode (base32-decode)",
		flags:       nil,
		name:        "base32-decode",
		title:       "Base32 Decode (base32-decode)",
	}
	p := Base32Decode{}
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

func TestBase32Decode_Transform(t *testing.T) {
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
			args: args{data: []byte("ORUGKIDROVUWG2ZAMJZG653OEBTG66BANJ2W24DTEBXXMZLSEBQSA3DBPJ4SAZDPM4======")},
			want: "the quick brown fox jumps over a lazy dog",
		}, {
			name: "Emoji",
			args: args{data: []byte("6CPZRA7QT6MIP4E7TGB7BH4ZQLYJ7GEJ6CPZRDHQT6MJT4E7TCL7BH4HV3YJ7B5T")},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("GEZDGMZUGUFGCYTDMQFDINJWBIYTEMYKMFRGGCRVGY3QUNZYHEYA====")},
			want: "123345\nabcd\n456\n123\nabc\n567\n7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base32Decode{}
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
