package initialization

import (
	"github.com/Michu8258/kangaroo/models"
)

func InitializeSudoku(sudoku *models.Sudoku, settings *models.Settings) []error {
	err := initializeRawData(sudoku)
	if err != nil {
		return []error{err}
	}

	errors := validateRawData(sudoku, settings)
	if len(errors) >= 1 {
		return errors
	}

	assignSudokuReferences(sudoku)

	return []error{}
}
