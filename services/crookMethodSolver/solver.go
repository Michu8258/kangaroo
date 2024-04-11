package crookMethodSolver

import (
	"errors"

	"github.com/Michu8258/kangaroo/models"
)

func SolveWithCrookMethod(sudoku *models.Sudoku, settings *models.Settings) (bool, []error) {
	errs := []error{}

	// simple sudokus that can be hamdled with pure elimination logic
	allCellsHaveValues, eliminationErrors := executeEliminationsLogic(sudoku, settings)
	errs = append(errs, eliminationErrors...)
	if len(errs) >= 1 {
		return false, errs
	}

	if allCellsHaveValues {
		return handleAllCellsFilledCase(sudoku, errs)
	}

	// preemptive sets (Crook)
	atLeastOnePreemptiveSetManaged := true
	for atLeastOnePreemptiveSetManaged {
		manageSuccess, err := executePreemptiveSetsLogic(sudoku)
		if err != nil {
			errs = append(errs, err)
			return false, errs
		}

		atLeastOnePreemptiveSetManaged = manageSuccess
		if atLeastOnePreemptiveSetManaged {
			atLeastOneValueAssigned := assignCertainValues(sudoku)
			if atLeastOneValueAssigned {
				SolveWithCrookMethod(sudoku, settings)
			}
		}
	}

	// at this point, we have exhausted simple elimination method
	// and there are no cells wirh single potential value that
	// vould not violate sudoku rules. So we are guessing now.

	// for now we assume that the sudoku is not solved
	return false, errs
}

// executeEliminationsLogic executes simple elimination logic that may solve sudoku,
// but will not in case of difficult ones. It returns a boolean flag indicating
// if all cells has assigned certain values and slice of errors
func executeEliminationsLogic(sudoku *models.Sudoku, settings *models.Settings) (bool, []error) {
	errs := []error{}
	noMoreValuesToEliminate := false

	for !noMoreValuesToEliminate {
		//assign potential values
		errs = append(errs, assignCellsPotentialValues(sudoku, settings)...)
		if len(errs) >= 1 {
			return false, errs
		}

		// try to assign certain values
		atLeastOneValueAssigned := assignCertainValues(sudoku)
		if atLeastOneValueAssigned {
			noMoreValuesToEliminate = false
			clearPossibleValues(sudoku)
			allCellsFilled := checkIfAllCellsHaveValues(sudoku, settings)
			if allCellsFilled {
				return true, errs
			}
			continue
		}

		noMoreValuesToEliminate = true
	}

	return false, errs
}

// handleAllCellsFilledCase verifies if sudoku is correctly solved. Precondition:
// this function expects that all cells of a sudoke are filled.
func handleAllCellsFilledCase(sudoku *models.Sudoku, errs []error) (bool, []error) {
	ruleValidationSuccess, err := validateSudokuRules(sudoku)
	if err != nil {
		errs = append(errs, err)
		return false, errs
	}

	if ruleValidationSuccess {
		return true, errs
	} else {
		errs = addUnsolvableSudokuError(errs)
		return false, errs
	}
}

// addUnsolvableSudokuError simply adds unsolvable sudoku error to errors slice.
func addUnsolvableSudokuError(errs []error) []error {
	return append(errs, errors.New("failed to solve the sudoku. Does it have a solution?"))
}
