package binarySudokuManager

import (
	"testing"

	"github.com/Michu8258/kangaroo/testHelpers"
)

type decodeErrorTestCase struct {
	name                 string
	dataBytesInvalidator func([]byte) []byte
}

var correctVersion1DataBytes []byte = []byte{
	0, 1,
	3,
	3, 3,
	255, 128,
	6, 0, 0, 0, 1, 0, 0, 0, 7,
	0, 0, 3, 2, 0, 5, 0, 0, 0,
	4, 0, 0, 0, 7, 0, 0, 0, 1,
	0, 0, 0, 9, 0, 0, 0, 4, 0,
	1, 8, 0, 0, 0, 0, 6, 0, 7,
	5, 0, 0, 0, 8, 0, 0, 0, 0,
	0, 0, 6, 0, 8, 0, 2, 0, 0,
	0, 0, 0, 0, 3, 0, 5, 6, 0,
	0, 0, 3, 0, 2, 0, 7, 0, 0,
}

var decodeErrorTestCases = []decodeErrorTestCase{
	{
		name: "No version data",
		dataBytesInvalidator: func(correctData []byte) []byte {
			return []byte{}
		},
	},
	{
		name: "Unsupported version",
		dataBytesInvalidator: func(correctData []byte) []byte {
			return []byte{255, 255}
		},
	},
	{
		name: "No box size",
		dataBytesInvalidator: func(correctData []byte) []byte {
			return []byte{0, 1}
		},
	},
	{
		name: "No layout data 1",
		dataBytesInvalidator: func(correctData []byte) []byte {
			return []byte{0, 1, 3}
		},
	},
	{
		name: "No layout data 2",
		dataBytesInvalidator: func(correctData []byte) []byte {
			return []byte{0, 1, 3, 3}
		},
	},
	{
		name: "No boxes enabled/disabled state data",
		dataBytesInvalidator: func(correctData []byte) []byte {
			return []byte{0, 1, 3, 3, 3}
		},
	},
	{
		name: "Not enough boxes enabled/disabled state data",
		dataBytesInvalidator: func(correctData []byte) []byte {
			return []byte{0, 1, 3, 3, 3, 255}
		},
	},
	{
		name: "Not enough cells data",
		dataBytesInvalidator: func(correctData []byte) []byte {
			return correctData[:30]
		},
	},
}

func TestReadFromBytes_Error(t *testing.T) {
	for _, testCase := range decodeErrorTestCases {
		settings := testHelpers.GetTestSettings()
		manager := GetNewBinarySudokuManager(settings)

		sudoku := testHelpers.GetTestSudokuDto()
		testValue := 1
		sudoku.Boxes[0].Cells[0].Value = &testValue
		input := testCase.dataBytesInvalidator(correctVersion1DataBytes)

		_, err := manager.ReadFromBytes(input)

		if err == nil {
			t.Errorf("%s - ReadFromBytes - expected error, but none returned", testCase.name)
		}
	}
}

func TestReadFromBytes_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	manager := GetNewBinarySudokuManager(settings)

	data := correctVersion1DataBytes[:]
	data[5] = 254 // making one box disabled
	sudokuDto, err := manager.ReadFromBytes(correctVersion1DataBytes)

	if err != nil {
		t.Errorf("ToBase64 - unexpected error: %s", err)
	}

	if sudokuDto.BoxSize != 3 || sudokuDto.Layout.Width != 3 ||
		sudokuDto.Layout.Height != 3 || len(sudokuDto.Boxes) != 9 {
		t.Error("ToBase64 - invalid sudoku DTO data")
	}
}

func TestReadFromBase64_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	manager := GetNewBinarySudokuManager(settings)

	input := "AAEDAwP/gAYAAAABAAAABwAAAwIABQAAAAQAAAAHAAAAAQAAAAkAAAAEAAEIAAAAAAYABwUAAAAIAAAAAAAABgAIAAIAAAAAAAADAAUGAAAAAwACAAcAAA=="
	sudokuDto, err := manager.ReadFromBase64(input)

	if err != nil {
		t.Errorf("ReadFromBase64 - unexpected error: %s", err)
	}

	if sudokuDto.BoxSize != 3 || sudokuDto.Layout.Width != 3 ||
		sudokuDto.Layout.Height != 3 || len(sudokuDto.Boxes) != 9 {
		t.Error("ReadFromBase64 - invalid sudoku DTO data")
	}
}
