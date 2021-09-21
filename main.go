package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/abhimanyu003/sttr/processors"
	"os"
	"strings"
  
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// version specify version of application using ldflags
var version string

var items = []list.Item{
	item{title: "Base64 Encoding", arg: "b64-enc", desc: "Encode your text to Base64", processor: processors.Base64Encode},
	item{title: "Base64 Decode", arg: "b64-dec", desc: "Decode your Base64 text", processor: processors.Base64Decode},

	item{title: "URL Encode", arg: "url-enc", desc: "Encode URL entities", processor: processors.URLEncode},
	item{title: "URL Decode", arg: "url-dec", desc: "Decode URL entities", processor: processors.URLDecode},

	item{title: "ROT13 Encode", arg: "rot13", desc: "Encode your text to ROT13", processor: processors.ROT13Encode},

	item{title: "To Title Case", arg: "title-case", desc: "Convert your text to Title Case", processor: processors.StringToTitle},
	item{title: "To Lower Case", arg: "lower-case", desc: "Convert your text to lower case", processor: processors.StringToLower},
	item{title: "To Upper Case", arg: "upper-case", desc: "Convert your text to UPPER CASE", processor: processors.StringToUpper},
	item{title: "To Snake Case", arg: "snake-case", desc: "Convert your text to snake_case", processor: processors.StringToSnakeCase},
	item{title: "To Kebab Case", arg: "kebab-case", desc: "Convert your text to keybab-case", processor: processors.StringToKebab},
	item{title: "To Slug Case", arg: "slug-case", desc: "Convert your text to slug-case", processor: processors.StringToSlug},
	item{title: "To Camel Case", arg: "camel-case", desc: "Convert your text to CamelCase", processor: processors.StringToCamel},
	item{title: "Reverse String", arg: "reverse", desc: "Reverse String ( gnirtS esreveR )", processor: processors.StringReverse},

	item{title: "Count Number of Characters", arg: "count-chars", desc: "Find the length of your text (including spaces)",
		processor: processors.CountNumberCharacters},
	item{title: "Count Number of Words", arg: "count-words", desc: "Count the number of words in your text",
		processor: processors.CountWords},
	item{title: "Count Number of Lines", arg: "count-lines", desc: "Count the number of lines in your text",
		processor: processors.CountLines},

	item{title: "MD5 Sum", arg: "md5", desc: "Get the MD5 checksum of your text", processor: processors.MD5Encode},
	item{title: "SHA1 Sum", arg: "sha1", desc: "Get the SHA1 checksum of your text", processor: processors.SHA1Encode},
	item{title: "SHA256 Sum", arg: "sha256", desc: "Get the SHA256 checksum of your text", processor: processors.SHA256Encode},
	item{title: "SHA512 Sum", arg: "sha512", desc: "Get the SHA512 checksum of your text", processor: processors.SHA512Encode},

	item{title: "Format JSON", arg: "json", desc: "Format your text as JSON", processor: processors.FormatJSON},
	item{title: "JSON To YAML", arg: "json-yaml", desc: "Convert JSON to YAML text", processor: processors.JSONToYAML},
	item{title: "YAML To JSON", arg: "yaml-json", desc: "Convert YAML to JSON text", processor: processors.YAMLToJSON},

	item{title: "Hex To RGB", arg: "hex-rgb", desc: "Convert a #Hex code to RGB", processor: processors.HexToRGB},

	item{title: "Sort Lines", arg: "sort-lines", desc: "Sort lines alphabetically", processor: processors.SortLines},
}

var appStyle = lipgloss.NewStyle().Width(80)
var borderStyle = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
var specialStyle = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

type item struct {
	title string
	arg   string
	desc  string
	processor func(string) string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list   list.Model
	input  string
	output string
}

func main() {
	var input string
	var proc string
	flag.StringVar(&input, "i", "", "string to process")
	flag.StringVar(&proc, "p", "", "process to be used")
	versionFlag := flag.Bool("v", false, "display version of application")
	flag.Parse()

	if *versionFlag == true {
		fmt.Println("version: ", version)
		return
	}

	if input == "" {
		divider := lipgloss.NewStyle().Padding(0, 1).Foreground(borderStyle).SetString("•").String()

		info := lipgloss.NewStyle().Foreground(specialStyle).Render
		welcome := strings.Builder{}

		title := lipgloss.NewStyle().
			Padding(0, 0, 0, 0).
			Width(80).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderTop(true).
			BorderBottom(true).
			BorderForeground(borderStyle).
			Render("Provide string to transform" + divider + info("[ Enter 2 empty lines to process ] v"+version))

		welcome.WriteString(title)

		fmt.Println(appStyle.Render(welcome.String()))

		input = readMultilineInput()
	}

	if proc != "" {
		for _, l := range items {
			if l.(item).arg == strings.TrimSpace(proc) {
				fmt.Println(l.(item).processor(input))
				os.Exit(0)
			}
		}
		fmt.Printf("Unknown processor: %s\n\n", proc)
		printProcessHelp()
		os.Exit(1)
	}

	m := model{
		list:  list.NewModel(items, list.NewDefaultDelegate(), 0, 0),
		input: input,
	}
	m.list.Title = "Select operation"

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func printProcessHelp() {
	fmt.Println("Available processors:")
	for _, l := range items {
		x := l.(item)
		fmt.Printf("\t%s - %s\n", x.arg, x.desc)
	}
}

// Init ...
func (m model) Init() tea.Cmd {
	return nil
}

// View view builder
func (m model) View() string {
	if m.output != "" {
		return lipgloss.NewStyle().
			Padding(1, 0, 1, 0).
			BorderTop(true).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(borderStyle).
			Width(80).
			Render(m.output)
	}

	return appStyle.Margin(1, 1).Render(m.list.View())
}

// Update handle all the input events and cases.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var data string

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if keypress := msg.String(); keypress == "enter" {
			i, ok := m.list.SelectedItem().(item)
			if ok {
				data = i.processor(m.input)
			}
			m.output = data
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := appStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func readMultilineInput() string {
	str := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			str = append(str, text)
		} else {
			break
		}
	}
	// Use collected inputs
	return strings.Join(str, "\n")
}
