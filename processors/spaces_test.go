package processors

import (
	"reflect"
	"testing"
)

func TestRemoveNewLines_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"remove-new-lines", "trim-newlines", "trim-new-lines"},
		description: "Remove all new lines",
		filterValue: "Remove all new lines (remove-newlines)",
		flags: []Flag{
			{
				Name:  "separator",
				Short: "s",
				Desc:  "Separator to split multiple lines",
				Value: "",
				Type:  FlagString,
			},
		},
		name:  "remove-newlines",
		title: "Remove all new lines (remove-newlines)",
	}
	p := RemoveNewLines{}
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

func TestRemoveNewLines_Transform(t *testing.T) {
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
			name: "Remove new lines",
			args: args{data: []byte("1\n2\n3\n4")},
			want: "1 2 3 4",
		},
		{
			name: "Remove multiple new lines",
			args: args{data: []byte("1\n\n2\n\n3\n\n4")},
			want: "1 2 3 4",
		},
		{
			name: "Remove return carriage lines",
			args: args{data: []byte("1\r\r2\r\r3\r\r4")},
			want: "1 2 3 4",
		},
		{
			name: "Remove return carriage and new lines",
			args: args{data: []byte("1\r\n2\r\n3\r\n4")},
			want: "1 2 3 4",
		},
		{
			name: "Remove newlines with separator",
			args: args{data: []byte("1\r\n2\r\n3\r\n4"), opts: []Flag{{Short: "s", Value: ","}}},
			want: "1,2,3,4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RemoveNewLines{}
			got, err := p.Transform(tt.args.data, tt.args.opts...)
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

func TestRemoveSpaces_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"remove-space", "trim-spaces", "trim-space"},
		description: "Remove all spaces + new lines",
		filterValue: "Remove all spaces + new lines (remove-spaces)",
		flags: []Flag{
			{
				Name:  "separator",
				Short: "s",
				Desc:  "Separator to split spaces",
				Value: "",
				Type:  FlagString,
			},
		},
		name:  "remove-spaces",
		title: "Remove all spaces + new lines (remove-spaces)",
	}
	p := RemoveSpaces{}
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

func TestRemoveSpaces_Transform(t *testing.T) {
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
			name: "Remove spaes",
			args: args{data: []byte("1 2 3 4")},
			want: "1234",
		},
		{
			name: "Remove new lines",
			args: args{data: []byte("1\n2\n3\n4")},
			want: "1234",
		},
		{
			name: "Remove multiple new lines",
			args: args{data: []byte("1\n\n2\n\n3\n\n4")},
			want: "1234",
		},
		{
			name: "Remove return carriage lines",
			args: args{data: []byte("1\r\r2\r\r3\r\r4")},
			want: "1234",
		},
		{
			name: "Remove return carriage and new lines",
			args: args{data: []byte("1\r\n2\r\n3\r\n4")},
			want: "1234",
		},
		{
			name: "Remove newlines with separator",
			args: args{data: []byte("1\r\n2\r\n3\r\n4"), opts: []Flag{{Short: "s", Value: ","}}},
			want: "1,2,3,4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := RemoveSpaces{}
			got, err := p.Transform(tt.args.data, tt.args.opts...)
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
