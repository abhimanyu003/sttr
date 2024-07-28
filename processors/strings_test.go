package processors

import (
	"reflect"
	"testing"
)

func TestLower_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Transform your text to lower case",
		filterValue: "To Lower case (lower)",
		flags:       nil,
		name:        "lower",
		title:       "To Lower case (lower)",
	}
	p := Lower{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
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

func TestUpper_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Transform your text to UPPER CASE",
		filterValue: "To Upper case (upper)",
		flags:       nil,
		name:        "upper",
		title:       "To Upper case (upper)",
	}
	p := Upper{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
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

func TestCountCharacters_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Find the length of your text (including spaces)",
		filterValue: "Count Number of Characters (count-chars)",
		flags:       nil,
		name:        "count-chars",
		title:       "Count Number of Characters (count-chars)",
	}
	p := CountCharacters{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
			want: "41",
		},
		{
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "10",
		},
		{
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

func TestTitle_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Transform your text to Title Case",
		filterValue: "To Title Case (title)",
		flags:       nil,
		name:        "title",
		title:       "To Title Case (title)",
	}
	p := Title{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
			want: "The Quick Brown Fox Jumps Over A Lazy Dog",
		},
		{
			name: "String Uppercase",
			args: args{data: []byte("THE QUICK BROWN FOX JUMPS OVER A LAZY DOG")},
			want: "THE QUICK BROWN FOX JUMPS OVER A LAZY DOG",
		},
		{
			name: "Camel Case Text",
			args: args{data: []byte("camelCaseText")},
			want: "CamelCaseText",
		},
		{
			name: "Underscore text",
			args: args{data: []byte("underscore_text")},
			want: "Underscore_text",
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

func TestSnake_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Transform your text to snake_case",
		filterValue: "To Snake case (snake)",
		flags:       nil,
		name:        "snake",
		title:       "To Snake case (snake)",
	}
	p := Snake{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
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
		// alias:       []string{"b64-enc", "b64-encode"},
		description: "Transform your text to kebab-case",
		filterValue: "To Kebab case (kebab)",
		flags:       nil,
		name:        "kebab",
		title:       "To Kebab case (kebab)",
	}
	p := Kebab{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
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

func TestSlug_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Transform your text to slug-case",
		filterValue: "To Slug case (slug)",
		flags:       nil,
		name:        "slug",
		title:       "To Slug case (slug)",
	}
	p := Slug{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
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

func TestReverse_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Reverse Text ( txeT esreveR )",
		filterValue: "Reverse text (reverse)",
		flags:       nil,
		name:        "reverse",
		title:       "Reverse text (reverse)",
	}
	p := Reverse{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
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

func TestCountWords_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Count the number of words in your text",
		filterValue: "Count Number of Words (count-words)",
		flags:       nil,
		name:        "count-words",
		title:       "Count Number of Words (count-words)",
	}
	p := CountWords{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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

func TestCamel_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Transform your text to camelCase",
		filterValue: "To Camel case (camel)",
		flags:       nil,
		name:        "camel",
		title:       "To Camel case (camel)",
	}
	p := Camel{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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
			name: "Normal String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
			want: "theQuickBrownFoxJumpsOverALazyDog",
		},
		{
			name: "String Uppercase",
			args: args{data: []byte("THE QUICK BROWN FOX JUMPS OVER A LAZY DOG")},
			want: "theQuickBrownFoxJumpsOverALazyDog",
		},
		{
			name: "Camel Case Text",
			args: args{data: []byte("camelCaseText")},
			want: "camelCaseText", // stable
		},
		{
			name: "Pascal Case Text",
			args: args{data: []byte("PascalCaseText")},
			want: "pascalCaseText",
		},
		{
			name: "Underscore text lowercase",
			args: args{data: []byte("underscore_text")},
			want: "underscoreText",
		},
		{
			name: "Underscore text uppercase",
			args: args{data: []byte("UNDERSCORE_TEXT_UPPER_CASE")},
			want: "underscoreTextUpperCase",
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

func TestPascal_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Transform your text to PascalCase",
		filterValue: "To Pascal case (pascal)",
		flags:       nil,
		name:        "pascal",
		title:       "To Pascal case (pascal)",
	}
	p := Pascal{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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

func TestStringToPascal(t *testing.T) {
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
			name: "Normal String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
			want: "TheQuickBrownFoxJumpsOverALazyDog",
		},
		{
			name: "String Uppercase",
			args: args{data: []byte("THE QUICK BROWN FOX JUMPS OVER A LAZY DOG")},
			want: "TheQuickBrownFoxJumpsOverALazyDog",
		},
		{
			name: "Camel Case Text",
			args: args{data: []byte("camelCaseText")},
			want: "CamelCaseText",
		},
		{
			name: "Pacal Case Text",
			args: args{data: []byte("PacalCaseText")},
			want: "PacalCaseText", // stable
		},
		{
			name: "Underscore text lowercase",
			args: args{data: []byte("underscore_text")},
			want: "UnderscoreText",
		},
		{
			name: "Underscore text uppercase",
			args: args{data: []byte("UNDERSCORE_TEXT_UPPER_CASE")},
			want: "UnderscoreTextUpperCase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pascal{}
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

func TestEscapeQuotes_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"esc-quotes", "escape-quotes"},
		description: "Escapes single and double quotes by default",
		filterValue: "Escape Quotes (escape-quotes)",
		flags: []Flag{
			{
				Name:  "double-quote",
				Short: "d",
				Desc:  "Escape double quote",
				Value: true,
				Type:  FlagBool,
			},
			{
				Name:  "single-quote",
				Short: "s",
				Desc:  "Escape single quote",
				Value: true,
				Type:  FlagBool,
			},
		},
		name:  "escape-quotes",
		title: "Escape Quotes (escape-quotes)",
	}
	p := EscapeQuotes{}
	if got := p.Alias(); !reflect.DeepEqual(got, test.alias) {
		t.Errorf("Alias() = %v, want %v", got, test.alias)
	}
	if got := p.Description(); got != test.description {
		t.Errorf("Description() = %v, want %v", got, test.description)
	}
	if got := p.FilterValue(); got != test.filterValue {
		t.Errorf("FilterValue() = %v, want %v", got, test.filterValue)
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

func TestEscapeQuotes(t *testing.T) {
	type args struct {
		data []byte
		f    []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Normal String double quote",
			args: args{data: []byte("this is great \"test\"")},
			want: "this is great \\\"test\\\"",
		},
		{
			name: "Normal String single quote",
			args: args{data: []byte("this is great 'test'")},
			want: "this is great \\'test\\'",
		},
		{
			name: "Both single and double quote",
			args: args{data: []byte("this is 'great' \"test\"")},
			want: "this is \\'great\\' \\\"test\\\"",
		},
		{
			name: "Normal String double quote",
			args: args{data: []byte("this is great \"test\"")},
			want: "this is great \\\"test\\\"",
		},
		{
			name: "Normal String double quote",
			args: args{data: []byte("this is great \"test\""), f: []Flag{{Short: "d", Value: true}}},
			want: "this is great \\\"test\\\"",
		},
		{
			name: "Normal String single quote",
			args: args{data: []byte("this is great 'test'"), f: []Flag{{Short: "s", Value: true}}},
			want: "this is great \\'test\\'",
		},
		{
			name: "Both single and double quote",
			args: args{data: []byte("this is 'great' \"test\""), f: []Flag{{Short: "d", Value: true}, {Short: "s", Value: true}}},
			want: "this is \\'great\\' \\\"test\\\"",
		},
		{
			name: "String with no quote",
			args: args{data: []byte("this is great test")},
			want: "this is great test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := EscapeQuotes{}
			got, err := p.Transform(tt.args.data, tt.args.f...)
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
