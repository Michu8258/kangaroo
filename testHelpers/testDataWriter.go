package testHelpers

import "github.com/Michu8258/kangaroo/models"

type TestDataWriter struct {
	FileWrittenFlag bool
	Error           error
}

func NewTestDataWriter(fileWritten bool, err error) *TestDataWriter {
	return &TestDataWriter{
		FileWrittenFlag: fileWritten,
		Error:           err,
	}
}

func (writer *TestDataWriter) SaveSudokuToJson(sudoku *models.Sudoku, path string, overwrite bool) (bool, error) {
	return writer.FileWrittenFlag, writer.Error
}

func (writer *TestDataWriter) SaveSudokuDtoToJson(sudokuDto *models.SudokuDTO, path string, overwrite bool) (bool, error) {
	return writer.FileWrittenFlag, writer.Error
}

func (writer *TestDataWriter) SaveSudokuToTxt(sudoku *models.Sudoku, path string, overwrite bool) (bool, error) {
	return writer.FileWrittenFlag, writer.Error
}
