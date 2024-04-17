package sudokuInit

import (
	"github.com/Michu8258/kangaroo/models"
)

// InitializeSudoku executes initialization of sudoku puzzle describing object.
// That includes: Precomputing initial data, assigning circular references in
// sudoku object, constructing subsudokus and validation of input data.
// Returns boolean flag indicating that sudoku is printable and collection of errors
func (init *SudokuInit) InitializeSudoku(sudoku *models.Sudoku) (bool, []error) {
	errs := []error{}

	errs = append(errs, init.validateRawData(sudoku)...)
	if len(errs) >= 1 {
		return false, errs
	}

	err := init.assignSudokuReferences(sudoku)
	if err != nil {
		errs = append(errs, err)
		return false, errs
	}

	errs = append(errs, init.validateSudokuValues(sudoku)...)
	if len(errs) >= 1 {
		return true, errs
	}

	return true, errs
}
