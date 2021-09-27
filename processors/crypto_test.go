package processors

import (
	"reflect"
	"testing"
)

func TestMD5Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"md5"},
		description: "Get the MD5 checksum of your text",
		filterValue: "MD5 Sum",
		flags:       nil,
		name:        "md5-sum",
		title:       "MD5 Sum",
	}
	p := MD5Encode{}
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

func TestMD5Encode_Transform(t *testing.T) {
	type args struct {
		data string
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
		}, {
			name: "Trimmed linebreak",
			args: args{data: "Hello World\n"},
			want: "b10a8db164e0754105b7a99be72e3fe5",
		}, {
			name: "With linebreak (from file)",
			args: args{data: "Hello World\n", in1: []Flag{ {Short: FlagFile, Value: true} }},
			want: "e59ff97941044f85df5297e1c302d260",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MD5Encode{}
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

func TestSHA1Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"sha1"},
		description: "Get the SHA1 checksum of your text",
		filterValue: "SHA1 Sum",
		flags:       nil,
		name:        "sha1-sum",
		title:       "SHA1 Sum",
	}
	p := SHA1Encode{}
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

func TestSHA1Encode_Transform(t *testing.T) {
	type args struct {
		data string
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
		}, {
			name: "Trimmed linebreak",
			args: args{data: "Hello World\n"},
			want: "0a4d55a8d778e5022fab701977c5d840bbc486d0",
		}, {
			name: "With linebreak (from file)",
			args: args{data: "Hello World\n", in1: []Flag{ {Short: FlagFile, Value: true} }},
			want: "648a6a6ffffdaa0badb23b8baf90b6168dd16b3a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SHA1Encode{}
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

func TestSHA256Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"sha256"},
		description: "Get the SHA256 checksum of your text",
		filterValue: "SHA256 Sum",
		flags:       nil,
		name:        "sha256-sum",
		title:       "SHA256 Sum",
	}
	p := SHA256Encode{}
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

func TestSHA256Encode_Transform(t *testing.T) {
	type args struct {
		data string
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
		}, {
			name: "Trimmed linebreak",
			args: args{data: "Hello World\n"},
			want: "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e",
		}, {
			name: "With linebreak (from file)",
			args: args{data: "Hello World\n", in1: []Flag{ {Short: FlagFile, Value: true} }},
			want: "d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SHA256Encode{}
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

func TestSHA512Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"sha512"},
		description: "Get the SHA512 checksum of your text",
		filterValue: "SHA512 Sum",
		flags:       nil,
		name:        "sha512-sum",
		title:       "SHA512 Sum",
	}
	p := SHA512Encode{}
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

func TestSHA512Encode_Transform(t *testing.T) {
	type args struct {
		data string
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
		}, {
			name: "Trimmed linebreak",
			args: args{data: "Hello World\n"},
			want: "2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b",
		}, {
			name: "With linebreak (from file)",
			args: args{data: "Hello World\n", in1: []Flag{ {Short: FlagFile, Value: true} }},
			want: "e1c112ff908febc3b98b1693a6cd3564eaf8e5e6ca629d084d9f0eba99247cacdd72e369ff8941397c2807409ff66be64be908da17ad7b8a49a2a26c0e8086aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SHA512Encode{}
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
