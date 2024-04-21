package sudokuInit

import (
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestAssignSudokuReferences_Error(t *testing.T) {
	testCases := []struct {
		name              string
		sudokuInvalidator func(sudoku *models.Sudoku)
	}{
		{
			name: "Missing box",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes[0].IndexRow = 10
				sudoku.Boxes[0].IndexColumn = 10
			},
		},
		{
			name: "Missing cell",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes[0].Cells[0].IndexRowInBox = 10
				sudoku.Boxes[0].Cells[0].IndexColumnInBox = 10
			},
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		init := SudokuInit{Settings: settings}

		sudoku := getTestSudoku(t)
		init.initializeSubSudokus(sudoku)
		testCase.sudokuInvalidator(sudoku)
		err := init.assignSudokuReferences(sudoku)

		if err == nil {
			t.Errorf("%s: no validation errors", testCase.name)
		}
	}
}

func TestAssignSudokuReferences_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	init := SudokuInit{Settings: settings}

	sudoku := getTestSudoku(t)
	init.initializeSubSudokus(sudoku)
	err := init.assignSudokuReferences(sudoku)

	if err != nil {
		t.Errorf("Successfull sudoku references assign should not return any error. "+
			"But the error was returned: %s.", err)
	}
}
