package processors

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var List = []list.Item{
	Adler32{},
	ASCII85Decoding{},
	ASCII85Encoding{},
	Base32Decode{},
	Base32Encoding{},
	Base58Decode{},
	Base58Encode{},
	Base62Decode{},
	Base62Encode{},
	Base64Decode{},
	Base64Encode{},
	Base64URLDecode{},
	Base64URLEncode{},
	Bcrypt{},
	BLAKE2b{},
	BLAKE2s{},
	Camel{},
	CountCharacters{},
	CountLines{},
	CountWords{},
	CRC32{},
	CrockfordBase32Decode{},
	CrockfordBase32Encode{},
	EscapeQuotes{},
	ExtractEmails{},
	ExtractURLs{},
	ExtractIPs{},
	FormatJSON{},
	HexDecode{},
	HexEncode{},
	HexToRGB{},
	HTMLDecode{},
	HTMLEncode{},
	JSONEscape{},
	JSONToMSGPACK{},
	JSONToYAML{},
	JSONUnescape{},
	Kebab{},
	Lower{},
	Markdown{},
	MorseCodeEncode{},
	MorseCodeDecode{},
	MD5{},
	MSGPACKToJSON{},
	NumberLines{},
	Pascal{},
	QRCode{},
	RemoveNewLines{},
	RemoveSpaces{},
	Reverse{},
	ReverseLines{},
	ROT13{},
	SHA1{},
	SHA224{},
	SHA256{},
	SHA384{},
	SHA512{},
	ShuffleLines{},
	Slug{},
	Snake{},
	SortLines{},
	Title{},
	UniqueLines{},
	Upper{},
	URLDecode{},
	URLEncode{},
	XXH32{},
	XXH64{},
	XXH128{},
	YAMLToJSON{},
	Zeropad{},
}

type Processor interface {

	// Name is the name of a processor used as the CLI command, must be one lowercase word,
	// hyphens are allowed
	Name() string

	// Alias is an optional array of alias names for the processor
	Alias() []string

	// Transform is the text transformation function, implemented by the processor
	Transform(data []byte, opts ...Flag) (string, error)

	// Flags are flags that could be used to transform the text
	Flags() []Flag
}

// StreamingConfig defines how a processor should handle streaming
type StreamingConfig struct {
	// ChunkSize defines the size of chunks to read from input (default: 64KB)
	ChunkSize int
	// BufferOutput whether to buffer output before writing (useful for processors that need full input)
	BufferOutput bool
	// LineByLine whether to process input line by line (useful for line-based processors)
	LineByLine bool
	// UTF8Chunking preserves rune boundaries when processing UTF-8 text in chunks.
	UTF8Chunking bool
}

// ConfigurableStreamingProcessor is an optional interface that processors can implement
// to customize their streaming behavior
type ConfigurableStreamingProcessor interface {
	Processor
	// GetStreamingConfig returns the streaming configuration for this processor
	GetStreamingConfig() StreamingConfig
}

type FlagType string

func (f FlagType) String() string {
	return string(f)
}

func (f FlagType) IsString() bool {
	return f == FlagString
}

const (
	FlagInt    = FlagType("Int")
	FlagUint   = FlagType("Uint")
	FlagBool   = FlagType("Bool")
	FlagString = FlagType("String")
)

type Flag struct {
	// Name - required (long version) of the flag, lowercase (with hyphens)
	Name string

	// Short - required (single character, lowercase) of the flag
	Short string

	// Desc - required, a short description of the flag
	Desc string
	// Type - required the type of the flag
	Type FlagType

	// Value - optional default value of the flag
	Value any
}

// DefaultStreamingConfig provides sensible defaults for streaming
var DefaultStreamingConfig = StreamingConfig{
	ChunkSize:    64 * 1024, // 64KB
	BufferOutput: false,
	LineByLine:   false,
}

// TransformStream provides a central streaming function that works with any processor
// It uses the processor's existing Transform method to handle streaming data
func TransformStream(processor Processor, reader io.Reader, writer io.Writer, opts ...Flag) error {
	// Get streaming configuration
	config := DefaultStreamingConfig
	if sp, ok := processor.(ConfigurableStreamingProcessor); ok {
		config = sp.GetStreamingConfig()
	}

	if config.LineByLine {
		return transformStreamLineByLine(processor, reader, writer, opts...)
	}

	if config.BufferOutput {
		return transformStreamBuffered(processor, reader, writer, opts...)
	}

	if config.UTF8Chunking {
		return transformStreamUTF8Chunked(processor, reader, writer, config.ChunkSize, opts...)
	}

	return transformStreamChunked(processor, reader, writer, config.ChunkSize, opts...)
}

// transformStreamLineByLine processes input line by line
func transformStreamLineByLine(processor Processor, reader io.Reader, writer io.Writer, opts ...Flag) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		result, err := processor.Transform(line, opts...)
		if err != nil {
			return err
		}
		if _, err := writer.Write([]byte(result + "\n")); err != nil {
			return err
		}
	}
	return scanner.Err()
}

// transformStreamBuffered reads all input, processes it, then writes output
func transformStreamBuffered(processor Processor, reader io.Reader, writer io.Writer, opts ...Flag) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	result, err := processor.Transform(data, opts...)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(result))
	return err
}

// transformStreamUTF8Chunked keeps a possibly incomplete UTF-8 rune for the next read.
func transformStreamUTF8Chunked(processor Processor, reader io.Reader, writer io.Writer, chunkSize int, opts ...Flag) error {
	buffer := make([]byte, chunkSize+utf8.UTFMax)
	pending := 0

	for {
		n, err := reader.Read(buffer[pending : pending+chunkSize])
		total := pending + n

		if err != nil {
			if total > 0 {
				result, transformErr := processor.Transform(buffer[:total], opts...)
				if transformErr != nil {
					return transformErr
				}
				if _, writeErr := writer.Write([]byte(result)); writeErr != nil {
					return writeErr
				}
			}
			if err == io.EOF {
				return nil
			}
			return err
		}

		complete := utf8CompletePrefix(buffer[:total])
		if complete > 0 {
			result, transformErr := processor.Transform(buffer[:complete], opts...)
			if transformErr != nil {
				return transformErr
			}
			if _, writeErr := writer.Write([]byte(result)); writeErr != nil {
				return writeErr
			}
		}

		pending = total - complete
		copy(buffer[:pending], buffer[complete:total])
	}
}

func utf8CompletePrefix(data []byte) int {
	start := len(data) - utf8.UTFMax + 1
	if start < 0 {
		start = 0
	}

	for i := start; i < len(data); i++ {
		if !utf8.FullRune(data[i:]) {
			return i
		}
	}
	return len(data)
}

// transformStreamChunked processes input in chunks (for processors that can handle partial data)
func transformStreamChunked(processor Processor, reader io.Reader, writer io.Writer, chunkSize int, opts ...Flag) error {
	buffer := make([]byte, chunkSize)

	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			result, transformErr := processor.Transform(buffer[:n], opts...)
			if transformErr != nil {
				return transformErr
			}
			if _, writeErr := writer.Write([]byte(result)); writeErr != nil {
				return writeErr
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return nil
}

// CanStream returns true if a processor can handle streaming
// All processors can stream using the central TransformStream function
func CanStream(processor Processor) bool {
	return true
}

// PreferStream returns true if a processor benefits from streaming
// This is useful for large files or processors that don't need full input
func PreferStream(processor Processor) bool {
	if sp, ok := processor.(ConfigurableStreamingProcessor); ok {
		config := sp.GetStreamingConfig()

		name := processor.Name()

		hashFunctions := []string{"md5", "sha1", "sha224", "sha256", "sha384", "sha512"}
		for _, hash := range hashFunctions {
			if name == hash {
				return true
			}
		}

		if !config.BufferOutput {
			return true
		}

		if config.LineByLine {
			return true
		}

		return false
	}

	// Check if processor implements the old StreamingProcessor interface
	if sp, ok := processor.(StreamingProcessor); ok {
		return sp.PreferStream()
	}

	// Default: prefer streaming for hash functions and encoders
	name := processor.Name()
	streamingFriendly := []string{
		"md5", "sha1", "sha224", "sha256", "sha384", "sha512",
		"hex-encode", "hex-decode", "base64-encode", "base64-decode",
		"base32-encode", "base32-decode", "upper", "lower",
	}

	for _, friendly := range streamingFriendly {
		if name == friendly {
			return true
		}
	}

	return false
}

// Zeropad is an Example processor to show how to add text processors,
// it checks if the data is a number and pads it with zeros
// Example implements 'Item' and 'DefaultItem' from package 'github.com/charmbracelet/bubbles/list'
// to work with the UI, and `Processor` from this package to do the text transformation and generation
// of the CLI commands
// After implementing add the struct to List.
type Zeropad struct{}

func (p Zeropad) Name() string {
	return "zeropad"
}

func (p Zeropad) Alias() []string {
	return nil
}

func (p Zeropad) Transform(data []byte, f ...Flag) (string, error) {
	strIn := strings.TrimSpace(string(data))
	neg := ""
	i, err := strconv.ParseFloat(strIn, 64)
	if err != nil {
		return "", fmt.Errorf("number expected: '%s'", data)
	}
	if i < 0 {
		neg = "-"
		data = data[1:]
	}

	var n int
	pre := ""
	for _, flag := range f {
		if flag.Short == "n" {
			x, ok := flag.Value.(uint)
			if ok {
				n = int(x)
			}
		} else if flag.Short == "p" {
			x, ok := flag.Value.(string)
			if ok {
				pre = x
			}
		}
	}
	return fmt.Sprintf("%s%s%s%s", pre, neg, strings.Repeat("0", n), data), nil
}

func (p Zeropad) Flags() []Flag {
	return []Flag{
		{
			Name:  "number-of-zeros",
			Short: "n",
			Desc:  "Number of zeros to be padded",
			Value: 5,
			Type:  FlagUint,
		},
		{
			Name:  "prefix",
			Short: "p",
			Desc:  "The number get prefixed with this",
			Value: "",
			Type:  FlagString,
		},
	}
}

func (p Zeropad) Title() string {
	title := cases.Title(language.Und, cases.NoLower).String(p.Name())
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Zeropad) Description() string {
	return "Pad a number with zeros"
}

func (p Zeropad) FilterValue() string {
	return p.Title()
}
