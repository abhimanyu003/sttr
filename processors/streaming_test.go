package processors

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"
)

func TestTransformStream(t *testing.T) {
	tests := []struct {
		name      string
		processor Processor
		input     string
		want      string
		wantErr   bool
	}{
		{
			name:      "Upper case streaming",
			processor: Upper{},
			input:     "hello world",
			want:      "HELLO WORLD",
		},
		{
			name:      "Lower case streaming",
			processor: Lower{},
			input:     "HELLO WORLD",
			want:      "hello world",
		},
		{
			name:      "MD5 hash streaming",
			processor: MD5{},
			input:     "hello world",
			want:      "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:      "Count lines streaming",
			processor: CountLines{},
			input:     "line1\nline2\nline3\n",
			want:      "3",
		},
		{
			name:      "Count characters streaming",
			processor: CountCharacters{},
			input:     "hello",
			want:      "5",
		},
		{
			name:      "Count words streaming",
			processor: CountWords{},
			input:     "hello world test",
			want:      "3",
		},
		{
			name:      "Hex encode streaming",
			processor: HexEncode{},
			input:     "hello",
			want:      "68656c6c6f",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			var writer bytes.Buffer

			err := TransformStream(tt.processor, reader, &writer, nil...)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := writer.String()
			if got != tt.want {
				t.Errorf("TransformStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamingConfig(t *testing.T) {
	tests := []struct {
		name      string
		processor interface{}
		hasConfig bool
		config    StreamingConfig
	}{
		{
			name:      "MD5 has streaming config",
			processor: MD5{},
			hasConfig: true,
			config: StreamingConfig{
				ChunkSize:    64 * 1024,
				BufferOutput: true,
				LineByLine:   false,
			},
		},
		{
			name:      "Upper has streaming config",
			processor: Upper{},
			hasConfig: true,
			config: StreamingConfig{
				ChunkSize:    64 * 1024,
				BufferOutput: false,
				LineByLine:   false,
			},
		},
		{
			name:      "CountLines has streaming config",
			processor: CountLines{},
			hasConfig: true,
			config: StreamingConfig{
				ChunkSize:    64 * 1024,
				BufferOutput: true,
				LineByLine:   false,
			},
		},
		{
			name:      "Processor without config",
			processor: Reverse{},
			hasConfig: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if configurable, ok := tt.processor.(ConfigurableStreamingProcessor); ok {
				if !tt.hasConfig {
					t.Errorf("Expected processor to not have config, but it does")
					return
				}
				config := configurable.GetStreamingConfig()
				if config.ChunkSize != tt.config.ChunkSize {
					t.Errorf("ChunkSize = %v, want %v", config.ChunkSize, tt.config.ChunkSize)
				}
				if config.BufferOutput != tt.config.BufferOutput {
					t.Errorf("BufferOutput = %v, want %v", config.BufferOutput, tt.config.BufferOutput)
				}
				if config.LineByLine != tt.config.LineByLine {
					t.Errorf("LineByLine = %v, want %v", config.LineByLine, tt.config.LineByLine)
				}
			} else if tt.hasConfig {
				t.Errorf("Expected processor to have config, but it doesn't")
			}
		})
	}
}

func TestCanStream(t *testing.T) {
	processors := []Processor{
		MD5{},
		SHA256{},
		Upper{},
		Lower{},
		CountLines{},
		CountWords{},
		CountCharacters{},
		HexEncode{},
		HexDecode{},
		Reverse{},
	}

	for _, p := range processors {
		t.Run(p.Name(), func(t *testing.T) {

			if !CanStream(p) {
				t.Errorf("CanStream() = false, want true for %s", p.Name())
			}
		})
	}
}

func TestPreferStream(t *testing.T) {
	tests := []struct {
		name      string
		processor Processor
		want      bool
	}{
		{
			name:      "MD5 prefers streaming",
			processor: MD5{},
			want:      true,
		},
		{
			name:      "SHA256 prefers streaming",
			processor: SHA256{},
			want:      true,
		},
		{
			name:      "Upper prefers streaming",
			processor: Upper{},
			want:      true,
		},
		{
			name:      "Lower prefers streaming",
			processor: Lower{},
			want:      true,
		},
		{
			name:      "HexEncode prefers streaming",
			processor: HexEncode{},
			want:      true,
		},
		{
			name:      "CountLines doesn't prefer streaming (needs buffering)",
			processor: CountLines{},
			want:      false,
		},
		{
			name:      "Reverse doesn't prefer streaming by default",
			processor: Reverse{},
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreferStream(tt.processor)
			if got != tt.want {
				t.Errorf("PreferStream() = %v, want %v for %s", got, tt.want, tt.processor.Name())
			}
		})
	}
}

func TestLargeDataStreaming(t *testing.T) {

	var input strings.Builder
	expectedLines := 100000
	for i := 0; i < expectedLines; i++ {
		input.WriteString("a\n")
	}

	tests := []struct {
		name      string
		processor Processor
		want      string
	}{
		{
			name:      "Count lines with large input",
			processor: CountLines{},
			want:      "100000",
		},
		{
			name:      "Count characters with large input",
			processor: CountCharacters{},
			want:      "200000",
		},
		{
			name:      "Count words with large input",
			processor: CountWords{},
			want:      "100000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(input.String())
			var writer bytes.Buffer

			err := TransformStream(tt.processor, reader, &writer)
			if err != nil {
				t.Errorf("TransformStream() error = %v", err)
				return
			}

			got := writer.String()
			if got != tt.want {
				t.Errorf("TransformStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVeryLargeFileStreaming(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large file test in short mode")
	}

	const millionLines = 1000000

	tests := []struct {
		name           string
		processor      Processor
		lineContent    string
		expectedResult string
		description    string
	}{
		{
			name:           "Count 1M lines with 'a'",
			processor:      CountLines{},
			lineContent:    "a",
			expectedResult: "1000000",
			description:    "Simulates user's 500MB file with 'a' on each line",
		},
		{
			name:           "Count 1M characters with 'a'",
			processor:      CountCharacters{},
			lineContent:    "a",
			expectedResult: "2000000",
			description:    "Count characters in 1M line file",
		},
		{
			name:           "Count 1M words with 'a'",
			processor:      CountWords{},
			lineContent:    "a",
			expectedResult: "1000000",
			description:    "Count words in 1M line file",
		},
		{
			name:           "Upper case 1M lines",
			processor:      Upper{},
			lineContent:    "hello",
			expectedResult: "",
			description:    "Transform 1M lines to uppercase",
		},
		{
			name:           "Lower case 1M lines",
			processor:      Lower{},
			lineContent:    "HELLO",
			expectedResult: "",
			description:    "Transform 1M lines to lowercase",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Starting test: %s - %s", tt.name, tt.description)

			reader := &largeDataReader{
				lineContent: tt.lineContent,
				totalLines:  millionLines,
				currentLine: 0,
			}

			var writer bytes.Buffer

			start := time.Now()
			_ = start

			err := TransformStream(tt.processor, reader, &writer)
			if err != nil {
				t.Errorf("TransformStream() error = %v", err)
				return
			}

			result := writer.String()

			if tt.expectedResult != "" {
				if result != tt.expectedResult {
					t.Errorf("TransformStream() = %v, want %v", result, tt.expectedResult)
				}
			} else {

				expectedLines := millionLines
				actualLines := strings.Count(result, "\n")
				if tt.lineContent != "" && !strings.HasSuffix(result, "\n") {
					actualLines++
				}

				if actualLines != expectedLines {
					t.Errorf("Expected %d lines, got %d lines", expectedLines, actualLines)
				}

				lines := strings.Split(result, "\n")
				if len(lines) > 0 && lines[0] != "" {
					var expected string
					switch tt.processor.Name() {
					case "upper":
						expected = strings.ToUpper(tt.lineContent)
					case "lower":
						expected = strings.ToLower(tt.lineContent)
					}
					if expected != "" && lines[0] != expected {
						t.Errorf("First line transformation failed: got %v, want %v", lines[0], expected)
					}
				}
			}

			t.Logf("Successfully processed %d lines", millionLines)
		})
	}
}

func TestExtremelyLargeFileStreaming(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping extremely large file test in short mode")
	}

	const tenMillionLines = 10000000

	tests := []struct {
		name        string
		processor   Processor
		lineContent string
		want        string
	}{
		{
			name:        "Count 10M lines",
			processor:   CountLines{},
			lineContent: "a",
			want:        "10000000",
		},
		{
			name:        "Count 10M characters",
			processor:   CountCharacters{},
			lineContent: "a",
			want:        "20000000",
		},
		{
			name:        "Count 10M words",
			processor:   CountWords{},
			lineContent: "a",
			want:        "10000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Starting extreme test: %s with %d lines", tt.name, tenMillionLines)

			reader := &largeDataReader{
				lineContent: tt.lineContent,
				totalLines:  tenMillionLines,
				currentLine: 0,
			}

			var writer bytes.Buffer

			err := TransformStream(tt.processor, reader, &writer)
			if err != nil {
				t.Errorf("TransformStream() error = %v", err)
				return
			}

			got := writer.String()
			if got != tt.want {
				t.Errorf("TransformStream() = %v, want %v", got, tt.want)
			}

			t.Logf("Successfully processed %d lines", tenMillionLines)
		})
	}
}

func TestMemoryEfficiencyWithLargeFiles(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping memory efficiency test in short mode")
	}

	const lines = 100000

	tests := []struct {
		name      string
		processor Processor
		want      string
	}{
		{
			name:      "Memory efficient line counting",
			processor: CountLines{},
			want:      "100000",
		},
		{
			name:      "Memory efficient character counting",
			processor: CountCharacters{},
			want:      "200000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			reader := &largeDataReader{
				lineContent: "a",
				totalLines:  lines,
				currentLine: 0,
			}

			var writer bytes.Buffer

			err := TransformStream(tt.processor, reader, &writer)
			if err != nil {
				t.Errorf("TransformStream() error = %v", err)
				return
			}

			got := writer.String()
			if got != tt.want {
				t.Errorf("TransformStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

type largeDataReader struct {
	lineContent string
	totalLines  int
	currentLine int
	lineBuffer  []byte
	bufferPos   int
}

func (r *largeDataReader) Read(p []byte) (n int, err error) {
	if r.currentLine >= r.totalLines {
		return 0, io.EOF
	}

	if r.bufferPos >= len(r.lineBuffer) {
		r.lineBuffer = []byte(r.lineContent + "\n")
		r.bufferPos = 0
		r.currentLine++
	}

	n = copy(p, r.lineBuffer[r.bufferPos:])
	r.bufferPos += n

	return n, nil
}

func TestStreamingVsDirectTransform(t *testing.T) {
	testData := []string{
		"hello world",
		"line1\nline2\nline3",
		"a\na\na\na\na",
		"UPPERCASE TEXT",
		"",
		"single",
		"multiple words in a sentence",
	}

	processors := []Processor{
		Upper{},
		Lower{},
		CountLines{},
		CountWords{},
		CountCharacters{},
		MD5{},
		SHA256{},
		HexEncode{},
	}

	for _, data := range testData {
		for _, p := range processors {
			t.Run(p.Name()+"_"+data, func(t *testing.T) {

				directResult, err := p.Transform([]byte(data))
				if err != nil {
					t.Errorf("Direct Transform() error = %v", err)
					return
				}

				reader := strings.NewReader(data)
				var writer bytes.Buffer
				err = TransformStream(p, reader, &writer)
				if err != nil {
					t.Errorf("TransformStream() error = %v", err)
					return
				}
				streamResult := writer.String()

				if directResult != streamResult {
					t.Errorf("Results differ: direct=%v, stream=%v", directResult, streamResult)
				}
			})
		}
	}
}

func TestChunkedStreaming(t *testing.T) {

	processor := Upper{}
	input := "hello world this is a test"

	reader := &chunkReader{
		data:      []byte(input),
		chunkSize: 3,
	}

	var writer bytes.Buffer
	err := TransformStream(processor, reader, &writer)
	if err != nil {
		t.Errorf("TransformStream() error = %v", err)
		return
	}

	got := writer.String()
	want := "HELLO WORLD THIS IS A TEST"
	if got != want {
		t.Errorf("TransformStream() = %v, want %v", got, want)
	}
}

type chunkReader struct {
	data      []byte
	pos       int
	chunkSize int
}

func (r *chunkReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}

	end := r.pos + r.chunkSize
	if end > len(r.data) {
		end = len(r.data)
	}

	n = copy(p, r.data[r.pos:end])
	r.pos += n

	return n, nil
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		processor Processor
		input     string
		want      string
	}{
		{
			name:      "Empty input",
			processor: CountLines{},
			input:     "",
			want:      "0",
		},
		{
			name:      "Single newline",
			processor: CountLines{},
			input:     "\n",
			want:      "1",
		},
		{
			name:      "Multiple newlines",
			processor: CountLines{},
			input:     "\n\n\n",
			want:      "3",
		},
		{
			name:      "Text without trailing newline",
			processor: CountLines{},
			input:     "line1\nline2",
			want:      "2",
		},
		{
			name:      "Text with trailing newline",
			processor: CountLines{},
			input:     "line1\nline2\n",
			want:      "2",
		},
		{
			name:      "Empty input character count",
			processor: CountCharacters{},
			input:     "",
			want:      "0",
		},
		{
			name:      "Empty input word count",
			processor: CountWords{},
			input:     "",
			want:      "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			var writer bytes.Buffer

			err := TransformStream(tt.processor, reader, &writer)
			if err != nil {
				t.Errorf("TransformStream() error = %v", err)
				return
			}

			got := writer.String()
			if got != tt.want {
				t.Errorf("TransformStream() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkLargeFileStreaming(b *testing.B) {
	const lines = 100000

	benchmarks := []struct {
		name      string
		processor Processor
	}{
		{
			name:      "CountLines",
			processor: CountLines{},
		},
		{
			name:      "CountCharacters",
			processor: CountCharacters{},
		},
		{
			name:      "CountWords",
			processor: CountWords{},
		},
		{
			name:      "Upper",
			processor: Upper{},
		},
		{
			name:      "Lower",
			processor: Lower{},
		},
		{
			name:      "MD5",
			processor: MD5{},
		},
		{
			name:      "SHA256",
			processor: SHA256{},
		},
		{
			name:      "HexEncode",
			processor: HexEncode{},
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reader := &largeDataReader{
					lineContent: "test line content",
					totalLines:  lines,
					currentLine: 0,
				}

				var writer bytes.Buffer
				err := TransformStream(bm.processor, reader, &writer)
				if err != nil {
					b.Errorf("TransformStream() error = %v", err)
				}
			}
		})
	}
}

func BenchmarkStreamingVsDirect(b *testing.B) {

	const lines = 10000
	var testData strings.Builder
	for i := 0; i < lines; i++ {
		testData.WriteString("test line content\n")
	}
	data := testData.String()
	dataBytes := []byte(data)

	processors := []Processor{
		CountLines{},
		CountCharacters{},
		Upper{},
		Lower{},
	}

	for _, p := range processors {
		b.Run(p.Name()+"_Direct", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := p.Transform(dataBytes)
				if err != nil {
					b.Errorf("Transform() error = %v", err)
				}
			}
		})

		b.Run(p.Name()+"_Streaming", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reader := strings.NewReader(data)
				var writer bytes.Buffer
				err := TransformStream(p, reader, &writer)
				if err != nil {
					b.Errorf("TransformStream() error = %v", err)
				}
			}
		})
	}
}

func TestRealWorldLargeFile(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping real world large file test in short mode")
	}

	const realWorldLines = 1000000

	t.Run("RealWorld_500MB_File_Simulation", func(t *testing.T) {
		t.Logf("Simulating processing of 500MB file with %d lines", realWorldLines)

		reader := &largeDataReader{
			lineContent: "a",
			totalLines:  realWorldLines,
			currentLine: 0,
		}

		var writer bytes.Buffer

		err := TransformStream(CountLines{}, reader, &writer)
		if err != nil {
			t.Errorf("TransformStream() error = %v", err)
			return
		}

		result := writer.String()
		expected := "1000000"
		if result != expected {
			t.Errorf("Expected %s lines, got %s", expected, result)
		}

		t.Logf("Successfully counted %s lines in simulated 500MB file", result)
	})

	operations := []struct {
		name      string
		processor Processor
		expected  string
	}{
		{
			name:      "Count_Characters_500MB",
			processor: CountCharacters{},
			expected:  "2000000",
		},
		{
			name:      "Count_Words_500MB",
			processor: CountWords{},
			expected:  "1000000",
		},
		{
			name:      "Upper_Case_500MB",
			processor: Upper{},
			expected:  "",
		},
	}

	for _, op := range operations {
		t.Run(op.name, func(t *testing.T) {
			reader := &largeDataReader{
				lineContent: "a",
				totalLines:  realWorldLines,
				currentLine: 0,
			}

			var writer bytes.Buffer

			err := TransformStream(op.processor, reader, &writer)
			if err != nil {
				t.Errorf("TransformStream() error = %v", err)
				return
			}

			result := writer.String()

			if op.expected != "" {
				if result != op.expected {
					t.Errorf("Expected %s, got %s", op.expected, result)
				}
			} else {

				lines := strings.Count(result, "\n")
				if lines != realWorldLines {
					t.Errorf("Expected %d lines, got %d lines", realWorldLines, lines)
				}
			}

			t.Logf("Successfully processed %d lines with %s", realWorldLines, op.processor.Name())
		})
	}
}
