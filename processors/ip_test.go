package processors

import (
	"reflect"
	"testing"
)

func TestExtractIPs_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"find-ips", "find-ip", "extract-ips"},
		description: "Extract IPv4 and IPv6 from your text",
		filterValue: "Extract IPs",
		flags:       nil,
		name:        "extract-ip",
		title:       "Extract IPs",
	}
	p := ExtractIPs{}
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

func TestExtractIPs_Transform(t *testing.T) {
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
			name: "Test for IPv6",
			args: args{data: []byte("Test for IPv4 185.141.205.123")},
			want: "185.141.205.123",
		},
		{
			name: "Test for IPv6",
			args: args{data: []byte("Test for IPv6 bb62:9bb8:46e2:640:e3a6:33b5:670a:a74")},
			want: "bb62:9bb8:46e2:640:e3a6:33b5:670a:a74",
		},
		{
			name: "Test for IPv4 and IPv6",
			args: args{data: []byte("IPv4 = 185.141.205.123 IPv6 = bb62:9bb8:46e2:640:e3a6:33b5:670a:a74")},
			want: "185.141.205.123\nbb62:9bb8:46e2:640:e3a6:33b5:670a:a74",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ExtractIPs{}
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
