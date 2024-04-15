package dataInputs

import (
	"errors"
	"fmt"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/prompts"
	"github.com/Michu8258/kangaroo/types"
)

// ReadFromConsole reads raw sudoku data based on user console inputs.
// Initial request config is respected and some questions are skipped
// if provided object has correct values (in range) assigned. Returns
// sudoku DTO and error if occures.
func ReadFromConsole(request *models.SolveCommandRequest, settings *models.Settings) (*models.SudokuDTO, error) {
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

	sudokuDto := buildEmptySudokuDTO(request)
	err = prompts.PromptSudokuValues(sudokuDto, settings)
	if err != nil {
		if settings.UseDebugPrints {
			fmt.Println(err)
		}
		return nil, readError
	}

	return nil, errors.New("fwiuehfiuweifuh")
	// return sudokuDto, nil
}

// buildEmptySudokuDTO builds sudokuDTO object based un user provided requirements
func buildEmptySudokuDTO(request *models.SolveCommandRequest) *models.SudokuDTO {
	sudokuDto := &models.SudokuDTO{
		BoxSize: *request.BoxSize,
		Layout: models.SudokuLayoutDTO{
			Width:  *request.LayoutWidth,
			Height: *request.LayoutHeight,
		},
		Boxes: types.GenericSlice[*models.SudokuBoxDTO]{},
	}

	var bowRowIndex int8 = 0
	var boxColumnIndex int8 = 0

	for bowRowIndex = 0; bowRowIndex < sudokuDto.Layout.Width; bowRowIndex++ {
		for boxColumnIndex = 0; boxColumnIndex < sudokuDto.Layout.Height; boxColumnIndex++ {
			sudokuBox := &models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    bowRowIndex,
				IndexColumn: boxColumnIndex,
				Cells:       types.GenericSlice[*models.SudokuCellDTO]{},
			}

			var cellRowIndex int8 = 0
			var cellColumnIndex int8 = 0

			for cellRowIndex = 0; cellRowIndex < sudokuDto.BoxSize; cellRowIndex++ {
				for cellColumnIndex = 0; cellColumnIndex < sudokuDto.BoxSize; cellColumnIndex++ {
					sudokuBox.Cells = append(sudokuBox.Cells, &models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    cellRowIndex,
						IndexColumnInBox: cellColumnIndex,
					})
				}
			}

			sudokuDto.Boxes = append(sudokuDto.Boxes, sudokuBox)
		}
	}

	return sudokuDto
}

// readLayoutSize prompts user for layout size - if wrong value pre-provided
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

	result, err := prompts.PromptMakeSelectChoice(question, options, defaultIndex)
	if err != nil {
		return 0, err
	}

	return result.Value, nil
}

// getBoxSizeSelectOptions generates slice of correct options for sudoku layout
func getLayoutSelectOptions(settings *models.Settings, direction string) ([]prompts.PromptSelectOption[int8], int) {
	options := []prompts.PromptSelectOption[int8]{}
	defaultElementIndex := 0

	var size int8 = 0
	for size = settings.MinimumLayoutSizeInclusive; size <= settings.MaximumLayoutSizeInclusive; size++ {
		options = append(options, prompts.PromptSelectOption[int8]{
			Label: fmt.Sprintf("Layout size %s %d", direction, size),
			Value: size,
		})

		if size == settings.DefaultLayoutSize {
			defaultElementIndex = int(size - settings.MinimumBoxSizeInclusive)
		}
	}

	return options, defaultElementIndex
}

// readBoxSize prompts user for box size - if wrong value pre-provided
func readBoxSize(request *models.SolveCommandRequest, settings *models.Settings) (int8, error) {
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

	result, err := prompts.PromptMakeSelectChoice(question, options, defaultIndex)
	if err != nil {
		return 0, err
	}

	return result.Value, nil
}

// getBoxSizeSelectOptions generates slice of correct options for box size
func getBoxSizeSelectOptions(settings *models.Settings) ([]prompts.PromptSelectOption[int8], int) {
	options := []prompts.PromptSelectOption[int8]{}
	defaultElementIndex := 0

	var size int8 = 0
	for size = settings.MinimumBoxSizeInclusive; size <= settings.MaximumBoxSizeInclusive; size++ {
		options = append(options, prompts.PromptSelectOption[int8]{
			Label: fmt.Sprintf("Box size %d", size),
			Value: size,
		})

		if size == settings.DefaultBoxSize {
			defaultElementIndex = int(size - settings.MinimumBoxSizeInclusive)
		}
	}

	return options, defaultElementIndex
}
