package helpers

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// IterateSubSudokusBoxesRowsCells iterates through all subsudokus and then throug
// every box, forst cell of each row, first cell of each column within the sub sudoku.
// You can provide all actions or only one for example for iterating through boxes.
// terminateOnError flag breaks loop execution if any error returned from any of the
// provided actions will not be nil.
func IterateSubSudokusBoxesRowsCells(sudoku *models.Sudoku,
	terminateOnError bool,
	boxAction *func(box *models.SudokuBox) error,
	rowAction *func(firstCellInRow *models.SudokuLine) error,
	columnAction *func(firstCellInColumn *models.SudokuLine) error) error {

	for _, subSudoku := range sudoku.SubSudokus {
		// first we are iterating through boxes
		if boxAction != nil {
			for _, subSudokuBox := range subSudoku.Boxes {
				action := *boxAction
				err := action(subSudokuBox)
				if err != nil && terminateOnError {
					return err
				}
			}
		}

		topLeftSubSudokuBoxAbsoluteRowIndex := subSudoku.TopLeftBoxRowIndex
		topLeftSubSudokuBoxAbsoluteColumnIndex := subSudoku.TopLeftBoxColumnIndex

		// then we walidate rows
		if rowAction != nil {
			action := *rowAction
			var boxRowIndex int8 = 0
			for boxRowIndex = 0; boxRowIndex < sudoku.BoxSize; boxRowIndex++ {
				// searching for sudoku boxes in top row of boxes of the subsudoku
				sudokuBox := subSudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
					return box.IndexRow == boxRowIndex+topLeftSubSudokuBoxAbsoluteRowIndex &&
						box.IndexColumn == topLeftSubSudokuBoxAbsoluteColumnIndex
				})

				if sudokuBox == nil {
					return fmt.Errorf("failed to locate sudoku box")
				}

				var cellRowIndex int8 = 0
				for cellRowIndex = 0; cellRowIndex < sudoku.BoxSize; cellRowIndex++ {
					// searching for sudoku cells in first left column of cells of the subsudoku
					// to have first cell for each row
					sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCell) bool {
						return cell.IndexRowInBox == cellRowIndex && cell.IndexColumnInBox == 0
					})

					if sudokuCell == nil {
						return fmt.Errorf("failed to locate sudoku cell")
					}

					for _, line := range sudokuCell.MemberOfLines {
						err := action(line)
						if err != nil && terminateOnError {
							return err
						}
					}
				}
			}
		}

		// and then columns
		if columnAction != nil {
			action := *columnAction
			var boxColumnIndex int8 = 0
			for boxColumnIndex = 0; boxColumnIndex < sudoku.BoxSize; boxColumnIndex++ {
				// searching for sudoku boxes in first left column of boxes of the subsudoku
				sudokuBox := subSudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
					return box.IndexRow == topLeftSubSudokuBoxAbsoluteRowIndex &&
						box.IndexColumn == boxColumnIndex+topLeftSubSudokuBoxAbsoluteColumnIndex
				})

				if sudokuBox == nil {
					return fmt.Errorf("failed to locate sudoku box")
				}

				var cellColumnIndex int8 = 0
				for cellColumnIndex = 0; cellColumnIndex < sudoku.BoxSize; cellColumnIndex++ {
					// searching for sudoku cells in top row of cells of the subsudoku
					// to have first cell for each column
					sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCell) bool {
						return cell.IndexRowInBox == 0 && cell.IndexColumnInBox == cellColumnIndex
					})

					if sudokuCell == nil {
						return fmt.Errorf("failed to locate sudoku cell")
					}

					for _, line := range sudokuCell.MemberOfLines {
						err := action(line)
						if err != nil && terminateOnError {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}
