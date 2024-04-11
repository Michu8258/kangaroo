package crookMethodSolver

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// assignCertainValues assigns certain values as final cell value (certain
// values is when there is only one potential value in slice of potential
// values in given cell). Returns true if at least one such case was found,
// otherwise returns false.
func assignCertainValues(sudoku *models.Sudoku) bool {
	atLeastOneAssigned := false

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				if subSudokuBoxCell.Value == nil && !subSudokuBoxCell.IsInputValue &&
					subSudokuBoxCell.PotentialValues != nil && len(*subSudokuBoxCell.PotentialValues) == 1 {
					subSudokuBoxCell.Value = &((*subSudokuBoxCell.PotentialValues)[0])
					atLeastOneAssigned = true
				}
			}
		}
	}

	return atLeastOneAssigned
}

// clearPossibleValues clear all previously assigned slices
// of potantial values in all sudoku cells.
func clearPossibleValues(sudoku *models.Sudoku) {
	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				subSudokuBoxCell.PotentialValues = nil
			}
		}
	}
}

// checkIfAllCellsHaveValues checks if all sudokou cells has values
// and return true if that is the case
func checkIfAllCellsHaveValues(sudoku *models.Sudoku, settings *models.Settings) bool {
	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				if subSudokuBoxCell.Value == nil {
					return false
				}
			}
		}
	}

	if settings.UseDebugPrints {
		fmt.Println("It appears that the Sudoku has all cells filled with values.")
	}

	return true
}
