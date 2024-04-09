package dataInputs

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Michu8258/kangaroo/models"
)

// ReadFromFile reads raw sudoku data object from file with specified path.
// The path can be either relative (to main.go) or absolute.
func ReadFromFile(path string) (*models.Sudoku, error) {
	sudokuDataBytes, err := os.ReadFile(path)
	if err != nil {
		log.Println("failed to read sudoku data file", err)
		return nil, err
	}
	sudoku := models.Sudoku{}

	err = json.Unmarshal(sudokuDataBytes, &sudoku)
	if err != nil {
		log.Println("failed to parse sudoku JSON data", err)
	}

	return &sudoku, nil
}
