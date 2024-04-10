package initialization

import (
	"github.com/Michu8258/kangaroo/models"
)

// InitializeSudoku executes initialization of sudoku puzzle describing object.
// That includes: Precomputing initial data, assigning circular references in
// sudoku object, constructing subsudokus and validation of input data.
func InitializeSudoku(sudoku *models.Sudoku, settings *models.Settings) []error {
	errs := []error{}

	errs = append(errs, validateRawData(sudoku, settings)...)
	if len(errs) >= 1 {
		return errs
	}

	err := assignSudokuReferences(sudoku)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	errs = append(errs, validateSudokuValues(sudoku)...)
	if len(errs) >= 1 {
		return errs
	}

	return errs
}
