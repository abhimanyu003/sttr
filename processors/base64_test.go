package processors

import (
	"reflect"
	"testing"
)

func TestBase64Encode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"b64-enc", "b64-encode"},
		description: "Encode your text to Base64",
		filterValue: "Base64 Encoding (base64-encode)",
		flags:       []Flag{base64RawFlag},
		name:        "base64-encode",
		title:       "Base64 Encoding (base64-encode)",
	}
	p := Base64Encode{}
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

func TestBase64Encode_Transform(t *testing.T) {
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
			want: "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIGEgbGF6eSBkb2c=",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳")},
			want: "8J+Yg/CfmIfwn5mD8J+ZgvCfmInwn5iM8J+YmfCfmJfwn4eu8J+Hsw==",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890")},
			want: "MTIzMzQ1CmFiY2QKNDU2CjEyMwphYmMKNTY3Cjc4OTA=",
		}, {
			name: "Test For baser64 standard Encoding",
			args: args{data: []byte("�")},
			want: "77+9",
		},
		{
			name: "String",
			args: args{data: []byte("the quick brown fox jumps over a lazy dog"), in1: []Flag{{Short: "r", Value: true}}},
			want: "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIGEgbGF6eSBkb2c",
		}, {
			name: "Emoji",
			args: args{data: []byte("😃😇🙃🙂😉😌😙😗🇮🇳"), in1: []Flag{{Short: "r", Value: true}}},
			want: "8J+Yg/CfmIfwn5mD8J+ZgvCfmInwn5iM8J+YmfCfmJfwn4eu8J+Hsw",
		}, {
			name: "Multi line string",
			args: args{data: []byte("123345\nabcd\n456\n123\nabc\n567\n7890"), in1: []Flag{{Short: "r", Value: true}}},
			want: "MTIzMzQ1CmFiY2QKNDU2CjEyMwphYmMKNTY3Cjc4OTA",
		}, {
			name: "Test For baser64 standard Encoding",
			args: args{data: []byte("�"), in1: []Flag{{Short: "r", Value: true}}},
			want: "77+9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base64Encode{}
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

func TestBase64Decode_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"b64-dec", "b64-decode"},
		description: "Decode your base64 text",
		filterValue: "Base64 Decode (base64-decode)",
		flags:       []Flag{base64RawFlag},
		name:        "base64-decode",
		title:       "Base64 Decode (base64-decode)",
	}
	p := Base64Decode{}
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

func TestBase64Decode_Transform(t *testing.T) {
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
			args: args{data: []byte("dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIGEgbGF6eSBkb2c=")},
			want: "the quick brown fox jumps over a lazy dog",
		}, {
			name: "Emoji",
			args: args{data: []byte("8J+Yg/CfmIfwn5mD8J+ZgvCfmInwn5iM8J+YmfCfmJfwn4eu8J+Hsw==")},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("MTIzMzQ1CmFiY2QKNDU2CjEyMwphYmMKNTY3Cjc4OTA=")},
			want: "123345\nabcd\n456\n123\nabc\n567\n7890",
		},
		{
			name: "Test baser64 standard decode",
			args: args{data: []byte("77+9")},
			want: "�",
		},
		{
			name: "String",
			args: args{data: []byte("dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIGEgbGF6eSBkb2c"), in1: []Flag{{Short: "r", Value: true}}},
			want: "the quick brown fox jumps over a lazy dog",
		}, {
			name: "Emoji",
			args: args{data: []byte("8J+Yg/CfmIfwn5mD8J+ZgvCfmInwn5iM8J+YmfCfmJfwn4eu8J+Hsw"), in1: []Flag{{Short: "r", Value: true}}},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: []byte("MTIzMzQ1CmFiY2QKNDU2CjEyMwphYmMKNTY3Cjc4OTA"), in1: []Flag{{Short: "r", Value: true}}},
			want: "123345\nabcd\n456\n123\nabc\n567\n7890",
		},
		{
			name: "Test baser64 standard decode",
			args: args{data: []byte("77+9")},
			want: "�",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base64Decode{}
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
