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
			name: "String", // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			args: args{data: []byte("Roger federar is not a tennis player")},
			want: "ef544ce741b215c2",
		}, {
			name: "empty string", // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			args: args{data: []byte("")},
			want: "ef46db3751d8e999",
		}, {
			name: "Short string - single character", // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			args: args{data: []byte("a")},
			want: "d24ec4f1a98c6e5b", // expected hash
		}, {
			name: "Numeric string",
			args: args{data: []byte("1234567890")}, // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			want: "a9d4d4132eff23b6",
		}, {
			name: "Special characters",
			args: args{data: []byte("!@#$%^&*()_+")}, // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			want: "05a23b68b8435433",
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
