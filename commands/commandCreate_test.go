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

func TestCreateCommand(t *testing.T) {
	testCases := []struct {
		name             string
		arguments        []string
		dataReaderResult *models.SudokuDTO
		dataReaderError  error
		sudokuInitResult bool
		sudokuInitErrors []error
		printContent     []string
	}{
		{
			name:             "Everything OK",
			arguments:        []string{"", "create", "-s", "3", "--lw", "4", "--lh", "5", "-r", "./relative/path/to/file.json"},
			dataReaderResult: testHelpers.GetTestSudokuDto(),
			dataReaderError:  nil,
			sudokuInitResult: true,
			sudokuInitErrors: []error{},
			printContent: []string{
				"Saving results:",
			},
		},
		{
			name:             "Everything OK - no flags",
			arguments:        []string{"", "create", "./relative/path/to/file.json"},
			dataReaderResult: testHelpers.GetTestSudokuDto(),
			dataReaderError:  nil,
			sudokuInitResult: true,
			sudokuInitErrors: []error{},
			printContent: []string{
				"Saving results:",
			},
		},
		{
			name:             "No destination path",
			arguments:        []string{"", "create", "-s", "3", "--lw", "4", "--lh", "5", "-r"},
			dataReaderResult: testHelpers.GetTestSudokuDto(),
			dataReaderError:  nil,
			sudokuInitResult: true,
			sudokuInitErrors: []error{},
			printContent: []string{
				"argument for output file",
			},
		},
		{
			name:             "Sudoku console read fail",
			arguments:        []string{"", "create", "-s", "3", "--lw", "4", "--lh", "5", "-r", "./relative/path/to/file.json"},
			dataReaderResult: nil,
			dataReaderError:  errors.New("failed to read sudoku from console"),
			sudokuInitResult: true,
			sudokuInitErrors: []error{},
			printContent: []string{
				"Invalid sudoku input",
			},
		},
		{
			name:             "No valid file save path",
			arguments:        []string{"", "create", "-s", "3", "--lw", "4", "--lh", "5", "-r", "./relative/path/to/file.badext"},
			dataReaderResult: testHelpers.GetTestSudokuDto(),
			dataReaderError:  nil,
			sudokuInitResult: true,
			sudokuInitErrors: []error{},
			printContent: []string{
				"files",
				"not supported",
			},
		},
		{
			name:             "Sudoku initialization error",
			arguments:        []string{"", "create", "-s", "3", "--lw", "4", "--lh", "5", "-r", "./relative/path/to/file.json"},
			dataReaderResult: testHelpers.GetTestSudokuDto(),
			dataReaderError:  nil,
			sudokuInitResult: false,
			sudokuInitErrors: []error{errors.New("failed to initialize sudoku")},
			printContent: []string{
				"ailed to initialize sudoku",
			},
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		testPrinter := testHelpers.NewTestPrinter()

		config := &CommandContext{
			Settings: settings,
			ServiceCollection: &services.ServiceCollection{
				DataPrinter:     dataPrinters.GetNewDataPrinter(settings, testPrinter),
				TerminalPrinter: testPrinter,
				DataReader:      testHelpers.NewTestDataReader(testCase.dataReaderResult, testCase.dataReaderError),
				SudokuInit:      testHelpers.NewTestSudokuInit(testCase.sudokuInitResult, testCase.sudokuInitErrors),
				DataWriter:      testHelpers.NewTestDataWriter(true, nil),
			},
		}

		app := &cli.App{
			Name: "Kangaroo",
			Commands: []*cli.Command{
				config.CreateCommand(),
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
