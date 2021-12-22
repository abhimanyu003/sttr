package ui

import (
	"fmt"
	"log"

	"github.com/abhimanyu003/sttr/processors"

	"github.com/abhimanyu003/sttr/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle     = lipgloss.NewStyle().Width(80)
	borderStyle  = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	specialStyle = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
)

type UI struct {
	list     list.Model
	input    string
	output   string
	quitting bool
}

func New(input string) UI {
	return UI{
		input: input,
	}
}

func (u UI) Render() {
	if u.input == "" {
		divider := lipgloss.NewStyle().Padding(0, 1).Foreground(borderStyle).SetString("•").String()
		info := lipgloss.NewStyle().Foreground(specialStyle).Render

		title := lipgloss.NewStyle().
			Padding(0, 0, 0, 0).
			Width(80).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderTop(true).
			BorderBottom(true).
			BorderForeground(borderStyle).
			Render("Provide string to transform" + divider + info("[ Enter 2 empty lines to process ]"))

		fmt.Println(appStyle.Render(title))
		u.input = utils.ReadMultilineInput()
	}

	u.list = list.NewModel(processors.List, list.NewDefaultDelegate(), 0, 0)
	u.list.Title = "Select transformation"

	if err := tea.NewProgram(u).Start(); err != nil {
		log.Fatalf("error running ui: %v", err)
	}
}

func (u UI) Init() tea.Cmd {
	return nil
}

func (u UI) View() string {
	if u.quitting {
		return ""
	}

	if u.output != "" {
		return lipgloss.NewStyle().
			Padding(1, 0, 1, 0).
			BorderTop(true).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(borderStyle).
			Width(80).
			Render(u.output)
	}

	return appStyle.Margin(1, 1).Render(u.list.View())
}

func (u UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			u.quitting = true
			return u, tea.Quit
		case "enter":
			var data string
			var err error
			i, ok := u.list.SelectedItem().(processors.Processor)
			if ok {
				data, err = i.Transform([]byte(u.input))
				if err != nil {
					data = fmt.Sprintf("error: %s", err.Error())
				}
			}
			u.output = data
			return u, tea.Quit
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := appStyle.GetMargin()
		u.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	u.list, cmd = u.list.Update(msg)
	return u, cmd
}
