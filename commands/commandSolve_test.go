package commands

import (
	"errors"
	"strings"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/testHelpers"
	"github.com/urfave/cli/v2"
)

func TestSolveCommand(t *testing.T) {
	testCases := []struct {
		name                 string
		arguments            []string
		dataReaderResult     *models.SudokuDTO
		dataReaderError      error
		sudokuInitResult     bool
		sudokuInitErrors     []error
		sudokuSolutionResult bool
		sudokuSolutionErrors []error
		printContent         []string
	}{
		{
			name:                 "Invalid file",
			arguments:            []string{"", "solve", "-i", "/path/to/sudoku/data/file.json"},
			dataReaderResult:     nil,
			dataReaderError:      errors.New("sudoku data file read error"),
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Invalid sudoku input"},
		},
		{
			name:                 "Invalid user input",
			arguments:            []string{"", "solve"},
			dataReaderResult:     nil,
			dataReaderError:      errors.New("sudoku data file read error"),
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Invalid sudoku input"},
		},
		{
			name:                 "Failed sudoku initialization",
			arguments:            []string{"", "solve", "-i", "/path/to/sudoku/data/file.json"},
			dataReaderResult:     testHelpers.GetTestSudokuDto(),
			dataReaderError:      nil,
			sudokuInitResult:     false,
			sudokuInitErrors:     []error{errors.New("Failed to initialize sudoku")},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Invalid sudoku configuration"},
		},
		{
			name:                 "Solution of sudoku failed",
			arguments:            []string{"", "solve", "-s", "3", "--lw", "3", "--lh", "3", "-i", "/path/to/sudoku/data/file.json"},
			dataReaderResult:     testHelpers.GetTestSudokuDto(),
			dataReaderError:      nil,
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: false,
			sudokuSolutionErrors: []error{errors.New("Failed to solve sudoku")},
			printContent:         []string{"Failed to solve the sudoku"},
		},
		{
			name:                 "All good - with flags",
			arguments:            []string{"", "solve", "-s", "3", "--lw", "3", "--lh", "3", "-i", "/path/to/sudoku/data/file.json"},
			dataReaderResult:     testHelpers.GetTestSudokuDto(),
			dataReaderError:      nil,
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Sudoku puzzle solution"},
		},
		{
			name:                 "All good - no flags",
			arguments:            []string{"", "solve"},
			dataReaderResult:     testHelpers.GetTestSudokuDto(),
			dataReaderError:      nil,
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Sudoku puzzle solution"},
		},
		{
			name:                 "Output file save success",
			arguments:            []string{"", "solve", "-s", "3", "--lw", "3", "--lh", "3", "-r", "-i", "/path/to/sudoku/data/file.json", "-o", "/path/to/sudoku/sulution/file.json"},
			dataReaderResult:     testHelpers.GetTestSudokuDto(),
			dataReaderError:      nil,
			sudokuInitResult:     true,
			sudokuInitErrors:     []error{},
			sudokuSolutionResult: true,
			sudokuSolutionErrors: []error{},
			printContent:         []string{"Saving results"},
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		testPrinter := testHelpers.NewTestPrinter()
		debugPrinter := testHelpers.NewTestPrinter()

		config := &CommandContext{
			Settings: settings,
			ServiceCollection: &services.ServiceCollection{
				DataPrinter:     dataPrinters.GetNewDataPrinter(settings, testPrinter),
				TerminalPrinter: testPrinter,
				DebugPrinter:    debugPrinter,
				DataReader: testHelpers.NewTestDataReader(
					testCase.dataReaderResult, testCase.dataReaderError),
				SudokuInit: testHelpers.NewTestSudokuInit(
					testCase.sudokuInitResult, testCase.sudokuInitErrors),
				DataWriter: testHelpers.NewTestDataWriter(true, nil),
				Solver: testHelpers.GetNewTestSolver(
					testCase.sudokuSolutionResult, testCase.sudokuSolutionErrors),
			},
		}

		app := &cli.App{
			Name: "Kangaroo",
			Commands: []*cli.Command{
				config.SolveCommand(),
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
