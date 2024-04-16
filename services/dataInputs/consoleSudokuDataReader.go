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
func ReadFromConsole(request *models.SudokuConfigRequest, settings *models.Settings) (*models.SudokuDTO, error) {
	readError := errors.New("failed to read sudoku user data inputs")

	boxSize, err := prompts.PromptGetBoxSize(request.BoxSize, settings)
	if err != nil {
		return nil, readError
	}

	layoutWidth, err := prompts.PromptGetLayoutSize(request.LayoutWidth, "width", settings)
	if err != nil {
		return nil, readError
	}

	layoutHeight, err := prompts.PromptGetLayoutSize(request.LayoutHeight, "height", settings)
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

	return sudokuDto, nil
}

// buildEmptySudokuDTO builds sudokuDTO object based un user provided requirements
func buildEmptySudokuDTO(request *models.SudokuConfigRequest) *models.SudokuDTO {
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