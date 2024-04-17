package sudokuInit

import "github.com/Michu8258/kangaroo/models"

type SudokuInit struct {
	Settings *models.Settings
}

type ISudokuInit interface {
	InitializeSudoku(sudoku *models.Sudoku) (bool, []error)
}

func GetNewSudokuInit(settings *models.Settings) ISudokuInit {
	return &SudokuInit{
		Settings: settings,
	}
}
