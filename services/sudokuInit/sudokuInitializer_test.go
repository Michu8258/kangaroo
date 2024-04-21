package sudokuInit

import (
	"testing"

	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestInitializeSudoku(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	sudoku := getTestSudoku(t)
	init := GetNewSudokuInit(settings)
	succes, errs := init.InitializeSudoku(sudoku)

	if len(errs) >= 1 {
		t.Errorf("Sudoku initialization errors count: %d", len(errs))
	}

	if !succes {
		t.Error("Sudoku initialization unsuccessfull.")
	}
}
