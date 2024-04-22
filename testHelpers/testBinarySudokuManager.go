package testHelpers

import "github.com/Michu8258/kangaroo/models"

type TestBinarySudokuManager struct {
	readResultError error
	toBase64Error   error
	toBytes         error
}

func NewTestBinarySudokuManager(readResultError error, toBase64Error error,
	toBytes error) *TestBinarySudokuManager {
	return &TestBinarySudokuManager{
		readResultError: readResultError,
		toBase64Error:   toBase64Error,
		toBytes:         toBytes,
	}
}

func (manager *TestBinarySudokuManager) ReadFromBase64(base64Data string) (*models.SudokuDTO, error) {
	if manager.readResultError == nil {
		return GetTestSudokuDto(), nil
	}

	return nil, manager.readResultError
}

func (manager *TestBinarySudokuManager) ReadFromBytes(sudokuData []byte) (*models.SudokuDTO, error) {
	if manager.readResultError == nil {
		return GetTestSudokuDto(), nil
	}

	return nil, manager.readResultError
}

func (manager *TestBinarySudokuManager) ToBase64(sudokuDto *models.SudokuDTO) (string, error) {
	if manager.toBase64Error == nil {
		return "", nil
	}

	return "base64", manager.toBase64Error
}

func (manager *TestBinarySudokuManager) ToBytes(sudokuDto *models.SudokuDTO) ([]byte, error) {
	if manager.toBytes == nil {
		return []byte{}, nil
	}

	return []byte{1, 2, 3}, manager.toBytes
}
