package prompts

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// PromptGetBoxSize prompts user for box size - if wrong value pre-provided
func PromptGetBoxSize(initialBoxSize *int8, settings *models.Settings) (int8, error) {
	if initialBoxSize != nil && *initialBoxSize >= settings.MinimumBoxSizeInclusive &&
		*initialBoxSize <= settings.MaximumBoxSizeInclusive {
		return *initialBoxSize, nil
	}

	question := ""
	options, defaultIndex := getBoxSizeSelectOptions(settings)

	if initialBoxSize == nil {
		question = "Please select a sudoku box size:"
	} else {
		question = fmt.Sprintf("Please select a sudoku box size in range %d to %d:",
			settings.MinimumBoxSizeInclusive, settings.MaximumBoxSizeInclusive)
	}

	result, err := PromptMakeSelectChoice(question, options, defaultIndex)
	if err != nil {
		return 0, err
	}

	return result.Value, nil
}

// getBoxSizeSelectOptions generates slice of correct options for box size
func getBoxSizeSelectOptions(settings *models.Settings) ([]PromptSelectOption[int8], int) {
	options := []PromptSelectOption[int8]{}
	defaultElementIndex := 0

	var size int8 = 0
	for size = settings.MinimumBoxSizeInclusive; size <= settings.MaximumBoxSizeInclusive; size++ {
		options = append(options, PromptSelectOption[int8]{
			Label: fmt.Sprintf("Box size %d", size),
			Value: size,
		})

		if size == settings.DefaultBoxSize {
			defaultElementIndex = int(size - settings.MinimumBoxSizeInclusive)
		}
	}

	return options, defaultElementIndex
}
