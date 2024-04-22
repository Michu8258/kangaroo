package binarySudokuManager

import (
	"bytes"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
)

type encodeErrorTestCase struct {
	name                 string
	sudokuDtoInvalidator func(*models.SudokuDTO, *models.Settings)
}

var encodeErrorTestCases = []encodeErrorTestCase{
	{
		name: "Unsuported version",
		sudokuDtoInvalidator: func(sudokuDto *models.SudokuDTO, settings *models.Settings) {
			settings.SudokuBinaryEncoderVersion = 10000
		},
	},
	{
		name: "Missing box",
		sudokuDtoInvalidator: func(sudokuDto *models.SudokuDTO, settings *models.Settings) {
			sudokuDto.Boxes = sudokuDto.Boxes[2:]
		},
	},
	{
		name: "Missing cell",
		sudokuDtoInvalidator: func(sudokuDto *models.SudokuDTO, settings *models.Settings) {
			sudokuDto.Boxes[0].Cells = sudokuDto.Boxes[0].Cells[2:]
		},
	},
}

func TestToBase64_Error(t *testing.T) {
	for _, testCase := range encodeErrorTestCases {
		settings := testHelpers.GetTestSettings()
		manager := GetNewBinarySudokuManager(settings)

		sudoku := testHelpers.GetTestSudokuDto()
		testValue := 1
		sudoku.Boxes[0].Cells[0].Value = &testValue
		testCase.sudokuDtoInvalidator(sudoku, settings)

		_, err := manager.ToBase64(sudoku)

		if err == nil {
			t.Errorf("%s - ToBase64 - expected error, but none returned", testCase.name)
		}
	}
}

func TestToBase64_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	manager := GetNewBinarySudokuManager(settings)

	sudoku := testHelpers.GetTestSudokuDto()
	testValue := 1
	sudoku.Boxes[0].Cells[0].Value = &testValue

	_, err := manager.ToBase64(sudoku)

	if err != nil {
		t.Errorf("ToBase64 - unexpected error: %s", err)
	}
}

func TestToBytes_Error(t *testing.T) {
	for _, testCase := range encodeErrorTestCases {
		settings := testHelpers.GetTestSettings()
		manager := GetNewBinarySudokuManager(settings)

		sudoku := testHelpers.GetTestSudokuDto()
		testValue := 1
		sudoku.Boxes[0].Cells[0].Value = &testValue
		testCase.sudokuDtoInvalidator(sudoku, settings)

		_, err := manager.ToBytes(sudoku)

		if err == nil {
			t.Errorf("%s - ToBytes - expected error, but none returned", testCase.name)
		}
	}
}

func TestToBytes_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	manager := GetNewBinarySudokuManager(settings)

	// read root/documentation/binaryFormat.md to find out why
	expectedBytesLength := 88
	expectedConfigBytes := []byte{0, 1, 3, 3, 3, 255, 128}

	sudoku := testHelpers.GetTestSudokuDto()
	testValue := 1
	sudoku.Boxes[0].Cells[0].Value = &testValue

	dataBytes, err := manager.ToBytes(sudoku)
	actualConfigBytes := dataBytes[:len(expectedConfigBytes)]

	if err != nil {
		t.Errorf("ToBytes - unexpected error: %s", err)
	}

	if len(dataBytes) != expectedBytesLength {
		t.Errorf("Invalid binary data length. Expected %d bytes, got %d.",
			expectedBytesLength, len(dataBytes))
	}

	if !bytes.Equal(expectedConfigBytes, actualConfigBytes) {
		t.Errorf("Invalid binary sudoku configuration. Expected %d, got %d.",
			expectedConfigBytes, actualConfigBytes)
	}

}
