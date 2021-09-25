package processors

import "testing"

func TestBase64Encode_Transform(t *testing.T) {
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
			want: "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIGEgbGF6eSBkb2c=",
		}, {
			name: "Emoji",
			args: args{data: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "8J+Yg/CfmIfwn5mD8J+ZgvCfmInwn5iM8J+YmfCfmJfwn4eu8J+Hsw==",
		}, {
			name: "Multi line string",
			args: args{data: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "MTIzMzQ1CmFiY2QKNDU2CjEyMwphYmMKNTY3Cjc4OTA=",
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

func TestBase64Decode_Transform(t *testing.T) {
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
			args: args{data: "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIGEgbGF6eSBkb2c="},
			want: "the quick brown fox jumps over a lazy dog",
		}, {
			name: "Emoji",
			args: args{data: "8J+Yg/CfmIfwn5mD8J+ZgvCfmInwn5iM8J+YmfCfmJfwn4eu8J+Hsw=="},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: "MTIzMzQ1CmFiY2QKNDU2CjEyMwphYmMKNTY3Cjc4OTA="},
			want: "123345\nabcd\n456\n123\nabc\n567\n7890",
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
