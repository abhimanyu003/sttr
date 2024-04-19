package processors

import (
	"reflect"
	"testing"
)

func TestMarkdown_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"md-html"},
		description: "Convert Markdown to HTML",
		filterValue: "Markdown to HTML (markdown-html)",
		flags:       nil,
		name:        "markdown-html",
		title:       "Markdown to HTML (markdown-html)",
	}
	p := Markdown{}
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

func TestMarkdown_Transform(t *testing.T) {
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
			name: "test empty string",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "test H1",
			args: args{data: []byte("# H1")},
			want: "<h1>H1</h1>\n",
		},
		{
			name: "test bold",
			args: args{data: []byte("**the quick brown fox jumps over a lazy dog**")},
			want: "<p><strong>the quick brown fox jumps over a lazy dog</strong></p>\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Markdown{}
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
