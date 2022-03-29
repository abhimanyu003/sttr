package processors

import (
	"reflect"
	"testing"
)

func TestURLEncode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"url-enc"},
		description: "Encode URL entities",
		filterValue: "URL Encode",
		flags:       nil,
		name:        "url-encode",
		title:       "URL Encode",
	}
	p := URLEncode{}
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

func TestURLEncode_Transform(t *testing.T) {
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
			name: "Special characters",
			args: args{data: []byte("http://example.com/test?test=something&value=*&^%$#@@#$%^&*(")},
			want: "http%3A%2F%2Fexample.com%2Ftest%3Ftest%3Dsomething%26value%3D%2A%26%5E%25%24%23%40%40%23%24%25%5E%26%2A%28",
		},
		{
			name: "Special characters",
			args: args{data: []byte(" ?&=#+%!<>#\\\"{}|\\\\^[]`☺\\t:/@$'()*,;")},
			want: "+%3F%26%3D%23%2B%25%21%3C%3E%23%5C%22%7B%7D%7C%5C%5C%5E%5B%5D%60%E2%98%BA%5Ct%3A%2F%40%24%27%28%29%2A%2C%3B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := URLEncode{}
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

func TestURLDecode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"url-dec"},
		description: "Decode URL entities",
		filterValue: "URL Decode",
		flags:       nil,
		name:        "url-decode",
		title:       "URL Decode",
	}
	p := URLDecode{}
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

func TestURLDecode_Transform(t *testing.T) {
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
			name: "Special characters",
			args: args{data: []byte("http%3A%2F%2Fexample.com%2Ftest%3Ftest%3Dsomething%26value%3D%2A%26%5E%25%24%23%40%40%23%24%25%5E%26%2A%28")},
			want: "http://example.com/test?test=something&value=*&^%$#@@#$%^&*(",
		},
		{
			name: "Special characters",
			args: args{data: []byte("+%3F%26%3D%23%2B%25%21%3C%3E%23%5C%22%7B%7D%7C%5C%5C%5E%5B%5D%60%E2%98%BA%5Ct%3A%2F%40%24%27%28%29%2A%2C%3B")},
			want: " ?&=#+%!<>#\\\"{}|\\\\^[]`☺\\t:/@$'()*,;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := URLDecode{}
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
