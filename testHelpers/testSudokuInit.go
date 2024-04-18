package testHelpers

import "github.com/Michu8258/kangaroo/models"

type TestSudokuInit struct {
	ResultFlag   bool
	ResultErrors []error
}

func NewTestSudokuInit(result bool, resultErrors []error) *TestSudokuInit {
	return &TestSudokuInit{
		ResultFlag:   result,
		ResultErrors: resultErrors,
	}
}

func (init *TestSudokuInit) InitializeSudoku(sudoku *models.Sudoku) (bool, []error) {
	return init.ResultFlag, init.ResultErrors
}
