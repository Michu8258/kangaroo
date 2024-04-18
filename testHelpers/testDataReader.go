package testHelpers

import "github.com/Michu8258/kangaroo/models"

type TestDataReader struct {
	SudokuResult *models.SudokuDTO
	ErrorResult  error
}

func NewTestDataReader(sudokuResult *models.SudokuDTO, errorResult error) *TestDataReader {
	return &TestDataReader{
		SudokuResult: sudokuResult,
		ErrorResult:  errorResult,
	}
}

func (reader *TestDataReader) ReadSudokuFromConsole(request *models.SudokuConfigRequest) (*models.SudokuDTO, error) {
	return reader.SudokuResult, reader.ErrorResult
}

func (reader *TestDataReader) ReadSudokuFromJsonFile(path string) (*models.SudokuDTO, error) {
	return reader.SudokuResult, reader.ErrorResult
}
