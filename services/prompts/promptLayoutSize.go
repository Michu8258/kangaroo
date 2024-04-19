package prompts

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// PromptGetLayoutSize prompts user for layout size - if wrong value pre-provided
func (prompter *Prompter) PromptGetLayoutSize(initialSize *int8, direction string) (int8, error) {
	if initialSize != nil && *initialSize >= prompter.Settings.MinimumLayoutSizeInclusive &&
		*initialSize <= prompter.Settings.MaximumLayoutSizeInclusive {
		return *initialSize, nil
	}

	question := ""
	options, defaultIndex := prompter.getLayoutSelectOptions(direction)

	if initialSize == nil {
		question = fmt.Sprintf("Please select a sudoku layout size (%s):", direction)
	} else {
		question = fmt.Sprintf("Please select a sudoku layout size (%s) in range %d to %d:",
			direction, prompter.Settings.MinimumLayoutSizeInclusive, prompter.Settings.MaximumLayoutSizeInclusive)
	}

	result, err := prompter.PromptMakeSelectChoice(question, options, defaultIndex)
	if err != nil {
		return 0, err
	}

	return result.Value.(int8), nil
}

// getBoxSizeSelectOptions generates slice of correct options for sudoku layout
func (prompter *Prompter) getLayoutSelectOptions(direction string) ([]models.PromptSelectOption, int) {
	options := []models.PromptSelectOption{}
	defaultElementIndex := 0

	var size int8 = 0
	for size = prompter.Settings.MinimumLayoutSizeInclusive; size <= prompter.Settings.MaximumLayoutSizeInclusive; size++ {
		options = append(options, models.PromptSelectOption{
			Label: fmt.Sprintf("Layout size %s %d", direction, size),
			Value: size,
		})

		if size == prompter.Settings.DefaultLayoutSize {
			defaultElementIndex = int(size - prompter.Settings.MinimumBoxSizeInclusive)
		}
	}

	return options, defaultElementIndex
}
