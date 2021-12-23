package processors

import (
	"reflect"
	"testing"
)

func TestExtractEmails_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"find-emails", "find-email", "extract-email"},
		description: "Extract emails from given text",
		filterValue: "Extract Emails",
		flags: []Flag{
			{
				Name:  "separator",
				Short: "s",
				Desc:  "Separator to split multiple emails",
				Value: "",
				Type:  FlagString,
			},
		},
		name:  "extract-emails",
		title: "Extract Emails",
	}
	p := ExtractEmails{}
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

func TestExtractEmails_Transform(t *testing.T) {
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
			name: "Email in single line string",
			args: args{data: []byte("this is example@gmail.com")},
			want: "example@gmail.com",
		},
		{
			name: "Multiple Emails in single line string",
			args: args{data: []byte("this is example@gmail.com and this is example2@gmail.com")},
			want: "example@gmail.com\nexample2@gmail.com",
		},
		{
			name: "No email in text",
			args: args{data: []byte("there is no email in text")},
			want: "",
		},
		{
			name: "Fake emails",
			args: args{data: []byte("this is @fake.com email")},
			want: "",
		},
		{
			name: "Multiple Emails with separator flag",
			args: args{data: []byte("this is example@gmail.com and this is example2@gmail.com"), opts: []Flag{{Short: "s", Value: ","}}},
			want: "example@gmail.com,example2@gmail.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ExtractEmails{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("ExtractEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
