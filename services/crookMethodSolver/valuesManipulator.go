package crookMethodSolver

import (
	"fmt"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
)

// assignCertainValues assigns certain values as final cell value (certain
// values is when there is only one potential value in slice of potential
// values in given cell). Returns true if at least one such case was found,
// otherwise returns false.
func (solver *CrookSolver) assignCertainValues(sudoku *models.Sudoku) bool {
	valuesAssigned := 0

	solver.DebugPrinter.PrintDefault("Starting certain values assignment - based of potential values.")
	solver.DebugPrinter.PrintNewLine()

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				if subSudokuBoxCell.Value == nil && !subSudokuBoxCell.IsInputValue &&
					subSudokuBoxCell.PotentialValues != nil && len(*subSudokuBoxCell.PotentialValues) == 1 {
					subSudokuBoxCell.Value = &((*subSudokuBoxCell.PotentialValues)[0])
					subSudokuBoxCell.PotentialValues = nil
					valuesAssigned += 1

					solver.DebugPrinter.PrintDefault(fmt.Sprintf(
						"Assigned certain cell value: %v. Cell %s.",
						*subSudokuBoxCell.Value,
						helpers.GetCellCoordinatesString(sudoku, subSudokuBoxCell.Box, subSudokuBoxCell, true)))
					solver.DebugPrinter.PrintNewLine()
				}
			}
		}
	}

	if solver.Settings.UseDebugPrints {
		if valuesAssigned >= 1 {
			solver.DebugPrinter.PrintDefault(fmt.Sprintf(
				"Certain values assignment finished - assigned values count: %v",
				valuesAssigned))
			solver.DebugPrinter.PrintNewLine()
		} else {
			solver.DebugPrinter.PrintDefault("Certain values assignment finished - no value assigned")
			solver.DebugPrinter.PrintNewLine()
		}
	}

	return valuesAssigned >= 1
}

// checkIfAllCellsHaveValues checks if all sudokou cells has values
// and return true if that is the case
func (solver *CrookSolver) checkIfAllCellsHaveValues(sudoku *models.Sudoku) bool {
	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				if subSudokuBoxCell.Value == nil {
					return false
				}
			}
		}
	}

	solver.DebugPrinter.PrintDefault("It appears that the Sudoku has all cells filled with values.")
	solver.DebugPrinter.PrintNewLine()

	return true
}
