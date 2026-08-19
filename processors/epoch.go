package processors

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// namedLayouts maps friendly layout names to Go reference layouts, so that
// users don't have to remember the "2006-01-02" reference time. Any value
// that is not a known name is used verbatim as a Go layout string.
var namedLayouts = map[string]string{
	"rfc3339":  time.RFC3339,
	"rfc1123":  time.RFC1123,
	"rfc822":   time.RFC822,
	"ansic":    time.ANSIC,
	"unixdate": time.UnixDate,
	"kitchen":  time.Kitchen,
	"datetime": "2006-01-02 15:04:05",
	"date":     "2006-01-02",
	"time":     "15:04:05",
}

// parseLayouts is the ordered set of layouts tried when decoding a human date
// without an explicit --format flag.
var parseLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02 15:04:05.999999999",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02",
	"2006/01/02 15:04:05",
	"2006/01/02",
	time.RFC1123Z,
	time.RFC1123,
	time.RFC822Z,
	time.RFC822,
	time.ANSIC,
	time.UnixDate,
}

// resolveLayout returns the Go layout for a --format value. A recognised name
// is mapped to its layout, anything else is treated as a raw Go layout.
func resolveLayout(format string) string {
	if l, ok := namedLayouts[strings.ToLower(strings.TrimSpace(format))]; ok {
		return l
	}
	return format
}

func stringFlag(short string, f []Flag) string {
	for _, flag := range f {
		if flag.Short == short {
			if v, ok := flag.Value.(string); ok {
				return v
			}
		}
	}
	return ""
}

func boolFlag(short string, f []Flag) bool {
	for _, flag := range f {
		if flag.Short == short {
			if v, ok := flag.Value.(bool); ok {
				return v
			}
		}
	}
	return false
}

// scaleToNanos converts an integer timestamp expressed in unit to nanoseconds.
func scaleToNanos(value int64, unit string) (int64, error) {
	switch unit {
	case "s":
		return value * int64(time.Second), nil
	case "ms":
		return value * int64(time.Millisecond), nil
	case "us", "µs":
		return value * int64(time.Microsecond), nil
	case "ns":
		return value, nil
	default:
		return 0, fmt.Errorf("unknown unit '%s', expected one of auto, s, ms, us, ns", unit)
	}
}

// detectUnit guesses the timestamp unit from the number of digits, using the
// common convention that a seconds timestamp is ~10 digits, milliseconds ~13,
// microseconds ~16 and nanoseconds ~19.
func detectUnit(value int64) string {
	digits := len(strconv.FormatInt(value, 10))
	if value < 0 {
		digits-- // don't count the sign
	}
	switch {
	case digits <= 11:
		return "s"
	case digits <= 14:
		return "ms"
	case digits <= 17:
		return "us"
	default:
		return "ns"
	}
}

// EpochToDate converts a Unix timestamp to a human-readable date.
type EpochToDate struct{}

func (p EpochToDate) Name() string {
	return "epoch-date"
}

func (p EpochToDate) Alias() []string {
	return []string{"epoch", "epoch-human", "from-epoch"}
}

func (p EpochToDate) Transform(data []byte, f ...Flag) (string, error) {
	in := strings.TrimSpace(string(data))
	value, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid epoch timestamp '%s'", in)
	}

	unit := stringFlag("u", f)
	if unit == "" || unit == "auto" {
		unit = detectUnit(value)
	}
	nanos, err := scaleToNanos(value, unit)
	if err != nil {
		return "", err
	}

	t := time.Unix(0, nanos).UTC()
	if boolFlag("l", f) {
		t = t.Local()
	}

	layout := time.RFC3339
	if format := stringFlag("f", f); format != "" {
		layout = resolveLayout(format)
	}
	return t.Format(layout), nil
}

func (p EpochToDate) Flags() []Flag {
	return []Flag{
		{Name: "format", Short: "f", Desc: "Output layout: a named layout (rfc3339, datetime, date, ...) or a Go time layout", Type: FlagString, Value: ""},
		{Name: "unit", Short: "u", Desc: "Input unit: auto, s, ms, us or ns", Type: FlagString, Value: "auto"},
		{Name: "local", Short: "l", Desc: "Render the date in the local timezone instead of UTC", Type: FlagBool, Value: false},
	}
}

func (p EpochToDate) Title() string {
	title := "Epoch to Date"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p EpochToDate) Description() string {
	return "Convert a Unix timestamp to a human-readable date"
}

func (p EpochToDate) FilterValue() string {
	return p.Title()
}

// DateToEpoch converts a human-readable date to a Unix timestamp.
type DateToEpoch struct{}

func (p DateToEpoch) Name() string {
	return "date-epoch"
}

func (p DateToEpoch) Alias() []string {
	return []string{"to-epoch", "human-epoch", "date-to-epoch"}
}

func (p DateToEpoch) parse(in string, f ...Flag) (time.Time, error) {
	loc := time.UTC
	if boolFlag("l", f) {
		loc = time.Local
	}

	if format := stringFlag("f", f); format != "" {
		return time.ParseInLocation(resolveLayout(format), in, loc)
	}

	for _, layout := range parseLayouts {
		if t, err := time.ParseInLocation(layout, in, loc); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("could not parse date '%s', pass --format to specify the layout", in)
}

func (p DateToEpoch) Transform(data []byte, f ...Flag) (string, error) {
	in := strings.TrimSpace(string(data))
	t, err := p.parse(in, f...)
	if err != nil {
		return "", err
	}

	unit := stringFlag("u", f)
	switch unit {
	case "", "s":
		return strconv.FormatInt(t.Unix(), 10), nil
	case "ms":
		return strconv.FormatInt(t.UnixMilli(), 10), nil
	case "us", "µs":
		return strconv.FormatInt(t.UnixMicro(), 10), nil
	case "ns":
		return strconv.FormatInt(t.UnixNano(), 10), nil
	default:
		return "", fmt.Errorf("unknown unit '%s', expected one of s, ms, us, ns", unit)
	}
}

func (p DateToEpoch) Flags() []Flag {
	return []Flag{
		{Name: "format", Short: "f", Desc: "Input layout: a named layout (rfc3339, datetime, date, ...) or a Go time layout", Type: FlagString, Value: ""},
		{Name: "unit", Short: "u", Desc: "Output unit: s, ms, us or ns", Type: FlagString, Value: "s"},
		{Name: "local", Short: "l", Desc: "Interpret a date without a timezone as local time instead of UTC", Type: FlagBool, Value: false},
	}
}

func (p DateToEpoch) Title() string {
	title := "Date to Epoch"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p DateToEpoch) Description() string {
	return "Convert a human-readable date to a Unix timestamp"
}

func (p DateToEpoch) FilterValue() string {
	return p.Title()
}
