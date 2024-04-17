package crookMethodSolver

import (
	"fmt"
	"time"

	"github.com/Michu8258/kangaroo/models"
)

type sudokuRecursionData struct {
	Sudoku         *models.Sudoku
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
func (solver *CrookSolver) Solve(sudoku *models.Sudoku) (result bool, errors []error) {

	startTime := time.Now()

	defer func() {
		duration := time.Since(startTime)
		solver.DebugPrinter.PrintDefault(fmt.Sprintf(
			"CROOK's method solution duration: %v", duration))
		solver.DebugPrinter.PrintNewLine()

		if err := recover(); err != nil {
			result = false
			errors = append(errors, fmt.Errorf("fatal error: failed to execute Crook's alrogithm. "+
				"Underlying error: %s", err))
		}
	}()

	solutionResult := solver.executeRecursiveSolution(sudokuRecursionData{
		Sudoku:         sudoku,
		IsGuessing:     false,
		RecursionDepth: 0,
	})

	sudoku.Result = solutionResult.ResultType

	return solutionResult.ResultType == models.SuccessfullSolution, solutionResult.Errors
}

// executeRecursiveSolution is the actual method that executes Sudoku puzzle solution with
// Crook's algorithm.  It returns and object with collections of errors and result status
// (successfull solution/failure/invalid guess/unsolvable sudoku)
func (solver *CrookSolver) executeRecursiveSolution(recursionData sudokuRecursionData) sudokuSolutionResult {
	defer func() {
		solver.DebugPrinter.PrintDefault(fmt.Sprintf(
			"REACHED END OF THE CROOK'S RECURSIVE METHOD. RETUNING UNSOLVABLE SUDOKU - DEPTH: %v",
			recursionData.RecursionDepth))
		solver.DebugPrinter.PrintNewLine()
	}()

	solver.DebugPrinter.PrintDefault(fmt.Sprintf(
		"RECURSIVE SOLUTION CROOK - DEPTH: %v", recursionData.RecursionDepth))
	solver.DebugPrinter.PrintNewLine()

	// simple sudokus that can be hamdled with pure elimination logic
	solved, shortCircuitResult, result := solver.executeSimpleAlgorithm(recursionData)
	if solved || shortCircuitResult || result.ResultType == models.InvalidGuess {
		return result
	}

	// preemptive sets (Crook)
	for {
		setManagedSuccessfully, atLeastOneCellWithNoPotentialValues, err :=
			solver.executePreemptiveSetsLogic(recursionData.Sudoku)
		if err != nil {
			return sudokuSolutionResult{
				ResultType: models.Failure,
				Errors:     []error{err},
			}
		}

		atLeastOneValueAssigned := setManagedSuccessfully && solver.assignCertainValues(recursionData.Sudoku)

		if atLeastOneCellWithNoPotentialValues {
			solver.DebugPrinter.PrintDefault("At least one cell with no potential value found.")
			solver.DebugPrinter.PrintNewLine()

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
			solver.DebugPrinter.PrintDefault("No preemptive set successfully processed (probably not found).")
			solver.DebugPrinter.PrintNewLine()
			break
		}

		if atLeastOneValueAssigned {
			return solver.executeRecursiveSolution(sudokuRecursionData{
				Sudoku:         recursionData.Sudoku,
				IsGuessing:     recursionData.IsGuessing,
				RecursionDepth: recursionData.RecursionDepth + 1,
			})
		}
	}

	// at this point, we have exhausted simple elimination method
	// and there are no cells with single potential value that
	// would not violate sudoku rules. So we are guessing now.
	for {
		cellToGuessExists, cellValueGuess, err := solver.designateSudokuGuess(
			recursionData.Sudoku)
		if err != nil {
			return sudokuSolutionResult{
				ResultType: models.Failure,
				Errors:     []error{err},
			}
		}

		// this means all cells have values assigned and we can validate sudoku rules and check if
		// we solved a sudoku
		if !cellToGuessExists {
			allCellsHaveValues := solver.checkIfAllCellsHaveValues(recursionData.Sudoku)
			ruleValidationNoError, err := solver.validateSudokuRules(recursionData.Sudoku)
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

		solver.applySudokuValueGuess(cellValueGuess)
		nestedIterationResult := solver.executeRecursiveSolution(sudokuRecursionData{
			Sudoku:         recursionData.Sudoku,
			IsGuessing:     true,
			RecursionDepth: recursionData.RecursionDepth + 1,
		})

		if nestedIterationResult.ResultType == models.InvalidGuess {
			err = solver.restoreSnapshotFromGuessedValue(recursionData.Sudoku, cellValueGuess)
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
func (solver *CrookSolver) executeSimpleAlgorithm(recursionData sudokuRecursionData) (
	bool, bool, sudokuSolutionResult) {

	allCellsHaveValues, anyCellWithNoPotentialValues, errs := solver.
		executeEliminationsLogic(recursionData.Sudoku)
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

	ruleValidationNoError, err := solver.validateSudokuRules(recursionData.Sudoku)
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
func (solver *CrookSolver) executeEliminationsLogic(sudoku *models.Sudoku) (bool, bool, []error) {

	assignmentsExhausted := false

	for !assignmentsExhausted {
		//assign potential values
		anyCellWithNoPotentialValues, errs := solver.assignCellsPotentialValues(sudoku)

		if len(errs) >= 1 {
			return false, false, errs
		}

		if anyCellWithNoPotentialValues {
			return false, true, []error{}
		}

		// try to assign certain values
		atLeastOneValueAssigned := solver.assignCertainValues(sudoku)
		if atLeastOneValueAssigned {
			assignmentsExhausted = false
			allCellsFilled := solver.checkIfAllCellsHaveValues(sudoku)
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
func (solver *CrookSolver) applySudokuValueGuess(cellValueGuess *models.SudokuValueGuess) {
	cellValueGuess.GuessedCell.PotentialValues = nil
	cellValueGuess.GuessedCell.Value = &cellValueGuess.GuessedValue
}
