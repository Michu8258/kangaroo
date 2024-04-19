package prompts

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// PromptGetBoxSize prompts user for box size - if wrong value pre-provided
func (prompter *Prompter) PromptGetBoxSize(initialBoxSize *int8) (int8, error) {
	if initialBoxSize != nil && *initialBoxSize >= prompter.Settings.MinimumBoxSizeInclusive &&
		*initialBoxSize <= prompter.Settings.MaximumBoxSizeInclusive {
		return *initialBoxSize, nil
	}

	question := ""
	options, defaultIndex := prompter.getBoxSizeSelectOptions()

	if initialBoxSize == nil {
		question = "Please select a sudoku box size:"
	} else {
		question = fmt.Sprintf("Please select a sudoku box size in range %d to %d:",
			prompter.Settings.MinimumBoxSizeInclusive, prompter.Settings.MaximumBoxSizeInclusive)
	}

	result, err := prompter.PromptMakeSelectChoice(question, options, defaultIndex)
	if err != nil {
		return 0, err
	}

	return result.Value.(int8), nil
}

// getBoxSizeSelectOptions generates slice of correct options for box size
func (prompter *Prompter) getBoxSizeSelectOptions() ([]models.PromptSelectOption, int) {
	options := []models.PromptSelectOption{}
	defaultElementIndex := 0

	var size int8 = 0
	for size = prompter.Settings.MinimumBoxSizeInclusive; size <= prompter.Settings.MaximumBoxSizeInclusive; size++ {
		options = append(options, models.PromptSelectOption{
			Label: fmt.Sprintf("Box size %d", size),
			Value: size,
		})

		if size == prompter.Settings.DefaultBoxSize {
			defaultElementIndex = int(size - prompter.Settings.MinimumBoxSizeInclusive)
		}
	}

	return options, defaultElementIndex
}
