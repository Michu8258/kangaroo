package prompts

import (
	"fmt"
	"strings"

	"github.com/Michu8258/kangaroo/models"
	tea "github.com/charmbracelet/bubbletea"
)

type promptSelect struct {
	cursor       int
	title        string
	activeChoice models.PromptSelectOption
	choices      []models.PromptSelectOption
	quit         bool
}

// PromptMakeSelectChoice wraps logic for promptin the user to select
// one of selected option (with default option index).
func (prompter *Prompter) PromptMakeSelectChoice(title string,
	options []models.PromptSelectOption,
	initialChoiceIndex int) (models.PromptSelectOption, error) {
	model, err := prompter.TeaProgramRunner(promptSelect{
		cursor:       initialChoiceIndex,
		activeChoice: options[initialChoiceIndex],
		choices:      options,
		title:        title,
	})

	if err != nil {
		return options[initialChoiceIndex], err
	}

	if m, ok := model.(promptSelect); ok {
		return m.activeChoice, nil
	}

	return options[initialChoiceIndex], fmt.Errorf("failed to get console select input value")
}

// Init iniitalizes tea model state
func (m promptSelect) Init() tea.Cmd {
	return nil
}

// Update updates tea model state
func (m promptSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		keyString := msg.String()
		switch keyString {
		case "ctrl+c", "esc":
			m.quit = true
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.activeChoice = m.choices[m.cursor]
			m.quit = true
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

// view renders output based on model state
func (m promptSelect) View() string {
	if m.quit {
		return ""
	}

	s := strings.Builder{}
	s.WriteString(models.TerminalStyles.PrimaryStyle.Render(m.title))
	s.WriteString("\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString(models.TerminalStyles.SuccessStyle.Render("[✓] "))
			s.WriteString(models.TerminalStyles.SuccessStyle.Render(m.choices[i].Label))
		} else {
			s.WriteString(models.TerminalStyles.DefaultStyle.Render("[ ] "))
			s.WriteString(models.TerminalStyles.DefaultStyle.Render(m.choices[i].Label))
		}
		s.WriteString("\n")
	}

	s.WriteString("\n")

	return s.String()
}
