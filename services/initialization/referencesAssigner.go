package initialization

import "github.com/Michu8258/kangaroo/models"

func assignSudokuReferences(sudoku *models.Sudoku) {
	assignBoxReferencesInCells(sudoku)
	buildMembersOfLines(sudoku)
}

func assignBoxReferencesInCells(sudoku *models.Sudoku) {
	for _, box := range sudoku.Boxes {
		for _, cell := range box.Cells {
			cell.Box = box
		}
	}
}

func buildMembersOfLines(sudoku *models.Sudoku) {

}
