package processors

import "testing"

func TestURLDecode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Special characters",
			args: args{input: "http%3A%2F%2Fexample.com%2Ftest%3Ftest%3Dsomething%26value%3D%2A%26%5E%25%24%23%40%40%23%24%25%5E%26%2A%28"},
			want: "http://example.com/test?test=something&value=*&^%$#@@#$%^&*(",
		},
		{
			name: "Special characters",
			args: args{input: "+%3F%26%3D%23%2B%25%21%3C%3E%23%5C%22%7B%7D%7C%5C%5C%5E%5B%5D%60%E2%98%BA%5Ct%3A%2F%40%24%27%28%29%2A%2C%3B"},
			want: " ?&=#+%!<>#\\\"{}|\\\\^[]`☺\\t:/@$'()*,;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLDecode(tt.args.input); got != tt.want {
				t.Errorf("URLDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLEncode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Special characters",
			args: args{input: "http://example.com/test?test=something&value=*&^%$#@@#$%^&*("},
			want: "http%3A%2F%2Fexample.com%2Ftest%3Ftest%3Dsomething%26value%3D%2A%26%5E%25%24%23%40%40%23%24%25%5E%26%2A%28",
		},
		{
			name: "Special characters",
			args: args{input: " ?&=#+%!<>#\\\"{}|\\\\^[]`☺\\t:/@$'()*,;"},
			want: "+%3F%26%3D%23%2B%25%21%3C%3E%23%5C%22%7B%7D%7C%5C%5C%5E%5B%5D%60%E2%98%BA%5Ct%3A%2F%40%24%27%28%29%2A%2C%3B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLEncode(tt.args.input); got != tt.want {
				t.Errorf("URLEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
