package processors

import "testing"

func TestHTMLEncode_Transform(t *testing.T) {
	type args struct {
		input string
		in1   []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should escape HTML",
			args: args{input: `<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

</body>
</html>`},
			want: `&lt;!DOCTYPE html&gt;
&lt;html&gt;
&lt;body&gt;

&lt;h1&gt;My First Heading&lt;/h1&gt;

&lt;p&gt;My first paragraph.&lt;/p&gt;

&lt;/body&gt;
&lt;/html&gt;`,
			wantErr: false,
		},
		{
			name:    "should escape xss string",
			args:    args{input: `<script>alert("XSS");</script>`},
			want:    `&lt;script&gt;alert(&#34;XSS&#34;);&lt;/script&gt;`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := HTMLEncode{}
			got, err := p.Transform(tt.args.input, tt.args.in1...)
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

func TestHTMLDecode_Transform(t *testing.T) {
	type args struct {
		input string
		in1   []Flag
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should unescape HTML",
			args: args{input: `&lt;!DOCTYPE html&gt;
&lt;html&gt;
&lt;body&gt;

&lt;h1&gt;My First Heading&lt;/h1&gt;

&lt;p&gt;My first paragraph.&lt;/p&gt;

&lt;/body&gt;
&lt;/html&gt;`},
			want: `<!DOCTYPE html>
<html>
<body>

<h1>My First Heading</h1>

<p>My first paragraph.</p>

</body>
</html>`,
			wantErr: false,
		},
		{
			name:    "should unescape xss string",
			args:    args{input: `&lt;script&gt;alert(&#34;XSS&#34;);&lt;/script&gt;`},
			want:    `<script>alert("XSS");</script>`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := HTMLDecode{}
			got, err := p.Transform(tt.args.input, tt.args.in1...)
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
