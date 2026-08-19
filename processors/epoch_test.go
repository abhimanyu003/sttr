package processors

import (
	"reflect"
	"testing"
)

func TestEpochToDate_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"epoch", "epoch-human", "from-epoch"},
		description: "Convert a Unix timestamp to a human-readable date",
		filterValue: "Epoch to Date (epoch-date)",
		flags: []Flag{
			{Name: "format", Short: "f", Desc: "Output layout: a named layout (rfc3339, datetime, date, ...) or a Go time layout", Type: FlagString, Value: ""},
			{Name: "unit", Short: "u", Desc: "Input unit: auto, s, ms, us or ns", Type: FlagString, Value: "auto"},
			{Name: "local", Short: "l", Desc: "Render the date in the local timezone instead of UTC", Type: FlagBool, Value: false},
		},
		name:  "epoch-date",
		title: "Epoch to Date (epoch-date)",
	}
	p := EpochToDate{}
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

func TestEpochToDate_Transform(t *testing.T) {
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
			name: "Seconds auto-detected",
			args: args{data: []byte("1710152465")},
			want: "2024-03-11T10:21:05Z",
		},
		{
			name: "Milliseconds auto-detected",
			args: args{data: []byte("1710152465311")},
			want: "2024-03-11T10:21:05Z",
		},
		{
			name: "Microseconds auto-detected",
			args: args{data: []byte("1710152465311000")},
			want: "2024-03-11T10:21:05Z",
		},
		{
			name: "Nanoseconds auto-detected",
			args: args{data: []byte("1710152465311000000")},
			want: "2024-03-11T10:21:05Z",
		},
		{
			name: "Surrounding whitespace is trimmed",
			args: args{data: []byte("  1710152465\n")},
			want: "2024-03-11T10:21:05Z",
		},
		{
			name: "Unix epoch zero",
			args: args{data: []byte("0")},
			want: "1970-01-01T00:00:00Z",
		},
		{
			name: "Negative timestamp before 1970",
			args: args{data: []byte("-1")},
			want: "1969-12-31T23:59:59Z",
		},
		{
			name: "Explicit unit overrides detection",
			args: args{data: []byte("1710152465"), in1: []Flag{{Short: "u", Value: "ms"}}},
			want: "1970-01-20T19:02:32Z",
		},
		{
			name: "Custom named layout",
			args: args{data: []byte("1710152465"), in1: []Flag{{Short: "f", Value: "date"}}},
			want: "2024-03-11",
		},
		{
			name: "Custom Go layout",
			args: args{data: []byte("1710152465"), in1: []Flag{{Short: "f", Value: "2006-01-02 15:04:05"}}},
			want: "2024-03-11 10:21:05",
		},
		{
			name:    "Non-numeric input errors",
			args:    args{data: []byte("not-a-number")},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Unknown unit errors",
			args:    args{data: []byte("1710152465"), in1: []Flag{{Short: "u", Value: "weeks"}}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := EpochToDate{}
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

func TestDateToEpoch_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		name        string
		title       string
	}{
		alias:       []string{"to-epoch", "human-epoch", "date-to-epoch"},
		description: "Convert a human-readable date to a Unix timestamp",
		filterValue: "Date to Epoch (date-epoch)",
		name:        "date-epoch",
		title:       "Date to Epoch (date-epoch)",
	}
	p := DateToEpoch{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
	}
	if got := p.Name(); got != test.name {
		t.Errorf("Name() = %v, want %v", got, test.name)
	}
	if got := p.Title(); got != test.title {
		t.Errorf("Title() = %v, want %v", got, test.title)
	}
}

func TestDateToEpoch_Transform(t *testing.T) {
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
			name: "RFC3339 with zone",
			args: args{data: []byte("2024-03-11T10:21:05Z")},
			want: "1710152465",
		},
		{
			name: "Date and time without zone treated as UTC",
			args: args{data: []byte("2024-03-11 10:21:05")},
			want: "1710152465",
		},
		{
			name: "Date only",
			args: args{data: []byte("2024-03-11")},
			want: "1710115200",
		},
		{
			name: "Slash separated date time",
			args: args{data: []byte("2024/03/11 10:21:05")},
			want: "1710152465",
		},
		{
			name: "Milliseconds output unit",
			args: args{data: []byte("2024-03-11T10:21:05Z"), in1: []Flag{{Short: "u", Value: "ms"}}},
			want: "1710152465000",
		},
		{
			name: "Explicit format flag",
			args: args{data: []byte("11-03-2024"), in1: []Flag{{Short: "f", Value: "02-01-2006"}}},
			want: "1710115200",
		},
		{
			name: "Round trips with epoch-date",
			args: args{data: []byte("2024-03-11T10:21:05Z")},
			want: "1710152465",
		},
		{
			name:    "Unparseable date errors",
			args:    args{data: []byte("sometime last tuesday")},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Unknown unit errors",
			args:    args{data: []byte("2024-03-11T10:21:05Z"), in1: []Flag{{Short: "u", Value: "fortnights"}}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := DateToEpoch{}
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
