package crookMethodSolver

import (
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
)

type CrookSolver struct {
	Settings     *models.Settings
	DebugPrinter printer.IPrinter
}

type ISudokuSolver interface {
	Solve(sudoku *models.Sudoku) (result bool, errors []error)
}

func GetNewSudokuSolver(settings *models.Settings, debugPrinter printer.IPrinter) ISudokuSolver {
	return &CrookSolver{
		Settings:     settings,
		DebugPrinter: debugPrinter,
	}
}
