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

func TestHexToString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{input: "7468697320697320737472696e67"},
			want: "this is string",
		}, {
			name: "Emoji",
			args: args{input: "f09f9883f09f9887f09f9983f09f9982f09f9889f09f988cf09f9899f09f9897f09f87aef09f87b3"},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{input: "48656c6c6f0a666f6f0a6261720a666f6f0a6630300a252a265e2a265e260a2a2a2a"},
			want: "Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HexToString(tt.args.input); got != tt.want {
				t.Errorf("HexToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToHex(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{input: "this is string"},
			want: "7468697320697320737472696e67",
		}, {
			name: "Emoji",
			args: args{input: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "f09f9883f09f9887f09f9983f09f9982f09f9889f09f988cf09f9899f09f9897f09f87aef09f87b3",
		}, {
			name: "Multi line string",
			args: args{input: "Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***"},
			want: "48656c6c6f0a666f6f0a6261720a666f6f0a6630300a252a265e2a265e260a2a2a2a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToHex(tt.args.input); got != tt.want {
				t.Errorf("StringToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
