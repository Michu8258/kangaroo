package dataReader

import (
	"testing"

	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestReadSudokuFromJsonFile(t *testing.T) {
	testCases := []struct {
		name         string
		filePath     string
		expectsError bool
	}{
		{
			name:         "Non existing file",
			filePath:     "./file/to/non/existing/file.jpg",
			expectsError: true,
		},
		{
			name:         "No sudoku JSON file",
			filePath:     "../../testConfigs/invalid.txt",
			expectsError: true,
		},
		{
			name:         "Correct case",
			filePath:     "../../testConfigs/simple1.json",
			expectsError: false,
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		terminalPrinter := testHelpers.NewTestPrinter()
		debugPrinter := testHelpers.NewTestPrinter()
		prompter := testHelpers.GetNewTestPrompter(&testHelpers.TestPrompterConfig{
			SelectError:              nil,
			SudokuPromptError:        nil,
			SelectPromptFailEnforcer: nil,
			BoxSizePromptFunc:        nil,
			LayoutSizePromptFunc:     nil,
		})

		dataReader := GetNewDataReader(settings, terminalPrinter, debugPrinter, prompter)
		_, err := dataReader.ReadSudokuFromJsonFile(testCase.filePath)
		hasError := err != nil

		if testCase.expectsError != hasError {
			t.Errorf("%s: invalid error state", testCase.name)
		}
	}
}
