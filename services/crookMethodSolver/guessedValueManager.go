package crookMethodSolver

import (
	"fmt"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
	guid "github.com/nu7hatch/gouuid"
)

// restoreSnapshotFromGuessedValue restores snapshot from guessed value for the sudoku,
// excludes invalid potential value from potential values collection for the cell
// referenced in sudoku value guess object.
func restoreSnapshotFromGuessedValue(sudoku *models.Sudoku, cellValueGuess *models.SudokuValueGuess,
	debugPrinter printer.Printer) error {
	snapshot := cellValueGuess.PotentialValuesSnapshot

	// so we iterate through every cell, restore snapshot for thet cell
	for _, sudokuBox := range sudoku.Boxes {
		for _, sudokuCell := range sudokuBox.Cells {
			snapshottedPotentialValues, exists := snapshot[sudokuCell.Id]
			if !exists {
				return fmt.Errorf("could not find snapshotted potential values for cell "+
					"with id of %s", sudokuCell.Id.String())
			}

			if snapshottedPotentialValues == nil {
				sudokuCell.PotentialValues = nil
				continue
			}

			potentialValues := models.GenericSlice[int]{}
			for _, pv := range *snapshottedPotentialValues {
				potentialValues = append(potentialValues, pv)
			}
			sudokuCell.PotentialValues = &potentialValues
			sudokuCell.Value = nil
		}
	}

	// and then permanently remove the value that was not right (from
	// potential values of the cell selected to be guessed value for)
	restoredPotentialValues := cellValueGuess.GuessedCell.PotentialValues
	updatedPotentialValues := restoredPotentialValues.Where(func(potVal int) bool {
		return potVal != cellValueGuess.GuessedValue
	})

	debugPrinter.PrintDefault(fmt.Sprintf("Restored potential values snapshot. "+
		"New potential values for the cell: %v", updatedPotentialValues))
	debugPrinter.PrintNewLine()

	// we can assign it in guess object, because it holds reference to the actual cell
	cellValueGuess.GuessedCell.PotentialValues = &updatedPotentialValues
	cellValueGuess.GuessedCell.Value = nil

	return nil
}

// designateSudokuGuess creates an object representing a value to guess in the sudoku puzzle.
// returns boolean flag indicating if suitable cell was found, an object containing a
// snapshot of current state of potential vales per sudoku cell, and error if offured.
func designateSudokuGuess(sudoku *models.Sudoku, debugPrinter printer.Printer) (
	bool, *models.SudokuValueGuess, error) {

	cell, subSudokuId, err := findCellWithLowestPotentialValues(sudoku, debugPrinter)
	if err != nil {
		return false, nil, err
	}

	if cell == nil || cell.PotentialValues == nil || len(*cell.PotentialValues) < 1 {
		return false, nil, nil
	}

	debugPrinter.PrintDefault("Found cell suitable for guessing potential value of.")
	debugPrinter.PrintNewLine()
	potentialValuesSnapshot := createPotentialValuesSnapshot(sudoku, debugPrinter)

	guess := &models.SudokuValueGuess{
		GuessedValue:            (*cell.PotentialValues)[0],
		GuessedCell:             cell,
		SubsudokuId:             *subSudokuId,
		PotentialValuesSnapshot: potentialValuesSnapshot,
	}

	debugPrinter.PrintDefault(fmt.Sprintf("Created sudoku guess object. "+
		"Cell %s, value selected as guess, %d.",
		helpers.GetCellCoordinatesString(sudoku, cell.Box, cell, true),
		guess.GuessedValue))
	debugPrinter.PrintNewLine()

	return true, guess, nil
}

// findCellWithLowestPotentialValues searches for cell most suitable for being selected as
// a cell we will guess value for. Returns reference to the cell itself (pointer), containing
// sub-sudoku id (pointer) and error (if occured)
func findCellWithLowestPotentialValues(sudoku *models.Sudoku,
	debugPrinter printer.Printer) (*models.SudokuCell, *guid.UUID, error) {

	var sudokuCell *models.SudokuCell
	var subSudokuId *guid.UUID

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {

				if subSudokuBoxCell.Value == nil && subSudokuBoxCell.PotentialValues != nil {
					if len(*subSudokuBoxCell.PotentialValues) == 0 {
						debugPrinter.PrintDefault("Found a cell with no potantial values " +
							"during sudoku cell guess selection.")
						debugPrinter.PrintNewLine()

						return nil, nil, nil
					}

					if len(*subSudokuBoxCell.PotentialValues) == 1 {
						debugPrinter.PrintDefault("Found a cell with exactly one potantial value " +
							"during sudoku cell guess selection.")
						debugPrinter.PrintNewLine()

						return subSudokuBoxCell, &subSudokuBoxCell.Box.Id, nil
					}

					if sudokuCell == nil {
						sudokuCell = subSudokuBoxCell
						subSudokuId = &subSudoku.Id
					}

					if len(*subSudokuBoxCell.PotentialValues) < len(*sudokuCell.PotentialValues) {
						sudokuCell = subSudokuBoxCell
						subSudokuId = &subSudoku.Id
					}

					// if we have cell with only 2 possible value - we have bigest chance to guess
					// correctly. Case of less than 2 possible values are invalid.
					if len(*sudokuCell.PotentialValues) <= 2 {
						return sudokuCell, subSudokuId, nil
					}
				}
			}
		}
	}

	if sudokuCell == nil {
		debugPrinter.PrintDefault("Did not find any cell suitable for guessing - " +
			"perhaps there is no cell with potential values.")
		debugPrinter.PrintNewLine()
	}

	return sudokuCell, subSudokuId, nil
}

// createPotentialValuesSnapshot creates a snapshot map of the state of potential values
// across all cells in the sudoku to
func createPotentialValuesSnapshot(sudoku *models.Sudoku, printer printer.Printer) map[guid.UUID]*[]int {
	snapshot := map[guid.UUID]*[]int{}

	for _, sudokuBox := range sudoku.Boxes {
		for _, sudokuCell := range sudokuBox.Cells {
			if sudokuCell.PotentialValues == nil {
				snapshot[sudokuCell.Id] = nil
				continue
			}

			potentialValues := []int{}
			for _, pv := range *sudokuCell.PotentialValues {
				potentialValues = append(potentialValues, pv)
			}
			snapshot[sudokuCell.Id] = &potentialValues
		}
	}

	printer.PrintDefault(fmt.Sprintf(
		"Built snapshot for sudoku potantial values. Snapshot length: %v\n",
		len(snapshot)))

	return snapshot
}
