package sudokuInit

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestValidateRawData_Error(t *testing.T) {
	testCases := []struct {
		name              string
		sudokuInvalidator func(sudoku *models.Sudoku)
	}{
		{
			name: "Box size to small",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.BoxSize = 1
			},
		},
		{
			name: "Box size to big",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.BoxSize = 100
			},
		},
		{
			name: "Layout width to small",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Layout.Width = 1
			},
		},
		{
			name: "Layout width to big",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Layout.Width = 100
			},
		},
		{
			name: "Layout height to small",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Layout.Height = 1
			},
		},
		{
			name: "Layout height to big",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Layout.Height = 100
			},
		},
		{
			name: "To little boxes count",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes = sudoku.Boxes[:len(sudoku.Boxes)-1]
			},
		},
		{
			name: "To much boxes count",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes = append(sudoku.Boxes, &models.SudokuBox{})
			},
		},
		{
			name: "To much boxes count",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes = append(sudoku.Boxes, &models.SudokuBox{})
			},
		},
		{
			name: "Box with invalid indexes",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes[8].IndexColumn = -1
				sudoku.Boxes[8].IndexRow = -1
			},
		},
		{
			name: "Two boxes with same indexes",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				// same as first box
				sudoku.Boxes[1].IndexColumn = 0
				sudoku.Boxes[1].IndexRow = 0
			},
		},
		{
			name: "To little cells in box",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes[0].Cells = sudoku.Boxes[0].Cells[:len(sudoku.Boxes[0].Cells)-1]
			},
		},
		{
			name: "To much cells in box",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes[0].Cells = append(sudoku.Boxes[0].Cells, &models.SudokuCell{})
			},
		},
		{
			name: "No expected cell (indexes)",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				sudoku.Boxes[0].Cells[4].IndexRowInBox = -1
				sudoku.Boxes[0].Cells[4].IndexColumnInBox = -1
			},
		},
		{
			name: "Invalid cell value",
			sudokuInvalidator: func(sudoku *models.Sudoku) {
				value := 100
				sudoku.Boxes[0].Cells[1].Value = &value
			},
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		init := SudokuInit{Settings: settings}

		sudoku := getTestSudoku(t)
		testCase.sudokuInvalidator(sudoku)

		errs := init.validateRawData(sudoku)

		if len(errs) < 1 {
			t.Errorf("%s: no validation errors", testCase.name)
		}
	}
}

func TestValidateRawData_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	init := SudokuInit{Settings: settings}

	sudoku := getTestSudoku(t)
	errs := init.validateRawData(sudoku)

	if len(errs) >= 1 {
		t.Errorf("Successfull sudoku raw data validation should not return any error. "+
			"Instead, %d errors were returned.", len(errs))
	}
}

func getTestSudoku(t *testing.T) *models.Sudoku {
	filePaths := "../../testConfigs/simple1.json"
	bytesData, err := os.ReadFile(filePaths)
	if err != nil {
		t.Errorf("Failed to read sudoku file '%s'", filePaths)
	}

	sudoku := models.Sudoku{}
	json.Unmarshal(bytesData, &sudoku)
	return &sudoku
}
