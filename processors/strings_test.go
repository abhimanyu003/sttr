package processors

import "testing"

func TestCountNumberCharacters(t *testing.T) {
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
			args: args{input: "the quick brown fox jumps over a lazy dog"},
			want: "41",
		}, {
			name: "Emoji",
			args: args{input: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "40",
		}, {
			name: "Multi line string",
			args: args{input: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "32",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountNumberCharacters(tt.args.input); got != tt.want {
				t.Errorf("CountNumberCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortLines(t *testing.T) {
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
			args: args{input: "Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***"},
			want: "%*&^*&^&\n***\nHello\nbar\nf00\nfoo\nfoo",
		},
		{
			name: "Numbers",
			args: args{input: "3\n1\n6\n5\n9\n10\n4\n8\n7\n2"},
			want: "1\n10\n2\n3\n4\n5\n6\n7\n8\n9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortLines(tt.args.input); got != tt.want {
				t.Errorf("SortLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToLower(t *testing.T) {
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
			args: args{input: "THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***"},
			want: "the quick brown fox jumps over a lazy dog hello foo bar foo f00 %*&^*&^& ***",
		}, {
			name: "Emoji",
			args: args{input: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{input: "Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***"},
			want: "hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToLower(tt.args.input); got != tt.want {
				t.Errorf("StringToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToTitle(t *testing.T) {
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
			args: args{input: "the quick brown fox jumps over a lazy dog"},
			want: "The Quick Brown Fox Jumps Over A Lazy Dog",
		}, {
			name: "String Uppercase",
			args: args{input: "THE QUICK BROWN FOX JUMPS OVER A LAZY DOG"},
			want: "THE QUICK BROWN FOX JUMPS OVER A LAZY DOG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToTitle(tt.args.input); got != tt.want {
				t.Errorf("StringToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToUpper(t *testing.T) {
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
			args: args{input: "THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***"},
			want: "THE QUICK BROWN FOX JUMPS OVER A LAZY DOG HELLO FOO BAR FOO F00 %*&^*&^& ***",
		}, {
			name: "Emoji",
			args: args{input: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{input: "Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***"},
			want: "HELLO\nFOO\nBAR\nFOO\nF00\n%*&^*&^&\n***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUpper(tt.args.input); got != tt.want {
				t.Errorf("StringToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToSnakeCase(t *testing.T) {
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
			args: args{input: "THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***"},
			want: "the_quick_brown_fox_jumps_over_a_lazy_dog_hello_foo_bar_foo_f_00_%*&^*&^&_***",
		},
		{
			name: "TestStringExample",
			args: args{input: "test_string_example"},
			want: "test_string_example",
		},
		{
			name: "Lots of Space",
			args: args{input: "Lots    Of      Space   "},
			want: "lots_of_space",
		},
		{
			name: "Multi Line String",
			args: args{input: "Multi\nLine\nString"},
			want: "multi_line_string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToSnakeCase(tt.args.input); got != tt.want {
				t.Errorf("StringToSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToKebab(t *testing.T) {
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
			args: args{input: "THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***"},
			want: "the-quick-brown-fox-jumps-over-a-lazy-dog-hello-foo-bar-foo-f-00-%*&^*&^&-***",
		},
		{
			name: "TestStringExample",
			args: args{input: "test_string_example"},
			want: "test-string-example",
		},
		{
			name: "Lots of Space",
			args: args{input: "Lots    Of      Space   "},
			want: "lots-of-space",
		},
		{
			name: "Multi Line String",
			args: args{input: "Multi\nLine\nString"},
			want: "multi-line-string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToKebab(tt.args.input); got != tt.want {
				t.Errorf("StringToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToSlug(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello world",
			args: args{input: "hello world"},
			want: "hello-world",
		},
		{
			name: "hello_world",
			args: args{input: "hello_world"},
			want: "hello-world",
		},
		{
			name: "Lots of Space",
			args: args{input: "Lots    Of      Space   "},
			want: "lots-of-space",
		},
		{
			name: "String",
			args: args{input: "THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***"},
			want: "the-quick-brown-fox-jumps-over-a-lazy-dog-hello-foo-bar-foo-f00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToSlug(tt.args.input); got != tt.want {
				t.Errorf("StringToSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringReverse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello world",
			args: args{input: "hello world"},
			want: "dlrow olleh",
		},
		{
			name: "hello_world",
			args: args{input: "hello_world"},
			want: "dlrow_olleh",
		},
		{
			name: "Lots of Space",
			args: args{input: "Lots    Of      Space   "},
			want: "   ecapS      fO    stoL",
		},
		{
			name: "String",
			args: args{input: "THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***"},
			want: "*** &^&*^&*% 00f oof rab oof olleH god yzal a revo spmuj xof nworb kciuq EHT",
		},
		{
			name: "Multi line string",
			args: args{input: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "0987\n765\ncba\n321\n654\ndcba\n543321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := StringReverse(tt.args.input); gotResult != tt.want {
				t.Errorf("StringReverse() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
