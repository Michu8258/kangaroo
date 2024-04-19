package prompts

import (
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
	tea "github.com/charmbracelet/bubbletea"
)

type Prompter struct {
	Settings         *models.Settings
	TerminalPrinter  printer.IPrinter
	TeaProgramRunner func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error)
}

type IPrompter interface {
	PromptMakeSelectChoice(title string, options []models.PromptSelectOption,
		initialChoiceIndex int) (models.PromptSelectOption, error)
	PromptSudokuValues(sudokuDto *models.SudokuDTO) error
	PromptGetBoxSize(initialBoxSize *int8) (int8, error)
	PromptGetLayoutSize(initialSize *int8, direction string) (int8, error)
}

func GetNewPrompter(settings *models.Settings,
	terminalPrinter printer.IPrinter,
	teaProgramRunner func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error)) IPrompter {
	return &Prompter{
		Settings:         settings,
		TerminalPrinter:  terminalPrinter,
		TeaProgramRunner: teaProgramRunner,
	}
}
