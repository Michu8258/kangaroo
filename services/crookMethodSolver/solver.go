package crookMethodSolver

import "github.com/Michu8258/kangaroo/models"

func SolveWithCrookMethod(sudoku *models.Sudoku, settings *models.Settings) []error {
	errs := []error{}

	errs = append(errs, assignCellsPotentialValues(sudoku, settings)...)
	if len(errs) >= 1 {
		return errs
	}

	return errs
}
