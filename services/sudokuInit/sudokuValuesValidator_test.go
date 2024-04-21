package sudokuInit

import (
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestValidateSudokuValues_Error(t *testing.T) {
	testCases := []struct {
		name              string
		sudokuInvalidator func(sudoku *models.Sudoku)
	}{
		{
			name: "Two same values",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				value := 5
				sudoku.Boxes[0].Cells[0].Value = &value
				sudoku.Boxes[0].Cells[1].Value = &value
			},
		},
		{
			name: "Value outside of range",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				value := 10
				sudoku.Boxes[0].Cells[0].Value = &value
			},
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		init := SudokuInit{Settings: settings}

		sudoku := getTestSudoku(t)
		init.initializeSubSudokus(sudoku)
		init.assignSudokuReferences(sudoku)
		testCase.sudokuInvalidator(sudoku)
		err := init.validateSudokuValues(sudoku)

		if err == nil {
			t.Errorf("%s: no validation errors", testCase.name)
		}
	}
}

func TestValidateSudokuValues_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	init := SudokuInit{Settings: settings}

	sudoku := getTestSudoku(t)
	init.initializeSubSudokus(sudoku)
	init.assignSudokuReferences(sudoku)
	errs := init.validateSudokuValues(sudoku)

	if len(errs) >= 1 {
		t.Errorf("Successfull sudoku values validation should not return any error. "+
			"But %d errors were returned.", len(errs))
	}
}
