package processors

import "testing"

func TestBase32Encode_Transform(t *testing.T) {
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
			want: "ORUGKIDROVUWG2ZAMJZG653OEBTG66BANJ2W24DTEBXXMZLSEBQSA3DBPJ4SAZDPM4======",
		}, {
			name: "Emoji",
			args: args{data: "😃😇🙃🙂😉😌😙😗🇮🇳"},
			want: "6CPZRA7QT6MIP4E7TGB7BH4ZQLYJ7GEJ6CPZRDHQT6MJT4E7TCL7BH4HV3YJ7B5T",
		}, {
			name: "Multi line string",
			args: args{data: "123345\nabcd\n456\n123\nabc\n567\n7890"},
			want: "GEZDGMZUGUFGCYTDMQFDINJWBIYTEMYKMFRGGCRVGY3QUNZYHEYA====",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base32Encoding{}
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

func TestBase32Decode_Transform(t *testing.T) {
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
			args: args{data: "ORUGKIDROVUWG2ZAMJZG653OEBTG66BANJ2W24DTEBXXMZLSEBQSA3DBPJ4SAZDPM4======"},
			want: "the quick brown fox jumps over a lazy dog",
		}, {
			name: "Emoji",
			args: args{data: "6CPZRA7QT6MIP4E7TGB7BH4ZQLYJ7GEJ6CPZRDHQT6MJT4E7TCL7BH4HV3YJ7B5T"},
			want: "😃😇🙃🙂😉😌😙😗🇮🇳",
		}, {
			name: "Multi line string",
			args: args{data: "GEZDGMZUGUFGCYTDMQFDINJWBIYTEMYKMFRGGCRVGY3QUNZYHEYA===="},
			want: "123345\nabcd\n456\n123\nabc\n567\n7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Base32Decode{}
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
