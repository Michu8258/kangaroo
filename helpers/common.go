package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Michu8258/kangaroo/models"
)

// IntToInt8Pointer converts integer to pinter of int8
func IntToInt8Pointer(i int) *int8 {
	i8 := int8(i)
	return &i8
}

// MakeFilePathAbsolute checks if provided file path is absolute
// or not. If provided path is absolute, it will be returned, if
// it is not, the absolute path will be calculated and returned.
func MakeFilePathAbsolute(path string) (string, error) {
	var absoluteFilePath = path
	if !filepath.IsAbs(path) {
		workingDirectory, err := os.Getwd()
		if err != nil {
			return "", err
		}
		absoluteFilePath = filepath.Join(workingDirectory, path)
	}

	return absoluteFilePath, nil
}

// GetBoxCoordinatesString returns user friendly box coordinates as string
func GetBoxCoordinatesString(box *models.SudokuBox, withParentheses bool) string {
	rowNumber := box.IndexRow + 1
	columnNumber := box.IndexColumn + 1

	return GetCoordinatesString(rowNumber, columnNumber, withParentheses)
}

// GetCellCoordinatesString returns user friendly cell coordinates as string
func GetCellCoordinatesString(sudoku *models.Sudoku, box *models.SudokuBox, cell *models.SudokuCell,
	withParentheses bool) string {

	rowNumber := GetCellNumber(sudoku.BoxSize, box.IndexRow, cell.IndexRowInBox)
	columnNumber := GetCellNumber(sudoku.BoxSize, box.IndexColumn, cell.IndexColumnInBox)

	return GetCoordinatesString(rowNumber, columnNumber, withParentheses)
}

// GetCellNumber returns user friendly cell number
func GetCellNumber(boxSize, boxIndex, cellIndex int8) int8 {
	return boxIndex*boxSize + cellIndex + 1
}

// GetCoordinatesString provides formatted coordinates string
func GetCoordinatesString(rowNumber int8, columnNumber int8, withParentheses bool) string {
	if withParentheses {
		return fmt.Sprintf("(row: %d, column: %d)", rowNumber, columnNumber)
	}

	return fmt.Sprintf("row: %d, column: %d", rowNumber, columnNumber)
}
