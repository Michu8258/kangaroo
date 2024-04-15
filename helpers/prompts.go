package helpers

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type PromptSelectOption[T comparable] struct {
	Label string
	Value T
}

type promptSelect[T comparable] struct {
	cursor       int
	title        string
	activeChoice PromptSelectOption[T]
	choices      []PromptSelectOption[T]
}

func PromptMakeSelectChoice[T comparable](title string, options []PromptSelectOption[T],
	initialChoiceIndex int) (PromptSelectOption[T], error) {

	model := promptSelect[T]{
		cursor:       initialChoiceIndex,
		activeChoice: options[initialChoiceIndex],
		choices:      options,
		title:        title,
	}

	prog := tea.NewProgram(model)

	defer func() {
		prog.Send(tea.ClearScreen)
	}()

	m, err := prog.Run()
	if err != nil {
		return options[initialChoiceIndex], err
	}

	if m, ok := m.(promptSelect[T]); ok {
		return m.activeChoice, nil
	}

	return options[initialChoiceIndex], fmt.Errorf("failed to get console select input value")
}

func (m promptSelect[T]) Init() tea.Cmd {
	return nil
}

func (m promptSelect[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.activeChoice = m.choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m promptSelect[T]) View() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%s\n", m.title))

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("[x] ")
		} else {
			s.WriteString("[ ] ")
		}
		s.WriteString(m.choices[i].Label)
		s.WriteString("\n")
	}

	s.WriteString("\n")

	return s.String()
}
