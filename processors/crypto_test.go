package processors

import "testing"

func TestMD5Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{data: "the quick brown fox jumps over a lazy dog"},
			want: "f3e7b7426c27a59d36cf9fe9db5a1a1b",
		}, {
			name: "Emoji",
			args: args{data: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "028990045a146ff3abbc0ba9ab772d84",
		}, {
			name: "Multi line string",
			args: args{data: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "4da14107f15ce261f2641ab7b8769466",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5Encode(tt.args.data); got != tt.want {
				t.Errorf("MD5Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA1Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{data: "the quick brown fox jumps over a lazy dog"},
			want: "dcab639dcb4bf4d577d396758ebf91b6939e732a",
		}, {
			name: "Emoji",
			args: args{data: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "c07a5913366f3fa980e2520089e1642d28edc116",
		}, {
			name: "Multi line string",
			args: args{data: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "6dfc74a3e7472a8ca5794962b62b144a82808e2c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1Encode(tt.args.data); got != tt.want {
				t.Errorf("SHA1Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA256Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{data: "the quick brown fox jumps over a lazy dog"},
			want: "883238e6c74b0b4838738c5117bef8660fb9207877603fbfd7fe5c8ab9e579a1",
		}, {
			name: "Emoji",
			args: args{data: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "29298957699f889a8dd88d6239ba927c01cca0dbdcccf2730cfc7533ed633f21",
		}, {
			name: "Multi line string",
			args: args{data: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "a5deaf4214a6cf73c20bfb6df0a0ec1d0ade6b3fe7d1845592e49d06082fa039",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256Encode(tt.args.data); got != tt.want {
				t.Errorf("SHA256Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA512Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String",
			args: args{data: "the quick brown fox jumps over a lazy dog"},
			want: "db9bf1e67167b9bd6573386cc212f3e0ad3f701f0c2e9779d0b752062bf38e62c205a3c02816b92ef3c4f9004f793ea9b92d99813134535ddc9cfde970f8131c",
		}, {
			name: "Emoji",
			args: args{data: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "e34072a7584c345d5baf2296a9e966b86329e8bee04a546f265f96f23e09152a9aedce87d36b7ef2859273d10eaa99ecac6261997c19b0d7858284aaa1e58056",
		}, {
			name: "Multi line string",
			args: args{data: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "aa53744b761ea00e61737ff65bee640519c21ce1850898a9dfd285057bba9a0cf2a9ba512dcdc1f5c8f6df0666336249495153b3875fa74f32b5e612f858f553",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA512Encode(tt.args.data); got != tt.want {
				t.Errorf("SHA512Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
