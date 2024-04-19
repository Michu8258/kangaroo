package dataPrinters

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestPrintSudoku(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()

	jsonBytes, _ := os.ReadFile("../../testConfigs/simple1.json")
	sudokuDto := models.SudokuDTO{}
	json.Unmarshal(jsonBytes, &sudokuDto)
	sudoku := sudokuDto.ToSudoku()

	// important part of the setup
	sudoku.Boxes[0].Disabled = true

	for _, cell := range sudoku.Boxes[1].Cells {
		cell.Box = &models.SudokuBox{
			ViolatesRule: true,
		}
		if cell.Value != nil {
			cell.IsInputValue = true
		}
	}

	for _, cell := range sudoku.Boxes[2].Cells {
		cell.IsInputValue = false
	}

	sudoku.Boxes[3].ViolatesRule = true
	for _, cell := range sudoku.Boxes[3].Cells {
		cell.IsInputValue = false
	}

	expectedLines := []string{
		"╔═══════════╦═══════════╦═══════════╗",
		"║           ║   │   │ 4 ║ 3 │ 1 │   ║",
		"║───────────║───────────║───────────║",
		"║           ║ 2 │ 7 │   ║ 5 │ 6 │   ║",
		"║───────────║───────────║───────────║",
		"║           ║   │ 5 │   ║   │   │   ║",
		"║═══════════╬═══════════╬═══════════║",
		"║ 9 │   │   ║ 5 │ 2 │ 7 ║   │   │ 1 ║",
		"║───────────║───────────║───────────║",
		"║ 5 │   │   ║   │   │ 6 ║   │   │   ║",
		"║───────────║───────────║───────────║",
		"║ 1 │ 7 │   ║   │   │   ║ 8 │ 5 │ 2 ║",
		"║═══════════╬═══════════╬═══════════║",
		"║   │ 3 │ 2 ║ 7 │   │   ║ 1 │ 4 │ 6 ║",
		"║───────────║───────────║───────────║",
		"║ 6 │   │   ║ 3 │ 4 │ 8 ║   │   │   ║",
		"║───────────║───────────║───────────║",
		"║   │   │ 4 ║   │   │   ║   │   │ 3 ║",
		"╚═══════════╩═══════════╩═══════════╝",
	}

	dataPrinter := GetNewDataPrinter(settings, testPrinter)
	dataPrinter.PrintSudoku(sudoku, testPrinter)

	for _, expectedLine := range expectedLines {
		if !strings.Contains(testPrinter.PrintedData, expectedLine) {
			t.Errorf(
				"Printed sudoku output does not contain required string: '%s'",
				expectedLine)
		}
	}
}
