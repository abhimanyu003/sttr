package processors

import (
	"reflect"
	"strings"
	"testing"
)

func TestCountLines_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Count the number of lines in your text",
		filterValue: "Count Number of Lines (count-lines)",
		flags:       nil,
		name:        "count-lines",
		title:       "Count Number of Lines (count-lines)",
	}
	p := CountLines{}
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

func TestCountLines_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Count empty",
			args: args{data: []byte("")},
			want: "0",
		},
		{
			name: "Count nil",
			args: args{data: nil},
			want: "0",
		},
		{
			name: "Count one line without newline",
			args: args{data: []byte("one line")},
			want: "1",
		},
		{
			name: "Count one line with newline",
			args: args{data: []byte("one line\n")},
			want: "1",
		},
		{
			name: "Count two lines without trailing newline",
			args: args{data: []byte("1st line\n2nd line")},
			want: "2",
		},
		{
			name: "Count two lines with trailing newline",
			args: args{data: []byte("1st line\n2nd line\n")},
			want: "2",
		},
		{
			name: "Count three empty lines",
			args: args{data: []byte("\n\n\n")},
			want: "3",
		},
		{
			name: "Count mixed content",
			args: args{data: []byte("1\n2\n3\n ")},
			want: "4",
		},
		{
			name: "Count single newline",
			args: args{data: []byte("\n")},
			want: "1",
		},
		{
			name: "Count multiple 'a' lines (like user's 500MB file)",
			args: args{data: []byte("a\na\na\na\na\n")},
			want: "5",
		},
		{
			name: "Count multiple 'a' lines without trailing newline",
			args: args{data: []byte("a\na\na\na\na")},
			want: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CountLines{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("CountLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleLines_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Shuffle lines randomly",
		filterValue: "Shuffle Lines (shuffle-lines)",
		flags:       nil,
		name:        "shuffle-lines",
		title:       "Shuffle Lines (shuffle-lines)",
	}
	p := ShuffleLines{}
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

func TestShuffleLines_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Should shuffle lines",
			args: args{data: []byte("1\n2")},
			want: 2,
		},
		{
			name: "Should return one line",
			args: args{data: []byte("there is no email in text")},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ShuffleLines{}

			got, _ := p.Transform(tt.args.data, tt.args.opts...)
			count := len(strings.Split(got, "\n"))
			if count != tt.want {
				t.Errorf("ShuffleLines() = %v, want %v", count, tt.want)
			}
		})
	}
}

func TestSortLines_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Sort lines alphabetically",
		filterValue: "Sort Lines (sort-lines)",
		flags:       nil,
		name:        "sort-lines",
		title:       "Sort Lines (sort-lines)",
	}
	p := SortLines{}
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

func TestSortLines_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "String",
			args: args{data: []byte("Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***")},
			want: "%*&^*&^&\n***\nHello\nbar\nf00\nfoo\nfoo",
		},
		{
			name: "Numbers",
			args: args{data: []byte("3\n1\n6\n5\n9\n10\n4\n8\n7\n2")},
			want: "1\n10\n2\n3\n4\n5\n6\n7\n8\n9",
		},
		{
			name: "Numbers with trailing new-line",
			args: args{data: []byte("3\n1\n6\n5\n9\n10\n4\n8\n7\n2\n")},
			want: "1\n10\n2\n3\n4\n5\n6\n7\n8\n9",
		},
		{
			name: "Numbers and an empty line with trailing new-line",
			args: args{data: []byte("3\n1\n6\n5\n9\n10\n4\n8\n7\n2\n\n")},
			want: "\n1\n10\n2\n3\n4\n5\n6\n7\n8\n9",
		},
		{
			name: "Trailing new-line only",
			args: args{data: []byte("\n")},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SortLines{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("SortLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueLines_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Unique Lines",
		filterValue: "Unique Lines (unique-lines)",
		flags:       nil,
		name:        "unique-lines",
		title:       "Unique Lines (unique-lines)",
	}
	p := UniqueLines{}
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

func TestUniqueLines_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Unique lines with numbers",
			args: args{data: []byte("1\n1\n2\n2\n3\n3")},
			want: "1\n2\n3",
		},
		{
			name: "Unique lines with numbers and dupes",
			args: args{data: []byte("1\n1\n2\n2\n3\n3\n1\n1\n2\n2\n3\n3")},
			want: "1\n2\n3",
		},
		{
			name: "Unique lines with alpha",
			args: args{data: []byte("a\na\nb\nb\nc\nc")},
			want: "a\nb\nc",
		},
		{
			name: "Empty lines",
			args: args{data: []byte("\n")},
			want: "",
		},
		{
			name: "Two empty lines",
			args: args{data: []byte("a\n\n\nb")},
			want: "a\n\nb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We are going run test 100 to make sure that list order is preserved.
			for i := 0; i < 100; i++ {
				p := UniqueLines{}
				if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
					t.Errorf("UniqueLines() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestReverseLines_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Reverse Lines",
		filterValue: "Reverse Lines (reverse-lines)",
		flags:       nil,
		name:        "reverse-lines",
		title:       "Reverse Lines (reverse-lines)",
	}
	p := ReverseLines{}
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

func TestReverseLines_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Reverse lines with numbers",
			args: args{data: []byte("1\n2\n3\n4")},
			want: "4\n3\n2\n1",
		},
		{
			name: "Reverse lines with numbers and alpha",
			args: args{data: []byte("1\n2\ntest")},
			want: "test\n2\n1",
		},
		{
			name: "Empty input",
			args: args{data: []byte("")},
			want: "",
		},
		{
			name: "Single input",
			args: args{data: []byte("1")},
			want: "1",
		},
		{
			name: "Single input with new line",
			args: args{data: []byte("1\n")},
			want: "\n1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We are going run test 100 to make sure that list order is preserved.
			for i := 0; i < 100; i++ {
				p := ReverseLines{}
				if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
					t.Errorf("ReverseLines() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
