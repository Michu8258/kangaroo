package dataReader

import (
	"errors"

	"github.com/Michu8258/kangaroo/models"
)

// ReadFromConsole reads raw sudoku data based on user console inputs.
// Initial request config is respected and some questions are skipped
// if provided object has correct values (in range) assigned. Returns
// sudoku DTO and error if occures.
func (reader *DataReader) ReadSudokuFromConsole(request *models.SudokuConfigRequest) (
	*models.SudokuDTO, error) {

	readError := errors.New("failed to read sudoku user data inputs")

	boxSize, err := reader.Prompter.PromptGetBoxSize(request.BoxSize)
	if err != nil {
		return nil, readError
	}

	layoutWidth, err := reader.Prompter.PromptGetLayoutSize(request.LayoutWidth, "width")
	if err != nil {
		return nil, readError
	}

	layoutHeight, err := reader.Prompter.PromptGetLayoutSize(request.LayoutHeight, "height")
	if err != nil {
		return nil, readError
	}

	request.BoxSize = &boxSize
	request.LayoutWidth = &layoutWidth
	request.LayoutHeight = &layoutHeight

	sudokuDto := reader.buildEmptySudokuDTO(request)
	err = reader.Prompter.PromptSudokuValues(sudokuDto)
	if err != nil {
		reader.DebugPrinter.PrintError(err.Error())
		reader.DebugPrinter.PrintNewLine()
		return nil, readError
	}

	return sudokuDto, nil
}

// buildEmptySudokuDTO builds sudokuDTO object based un user provided requirements
func (reader *DataReader) buildEmptySudokuDTO(request *models.SudokuConfigRequest) *models.SudokuDTO {
	sudokuDto := &models.SudokuDTO{
		BoxSize: *request.BoxSize,
		Layout: models.SudokuLayoutDTO{
			Width:  *request.LayoutWidth,
			Height: *request.LayoutHeight,
		},
		Boxes: models.GenericSlice[*models.SudokuBoxDTO]{},
	}

	var bowRowIndex int8 = 0
	var boxColumnIndex int8 = 0

	for bowRowIndex = 0; bowRowIndex < sudokuDto.Layout.Height; bowRowIndex++ {
		for boxColumnIndex = 0; boxColumnIndex < sudokuDto.Layout.Width; boxColumnIndex++ {
			sudokuBox := &models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    bowRowIndex,
				IndexColumn: boxColumnIndex,
				Cells:       models.GenericSlice[*models.SudokuCellDTO]{},
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
