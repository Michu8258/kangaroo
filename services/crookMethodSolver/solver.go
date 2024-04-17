package crookMethodSolver

import (
	"fmt"
	"time"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
)

type sudokuRecursionData struct {
	Sudoku         *models.Sudoku
	Settings       *models.Settings
	IsGuessing     bool
	RecursionDepth int
}

type sudokuSolutionResult struct {
	ResultType models.SudokuResultType
	Errors     []error
}

// SolveWithCrookMethod tries to solve the sudoku puzzle by altering references which soduku model
// (a parameter) is build with. Returns a boolean flag indicating if solution was found and is
// correct, and slice of errors. Errors should not be printed to the user, they are actualy an
// errors.
func SolveWithCrookMethod(sudoku *models.Sudoku, settings *models.Settings,
	debugPrinter printer.Printer) (result bool, errors []error) {

	startTime := time.Now()

	defer func() {
		duration := time.Since(startTime)
		debugPrinter.PrintDefault(fmt.Sprintf(
			"CROOK's method solution duration: %v", duration))
		debugPrinter.PrintNewLine()

		if err := recover(); err != nil {
			result = false
			errors = append(errors, fmt.Errorf("fatal error: failed to execute Crook's alrogithm. "+
				"Underlying error: %s", err))
		}
	}()

	solutionResult := executeRecursiveSolution(sudokuRecursionData{
		Sudoku:         sudoku,
		Settings:       settings,
		IsGuessing:     false,
		RecursionDepth: 0,
	}, debugPrinter)

	sudoku.Result = solutionResult.ResultType

	return solutionResult.ResultType == models.SuccessfullSolution, solutionResult.Errors
}

// executeRecursiveSolution is the actual method that executes Sudoku puzzle solution with
// Crook's algorithm.  It returns and object with collections of errors and result status
// (successfull solution/failure/invalid guess/unsolvable sudoku)
func executeRecursiveSolution(recursionData sudokuRecursionData, debugPrinter printer.Printer) sudokuSolutionResult {
	defer func() {
		debugPrinter.PrintDefault(fmt.Sprintf(
			"REACHED END OF THE CROOK'S RECURSIVE METHOD. RETUNING UNSOLVABLE SUDOKU - DEPTH: %v",
			recursionData.RecursionDepth))
		debugPrinter.PrintNewLine()
	}()

	debugPrinter.PrintDefault(fmt.Sprintf(
		"RECURSIVE SOLUTION CROOK - DEPTH: %v", recursionData.RecursionDepth))
	debugPrinter.PrintNewLine()

	// simple sudokus that can be hamdled with pure elimination logic
	solved, shortCircuitResult, result := executeSimpleAlgorithm(recursionData, debugPrinter)
	if solved || shortCircuitResult || result.ResultType == models.InvalidGuess {
		return result
	}

	// preemptive sets (Crook)
	for {
		setManagedSuccessfully, atLeastOneCellWithNoPotentialValues, err :=
			executePreemptiveSetsLogic(recursionData.Sudoku, recursionData.Settings, debugPrinter)
		if err != nil {
			return sudokuSolutionResult{
				ResultType: models.Failure,
				Errors:     []error{err},
			}
		}

		atLeastOneValueAssigned := setManagedSuccessfully && assignCertainValues(
			recursionData.Sudoku, recursionData.Settings, debugPrinter)

		if atLeastOneCellWithNoPotentialValues {
			debugPrinter.PrintDefault("At least one cell with no potential value found.")
			debugPrinter.PrintNewLine()

			var result models.SudokuResultType = models.InvalidGuess
			if !recursionData.IsGuessing {
				result = models.UnsolvableSudoku
			}

			return sudokuSolutionResult{
				ResultType: result,
				Errors:     []error{fmt.Errorf("there is a call with no possible value to fill")},
			}
		}

		if !setManagedSuccessfully && !atLeastOneValueAssigned {
			debugPrinter.PrintDefault("No preemptive set successfully processed (probably not found).")
			debugPrinter.PrintNewLine()
			break
		}

		if atLeastOneValueAssigned {
			return executeRecursiveSolution(sudokuRecursionData{
				Sudoku:         recursionData.Sudoku,
				Settings:       recursionData.Settings,
				IsGuessing:     recursionData.IsGuessing,
				RecursionDepth: recursionData.RecursionDepth + 1,
			}, debugPrinter)
		}
	}

	// at this point, we have exhausted simple elimination method
	// and there are no cells with single potential value that
	// would not violate sudoku rules. So we are guessing now.
	for {
		cellToGuessExists, cellValueGuess, err := designateSudokuGuess(
			recursionData.Sudoku, debugPrinter)
		if err != nil {
			return sudokuSolutionResult{
				ResultType: models.Failure,
				Errors:     []error{err},
			}
		}

		// this means all cells have values assigned and we can validate sudoku rules and check if
		// we solved a sudoku
		if !cellToGuessExists {
			allCellsHaveValues := checkIfAllCellsHaveValues(recursionData.Sudoku, debugPrinter)
			ruleValidationNoError, err := validateSudokuRules(recursionData.Sudoku)
			if err != nil {
				return sudokuSolutionResult{
					ResultType: models.Failure,
					Errors:     []error{err},
				}
			}

			// if rule validation is successfull, we can assume sudoku is completely solved
			// becuase all cells have a values assigned.
			if allCellsHaveValues && ruleValidationNoError {
				return sudokuSolutionResult{
					ResultType: models.SuccessfullSolution,
					Errors:     []error{},
				}
			}

			var result models.SudokuResultType

			if recursionData.IsGuessing {
				result = models.InvalidGuess
			} else {
				result = models.Failure
			}

			return sudokuSolutionResult{
				ResultType: result,
				Errors:     []error{err},
			}
		}

		applySudokuValueGuess(cellValueGuess)
		nestedIterationResult := executeRecursiveSolution(sudokuRecursionData{
			Sudoku:         recursionData.Sudoku,
			Settings:       recursionData.Settings,
			IsGuessing:     true,
			RecursionDepth: recursionData.RecursionDepth + 1,
		}, debugPrinter)

		if nestedIterationResult.ResultType == models.InvalidGuess {
			err = restoreSnapshotFromGuessedValue(recursionData.Sudoku,
				cellValueGuess, debugPrinter)

			if err != nil {
				return sudokuSolutionResult{
					ResultType: models.Failure,
					Errors:     []error{err},
				}
			}
		} else {
			return nestedIterationResult
		}
	}
}

// executeSimpleAlgorithm executes single algorighm based on potential values
// eliminations. It is suitable only for simplest sudokus, it will not handle
// harder ones. Returns twoo booleans and a solution result. FIRST boolean
// indicates if successfull solution was found, SECONDS indicates wheather
// the result (third returned value) should be short circuited and returned
// immediately from calling function.
func executeSimpleAlgorithm(recursionData sudokuRecursionData,
	debugPrinter printer.Printer) (bool, bool, sudokuSolutionResult) {

	allCellsHaveValues, anyCellWithNoPotentialValues, errs := executeEliminationsLogic(
		recursionData.Sudoku, recursionData.Settings, debugPrinter)
	if len(errs) >= 1 {
		return false, true, sudokuSolutionResult{
			ResultType: models.Failure,
			Errors:     errs,
		}
	}

	if anyCellWithNoPotentialValues {
		if !recursionData.IsGuessing {
			return false, true, sudokuSolutionResult{
				ResultType: models.UnsolvableSudoku,
				Errors:     errs,
			}
		}

		return false, false, sudokuSolutionResult{
			ResultType: models.InvalidGuess,
			Errors:     errs,
		}
	}

	ruleValidationNoError, err := validateSudokuRules(recursionData.Sudoku)
	if err != nil {
		return false, true, sudokuSolutionResult{
			ResultType: models.Failure,
			Errors:     errs,
		}
	}

	// there is a sudoku rule validation error
	if !ruleValidationNoError {
		var result models.SudokuResultType

		if recursionData.IsGuessing {
			result = models.InvalidGuess
		} else {
			result = models.Failure
		}

		return false, result == models.Failure, sudokuSolutionResult{
			ResultType: result,
			Errors:     []error{err},
		}
	}

	if allCellsHaveValues && ruleValidationNoError {
		return true, true, sudokuSolutionResult{
			ResultType: models.SuccessfullSolution,
			Errors:     []error{},
		}
	}

	return false, false, sudokuSolutionResult{}
}

// executeEliminationsLogic executes simple elimination logic that may solve sudoku,
// but will not in case of difficult ones. It returns a pair of bools where FIRST
// boolean flag indicates if all cells has assigned certain values, SECOND indicates
// if there is at leas one cell with no potential values, and slice of errors
func executeEliminationsLogic(sudoku *models.Sudoku, settings *models.Settings,
	debugPrinter printer.Printer) (bool, bool, []error) {

	assignmentsExhausted := false

	for !assignmentsExhausted {
		//assign potential values
		anyCellWithNoPotentialValues, errs := assignCellsPotentialValues(
			sudoku, settings, debugPrinter)

		if len(errs) >= 1 {
			return false, false, errs
		}

		if anyCellWithNoPotentialValues {
			return false, true, []error{}
		}

		// try to assign certain values
		atLeastOneValueAssigned := assignCertainValues(
			sudoku, settings, debugPrinter)

		if atLeastOneValueAssigned {
			assignmentsExhausted = false
			allCellsFilled := checkIfAllCellsHaveValues(sudoku, debugPrinter)
			if allCellsFilled {
				return true, false, []error{}
			}
			continue
		}

		assignmentsExhausted = true
	}

	return false, false, []error{}
}

// applySudokuValueGuess applies guess sudoku value to the cell
func applySudokuValueGuess(cellValueGuess *models.SudokuValueGuess) {
	cellValueGuess.GuessedCell.PotentialValues = nil
	cellValueGuess.GuessedCell.Value = &cellValueGuess.GuessedValue
}
