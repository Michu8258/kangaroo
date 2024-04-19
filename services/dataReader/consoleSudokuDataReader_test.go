package dataReader

import (
	"errors"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestReadSudokuFromConsole(t *testing.T) {
	testCases := []struct {
		name                    string
		requestConfig           *models.SudokuConfigRequest
		boxSizePromptError      bool
		layoutWidthPromptError  bool
		layoutHeightPromptError bool
		sudokuPromptError       bool
		expectsError            bool
	}{
		{
			name:                    "Box size prompt failure",
			requestConfig:           &models.SudokuConfigRequest{},
			boxSizePromptError:      true,
			layoutWidthPromptError:  false,
			layoutHeightPromptError: false,
			sudokuPromptError:       false,
			expectsError:            true,
		},
		{
			name:                    "Layout width size prompt failure",
			requestConfig:           &models.SudokuConfigRequest{},
			boxSizePromptError:      false,
			layoutWidthPromptError:  true,
			layoutHeightPromptError: false,
			sudokuPromptError:       false,
			expectsError:            true,
		},
		{
			name:                    "Layout height size prompt failure",
			requestConfig:           &models.SudokuConfigRequest{},
			boxSizePromptError:      false,
			layoutWidthPromptError:  false,
			layoutHeightPromptError: true,
			sudokuPromptError:       false,
			expectsError:            true,
		},
		{
			name:                    "Sudoku values prompt failure",
			requestConfig:           &models.SudokuConfigRequest{},
			boxSizePromptError:      false,
			layoutWidthPromptError:  false,
			layoutHeightPromptError: false,
			sudokuPromptError:       true,
			expectsError:            true,
		},
		{
			name:                    "Successfull case",
			requestConfig:           &models.SudokuConfigRequest{},
			boxSizePromptError:      false,
			layoutWidthPromptError:  false,
			layoutHeightPromptError: false,
			sudokuPromptError:       false,
			expectsError:            false,
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		terminalPrinter := testHelpers.NewTestPrinter()
		debugPrinter := testHelpers.NewTestPrinter()

		var sudokuError error = nil
		if testCase.sudokuPromptError {
			sudokuError = errors.New("sudoku prompt error")
		}

		boxSizePrompt := func(callIndex int) (int8, error) {
			if testCase.boxSizePromptError {
				return 0, errors.New("Box Size prompt error")
			}

			if testCase.requestConfig.BoxSize != nil {
				return *testCase.requestConfig.BoxSize, nil
			}

			return 3, nil
		}

		layoutPromptError := func(callIndex int) (int8, error) {
			if callIndex == 0 && testCase.layoutWidthPromptError {
				return 0, errors.New("Layout prompt error")
			}

			if callIndex == 0 && testCase.requestConfig.LayoutWidth != nil {
				return *testCase.requestConfig.LayoutWidth, nil
			}

			if callIndex == 1 && testCase.layoutHeightPromptError {
				return 0, errors.New("Layout prompt error")
			}

			if callIndex == 1 && testCase.requestConfig.LayoutHeight != nil {
				return *testCase.requestConfig.LayoutHeight, nil
			}

			return 3, nil
		}

		prompter := testHelpers.GetNewTestPrompter(&testHelpers.TestPrompterConfig{
			SelectError:              nil,
			SudokuPromptError:        sudokuError,
			SelectPromptFailEnforcer: nil,
			BoxSizePromptFunc:        &boxSizePrompt,
			LayoutSizePromptFunc:     &layoutPromptError,
		})

		dataReader := GetNewDataReader(settings, terminalPrinter, debugPrinter, prompter)
		sudokuDTO, error := dataReader.ReadSudokuFromConsole(testCase.requestConfig)

		hasError := error != nil
		if hasError != testCase.expectsError {
			t.Errorf("%s: invalid error state", testCase.name)
		}

		if !testCase.expectsError {
			if sudokuDTO == nil {
				t.Errorf("Sudoku DTO should not be nil pointer.")
			}
		}
	}
}
