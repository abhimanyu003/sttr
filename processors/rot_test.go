package processors

import "testing"

func TestROT13Encode_Transform(t *testing.T) {
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
			name: "String",
			args: args{input: "the quick brown fox jumps over a lazy dog"},
			want: "gur dhvpx oebja sbk wh`cf bire n ynml qbt",
		}, {
			name: "String Uppercase",
			args: args{input: "THE QUICK BROWN FOX JUMPS OVER A LAZY DOG"},
			want: "GUR DHVPX OEBJA SBK WH@CF BIRE N YNML QBT",
		}, {
			name: "Emoji",
			args: args{input: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{input: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "123345\nnopq\n456\n123\nnop\n567\n7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ROT13Encode{}
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
