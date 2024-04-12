package crookMethodSolver

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// assignCertainValues assigns certain values as final cell value (certain
// values is when there is only one potential value in slice of potential
// values in given cell). Returns true if at least one such case was found,
// otherwise returns false.
func assignCertainValues(sudoku *models.Sudoku, settings *models.Settings) bool {
	valuesAssigned := 0

	if settings.UseDebugPrints {
		fmt.Println("Starting certain values assignment - based of potential values")
	}

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				if subSudokuBoxCell.Value == nil && !subSudokuBoxCell.IsInputValue &&
					subSudokuBoxCell.PotentialValues != nil && len(*subSudokuBoxCell.PotentialValues) == 1 {
					subSudokuBoxCell.Value = &((*subSudokuBoxCell.PotentialValues)[0])
					subSudokuBoxCell.PotentialValues = nil
					valuesAssigned += 1

					if settings.UseDebugPrints {
						fmt.Printf("Assigned certain cell value: %v. Box absolute indexes(row: %d, column: %d), "+
							"cell indexes(row: %d, column: %d).\n",
							*subSudokuBoxCell.Value,
							subSudokuBoxCell.Box.IndexRow, subSudokuBoxCell.Box.IndexColumn,
							subSudokuBoxCell.IndexRowInBox, subSudokuBoxCell.IndexColumnInBox)
					}
				}
			}
		}
	}

	if settings.UseDebugPrints {
		if valuesAssigned >= 1 {
			fmt.Println("Certain values assignment finished - assigned values count: ", valuesAssigned)
		} else {
			fmt.Println("Certain values assignment finished - no value assigned")
		}
	}

	return valuesAssigned >= 1
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
