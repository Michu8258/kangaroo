package dataWriter

import (
	"fmt"
	"os"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/services/printer"
	"github.com/Michu8258/kangaroo/testHelpers"
)

type fileWriteTestData[T interface{}] struct {
	name             string
	testData         T
	fileName         string
	overwrite        bool
	precreateTheFile bool
	expectedResult   bool
	expectsError     bool
}

func TestSaveSudokuToJson(t *testing.T) {
	testCases := []fileWriteTestData[*models.Sudoku]{
		{
			name:             "Success JSON new file",
			testData:         testHelpers.GetTestSudokuDto().ToSudoku(),
			fileName:         "test.json",
			overwrite:        false,
			precreateTheFile: false,
			expectedResult:   true,
			expectsError:     false,
		},
		{
			name:             "Success JSON file overwrite",
			testData:         testHelpers.GetTestSudokuDto().ToSudoku(),
			fileName:         "test.json",
			overwrite:        true,
			precreateTheFile: true,
			expectedResult:   true,
			expectsError:     false,
		},
		{
			name:             "JSON file already exists",
			testData:         testHelpers.GetTestSudokuDto().ToSudoku(),
			fileName:         "test.json",
			overwrite:        false,
			precreateTheFile: true,
			expectedResult:   false,
			expectsError:     false,
		},
	}

	for _, testCase := range testCases {
		genericWriteTest(t, "SaveSudokuToJson", testCase,
			func(writer IDataWriter, testData *models.Sudoku, path string, overwrite bool) (bool, error) {
				return writer.SaveSudokuToJson(testData, path, overwrite)
			})
	}
}

func TestSaveSudokuDtoToJson(t *testing.T) {
	testCases := []fileWriteTestData[*models.SudokuDTO]{
		{
			name:             "Success JSON new file",
			testData:         testHelpers.GetTestSudokuDto(),
			fileName:         "test.json",
			overwrite:        false,
			precreateTheFile: false,
			expectedResult:   true,
			expectsError:     false,
		},
		{
			name:             "Success JSON file overwrite",
			testData:         testHelpers.GetTestSudokuDto(),
			fileName:         "test.json",
			overwrite:        true,
			precreateTheFile: true,
			expectedResult:   true,
			expectsError:     false,
		},
		{
			name:             "JSON file already exists",
			testData:         testHelpers.GetTestSudokuDto(),
			fileName:         "test.json",
			overwrite:        false,
			precreateTheFile: true,
			expectedResult:   false,
			expectsError:     false,
		},
		{
			name:             "Invalid sudoku data",
			testData:         nil,
			fileName:         "test.json",
			overwrite:        false,
			precreateTheFile: false,
			expectedResult:   false,
			expectsError:     true,
		},
	}

	for _, testCase := range testCases {
		genericWriteTest(t, "SaveSudokuDtoToJson", testCase,
			func(writer IDataWriter, testData *models.SudokuDTO, path string, overwrite bool) (bool, error) {
				return writer.SaveSudokuDtoToJson(testData, path, overwrite)
			})
	}
}

func TestSaveSudokuToTxt(t *testing.T) {
	testCases := []fileWriteTestData[*models.Sudoku]{
		{
			name:             "Success TXT new file",
			testData:         testHelpers.GetTestSudokuDto().ToSudoku(),
			fileName:         "test.txt",
			overwrite:        false,
			precreateTheFile: false,
			expectedResult:   true,
			expectsError:     false,
		},
		{
			name:             "Success TXT file overwrite",
			testData:         testHelpers.GetTestSudokuDto().ToSudoku(),
			fileName:         "test.txt",
			overwrite:        true,
			precreateTheFile: true,
			expectedResult:   true,
			expectsError:     false,
		},
		{
			name:             "TXT file already exists",
			testData:         testHelpers.GetTestSudokuDto().ToSudoku(),
			fileName:         "test.txt",
			overwrite:        false,
			precreateTheFile: true,
			expectedResult:   false,
			expectsError:     false,
		},
	}

	for _, testCase := range testCases {
		genericWriteTest(t, "SaveSudokuToTxt", testCase,
			func(writer IDataWriter, testData *models.Sudoku, path string, overwrite bool) (bool, error) {
				return writer.SaveSudokuToTxt(testData, path, overwrite)
			})
	}
}

func genericWriteTest[T interface{}](t *testing.T, testGroupName string, testCaseData fileWriteTestData[T],
	testedFunc func(writer IDataWriter, testData T, path string, overwrite bool) (bool, error)) {

	setting := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	dataPrinter := dataPrinters.GetNewDataPrinter(setting, testPrinter)
	txtPrinter := testHelpers.NewTestPrinter()

	path := fmt.Sprintf("../../testFileWrites/%s", testCaseData.fileName)

	removeFileIfExists(path)
	defer removeFileIfExists(path)

	if testCaseData.precreateTheFile {
		file, err := os.Create(path)
		if err != nil {
			t.Errorf("%s - %s: failed to precreate the file '%s'.",
				testGroupName, testCaseData.name, path)
		}

		if file != nil {
			file.Close()
		}
	}

	writer := GetNewDataWriter(setting, dataPrinter,
		func(file *os.File) printer.IPrinter {
			return txtPrinter
		})

	result, err := testedFunc(writer, testCaseData.testData, path, testCaseData.overwrite)

	if testCaseData.expectedResult != result {
		t.Errorf("%s - %s: expected result %v, got %v.",
			testGroupName, testCaseData.name, testCaseData.expectedResult, result)
	}

	if testCaseData.expectsError && err == nil {
		t.Errorf("%s - %s: expected error but did not get one",
			testGroupName, testCaseData.name)
	}
}

func removeFileIfExists(path string) {
	if _, err := os.Stat(path); err == nil {
		os.Remove(path)
	}
}
