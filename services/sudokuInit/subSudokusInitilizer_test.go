package sudokuInit

import (
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
	guid "github.com/nu7hatch/gouuid"
)

func TestInitializeSubSudokus_Error(t *testing.T) {
	testCases := []struct {
		name              string
		sudokuInvalidator func(sudoku *models.Sudoku)
	}{
		{
			name: "box size to big",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.BoxSize = 100
			},
		},
		{
			name: "No top left subsudoku expected box",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes[0].IndexRow = 10
				sudoku.Boxes[0].IndexColumn = 10
			},
		},
		{
			name: "Missing box",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes = sudoku.Boxes[:len(sudoku.Boxes)-1]
			},
		},
		{
			name: "Additional box - not part of any subsudoku",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				id, _ := guid.NewV4()
				sudoku.Boxes = append(sudoku.Boxes, &models.SudokuBox{
					Disabled: false,
					Id:       *id,
				})
			},
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		init := SudokuInit{Settings: settings}

		sudoku := getTestSudoku(t)
		testCase.sudokuInvalidator(sudoku)

		errs := init.initializeSubSudokus(sudoku)

		if len(errs) < 1 {
			t.Errorf("%s: no validation errors", testCase.name)
		}
	}
}

func TestInitializeSubSudokus_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	init := SudokuInit{Settings: settings}

	sudoku := getTestSudoku(t)
	errs := init.initializeSubSudokus(sudoku)

	if len(errs) >= 1 {
		t.Errorf("Successfull subsudokus unitialization should not return any error. "+
			"Instead, %d errors were returned.", len(errs))
	}

	if len(sudoku.SubSudokus) != 1 {
		t.Errorf("Expected 1 subsudoku to be created, but none actualy were.")
	}
}
