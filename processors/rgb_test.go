package processors

import (
	"reflect"
	"testing"
)

func TestRGB_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		//alias:       []string{},
		description: "Convert a #hex-color code to RGB",
		filterValue: "Hex To RGB (hex-rgb)",
		flags:       nil,
		name:        "hex-rgb",
		title:       "Hex To RGB (hex-rgb)",
	}
	p := HexToRGB{}
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

func TestHexToRGB_Transform(t *testing.T) {
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
			name: "Hex with # string",
			args: args{data: []byte("#FF5733")},
			want: "255, 87, 51",
		},
		{
			name:    "HEX string with wrong string",
			args:    args{data: []byte("#PPPPP")},
			want:    "0, 0, 0",
			wantErr: true,
		},
		{
			name:    "HEX string with wrong string",
			args:    args{data: []byte("FF5733")},
			want:    "0, 0, 0",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := HexToRGB{}
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
