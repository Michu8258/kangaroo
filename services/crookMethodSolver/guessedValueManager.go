package crookMethodSolver

import (
	"fmt"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	guid "github.com/nu7hatch/gouuid"
)

// restoreSnapshotFromGuessedValue restores snapshot from guessed value for the sudoku,
// excludes invalid potential value from potential values collection for the cell
// referenced in sudoku value guess object.
func (solver *CrookSolver) restoreSnapshotFromGuessedValue(sudoku *models.Sudoku,
	cellValueGuess *models.SudokuValueGuess) error {

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

	solver.DebugPrinter.PrintDefault(fmt.Sprintf("Restored potential values snapshot. "+
		"New potential values for the cell: %v", updatedPotentialValues))
	solver.DebugPrinter.PrintNewLine()

	// we can assign it in guess object, because it holds reference to the actual cell
	cellValueGuess.GuessedCell.PotentialValues = &updatedPotentialValues
	cellValueGuess.GuessedCell.Value = nil

	return nil
}

// designateSudokuGuess creates an object representing a value to guess in the sudoku puzzle.
// returns boolean flag indicating if suitable cell was found, an object containing a
// snapshot of current state of potential vales per sudoku cell, and error if offured.
func (solver *CrookSolver) designateSudokuGuess(sudoku *models.Sudoku) (
	bool, *models.SudokuValueGuess, error) {

	cell, subSudokuId, err := solver.findCellWithLowestPotentialValues(sudoku)
	if err != nil {
		return false, nil, err
	}

	if cell == nil || cell.PotentialValues == nil || len(*cell.PotentialValues) < 1 {
		return false, nil, nil
	}

	solver.DebugPrinter.PrintDefault("Found cell suitable for guessing potential value of.")
	solver.DebugPrinter.PrintNewLine()
	potentialValuesSnapshot := solver.createPotentialValuesSnapshot(sudoku)

	guess := &models.SudokuValueGuess{
		GuessedValue:            (*cell.PotentialValues)[0],
		GuessedCell:             cell,
		SubsudokuId:             *subSudokuId,
		PotentialValuesSnapshot: potentialValuesSnapshot,
	}

	solver.DebugPrinter.PrintDefault(fmt.Sprintf("Created sudoku guess object. "+
		"Cell %s, value selected as guess, %d.",
		helpers.GetCellCoordinatesString(sudoku, cell.Box, cell, true),
		guess.GuessedValue))
	solver.DebugPrinter.PrintNewLine()

	return true, guess, nil
}

// findCellWithLowestPotentialValues searches for cell most suitable for being selected as
// a cell we will guess value for. Returns reference to the cell itself (pointer), containing
// sub-sudoku id (pointer) and error (if occured)
func (solver *CrookSolver) findCellWithLowestPotentialValues(sudoku *models.Sudoku) (
	*models.SudokuCell, *guid.UUID, error) {

	var sudokuCell *models.SudokuCell
	var subSudokuId *guid.UUID

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {

				if subSudokuBoxCell.Value == nil && subSudokuBoxCell.PotentialValues != nil {
					if len(*subSudokuBoxCell.PotentialValues) == 0 {
						solver.DebugPrinter.PrintDefault("Found a cell with no potantial values " +
							"during sudoku cell guess selection.")
						solver.DebugPrinter.PrintNewLine()

						return nil, nil, nil
					}

					if len(*subSudokuBoxCell.PotentialValues) == 1 {
						solver.DebugPrinter.PrintDefault("Found a cell with exactly one potantial value " +
							"during sudoku cell guess selection.")
						solver.DebugPrinter.PrintNewLine()

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
		solver.DebugPrinter.PrintDefault("Did not find any cell suitable for guessing - " +
			"perhaps there is no cell with potential values.")
		solver.DebugPrinter.PrintNewLine()
	}

	return sudokuCell, subSudokuId, nil
}

// createPotentialValuesSnapshot creates a snapshot map of the state of potential values
// across all cells in the sudoku to
func (solver *CrookSolver) createPotentialValuesSnapshot(sudoku *models.Sudoku) map[guid.UUID]*[]int {
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

	solver.DebugPrinter.PrintDefault(fmt.Sprintf(
		"Built snapshot for sudoku potantial values. Snapshot length: %v",
		len(snapshot)))
	solver.DebugPrinter.PrintNewLine()

	return snapshot
}
