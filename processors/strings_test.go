package processors

import (
	"reflect"
	"strings"
	"testing"
)

func TestCountCharacters_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
			want: "41",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "10",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "32",
		},
		{
			name: "Double-byte characters",
			args: args{data: []byte("你好")},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CountCharacters{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("CountCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortLines_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "String",
			args: args{data: []byte("Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***")},
			want: "%*&^*&^&\n***\nHello\nbar\nf00\nfoo\nfoo",
		},
		{
			name: "Numbers",
			args: args{data: []byte("3\n1\n6\n5\n9\n10\n4\n8\n7\n2")},
			want: "1\n10\n2\n3\n4\n5\n6\n7\n8\n9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SortLines{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("SortLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffleLines_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Shuffle lines randomly",
		filterValue: "Shuffle Lines",
		flags:       nil,
		name:        "shuffle-lines",
		title:       "Shuffle Lines",
	}
	p := ShuffleLines{}
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

func TestShuffleLines_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Should shuffle lines",
			args: args{data: []byte("1\n2")},
			want: 2,
		},
		{
			name: "Should return one line",
			args: args{data: []byte("there is no email in text")},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ShuffleLines{}

			got, _ := p.Transform(tt.args.data, tt.args.opts...)
			count := len(strings.Split(got, "\n"))
			if count != tt.want {
				t.Errorf("ExtractEmails() = %v, want %v", count, tt.want)
			}
		})
	}
}

func TestLower_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "String",
			args: args{data: []byte("THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***")},
			want: "the quick brown fox jumps over a lazy dog hello foo bar foo f00 %*&^*&^& ***",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***")},
			want: "hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Lower{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("StringToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTitle_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
			want: "The Quick Brown Fox Jumps Over A Lazy Dog",
		}, {
			name: "String Uppercase",
			args: args{data: []byte("THE QUICK BROWN FOX JUMPS OVER A LAZY DOG")},
			want: "THE QUICK BROWN FOX JUMPS OVER A LAZY DOG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Title{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("StringToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpper_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "String",
			args: args{data: []byte("THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***")},
			want: "THE QUICK BROWN FOX JUMPS OVER A LAZY DOG HELLO FOO BAR FOO F00 %*&^*&^& ***",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("Hello\nfoo\nbar\nfoo\nf00\n%*&^*&^&\n***")},
			want: "HELLO\nFOO\nBAR\nFOO\nF00\n%*&^*&^&\n***",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Upper{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("StringToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnakeCase_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "String",
			args: args{data: []byte("THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***")},
			want: "the_quick_brown_fox_jumps_over_a_lazy_dog_hello_foo_bar_foo_f_00_%*&^*&^&_***",
		},
		{
			name: "TestStringExample",
			args: args{data: []byte("test_string_example")},
			want: "test_string_example",
		},
		{
			name: "Lots of Space",
			args: args{data: []byte("Lots    Of      Space   ")},
			want: "lots_of_space",
		},
		{
			name: "Multi Line String",
			args: args{data: []byte("Multi\nLine\nString")},
			want: "multi_line_string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Snake{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("StringToSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKebab_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		//alias:       []string{"b64-enc", "b64-encode"},
		description: "Transform your text to kebab-case",
		filterValue: "To Kebab case",
		flags:       nil,
		name:        "kebab",
		title:       "To Kebab case",
	}
	p := Kebab{}
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

func TestKebab_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "String",
			args: args{data: []byte("THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***")},
			want: "the-quick-brown-fox-jumps-over-a-lazy-dog-hello-foo-bar-foo-f-00-%*&^*&^&-***",
		},
		{
			name: "TestStringExample",
			args: args{data: []byte("test_string_example")},
			want: "test-string-example",
		},
		{
			name: "Lots of Space",
			args: args{data: []byte("Lots    Of      Space   ")},
			want: "lots-of-space",
		},
		{
			name: "Multi Line String",
			args: args{data: []byte("Multi\nLine\nString")},
			want: "multi-line-string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Kebab{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("StringToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlug_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "hello world",
			args: args{data: []byte("hello world")},
			want: "hello-world",
		},
		{
			name: "hello_world",
			args: args{data: []byte("hello_world")},
			want: "hello-world",
		},
		{
			name: "Lots of Space",
			args: args{data: []byte("Lots    Of      Space   ")},
			want: "lots-of-space",
		},
		{
			name: "String",
			args: args{data: []byte("THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***")},
			want: "the-quick-brown-fox-jumps-over-a-lazy-dog-hello-foo-bar-foo-f00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Slug{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("StringToSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringReverse_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "hello world",
			args: args{data: []byte("hello world")},
			want: "dlrow olleh",
		},
		{
			name: "hello_world",
			args: args{data: []byte("hello_world")},
			want: "dlrow_olleh",
		},
		{
			name: "Lots of Space",
			args: args{data: []byte("Lots    Of      Space   ")},
			want: "   ecapS      fO    stoL",
		},
		{
			name: "String",
			args: args{data: []byte("THE quick brown fox jumps over a lazy dog Hello foo bar foo f00 %*&^*&^& ***")},
			want: "*** &^&*^&*% 00f oof rab oof olleH god yzal a revo spmuj xof nworb kciuq EHT",
		},
		{
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "0987\n765\ncba\n321\n654\ndcba\n543321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Reverse{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountWords_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "count number of words in string",
			args: args{data: []byte("hello world")},
			want: "2",
		},
		{
			name: "count number of words in string contains spaces",
			args: args{data: []byte(" This  is string having spaces?")},
			want: "5",
		},
		{
			name: "count number of words in comma separated string",
			args: args{data: []byte("word1, word2, word3")},
			want: "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CountWords{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("CountWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLines_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Count empty",
			args: args{data: []byte(nil)},
			want: "0",
		},
		{
			name: "Count one line",
			args: args{data: []byte("one line")},
			want: "1",
		},
		{
			name: "Count two line",
			args: args{data: []byte("1st line\n 2nd line")},
			want: "2",
		},
		{
			name: "Count empty line",
			args: args{data: []byte("\n\n\n")},
			want: "4",
		},
		{
			name: "Count empty + text line",
			args: args{data: []byte(`1
2
3
 `)},
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CountLines{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("CountLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToCamel(t *testing.T) {
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
			name: "String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
			want: "TheQuickBrownFoxJumpsOverALazyDog",
		}, {
			name: "String Uppercase",
			args: args{data: []byte("THE QUICK BROWN FOX JUMPS OVER A LAZY DOG")},
			want: "THEQUICKBROWNFOXJUMPSOVERALAZYDOG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Camel{}
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

func TestExtractEmails_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"find-emails", "find-email", "extract-email"},
		description: "Extract emails from given text",
		filterValue: "Extract Emails",
		flags: []Flag{
			{
				Name:  "separator",
				Short: "s",
				Desc:  "Separator to split multiple emails",
				Value: "",
				Type:  FlagString,
			},
		},
		name:  "extract-emails",
		title: "Extract Emails",
	}
	p := ExtractEmails{}
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

func TestExtractEmails_Transform(t *testing.T) {
	type args struct {
		data []byte
		opts []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Email in single line string",
			args: args{data: []byte("this is example@gmail.com")},
			want: "example@gmail.com",
		},
		{
			name: "Multiple Emails in single line string",
			args: args{data: []byte("this is example@gmail.com and this is example2@gmail.com")},
			want: "example@gmail.com\nexample2@gmail.com",
		},
		{
			name: "No email in text",
			args: args{data: []byte("there is no email in text")},
			want: "",
		},
		{
			name: "Fake emails",
			args: args{data: []byte("this is @fake.com email")},
			want: "",
		},
		{
			name: "Multiple Emails with separator flag",
			args: args{data: []byte("this is example@gmail.com and this is example2@gmail.com"), opts: []Flag{{Short: "s", Value: ","}}},
			want: "example@gmail.com,example2@gmail.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ExtractEmails{}
			if got, _ := p.Transform(tt.args.data, tt.args.opts...); got != tt.want {
				t.Errorf("ExtractEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
