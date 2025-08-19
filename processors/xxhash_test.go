package processors

import (
	"reflect"
	"testing"
)

func TestXXHash64_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"xxhash-64"},
		description: "Get the XXHash64 checksum of your text",
		filterValue: "XXhash - 64 (xxhash-64)",
		flags:       nil,
		name:        "xxhash-64",
		title:       "XXhash - 64 (xxhash-64)",
	}

	p := XXHash64{}
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

func TestXXHash64_Transform(t *testing.T) {
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
			name: "String", // test values form https://asecuritysite.com/encryption/xxhash
			args: args{data: []byte("Roger federar is not a tennis player")},
			want: "ef544ce741b215c2",
		}, {
			name: "empty string", // test values form https://asecuritysite.com/encryption/xxhash
			args: args{data: []byte("")},
			want: "ef46db3751d8e999",
		}, {
			name: "a rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("a")},
			want: "0cc175b9c0f1b6a831c399e269772661",
		}, {
			name: "abc rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("abc")},
			want: "900150983cd24fb0d6963f7d28e17f72",
		}, {
			name: "message digest rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("message digest")},
			want: "f96b697d7cb7938d525a2f31aaf161d0",
		}, {
			name: "abcdefghijklmnopqrstuvwxyz rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("abcdefghijklmnopqrstuvwxyz")},
			want: "c3fcd3d76192e4007dfb496cca67e13b",
		}, {
			name: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")},
			want: "d174ab98d277d9f5a5611c2c9f419d9f",
		}, {
			name: "12345678901234567890123456789012345678901234567890123456789012345678901234567890 rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890")},
			want: "57edf4a22be3c955ac49da2e2107b67a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := XXHash64{}
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
