package processors

import "testing"

func TestHexToRGB(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Hex with # string",
			args: args{input: "#FF5733"},
			want: "255, 87, 51",
		},
		{
			name: "HEX string with wrong string",
			args: args{input: "#PPPPP"},
			want: "0, 0, 0",
		},
		{
			name: "HEX string with wrong string",
			args: args{input: "FF5733"},
			want: "0, 0, 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToRGB(tt.args.input); got != tt.want {
				t.Errorf("HexToRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}
