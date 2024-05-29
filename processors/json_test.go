package processors

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func TestJSON_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       nil,
		description: "Format your text as JSON ( json decode )",
		filterValue: "Format JSON (json)",
		flags: []Flag{
			{
				Name:  "indent",
				Short: "i",
				Desc:  "Indent the output (prettyprint)",
				Value: false,
				Type:  FlagBool,
			},
		},
		name:  "json",
		title: "Format JSON (json)",
	}
	p := FormatJSON{}
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

func TestJSON_Transform(t *testing.T) {
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
			name:    "Should generate correct JSON for object input",
			args:    args{data: []byte(`{"name":"sttr"}`)},
			want:    `{"name":"sttr"}`,
			wantErr: false,
		},
		{
			name:    "Should generate correct JSON for array input",
			args:    args{data: []byte(`[{"name":"sttr"}]`)},
			want:    `[{"name":"sttr"}]`,
			wantErr: false,
		},
		{
			name:    "Should generate correct JSON having right space",
			args:    args{data: []byte(`{"name":"sttr"}       `)},
			want:    `{"name":"sttr"}`,
			wantErr: false,
		},
		{
			name:    "Should generate correct JSON having right space",
			args:    args{data: []byte(`       {"name":"sttr"}`)},
			want:    `{"name":"sttr"}`,
			wantErr: false,
		},
		{
			name:    "Should preserver order of object input",
			args:    args{data: []byte(`{"c":"c","b":"b","a":"a"}`)},
			want:    `{"c":"c","b":"b","a":"a"}`,
			wantErr: false,
		},
		{
			name:    "Should preserver order of array input",
			args:    args{data: []byte(`[{"c":"c","b":"b","a":"a"}]`)},
			want:    `[{"c":"c","b":"b","a":"a"}]`,
			wantErr: false,
		},
		{
			name: "Should preserver order of array input and indent on flag",
			args: args{
				data: []byte(`[{"c":"c","b":"b","a":"a"}]`),
				f: []Flag{
					{
						Name:  "indent",
						Short: "i",
						Desc:  "Indent the output (prettyprint)",
						Value: true,
						Type:  FlagBool,
					},
				},
			},
			want: `[
  {
    "c": "c",
    "b": "b",
    "a": "a"
  }
]`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := FormatJSON{}
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

func TestJSONToMSGPACK_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{},
		description: "Convert JSON to MSGPACK text",
		filterValue: "JSON To MSGPACK (json-msgpack)",
		flags:       nil,
		name:        "json-msgpack",
		title:       "JSON To MSGPACK (json-msgpack)",
	}
	p := JSONToMSGPACK{}
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

func TestJSONToMSGPACK_Transform(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		msgpack []byte
	}{
		{
			"Simple string",
			"\"Hello\"",
			[]byte{165, 72, 101, 108, 108, 111},
		},
		{
			"Map",
			"{\"id\":\"1\",\"user\":\"name\"}",
			[]byte{223, 0, 0, 0, 2, 162, 105, 100, 161, 49, 164, 117, 115, 101, 114, 164, 110, 97, 109, 101},
		},
		{
			"List",
			"{\"data\":[\"1\", \"2\", \"3\"]}",
			[]byte{
				223, 0, 0, 0, 1, 164, 100, 97, 116, 97, 221, 0, 0, 0, 3, 161, 49, 161, 50, 161,
				51,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := JSONToMSGPACK{}
			result, err := p.Transform([]byte(tt.json))
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			var resultInterface any
			err = msgpack.Unmarshal([]byte(result), &resultInterface)
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			var wantInterface any
			err = json.Unmarshal([]byte(tt.json), &wantInterface)
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			if !reflect.DeepEqual(wantInterface, resultInterface) {
				if err != nil {
					t.Errorf("Transform() got = %v, want %v", resultInterface, wantInterface)
				}
			}
		})
	}
}

func TestMSGPACKToJSON_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{},
		description: "Convert MSGPACK to JSON text",
		filterValue: "MSGPACK to JSON (msgpack-json)",
		flags:       nil,
		name:        "msgpack-json",
		title:       "MSGPACK to JSON (msgpack-json)",
	}
	p := MSGPACKToJSON{}
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

func TestMSGPACKToJSON_Transform(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		msgpack []byte
	}{
		{
			"Simple string",
			"\"Hello\"",
			[]byte{165, 72, 101, 108, 108, 111},
		},
		{
			"Map",
			"{\"id\":\"1\",\"user\":\"name\"}",
			[]byte{223, 0, 0, 0, 2, 162, 105, 100, 161, 49, 164, 117, 115, 101, 114, 164, 110, 97, 109, 101},
		},
		{
			"List",
			"{\"data\":[\"1\", \"2\", \"3\"]}",
			[]byte{
				223, 0, 0, 0, 1, 164, 100, 97, 116, 97, 221, 0, 0, 0, 3, 161, 49, 161, 50, 161,
				51,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MSGPACKToJSON{}
			result, err := p.Transform(tt.msgpack)
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			var resultInterface any
			err = json.Unmarshal([]byte(result), &resultInterface)
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			var wantInterface any
			err = msgpack.Unmarshal(tt.msgpack, &wantInterface)
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}

			if !reflect.DeepEqual(wantInterface, resultInterface) {
				if err != nil {
					t.Errorf("Transform() got = %v, want %v", resultInterface, wantInterface)
				}
			}
		})
	}
}

func TestJSONUnescape_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"json-unesc"},
		description: "JSON Unescape",
		filterValue: "JSON Unescape (json-unescape)",
		flags: []Flag{
			{
				Name:  "indent",
				Short: "i",
				Desc:  "Indent the output (prettyprint)",
				Value: false,
				Type:  FlagBool,
			},
		},
		name:  "json-unescape",
		title: "JSON Unescape (json-unescape)",
	}
	p := JSONUnescape{}
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

func TestJSONUnescape_Transform(t *testing.T) {
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
			name:    "Should generate correct JSON",
			args:    args{data: []byte(`{\n  \"name\": \"sttr\"\n}`)},
			want:    `{"name":"sttr"}`,
			wantErr: false,
		},
		{
			name: "Should generate correct JSON with indent",
			args: args{data: []byte(`{\n  \"name\": \"sttr\"\n}`), f: []Flag{
				{
					Short: "i",
					Value: true,
				},
			}},
			want: `{
  "name": "sttr"
}`,
			wantErr: false,
		},
		{
			name:    "Should generate correct JSON having right space",
			args:    args{data: []byte(`{\n  \"name\": \"sttr\"\n}       `)},
			want:    `{"name":"sttr"}`,
			wantErr: false,
		},
		{
			name:    "Should generate correct JSON having having let space",
			args:    args{data: []byte(`   {\n  \"name\": \"sttr\"\n}`)},
			want:    `{"name":"sttr"}`,
			wantErr: false,
		},
		{
			name:    "Should return error on invalid input",
			args:    args{data: []byte(`Invalid Input`)},
			want:    ``,
			wantErr: true,
		},
		{
			name:    "Should return error on invalid JSON",
			args:    args{data: []byte(`{\n  \"name\: \"name is missing quote\"\n}`)},
			want:    ``,
			wantErr: true,
		},
		{
			name: "Should preserver order of array input and indent on flag",
			args: args{
				data: []byte(`[{\"c\":\"c\",\"b\":\"b\",\"a\":\"a\"}]`),
				f: []Flag{
					{
						Name:  "indent",
						Short: "i",
						Desc:  "Indent the output (prettyprint)",
						Value: true,
						Type:  FlagBool,
					},
				},
			},
			want: `[
  {
    "c": "c",
    "b": "b",
    "a": "a"
  }
]`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := JSONUnescape{}
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

func TestJSONEscape_Command(t *testing.T) {
	test := struct {
		alias       []string
		description string
		filterValue string
		flags       []Flag
		name        string
		title       string
	}{
		alias:       []string{"json-esc"},
		description: "JSON Escape",
		filterValue: "JSON Escape (json-escape)",
		flags:       nil,
		name:        "json-escape",
		title:       "JSON Escape (json-escape)",
	}
	p := JSONEscape{}
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

func TestJSONEscape_Transform(t *testing.T) {
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
			name:    "Should generate correct JSON",
			args:    args{data: []byte(`{"name":"sttr"}`)},
			want:    `{\"name\":\"sttr\"}`,
			wantErr: false,
		},
		{
			name:    "Should generate correct JSON having right space",
			args:    args{data: []byte(`{"name":"sttr"}       `)},
			want:    `{\"name\":\"sttr\"}`,
			wantErr: false,
		},
		{
			name:    "Should generate correct JSON having having let space",
			args:    args{data: []byte(`   {"name":"sttr"}`)},
			want:    `{\"name\":\"sttr\"}`,
			wantErr: false,
		},
		{
			name:    "Should return error on invalid input",
			args:    args{data: []byte(`Invalid Input`)},
			want:    ``,
			wantErr: true,
		},
		{
			name:    "Should return error on invalid JSON",
			args:    args{data: []byte(`{\n  \"name\: \"name is missing quote\"\n}`)},
			want:    ``,
			wantErr: true,
		},
		{
			name: "Should preserver order of array input and indent on flag",
			args: args{
				data: []byte(`[{"c":"c","b":"b","a":"a"}]`),
			},
			want:    `[{\"c\":\"c\",\"b\":\"b\",\"a\":\"a\"}]`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := JSONEscape{}
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
