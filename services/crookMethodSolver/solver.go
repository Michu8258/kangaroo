package crookMethodSolver

import (
	"errors"
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// TODO - add documentation to all functions/methods
// TODO - add type for solution result
// TODO - add type for representing quess

// TODO - add documentation string
func SolveWithCrookMethod(sudoku *models.Sudoku, settings *models.Settings) (bool, []error) {
	return executeRecursiveSolution(sudoku, settings, false)
}

// TODO - add documentation string
func executeRecursiveSolution(sudoku *models.Sudoku, settings *models.Settings, isGuessing bool) (bool, []error) {
	errs := []error{}

	// simple sudokus that can be hamdled with pure elimination logic
	allCellsHaveValues, eliminationErrors := executeEliminationsLogic(sudoku, settings)
	errs = append(errs, eliminationErrors...)
	if len(errs) >= 1 {
		return false, errs // TODO result type
	}

	if allCellsHaveValues {
		solved, e := handleAllCellsFilledCase(sudoku, errs, isGuessing)
		if len(e) >= 1 {
			errs = append(errs, e...)
			return false, errs // TODO result type
		}

		if solved {
			return true, errs // TODO result type
		}
	}

	// preemptive sets (Crook)
	for {
		setManagedSuccessfully, atLeastOneCellWithNoPotentialValues, err :=
			executePreemptiveSetsLogic(sudoku, settings)
		if err != nil {
			errs = append(errs, err)
			return false, errs // TODO result type
		}

		atLeastOneValueAssigned := setManagedSuccessfully && assignCertainValues(sudoku, settings)

		if atLeastOneCellWithNoPotentialValues {
			if settings.UseDebugPrints {
				fmt.Println("At least one cell with no potential value found.")
			}
			break // todo add return with error here
			// TODO result type
		}

		if !setManagedSuccessfully && !atLeastOneValueAssigned {
			if settings.UseDebugPrints {
				fmt.Println("No preemptive set successfully processed (probably not found).")
			}
			break
		}

		if atLeastOneValueAssigned {
			return executeRecursiveSolution(sudoku, settings, false)
		}
	}

	// at this point, we have exhausted simple elimination method
	// and there are no cells wirh single potential value that
	// vould not violate sudoku rules. So we are guessing now.

	// for now we assume that the sudoku is not solved
	return false, errs // TODO result type
}

// executeEliminationsLogic executes simple elimination logic that may solve sudoku,
// but will not in case of difficult ones. It returns a boolean flag indicating
// if all cells has assigned certain values and slice of errors
func executeEliminationsLogic(sudoku *models.Sudoku, settings *models.Settings) (bool, []error) {
	errs := []error{}
	assignmentsExhausted := false

	for !assignmentsExhausted {
		//assign potential values
		errs = append(errs, assignCellsPotentialValues(sudoku, settings)...)
		if len(errs) >= 1 {
			return false, errs
		}

		// try to assign certain values
		atLeastOneValueAssigned := assignCertainValues(sudoku, settings)
		if atLeastOneValueAssigned {
			assignmentsExhausted = false
			allCellsFilled := checkIfAllCellsHaveValues(sudoku, settings)
			if allCellsFilled {
				return true, errs
			}
			continue
		}

		assignmentsExhausted = true
	}

	return false, errs
}

// handleAllCellsFilledCase verifies if sudoku is correctly solved. Precondition:
// this function expects that all cells of a sudoke are filled.
func handleAllCellsFilledCase(sudoku *models.Sudoku, errs []error, wasGuessing bool) (bool, []error) {
	ruleValidationSuccess, err := validateSudokuRules(sudoku)
	if err != nil {
		errs = append(errs, err)
		return false, errs
	}

	if ruleValidationSuccess {
		return true, errs
	} else {
		if !wasGuessing {
			errs = addUnsolvableSudokuError(errs)
		}
		return false, errs
	}
}

// addUnsolvableSudokuError simply adds unsolvable sudoku error to errors slice.
func addUnsolvableSudokuError(errs []error) []error {
	return append(errs, errors.New("failed to solve the sudoku. Does it have a solution?"))
}
