package processors

import (
	"io"
)

type StreamingProcessor interface {
	Processor

	TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error

	CanStream() bool

	PreferStream() bool
}

type StreamingCapable interface {
	CanStream() bool
}

type DefaultStreamingProcessor struct{}

func (d DefaultStreamingProcessor) CanStream() bool {
	return false
}

func (d DefaultStreamingProcessor) PreferStream() bool {
	return false
}

func (d DefaultStreamingProcessor) TransformStream(reader io.Reader, writer io.Writer, opts ...Flag) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	result, err := d.Transform(data, opts...)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(result))
	return err
}
func (d DefaultStreamingProcessor) Transform(data []byte, opts ...Flag) (string, error) {
	return string(data), nil
}

func (d DefaultStreamingProcessor) Name() string        { return "default" }
func (d DefaultStreamingProcessor) Alias() []string     { return nil }
func (d DefaultStreamingProcessor) Flags() []Flag       { return nil }
func (d DefaultStreamingProcessor) Title() string       { return "Default" }
func (d DefaultStreamingProcessor) Description() string { return "Default processor" }
func (d DefaultStreamingProcessor) FilterValue() string { return "Default" }
