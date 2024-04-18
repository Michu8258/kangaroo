package commands

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestValidateDestinationFilePaths_AtLeastOnePathIsOk(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()

	config := &CommandContext{
		Settings: settings,
		ServiceCollection: &services.ServiceCollection{
			DataPrinter: dataPrinters.GetNewDataPrinter(settings, testPrinter),
		},
	}

	destinationFilePaths := []string{
		"./directory/validFile.json",
		"./directory/validFile.txt",
		"./directory/validFile",
		"/root/ok/path/valid",
		"/root/ok/path/invalid.extension",
	}

	expectedResult := []string{
		"./directory/validFile.json",
		"./directory/validFile.txt",
		"./directory/validFile.json",
		"/root/ok/path/valid.json",
	}

	result := config.validateDestinationFilePaths(destinationFilePaths...)

	if len(result) != len(expectedResult) {
		t.Errorf("Validation failed. Expected %d valid file paths, got %d.",
			len(expectedResult), len(result))
	}

	for _, expectedPath := range expectedResult {
		if !slices.Contains(result, expectedPath) {
			t.Errorf("Validation failed. Expected '%s' file path to be in the result.", expectedPath)
		}
	}
}

func TestValidateDestinationFilePaths_AllPathsInvalid(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()

	config := &CommandContext{
		Settings: settings,
		ServiceCollection: &services.ServiceCollection{
			DataPrinter:     dataPrinters.GetNewDataPrinter(settings, testPrinter),
			TerminalPrinter: testPrinter,
		},
	}

	destinationFilePaths := []string{
		"./directory/validFile.bad",
		"/root/directory/validFile.verybad",
	}

	result := config.validateDestinationFilePaths(destinationFilePaths...)

	if len(result) != 0 {
		t.Errorf("Validation failed. Expected %d valid file paths, got %d.", 0, len(result))
	}

	for _, invalidPath := range result {
		if !strings.Contains(testPrinter.PrintedData, invalidPath) {
			t.Errorf("Invalid file name '%s' should be printed out, nut it is not.", invalidPath)
		}
	}
}

func TestExecuteSudokuFilesSave(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()

	config := &CommandContext{
		Settings: settings,
		ServiceCollection: &services.ServiceCollection{
			TerminalPrinter: testPrinter,
			DataPrinter:     dataPrinters.GetNewDataPrinter(settings, testPrinter),
			DataWriter:      testHelpers.NewTestDataWriter(true, nil),
		},
	}

	path := "/some/path/to/file.txt"

	config.executeSudokuFilesSave(&models.Sudoku{}, &models.SudokuConfigRequest{}, []string{path})
	if !strings.Contains(testPrinter.PrintedData, "results") ||
		!strings.Contains(testPrinter.PrintedData, path) {
		t.Errorf("Failed to save files. Expected file '%s' to be in console printout but it is not.", path)
	}
}

func TestSaveSudokuToFile(t *testing.T) {
	cases := []struct {
		path        string
		writeResult bool
		writeError  error
	}{
		{
			path:        "./file.json",
			writeResult: true,
			writeError:  nil,
		},
		{
			path:        "./file.txt",
			writeResult: true,
			writeError:  nil,
		},
		{
			path:        "./file.json",
			writeResult: false,
			writeError:  nil,
		},
		{
			path:        "./file.json",
			writeResult: false,
			writeError:  errors.New("file write unexpected error"),
		},
	}

	for testIndex, testCase := range cases {
		settings := testHelpers.GetTestSettings()
		testPrinter := testHelpers.NewTestPrinter()

		config := &CommandContext{
			Settings: settings,
			ServiceCollection: &services.ServiceCollection{
				TerminalPrinter: testPrinter,
				DataPrinter:     dataPrinters.GetNewDataPrinter(settings, testPrinter),
				DataWriter:      testHelpers.NewTestDataWriter(testCase.writeResult, testCase.writeError),
			},
		}

		config.saveSudokuToFile(&models.Sudoku{}, &models.SudokuConfigRequest{}, testCase.path)

		if testCase.writeError != nil {
			errorMsg := testCase.writeError.Error()
			if !strings.Contains(testPrinter.PrintedData, errorMsg) {
				t.Errorf("%d: Expected error message '%s' to be printed to console but it is not.",
					testIndex, errorMsg)
			}
		}

		if testCase.writeError == nil && testCase.writeResult {
			if !strings.Contains(testPrinter.PrintedData, "success") ||
				!strings.Contains(testPrinter.PrintedData, testCase.path) {
				t.Errorf(
					"%d: Expected successfull write message to be printed for file '%s' but it is not.",
					testIndex, testCase.path)
			}
		}

		if testCase.writeError == nil && !testCase.writeResult {
			if !strings.Contains(testPrinter.PrintedData, "exists") ||
				!strings.Contains(testPrinter.PrintedData, testCase.path) {
				t.Errorf(
					"%d: Expected ommited message to be printed for file '%s' but it is not.",
					testIndex, testCase.path)
			}
		}
	}
}

func TestExecuteSudokuInitialization(t *testing.T) {
	testCases := []struct {
		sudokuPrintable bool
		initErrors      []error
	}{
		{
			sudokuPrintable: true,
			initErrors:      []error{},
		},
		{
			sudokuPrintable: false,
			initErrors: []error{
				errors.New("some unexpected error"),
			},
		},
		{
			sudokuPrintable: true,
			initErrors: []error{
				errors.New("some unexpected error"),
			},
		},
	}

	for testIndex, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		testPrinter := testHelpers.NewTestPrinter()

		config := &CommandContext{
			Settings: settings,
			ServiceCollection: &services.ServiceCollection{
				TerminalPrinter: testPrinter,
				DataPrinter:     dataPrinters.GetNewDataPrinter(settings, testPrinter),
				SudokuInit:      testHelpers.NewTestSudokuInit(testCase.sudokuPrintable, testCase.initErrors),
			},
		}

		sudoku, ok := config.executeSudokuInitialization(testHelpers.GetTestSudokuDto())

		if sudoku == nil {
			t.Errorf("%v: Sudoku pointer is nil. Expected non nil pointer to sudoku object.",
				testIndex)
		}

		if len(testCase.initErrors) > 0 {
			if !strings.Contains(testPrinter.PrintedData, "Invalid sudoku configuration") {
				t.Errorf("%v: There should be sudoku initialization errors printed to the console",
					testIndex)
			}

			if testCase.sudokuPrintable && !strings.Contains(testPrinter.PrintedData, "Invalid sudoku values") {
				t.Errorf("%v: There should be invalid sudoku data printed to the console",
					testIndex)
			}

			if ok {
				t.Errorf("%v: Sudoku initialization should fail, but it did not.", testIndex)
			}
		} else {
			if !ok {
				t.Errorf("%v: Sudoku initialization should pass, but it failed.", testIndex)
			}

			if !strings.Contains(testPrinter.PrintedData, "Provided sudoku input") {
				t.Errorf("%v: There should be sudoku input data printed to the console",
					testIndex)
			}

			if !strings.Contains(testPrinter.PrintedData, "Selected sudoku puzzle configuration") {
				t.Errorf("%v: There should be sudoku configuration printed to the console",
					testIndex)
			}
		}
	}
}
