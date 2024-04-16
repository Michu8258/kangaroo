package prompts

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// PromptGetLayoutSize prompts user for layout size - if wrong value pre-provided
func PromptGetLayoutSize(initialSize *int8, direction string, settings *models.Settings) (int8, error) {
	if initialSize != nil && *initialSize >= settings.MinimumLayoutSizeInclusive &&
		*initialSize <= settings.MaximumLayoutSizeInclusive {
		return *initialSize, nil
	}

	question := ""
	options, defaultIndex := getLayoutSelectOptions(settings, direction)

	if initialSize == nil {
		question = fmt.Sprintf("Please select a sudoku layout size (%s):", direction)
	} else {
		question = fmt.Sprintf("Please select a sudoku layout size (%s) in range %d to %d:",
			direction, settings.MinimumLayoutSizeInclusive, settings.MaximumLayoutSizeInclusive)
	}

	result, err := PromptMakeSelectChoice(question, options, defaultIndex)
	if err != nil {
		return 0, err
	}

	return result.Value, nil
}

// getBoxSizeSelectOptions generates slice of correct options for sudoku layout
func getLayoutSelectOptions(settings *models.Settings, direction string) ([]PromptSelectOption[int8], int) {
	options := []PromptSelectOption[int8]{}
	defaultElementIndex := 0

	var size int8 = 0
	for size = settings.MinimumLayoutSizeInclusive; size <= settings.MaximumLayoutSizeInclusive; size++ {
		options = append(options, PromptSelectOption[int8]{
			Label: fmt.Sprintf("Layout size %s %d", direction, size),
			Value: size,
		})

		if size == settings.DefaultLayoutSize {
			defaultElementIndex = int(size - settings.MinimumBoxSizeInclusive)
		}
	}

	return options, defaultElementIndex
}
