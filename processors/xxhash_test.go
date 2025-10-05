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
		alias:       []string{"xxh64", "xxhash64", "xxhash-64"},
		description: "Get the XXH64 checksum of your text",
		filterValue: "xxHash - XXH64 (xxh-64)",
		flags:       nil,
		name:        "xxh-64",
		title:       "xxHash - XXH64 (xxh-64)",
	}

	p := XXH64{}
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
			p := XXH64{}
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

func TestXXHash32_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"xxh32", "xxhash32", "xxhash-32"},
		description: "Get the XXH32 checksum of your text",
		filterValue: "xxHash - XXH32 (xxh-32)",
		flags:       nil,
		name:        "xxh-32",
		title:       "xxHash - XXH32 (xxh-32)",
	}

	p := XXH32{}
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

func TestXXHash32_Transform(t *testing.T) {
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
			want: "21276298",
		}, {
			name: "empty string", // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			args: args{data: []byte("")},
			want: "02cc5d05",
		}, {
			name: "Short string - single character", // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			args: args{data: []byte("a")},
			want: "550d7456", // expected hash
		}, {
			name: "Numeric string",
			args: args{data: []byte("1234567890")}, // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			want: "e8412d73",
		}, {
			name: "Special characters",
			args: args{data: []byte("!@#$%^&*()_+")}, // test values form https://asecuritysite.com/encryption/xxhash test with zero seed value
			want: "551afd45",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := XXH32{}
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

func TestXXHash128_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"xxh128", "xxhash128", "xxhash-128"},
		description: "Get the XXH128 checksum of your text",
		filterValue: "xxHash - XXH128 (xxh-128)",
		flags:       nil,
		name:        "xxh-128",
		title:       "xxHash - XXH128 (xxh-128)",
	}

	p := XXH128{}
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

func TestXXHash128_Transform(t *testing.T) {
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
			name: "String", // test values form https://github.com/zeebo/xxh3 test with zero seed value
			args: args{data: []byte("Roger federar is not a tennis player")},
			want: "85823e339aafe1cf06d303927cfbd90d",
		}, {
			name: "empty string", // test values form https://github.com/zeebo/xxh3 test with zero seed value
			args: args{data: []byte("")},
			want: "99aa06d3014798d86001c324468d497f",
		}, {
			name: "Short string - single character", // test values form https://github.com/zeebo/xxh3 test with zero seed value
			args: args{data: []byte("a")},
			want: "a96faf705af16834e6c632b61e964e1f", // expected hash
		}, {
			name: "Numeric string",
			args: args{data: []byte("1234567890")}, // test values form https://github.com/zeebo/xxh3 test with zero seed value
			want: "82d9f70aeb974c48565e705734e91277",
		}, {
			name: "Special characters",
			args: args{data: []byte("!@#$%^&*()_+")}, // test values form https://github.com/zeebo/xxh3 test with zero seed value
			want: "7086571ffa7bdd87635c94f646d2a43a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := XXH128{}
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
