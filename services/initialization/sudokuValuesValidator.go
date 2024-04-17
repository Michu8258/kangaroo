package initialization

import (
	"fmt"
	"slices"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
)

// validateSudokuValues checks if all sub sudokus boxes, rows and columns
// contain values in permitted values range (if any value provided),
// and values duplications
func validateSudokuValues(sudoku *models.Sudoku) []error {
	errs := []error{}

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			errs = append(errs, validateCellsCollection(
				sudoku,
				sudoku.BoxSize,
				subSudokuBox.Cells,
				"box",
				func() {
					subSudokuBox.ViolatesRule = true
				},
			)...)
		}

		for _, subSudokuLine := range subSudoku.ChildLines {
			errs = append(errs, validateCellsCollection(
				sudoku,
				sudoku.BoxSize,
				subSudokuLine.Cells,
				subSudokuLine.LineType,
				func() {
					subSudokuLine.ViolatesRule = true
				},
			)...)
		}
	}

	return errs
}

// validateCellsCollection check if every cell with value has a value within an expected range,
// and if the value is not duplicated within cells collection (box, row, columns).
func validateCellsCollection(sudoku *models.Sudoku, boxSize int8, cells models.GenericSlice[*models.SudokuCell],
	collectionType string, cellsErrorSetter func()) []error {

	errs := []error{}
	minimumCellValue := 1
	maximumCellValue := int(boxSize * boxSize)
	alreadyExistingValues := []int{}

	for _, cell := range cells {
		value := cell.Value
		if value == nil {
			continue
		}

		if *value < minimumCellValue || *value > maximumCellValue {
			errs = append(errs, fmt.Errorf(
				"invalid cell %s value. Value is %d, but must be in range %d to %d inclusively",
				helpers.GetCellCoordinatesString(sudoku, cell.Box, cell, true),
				*value,
				minimumCellValue, maximumCellValue))

			cellsErrorSetter()
		}

		if slices.Contains(alreadyExistingValues, *value) {
			errs = append(errs, fmt.Errorf(
				"duplicated cell %s value %d in %s",
				helpers.GetCellCoordinatesString(sudoku, cell.Box, cell, true),
				*value,
				collectionType))
			cellsErrorSetter()
		}

		alreadyExistingValues = append(alreadyExistingValues, *value)
	}

	return errs
}
