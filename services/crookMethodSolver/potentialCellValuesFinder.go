package crookMethodSolver

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
)

// assignCellsPotentialValues assigns potential sudoku cell values.
// Potential cell values are also known and referred tu under the name
// of sudoku cell mark up. Returns a flag indicating if any of the cells
// has empty slice of potential values, and errir if any ocured.
func (solver *CrookSolver) assignCellsPotentialValues(sudoku *models.Sudoku) (bool, []error) {

	errs := []error{}
	anyPotentialValuesSliceIsEmpty := false
	minimumValue := 1
	maximumValue := int(sudoku.BoxSize * sudoku.BoxSize)

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				// if cell value is a input one, we can skip checking
				if subSudokuBoxCell.Value != nil {
					subSudokuBoxCell.PotentialValues = nil
					continue
				}

				// looking in box containing given cell
				emptyPotVal, err := solver.findPotentialValuesForCell(
					sudoku,
					subSudokuBoxCell,
					subSudokuBoxCell.Box.Cells,
					minimumValue,
					maximumValue)

				if err != nil {
					errs = append(errs, err)
				}

				if emptyPotVal {
					anyPotentialValuesSliceIsEmpty = true
				}

				// then looking for every row/column (line) containing given cell in current subsudoku
				linesWithinSubsudoku := subSudokuBoxCell.MemberOfLines.Where(func(l *models.SudokuLine) bool {
					return l.SubsudokuId == subSudoku.Id
				})

				for _, subSudokuLine := range linesWithinSubsudoku {
					emptyPotVal, err = solver.findPotentialValuesForCell(
						sudoku,
						subSudokuBoxCell,
						subSudokuLine.Cells,
						minimumValue,
						maximumValue)

					if err != nil {
						errs = append(errs, err)
					}

					if emptyPotVal {
						anyPotentialValuesSliceIsEmpty = true
					}
				}
			}
		}

	}

	if solver.Settings.UseDebugPrints {
		solver.printPotentialValues(sudoku, "POTENTIAL VALUES FINDER")
	}

	return anyPotentialValuesSliceIsEmpty, errs
}

// findPotentialValuesForCell searches for potential values that could be assigned to the
// cell and stores those value as a slice reference inside cell object. Possible values merge
// is performed if the same cell will be iterated for the second and nth time. Returns
// boolean flag indicating if the given cell has empty potential values slice, and
// error if any occured during processing
func (solver *CrookSolver) findPotentialValuesForCell(sudoku *models.Sudoku, cell *models.SudokuCell,
	cellsCollection models.GenericSlice[*models.SudokuCell],
	minimumCellValue int, maximumCellValue int) (
	emptyPotentialValues bool, errorResult error) {

	// in cas something went wrong
	defer func() {
		if err := recover(); err != nil {
			errorResult = errors.New("fatal error, failed to find potential values for a cell")
		}
	}()

	siblingCellsWithValues := cellsCollection.Where(func(c *models.SudokuCell) bool {
		return c.Id != cell.Id && c.Value != nil
	})

	takenValues := make(models.GenericSlice[int], 0, len(siblingCellsWithValues))
	for _, siblingCell := range siblingCellsWithValues {
		takenValues = append(takenValues, *siblingCell.Value)
	}

	potentialValues := make(models.GenericSlice[int], 0, maximumCellValue+1-minimumCellValue-len(takenValues))
	for i := minimumCellValue; i <= maximumCellValue; i++ {
		if !slices.Contains(takenValues, i) {
			potentialValues = append(potentialValues, i)
		}
	}

	if cell.PotentialValues == nil {
		// this is first iteration for this cell
		cell.PotentialValues = &potentialValues
		solver.logNoPotentialValues(sudoku, cell)
		return len(potentialValues) == 0, nil
	}

	// in case of another iteration for same cell, we need to merge potential values
	// by taking a common items in both slices
	intersection := cell.PotentialValues.Intersect(potentialValues)
	cell.PotentialValues = &intersection
	solver.logNoPotentialValues(sudoku, cell)

	return len(intersection) == 0, nil
}

// logNoPotentialValues log information about no potential values in
func (solver *CrookSolver) logNoPotentialValues(sudoku *models.Sudoku, cell *models.SudokuCell) {
	if cell.PotentialValues != nil && len(*cell.PotentialValues) == 0 {
		solver.DebugPrinter.PrintDefault(fmt.Sprintf(
			"Found a cell %s with no potential values during assigning potential values.",
			helpers.GetCellCoordinatesString(sudoku, cell.Box, cell, true)))
		solver.DebugPrinter.PrintNewLine()
	}
}
