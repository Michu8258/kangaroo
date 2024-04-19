package dataReader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
)

// ReadFromJsonFile reads raw sudoku data object from file with specified path.
// The path can be either relative (to main.go) or absolute.
func (reader *DataReader) ReadSudokuFromJsonFile(path string) (*models.SudokuDTO, error) {
	absolutePath, err := helpers.MakeFilePathAbsolute(path)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(absolutePath); err != nil {
		return nil, fmt.Errorf("sudoku input data file '%s' does not exist", absolutePath)
	}

	sudokuDataBytes, err := os.ReadFile(absolutePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read sudoku input data file '%s'", absolutePath)
	}

	sudoku := models.SudokuDTO{}
	err = json.Unmarshal(sudokuDataBytes, &sudoku)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sudoku input data file '%s'", absolutePath)
	}

	return &sudoku, nil
}
