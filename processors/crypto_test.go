package processors

import (
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
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
		alias:       []string{"md5-sum"},
		description: "Get the MD5 checksum of your text",
		filterValue: "MD5 Sum (md5)",
		flags:       nil,
		name:        "md5",
		title:       "MD5 Sum (md5)",
	}
	p := MD5{}
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
			want: "f3e7b7426c27a59d36cf9fe9db5a1a1b",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "028990045a146ff3abbc0ba9ab772d84",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "4da14107f15ce261f2641ab7b8769466",
		}, {
			name: "empty rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("")},
			want: "d41d8cd98f00b204e9800998ecf8427e",
		}, {
			name: "a rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("a")},
			want: "0cc175b9c0f1b6a831c399e269772661",
		}, {
			name: "abc rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("abc")},
			want: "900150983cd24fb0d6963f7d28e17f72",
		}, {
			name: "message digest rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("message digest")},
			want: "f96b697d7cb7938d525a2f31aaf161d0",
		}, {
			name: "abcdefghijklmnopqrstuvwxyz rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("abcdefghijklmnopqrstuvwxyz")},
			want: "c3fcd3d76192e4007dfb496cca67e13b",
		}, {
			name: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")},
			want: "d174ab98d277d9f5a5611c2c9f419d9f",
		}, {
			name: "12345678901234567890123456789012345678901234567890123456789012345678901234567890 rfc1321", // test values from https://datatracker.ietf.org/doc/html/rfc1321
			args: args{data: []byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890")},
			want: "57edf4a22be3c955ac49da2e2107b67a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MD5{}
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
		alias:       []string{"sha1-sum"},
		description: "Get the SHA1 checksum of your text",
		filterValue: "SHA1 Sum (sha1)",
		flags:       nil,
		name:        "sha1",
		title:       "SHA1 Sum (sha1)",
	}
	p := SHA1{}
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
			want: "dcab639dcb4bf4d577d396758ebf91b6939e732a",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "c07a5913366f3fa980e2520089e1642d28edc116",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "6dfc74a3e7472a8ca5794962b62b144a82808e2c",
		}, {
			name: "Trimmed linebreak",
			args: args{data: []byte("Hello World\n")},
			want: "648a6a6ffffdaa0badb23b8baf90b6168dd16b3a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SHA1{}
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
		alias:       []string{"sha256-sum"},
		description: "Get the SHA256 checksum of your text",
		filterValue: "SHA256 Sum (sha256)",
		flags:       nil,
		name:        "sha256",
		title:       "SHA256 Sum (sha256)",
	}
	p := SHA256{}
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
			want: "883238e6c74b0b4838738c5117bef8660fb9207877603fbfd7fe5c8ab9e579a1",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "29298957699f889a8dd88d6239ba927c01cca0dbdcccf2730cfc7533ed633f21",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "a5deaf4214a6cf73c20bfb6df0a0ec1d0ade6b3fe7d1845592e49d06082fa039",
		}, {
			name: "Trimmed linebreak",
			args: args{data: []byte("Hello World\n")},
			want: "d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SHA256{}
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

func TestSHA224Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"SHA224-sum"},
		description: "Get the SHA224 checksum of your text",
		filterValue: "SHA224 Sum (SHA224)",
		flags:       nil,
		name:        "SHA224",
		title:       "SHA224 Sum (SHA224)",
	}
	p := SHA224{}
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

func TestSHA224Encode_Transform(t *testing.T) {
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
			want: "4aad218ea5ec38461e574a78c59bf0eae149fce2d546bc31c5c46f83",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "b50e56ecce3ce8eaafc7e73b4dabd71a27ea22eee27507bc35b1eb81",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "ae876c55086958d4733fad2908a83c8f45a7f53bd460bbbcf571043b",
		}, {
			name: "Trimmed linebreak",
			args: args{data: []byte("Hello World\n")},
			want: "e53ee97e5e0a2a4d359b5b461409dc44d9315afbc3b7d6bc5cd598e6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SHA224{}
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

func TestSHA384Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"SHA384-sum"},
		description: "Get the SHA384 checksum of your text",
		filterValue: "SHA384 Sum (SHA384)",
		flags:       nil,
		name:        "SHA384",
		title:       "SHA384 Sum (SHA384)",
	}
	p := SHA384{}
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

func TestSHA384Encode_Transform(t *testing.T) {
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
			want: "40c4474cc2ce8d43fcea7ae23974078b970d984bed88ac5920499b3b3634f49e4fc1e689434c423508138d5d0972561b",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "8557db557fc52550ff148ffded2b3c39726aafe6159431aee9844561a83ec83ac5872cd9bdac9deba46e7f6d75d6d5c8",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "71bd688b9c316981c5fa511dbdded308325cff9b90b85b94eb261ab0ac83041714caec6a5751d27f53c0ee8a12135eb5",
		}, {
			name: "Trimmed linebreak",
			args: args{data: []byte("Hello World\n")},
			want: "acbfd470c22c0d95a1d10a087dc31988b9f7bfeb13be70b876a73558be664e5858d11f9459923e6e5fd838cb5708b969",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SHA384{}
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
		alias:       []string{"sha512-sum"},
		description: "Get the SHA512 checksum of your text",
		filterValue: "SHA512 Sum (sha512)",
		flags:       nil,
		name:        "sha512",
		title:       "SHA512 Sum (sha512)",
	}
	p := SHA512{}
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
			want: "db9bf1e67167b9bd6573386cc212f3e0ad3f701f0c2e9779d0b752062bf38e62c205a3c02816b92ef3c4f9004f793ea9b92d99813134535ddc9cfde970f8131c",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "e34072a7584c345d5baf2296a9e966b86329e8bee04a546f265f96f23e09152a9aedce87d36b7ef2859273d10eaa99ecac6261997c19b0d7858284aaa1e58056",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "aa53744b761ea00e61737ff65bee640519c21ce1850898a9dfd285057bba9a0cf2a9ba512dcdc1f5c8f6df0666336249495153b3875fa74f32b5e612f858f553",
		}, {
			name: "Trimmed linebreak",
			args: args{data: []byte("Hello World\n")},
			want: "e1c112ff908febc3b98b1693a6cd3564eaf8e5e6ca629d084d9f0eba99247cacdd72e369ff8941397c2807409ff66be64be908da17ad7b8a49a2a26c0e8086aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SHA512{}
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

func TestBcrypt_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"bcrypt-hash"},
		description: "Get the Bcrypt hash of your text",
		filterValue: "Bcrypt Hash (bcrypt)",
		flags: []Flag{
			{
				Name:  "number-of-rounds",
				Short: "r",
				Desc:  "Number of rounds",
				Value: 10,
				Type:  FlagUint,
			},
		},
		name:  "bcrypt",
		title: "Bcrypt Hash (bcrypt)",
	}
	p := Bcrypt{}
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

func TestBcrypt_Transform(t *testing.T) {
	type args struct {
		data []byte
		in1  []Flag
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		rounds  uint
	}{
		{
			name: "String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog")},
		},
		{
			name: "String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog"), in1: []Flag{{Short: "r", Value: 12}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Bcrypt{}
			got, err := p.Transform(tt.args.data, tt.args.in1...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = bcrypt.CompareHashAndPassword([]byte(got), tt.args.data)
			if err != nil {
				t.Errorf("Bcrypt validation failed")
			}
		})
	}
}
