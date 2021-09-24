package processors

import "testing"

func TestHexToRGB_Transform(t *testing.T) {
	type args struct {
		input string
		in1   []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Hex with # string",
			args: args{input: "#FF5733"},
			want: "255, 87, 51",
		},
		{
			name:    "HEX string with wrong string",
			args:    args{input: "#PPPPP"},
			want:    "0, 0, 0",
			wantErr: true,
		},
		{
			name:    "HEX string with wrong string",
			args:    args{input: "FF5733"},
			want:    "0, 0, 0",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := HexToRGB{}
			got, err := p.Transform(tt.args.input, tt.args.in1...)
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
