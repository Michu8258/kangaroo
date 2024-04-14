package crookMethodSolver

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/types"
	"github.com/beevik/guid"
)

// restoreSnapshotFromGuessedValue restores snapshot from guessed value for the sudoku,
// excludes invalid potential value from potential values collection for the cell
// referenced in sudoku value guess object.
func restoreSnapshotFromGuessedValue(sudoku *models.Sudoku, cellValueGuess *models.SudokuValueGuess,
	settings *models.Settings) error {
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

			potentialValues := types.GenericSlice[int]{}
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

	if settings.UseDebugPrints {
		fmt.Printf("Restored potential values snapshot. New potential values for "+
			"the cell: %v\n", updatedPotentialValues)
	}

	// we can assign it in guess object, because it holds reference to the actual cell
	cellValueGuess.GuessedCell.PotentialValues = &updatedPotentialValues
	cellValueGuess.GuessedCell.Value = nil

	return nil
}

// designateSudokuGuess creates an object representing a value to guess in the sudoku puzzle.
// returns boolean flag indicating if suitable cell was found, an object containing a
// snapshot of current state of potential vales per sudoku cell, and error if offured.
func designateSudokuGuess(sudoku *models.Sudoku, settings *models.Settings) (bool, *models.SudokuValueGuess, error) {
	cell, subSudokuId, err := findCellWithLowestPotentialValues(sudoku, settings)
	if err != nil {
		return false, nil, err
	}

	if cell == nil || cell.PotentialValues == nil || len(*cell.PotentialValues) < 1 {
		return false, nil, nil
	}

	if settings.UseDebugPrints {
		fmt.Println("Found cell suitable for guessing potential value of")
	}

	potentialValuesSnapshot := createPotentialValuesSnapshot(sudoku, settings)

	guess := &models.SudokuValueGuess{
		GuessedValue:            (*cell.PotentialValues)[0],
		GuessedCell:             cell,
		SubsudokuId:             *subSudokuId,
		PotentialValuesSnapshot: potentialValuesSnapshot,
	}

	if settings.UseDebugPrints {
		fmt.Printf("Created sudoku guess object. Box absolute indexes (row: %d, column: %d), "+
			"cell in box indexes (row: %d, column: %d), value selected as guess, %d.\n",
			guess.GuessedCell.Box.IndexRow, guess.GuessedCell.Box.IndexColumn,
			guess.GuessedCell.IndexRowInBox, guess.GuessedCell.IndexColumnInBox,
			guess.GuessedValue)
	}

	return true, guess, nil
}

// findCellWithLowestPotentialValues searches for cell most suitable for being selected as
// a cell we will guess value for. Returns reference to the cell itself (pointer), containing
// sub-sudoku id (pointer) and error (if occured)
func findCellWithLowestPotentialValues(sudoku *models.Sudoku, settings *models.Settings) (
	*models.SudokuCell, *guid.Guid, error) {
	var sudokuCell *models.SudokuCell
	var subSudokuId *guid.Guid

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {

				if subSudokuBoxCell.Value == nil && subSudokuBoxCell.PotentialValues != nil {
					if len(*subSudokuBoxCell.PotentialValues) == 0 {
						if settings.UseDebugPrints {
							fmt.Println("Found a cell with no potantial values " +
								"during sudoku cell guess selection.")
						}

						return nil, nil, nil
					}

					if len(*subSudokuBoxCell.PotentialValues) == 1 {
						if settings.UseDebugPrints {
							fmt.Println("Found a cell with exactly one potantial value " +
								"during sudoku cell guess selection.")
						}

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

	if settings.UseDebugPrints && sudokuCell == nil {
		fmt.Println("Did not find any cell suitable for guessing - perhaps there is no cell with potential values.")
	}

	return sudokuCell, subSudokuId, nil
}

// createPotentialValuesSnapshot creates a snapshot map of the state of potential values
// across all cells in the sudoku to
func createPotentialValuesSnapshot(sudoku *models.Sudoku, settings *models.Settings) map[guid.Guid]*[]int {
	snapshot := map[guid.Guid]*[]int{}

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

	if settings.UseDebugPrints {
		fmt.Println("Built snapshot for sudoku potantial values. Snapshot length:",
			len(snapshot))
	}

	return snapshot
}
