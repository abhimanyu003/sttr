package processors

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"reflect"
	"testing"
)

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
		filterValue: "JSON To MSGPACK",
		flags:       nil,
		name:        "json-msgpack",
		title:       "JSON To MSGPACK",
	}
	p := JSONToMSGPACK{}
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
			[]byte{223, 0, 0, 0, 1, 164, 100, 97, 116, 97, 221, 0, 0, 0, 3, 161, 49, 161, 50, 161,
				51},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := JSONToMSGPACK{}
			result, err := p.Transform([]byte(tt.json))
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			var resultInterface interface{}
			err = msgpack.Unmarshal([]byte(result), &resultInterface)
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			var wantInterface interface{}
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
		filterValue: "MSGPACK To JSON",
		flags:       nil,
		name:        "msgpack-json",
		title:       "MSGPACK To JSON",
	}
	p := MSGPACKToJSON{}
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
			[]byte{223, 0, 0, 0, 1, 164, 100, 97, 116, 97, 221, 0, 0, 0, 3, 161, 49, 161, 50, 161,
				51},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MSGPACKToJSON{}
			result, err := p.Transform(tt.msgpack)
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			var resultInterface interface{}
			err = json.Unmarshal([]byte(result), &resultInterface)
			if err != nil {
				t.Errorf("Transform() error = %v, wantErr %v", err, nil)
			}
			var wantInterface interface{}
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
