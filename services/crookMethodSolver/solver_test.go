package crookMethodSolver

import (
	"encoding/json"
	"os"
	"slices"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/sudokuInit"
	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		sourceFilePath  string
		resultsFilePath string
	}{
		{
			sourceFilePath:  "../../testConfigs/simple1.json",
			resultsFilePath: "../../testConfigs/simple1_solution.json",
		},
		{
			sourceFilePath:  "../../testConfigs/medium1.json",
			resultsFilePath: "../../testConfigs/medium1_solution.json",
		},
		{
			sourceFilePath:  "../../testConfigs/hard1.json",
			resultsFilePath: "../../testConfigs/hard1_solution.json",
		},
		{
			sourceFilePath:  "../../testConfigs/hard2.json",
			resultsFilePath: "../../testConfigs/hard2_solution.json",
		},
		{
			sourceFilePath:  "../../testConfigs/5x5boxes.json",
			resultsFilePath: "../../testConfigs/5x5boxes_solution.json",
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		settings.UseDebugPrints = true
		debugPrinter := testHelpers.NewTestPrinter()

		source := getSudoku(t, testCase.sourceFilePath)
		expectedResult := getSudoku(t, testCase.resultsFilePath)

		initializer := sudokuInit.GetNewSudokuInit(settings)
		initializer.InitializeSudoku(source)

		solver := GetNewSudokuSolver(settings, debugPrinter)

		result, errors := solver.Solve(source)

		if !result {
			t.Errorf("Failed to solve the sudoku in file '%s'.", testCase.sourceFilePath)
		}

		if len(errors) > 0 {
			t.Errorf("Sudoku solution failed for file '%s'.", testCase.sourceFilePath)
			for _, err := range errors {
				t.Errorf("Sudoku sulution error: '%s'", err)
			}
		}

		areEqual := compareSudokus(t, expectedResult, source)
		if !areEqual {
			t.Errorf(
				"Sudoku has invalid solition. Source file: '%s', Expected result file: '%s'.",
				testCase.sourceFilePath, testCase.resultsFilePath)
		}
	}
}

func compareSudokus(t *testing.T, expected *models.Sudoku, actual *models.Sudoku) bool {
	for _, expectedBox := range expected.Boxes {
		boxIndex := slices.Index(expected.Boxes, expectedBox)
		actualBox := actual.Boxes[boxIndex]
		for _, expectedCell := range expectedBox.Cells {
			cellIndex := slices.Index(expectedBox.Cells, expectedCell)
			actualCell := actualBox.Cells[cellIndex]

			if actualCell.Value == nil && expectedCell.Value == nil {
				continue
			}

			if *expectedCell.Value != *actualCell.Value {
				t.Errorf("Not equal values. Expected %v, got %v. Box index: %d, cell index: %d.",
					expectedCell.Value, actualCell.Value,
					boxIndex, cellIndex)
				return false
			}
		}
	}

	return true
}

func getSudoku(t *testing.T, path string) *models.Sudoku {
	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("Failed to read sudoku test file '%s', err: '%s'.",
			path, err)
	}
	sudokuDto := models.SudokuDTO{}
	json.Unmarshal(jsonBytes, &sudokuDto)
	sudoku := sudokuDto.ToSudoku()
	return sudoku
}
