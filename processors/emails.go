package processors

import (
	"github.com/mcnijman/go-emailaddress"
	"strings"
)

// ExtractEmails will pluck all the valid emails from a given text.
type ExtractEmails struct{}

func (p ExtractEmails) Name() string {
	return "extract-emails"
}

func (p ExtractEmails) Alias() []string {
	return []string{"find-emails", "find-email", "extract-email"}
}

func (p ExtractEmails) Transform(data []byte, f ...Flag) (string, error) {
	var emails []string
	extracted := emailaddress.FindWithIcannSuffix(data, false)
	for _, e := range extracted {
		emails = append(emails, e.String())
	}

	separator := "\n"
	for _, flag := range f {
		if flag.Short == "s" {
			x, ok := flag.Value.(string)
			if ok {
				separator = x
			}
		}
	}
	return strings.Join(emails, separator), nil
}

func (p ExtractEmails) Flags() []Flag {
	return []Flag{
		{
			Name:  "separator",
			Short: "s",
			Desc:  "Separator to split multiple emails",
			Value: "",
			Type:  FlagString,
		},
	}
}

func (p ExtractEmails) Title() string {
	return "Extract Emails"
}

func (p ExtractEmails) Description() string {
	return "Extract emails from given text"
}

func (p ExtractEmails) FilterValue() string {
	return p.Title()
}
