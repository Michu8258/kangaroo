package crookMethodSolver

import (
	"errors"
	"slices"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
)

// validateSudokuRules validates sudoku cells values against sudoku rules.
// If no rule is brokn, then true will be returned, otherwise, false.
// In case a sudoku with all filled cells values is provided, true as a result
// of this function will indicate correct sudoku solution.
func validateSudokuRules(sudoku *models.Sudoku) (bool, error) {
	validationError := errors.New("validation error")

	boxValidator := func(subSudokuBox *models.SudokuBox) error {
		boxRuleViolated := checkRuleViolation(sudoku, subSudokuBox.Cells)
		if boxRuleViolated {
			return validationError
		}

		return nil
	}

	lineValidator := func(firstCellInLine *models.SudokuLine) error {
		ruleViolated := checkRuleViolation(sudoku, firstCellInLine.Cells)
		if ruleViolated {
			return validationError
		}

		return nil
	}

	iterationError := helpers.IterateSubSudokusBoxesRowsCells(
		sudoku,
		true,
		&boxValidator,
		&lineValidator,
		&lineValidator)

	if iterationError != nil && iterationError == validationError {
		return false, nil
	} else if iterationError != nil {
		return false, iterationError
	} else {
		return true, nil
	}
}

// checkRuleViolation returns true if the rule is violated (broken)
func checkRuleViolation(sudoku *models.Sudoku, cells models.GenericSlice[*models.SudokuCell]) bool {
	minValue := 1
	maxValue := int(sudoku.BoxSize * sudoku.BoxSize)

	cellsWithValues := cells.Where(func(cell *models.SudokuCell) bool {
		return cell.Value != nil
	})

	existingValues := make([]int, 0, len(cellsWithValues))
	for _, cell := range cellsWithValues {
		val := *cell.Value

		// value in range and did not appear in cells collection yet
		if val >= minValue && val <= maxValue && !slices.Contains(existingValues, val) {
			existingValues = append(existingValues, val)
		} else {
			// violation
			return true
		}
	}

	return false
}
