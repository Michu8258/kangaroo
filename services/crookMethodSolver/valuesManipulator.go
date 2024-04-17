package crookMethodSolver

import (
	"fmt"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
)

// assignCertainValues assigns certain values as final cell value (certain
// values is when there is only one potential value in slice of potential
// values in given cell). Returns true if at least one such case was found,
// otherwise returns false.
func assignCertainValues(sudoku *models.Sudoku, settings *models.Settings,
	debugPrinter printer.Printer) bool {

	valuesAssigned := 0

	debugPrinter.PrintDefault("Starting certain values assignment - based of potential values.")
	debugPrinter.PrintNewLine()

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				if subSudokuBoxCell.Value == nil && !subSudokuBoxCell.IsInputValue &&
					subSudokuBoxCell.PotentialValues != nil && len(*subSudokuBoxCell.PotentialValues) == 1 {
					subSudokuBoxCell.Value = &((*subSudokuBoxCell.PotentialValues)[0])
					subSudokuBoxCell.PotentialValues = nil
					valuesAssigned += 1

					debugPrinter.PrintDefault(fmt.Sprintf(
						"Assigned certain cell value: %v. Cell %s.",
						*subSudokuBoxCell.Value,
						helpers.GetCellCoordinatesString(sudoku, subSudokuBoxCell.Box, subSudokuBoxCell, true)))
					debugPrinter.PrintNewLine()
				}
			}
		}
	}

	if settings.UseDebugPrints {
		if valuesAssigned >= 1 {
			debugPrinter.PrintDefault(fmt.Sprintf(
				"Certain values assignment finished - assigned values count: %v",
				valuesAssigned))
			debugPrinter.PrintNewLine()
		} else {
			debugPrinter.PrintDefault("Certain values assignment finished - no value assigned")
			debugPrinter.PrintNewLine()
		}
	}

	return valuesAssigned >= 1
}

// checkIfAllCellsHaveValues checks if all sudokou cells has values
// and return true if that is the case
func checkIfAllCellsHaveValues(sudoku *models.Sudoku, debugPrinter printer.Printer) bool {
	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				if subSudokuBoxCell.Value == nil {
					return false
				}
			}
		}
	}

	debugPrinter.PrintDefault("It appears that the Sudoku has all cells filled with values.")
	debugPrinter.PrintNewLine()

	return true
}
