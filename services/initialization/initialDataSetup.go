package initialization

import (
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/types"
	"github.com/beevik/guid"
)

// InitializeRawData initializes raw sudoku object - potentialy straight from the
// user input. It assigns IDs to boxes and cells and clears data (sets to nil)
// that could be included in user data but are expected to be in program
// access only.
func initializeRawData(sudoku *models.Sudoku) error {
	assignIds(sudoku)
	clearUnwantedData(sudoku)

	return nil
}

func clearUnwantedData(sudoku *models.Sudoku) {
	sudoku.SubSudokus = []*models.SubSudoku{}
	for _, sudokuBox := range sudoku.Boxes {
		for _, sudokuCell := range sudokuBox.Cells {
			sudokuCell.PotentialValues = &types.GenericSlice[int]{}
			sudokuCell.Box = nil
			sudokuCell.MemberOfLines = types.GenericSlice[*models.SudokuLine]{}
		}
	}
}

func assignIds(sudoku *models.Sudoku) {
	for _, sudokuBox := range sudoku.Boxes {
		sudokuBox.Id = *guid.New()
		for _, sudokuCell := range sudokuBox.Cells {
			sudokuCell.Id = *guid.New()
		}
	}
}
