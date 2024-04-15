package prompts

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
	tea "github.com/charmbracelet/bubbletea"
)

// https://github.com/charmbracelet/bubbletea/blob/master/examples/result/main.go

type sudokuValuesPrompt struct {
	sudokuDTO *models.SudokuDTO
	settings  *models.Settings
}

func PromptSudokuValues(sudokuDto *models.SudokuDTO, settings *models.Settings) error {
	failError := fmt.Errorf("failed to get sudoku values from manual input")

	program := tea.NewProgram(sudokuValuesPrompt{
		sudokuDTO: sudokuDto,
		settings:  settings,
	})

	model, err := program.Run()
	if err != nil {
		if settings.UseDebugPrints {
			fmt.Println(err)
		}
		return failError
	}

	if _, ok := model.(sudokuValuesPrompt); ok {
		return nil
	}

	return failError
}

func (m sudokuValuesPrompt) Init() tea.Cmd {
	return nil
}

func (m sudokuValuesPrompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO
	return m, nil
}

func (m sudokuValuesPrompt) View() string {
	// TODO
	return ""
}
