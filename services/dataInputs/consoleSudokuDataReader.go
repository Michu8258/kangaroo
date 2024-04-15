package dataInputs

import (
	"errors"
	"fmt"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
)

// https://github.com/charmbracelet/bubbletea/blob/master/examples/result/main.go

// TODO add docs
func ReadFromConsole(request models.SolveCommandRequest, settings *models.Settings) (*models.SudokuDTO, error) {
	readError := errors.New("failed to read sudoku user data inputs")

	boxSize, err := readBoxSize(request, settings)
	if err != nil {
		return nil, readError
	}

	layoutWidth, err := readLayoutSize(settings, "width", request.LayoutWidth)
	if err != nil {
		return nil, readError
	}

	layoutHeight, err := readLayoutSize(settings, "height", request.LayoutHeight)
	if err != nil {
		return nil, readError
	}

	request.BoxSize = &boxSize
	request.LayoutWidth = &layoutWidth
	request.LayoutHeight = &layoutHeight
	return nil, errors.New("fwiuehfiuweifuh")
}

func readLayoutSize(settings *models.Settings, direction string, layoutDirectionSize *int8) (int8, error) {
	if layoutDirectionSize != nil && *layoutDirectionSize >= settings.MinimumLayoutSizeInclusive &&
		*layoutDirectionSize <= settings.MaximumLayoutSizeInclusive {
		return *layoutDirectionSize, nil
	}

	question := ""
	options, defaultIndex := getLayoutSelectOptions(settings, direction)

	if layoutDirectionSize == nil {
		question = fmt.Sprintf("Please select a sudoku layout size (%s):", direction)
	} else {
		question = fmt.Sprintf("Please select a sudoku layout size (%s) in range %d to %d:",
			direction, settings.MinimumLayoutSizeInclusive, settings.MaximumLayoutSizeInclusive)
	}

	result, err := helpers.PromptMakeSelectChoice(question, options, defaultIndex)
	if err != nil {
		return 0, err
	}

	return result.Value, nil
}

func getLayoutSelectOptions(settings *models.Settings, direction string) ([]helpers.PromptSelectOption[int8], int) {
	options := []helpers.PromptSelectOption[int8]{}
	defaultElementIndex := 0

	var size int8 = 0
	for size = settings.MinimumLayoutSizeInclusive; size <= settings.MaximumLayoutSizeInclusive; size++ {
		options = append(options, helpers.PromptSelectOption[int8]{
			Label: fmt.Sprintf("Layout size %s %d", direction, size),
			Value: size,
		})

		if size == settings.DefaultLayoutSize {
			defaultElementIndex = int(size - settings.MinimumBoxSizeInclusive)
		}
	}

	return options, defaultElementIndex
}

func readBoxSize(request models.SolveCommandRequest, settings *models.Settings) (int8, error) {
	if request.BoxSize != nil && *request.BoxSize >= settings.MinimumBoxSizeInclusive &&
		*request.BoxSize <= settings.MaximumBoxSizeInclusive {
		return *request.BoxSize, nil
	}

	question := ""
	options, defaultIndex := getBoxSizeSelectOptions(settings)

	if request.BoxSize == nil {
		question = "Please select a sudoku box size:"
	} else {
		question = fmt.Sprintf("Please select a sudoku box size in range %d to %d:",
			settings.MinimumBoxSizeInclusive, settings.MaximumBoxSizeInclusive)
	}

	result, err := helpers.PromptMakeSelectChoice(question, options, defaultIndex)
	if err != nil {
		return 0, err
	}

	return result.Value, nil
}

func getBoxSizeSelectOptions(settings *models.Settings) ([]helpers.PromptSelectOption[int8], int) {
	options := []helpers.PromptSelectOption[int8]{}
	defaultElementIndex := 0

	var size int8 = 0
	for size = settings.MinimumBoxSizeInclusive; size <= settings.MaximumBoxSizeInclusive; size++ {
		options = append(options, helpers.PromptSelectOption[int8]{
			Label: fmt.Sprintf("Box size %d", size),
			Value: size,
		})

		if size == settings.DefaultBoxSize {
			defaultElementIndex = int(size - settings.MinimumBoxSizeInclusive)
		}
	}

	return options, defaultElementIndex
}
