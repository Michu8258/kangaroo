package testHelpers

import "github.com/Michu8258/kangaroo/models"

type TestPrompterConfig struct {
	SelectError              error
	SudokuPromptError        error
	SelectPromptFailEnforcer *func(callIndex int) bool
	BoxSizePromptFunc        *func(callIndex int) (int8, error)
	LayoutSizePromptFunc     *func(callIndex int) (int8, error)
}

type TestPrompter struct {
	selectCallIndex           int
	boxPromptCallIndex        int
	layoutSizePromptCallIndex int
	Config                    *TestPrompterConfig
}

func GetNewTestPrompter(config *TestPrompterConfig) *TestPrompter {
	return &TestPrompter{
		Config: config,
	}
}

func (prompter *TestPrompter) PromptMakeSelectChoice(title string,
	options []models.PromptSelectOption,
	initialChoiceIndex int) (models.PromptSelectOption, error) {
	defer func() {
		prompter.selectCallIndex++
	}()

	if prompter.Config.SelectPromptFailEnforcer != nil {
		f := *prompter.Config.SelectPromptFailEnforcer
		shouldFail := f(prompter.selectCallIndex)
		if shouldFail {
			return options[initialChoiceIndex], prompter.Config.SelectError
		}
	}

	return options[initialChoiceIndex], nil
}

func (prompter *TestPrompter) PromptSudokuValues(sudokuDto *models.SudokuDTO) error {
	return prompter.Config.SudokuPromptError
}

func (prompter *TestPrompter) PromptGetBoxSize(initialBoxSize *int8) (int8, error) {
	defer func() {
		prompter.boxPromptCallIndex++
	}()

	if prompter.Config.BoxSizePromptFunc != nil {
		f := *prompter.Config.BoxSizePromptFunc
		return f(prompter.boxPromptCallIndex)
	}

	return 0, nil
}

func (prompter *TestPrompter) PromptGetLayoutSize(initialSize *int8, direction string) (int8, error) {
	defer func() {
		prompter.layoutSizePromptCallIndex++
	}()

	if prompter.Config.LayoutSizePromptFunc != nil {
		f := *prompter.Config.LayoutSizePromptFunc
		return f(prompter.layoutSizePromptCallIndex)
	}

	return 0, nil
}
