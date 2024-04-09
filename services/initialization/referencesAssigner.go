package initialization

import (
	"fmt"
	"math"

	"github.com/Michu8258/kangaroo/models"
)

type cellSearchParams struct {
	overallRowIndex    int8
	overallColumnIndex int8
}

// assignSudokuReferences assigns box references inside cells references so
// there is always a possibility to reference box having a cell reference.
// It also builds up sudoku lines object so it is possible to reference
// other cells in the same sudoku line from the cell
func assignSudokuReferences(sudoku *models.Sudoku) error {
	assignBoxReferencesInCells(sudoku)
	return buildMembersOfLines(sudoku)
}

func assignBoxReferencesInCells(sudoku *models.Sudoku) {
	for _, box := range sudoku.Boxes {
		for _, cell := range box.Cells {
			cell.Box = box
		}
	}
}

func buildMembersOfLines(sudoku *models.Sudoku) error {
	cellsInLineCount := sudoku.BoxSize * sudoku.BoxSize

	for _, subSudoku := range sudoku.SubSudokus {

		// first, we iterate through columns
		err := iterateRowsColumnsLines(
			sudoku,
			subSudoku,
			cellsInLineCount,
			"column",
			func(firstDimensionIndex, secondDimensionIndex int8) cellSearchParams {
				return cellSearchParams{
					overallRowIndex:    secondDimensionIndex,
					overallColumnIndex: firstDimensionIndex,
				}
			})

		if err != nil {
			return err
		}

		// second, we iterate through rows
		err = iterateRowsColumnsLines(
			sudoku,
			subSudoku,
			cellsInLineCount,
			"row",
			func(firstDimensionIndex, secondDimensionIndex int8) cellSearchParams {
				return cellSearchParams{
					overallRowIndex:    firstDimensionIndex,
					overallColumnIndex: secondDimensionIndex,
				}
			})

		if err != nil {
			return err
		}
	}

	return nil
}

func iterateRowsColumnsLines(sudoku *models.Sudoku, subSudoku *models.SubSudoku, cellsInLineCount int8, lineType string,
	searchParamsProvider func(firstDimensionIndex, secondDimensionIndex int8) cellSearchParams) error {

	var firstDimensionIndex int8 = 0
	for firstDimensionIndex = 0; firstDimensionIndex < cellsInLineCount; firstDimensionIndex++ {
		sudokuColumn := &models.SudokuLine{}
		var secondDimensionIndex int8 = 0

		for secondDimensionIndex = 0; secondDimensionIndex < cellsInLineCount; secondDimensionIndex++ {
			cellSearchParams := searchParamsProvider(firstDimensionIndex, secondDimensionIndex)

			cellReference, err := getSudokuCellReference(sudoku, subSudoku, cellSearchParams, lineType)
			if err != nil {
				return err
			}

			// adding cell to the line
			sudokuColumn.Cells = append(sudokuColumn.Cells, cellReference)
			// adding line data to each cell
			cellReference.MemberOfLines = append(cellReference.MemberOfLines, sudokuColumn)
		}
	}

	return nil
}

func getSudokuCellReference(sudoku *models.Sudoku, subSudoku *models.SubSudoku, searchaParams cellSearchParams, lineType string) (*models.SudokuCell, error) {
	containingBoxAbsoluteRowIndex := subSudoku.TopLeftBoxRowIndex
	containingBoxAbsoluteColumnIndex := subSudoku.TopLeftBoxColumnIndex

	containingBoxRowIndexOffset := int8(math.Floor(float64(searchaParams.overallRowIndex) / float64(sudoku.BoxSize)))
	containingBoxColumnIndexOffset := int8(math.Floor(float64(searchaParams.overallColumnIndex) / float64(sudoku.BoxSize)))

	containingBoxCellRowIndex := searchaParams.overallRowIndex - (sudoku.BoxSize * containingBoxRowIndexOffset)
	containingBoxCellColumnIndex := searchaParams.overallColumnIndex - (sudoku.BoxSize * containingBoxColumnIndexOffset)
	containingBoxAbsoluteRowIndex += containingBoxRowIndexOffset
	containingBoxAbsoluteColumnIndex += containingBoxColumnIndexOffset

	sudokuBox := sudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
		return box.IndexRow == containingBoxAbsoluteRowIndex && box.IndexColumn == containingBoxAbsoluteColumnIndex
	})

	if sudokuBox == nil {
		return nil, fmt.Errorf(
			"could not find sudoku box when constructing a sudoku line (%s). Was looking for box with "+
				"row index of %d and column index of %d",
			lineType, containingBoxAbsoluteRowIndex, containingBoxAbsoluteColumnIndex)
	}

	sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCell) bool {
		return cell.IndexRowInBox == containingBoxCellRowIndex && cell.IndexColumnInBox == containingBoxCellColumnIndex
	})

	if sudokuCell == nil {
		return nil, fmt.Errorf(
			"sudoku box was found when constructing a sudoku line (%s). Box row index %d and column index %d. "+
				"The box does not contain a cell with row index %d and column index %d",
			lineType,
			containingBoxAbsoluteRowIndex, containingBoxAbsoluteColumnIndex,
			containingBoxCellRowIndex, containingBoxCellColumnIndex)
	}

	return sudokuCell, nil
}
