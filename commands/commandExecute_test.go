package commands

import (
	"errors"
	"strings"
	"testing"

	"github.com/Michu8258/kangaroo/services"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/testHelpers"
	"github.com/urfave/cli/v2"
)

func TestExecuteCommand(t *testing.T) {
	testCases := []struct {
		name                 string
		arguments            []string
		decodeHasError       error
		encodeToBytesError   error
		encodeToBase64Error  error
		sudokuInitResult     bool
		sudokuInitErrors     []error
		sudokuSolutionResult bool
		sudokuSolutionErrors []error
		printContent         []string
	}{
		{
			name:                 "No data argument",
			arguments:            []string{"", "exec"},
			decodeHasError:       nil,
			encodeToBytesError:   nil,
			encodeToBase64Error:  nil,
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Please provide"},
		},
		{
			name:                 "Base64 parse error",
			arguments:            []string{"", "exec", "base64Config"},
			decodeHasError:       errors.New("decode error"),
			encodeToBytesError:   nil,
			encodeToBase64Error:  nil,
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Failed to parse"},
		},
		{
			name:                 "Sudoku init error",
			arguments:            []string{"", "exec", "base64Config"},
			decodeHasError:       nil,
			encodeToBytesError:   nil,
			encodeToBase64Error:  nil,
			sudokuInitResult:     false,
			sudokuInitErrors:     []error{errors.New("sudoku init error")},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{},
		},
		{
			name:                 "Sudoku solution error",
			arguments:            []string{"", "exec", "base64Config"},
			decodeHasError:       nil,
			encodeToBytesError:   nil,
			encodeToBase64Error:  nil,
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: false,
			sudokuSolutionErrors: []error{errors.New("sudoku solve error")},
			printContent:         []string{"Failed to solve"},
		},
		{
			name:                 "Solution encode error",
			arguments:            []string{"", "exec", "base64Config"},
			decodeHasError:       nil,
			encodeToBytesError:   nil,
			encodeToBase64Error:  errors.New("solution encoding error"),
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Failed to encode output"},
		},
		{
			name:                 "Success",
			arguments:            []string{"", "exec", "base64Config"},
			decodeHasError:       nil,
			encodeToBytesError:   nil,
			encodeToBase64Error:  nil,
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{""},
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		settings.UseDebugPrints = true
		testPrinter := testHelpers.NewTestPrinter()

		config := &CommandContext{
			Settings: settings,
			ServiceCollection: &services.ServiceCollection{
				DataPrinter:     dataPrinters.GetNewDataPrinter(settings, testPrinter),
				TerminalPrinter: testPrinter,
				SudokuInit: testHelpers.NewTestSudokuInit(
					testCase.sudokuInitResult, testCase.sudokuInitErrors),
				SudokuEncoder: testHelpers.NewTestBinarySudokuManager(
					testCase.decodeHasError, testCase.encodeToBase64Error, testCase.encodeToBytesError),
				Solver: testHelpers.GetNewTestSolver(
					testCase.sudokuSolutionResult, testCase.sudokuSolutionErrors),
			},
		}

		app := &cli.App{
			Name: "Kangaroo",
			Commands: []*cli.Command{
				config.ExecuteCommand(),
			},
		}

		err := app.Run(testCase.arguments)
		if err != nil {
			t.Error(err)
		}

		printed := false
		for _, expectedPrintout := range testCase.printContent {
			if !strings.Contains(testPrinter.PrintedData, expectedPrintout) {
				t.Errorf("%s: Console printout is missing the following: '%s'",
					testCase.name, expectedPrintout)
				if !printed {
					printed = true
					t.Error(testCase.printContent)
				}
			}
		}
	}
}
