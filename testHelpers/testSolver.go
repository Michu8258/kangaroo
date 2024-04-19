package testHelpers

import "github.com/Michu8258/kangaroo/models"

type TestSolver struct {
	Result bool
	Errors []error
}

func GetNewTestSolver(result bool, errors []error) *TestSolver {
	return &TestSolver{
		Result: result,
		Errors: errors,
	}
}

func (solver *TestSolver) Solve(sudoku *models.Sudoku) (result bool, errors []error) {
	return solver.Result, solver.Errors
}
