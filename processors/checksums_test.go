package processors

import (
	"reflect"
	"testing"
)

func TestCRC32_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"crc32-sum", "crc32-checksum"},
		description: "Get the CRC32 checksum of your text",
		filterValue: "CRC32 Checksum (crc32)",
		flags: []Flag{
			{
				Name:  "polynomial",
				Short: "p",
				Desc:  "CRC32 polynomial (ieee, castagnoli, koopman)",
				Value: "ieee",
				Type:  FlagString,
			},
		},
		name:  "crc32",
		title: "CRC32 Checksum (crc32)",
	}
	p := CRC32{}
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

func TestCRC32_Transform(t *testing.T) {
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
			name: "IEEE polynomial (default)",
			args: args{data: []byte("hello world")},
			want: "0d4a1185",
		},
		{
			name: "Castagnoli polynomial",
			args: args{
				data: []byte("hello world"),
				in1:  []Flag{{Short: "p", Value: "castagnoli"}},
			},
			want: "c99465aa",
		},
		{
			name: "Koopman polynomial",
			args: args{
				data: []byte("hello world"),
				in1:  []Flag{{Short: "p", Value: "koopman"}},
			},
			want: "df373d3c",
		},
		{
			name: "Empty string",
			args: args{data: []byte("")},
			want: "00000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CRC32{}
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

func TestAdler32_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"adler32-sum", "adler32-checksum"},
		description: "Get the Adler32 checksum of your text",
		filterValue: "Adler32 Checksum (adler32)",
		flags:       nil,
		name:        "adler32",
		title:       "Adler32 Checksum (adler32)",
	}
	p := Adler32{}
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

func TestAdler32_Transform(t *testing.T) {
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
			want: "1a0b045d",
		},
		{
			name: "Empty string",
			args: args{data: []byte("")},
			want: "00000001",
		},
		{
			name: "Single character",
			args: args{data: []byte("a")},
			want: "00620062",
		},
		{
			name: "Numbers",
			args: args{data: []byte("123456789")},
			want: "091e01de",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Adler32{}
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
