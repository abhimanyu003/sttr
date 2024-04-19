package processors

import (
	"reflect"
	"testing"
)

func TestZeropad_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		//alias:       []string{"md-html"},
		description: "Pad a number with zeros",
		filterValue: "Zeropad (zeropad)",
		flags: []Flag{
			{
				Name:  "number-of-zeros",
				Short: "n",
				Desc:  "Number of zeros to be padded",
				Value: 5,
				Type:  FlagUint,
			},
			{
				Name:  "prefix",
				Short: "p",
				Desc:  "The number get prefixed with this",
				Value: "",
				Type:  FlagString,
			},
		},
		name:  "zeropad",
		title: "Zeropad (zeropad)",
	}
	p := Zeropad{}
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

func TestZeropad_Transform(t *testing.T) {
	type args struct {
		input string
		f     []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should return 5 zero's on UNIT 5",
			args: args{input: "12", f: []Flag{{Short: "n", Value: uint(5)}}},
			want: "0000012",
		},
		{
			name: "Testing negative number",
			args: args{input: "-12", f: []Flag{{Short: "n", Value: uint(5)}}},
			want: "-0000012",
		},
		{
			name: "Should return no zero's INT 5",
			args: args{input: "12", f: []Flag{{Short: "n", Value: 5}}},
			want: "12",
		},
		{
			name: "Should return no zero's -1",
			args: args{input: "12", f: []Flag{{Short: "n", Value: -1}}},
			want: "12",
		},
		{
			name: "Should return 5 with prefix zero's on UNIT 5",
			args: args{input: "12", f: []Flag{
				{Short: "n", Value: uint(5)},
				{Short: "p", Value: "A"},
			}},
			want: "A0000012",
		},
		{
			name: "Testing negative number with prefix",
			args: args{input: "-12", f: []Flag{
				{Short: "n", Value: uint(5)},
				{Short: "p", Value: "A"},
			}},
			want: "A-0000012",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Zeropad{}
			got, err := p.Transform([]byte(tt.args.input), tt.args.f...)
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
