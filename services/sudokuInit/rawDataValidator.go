package sudokuInit

import (
	"fmt"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
)

// ValidateRawData validates raw data - that is before constructing full Sudoku
// object with SudokuLines and references assignment.
func (init *SudokuInit) validateRawData(sudoku *models.Sudoku) []error {
	errs := init.validateLayout(sudoku)
	if len(errs) >= 1 {
		return errs
	}

	errs = init.validateBoxesPresence(sudoku)
	if len(errs) >= 1 {
		return errs
	}

	errs = init.validateCellsInitialValues(sudoku)
	if len(errs) >= 1 {
		return errs
	}

	errs = init.initializeSubSudokus(sudoku)
	if len(errs) >= 1 {
		return errs
	}

	return errs
}

// validateLayout checks sudoku layout requirements - if box size is within the accepted
// range, layout shape, sudoku boxes and cells presence.
func (init *SudokuInit) validateLayout(sudoku *models.Sudoku) []error {
	errs := []error{}

	if sudoku.BoxSize < init.Settings.MinimumBoxSizeInclusive ||
		sudoku.BoxSize > init.Settings.MaximumBoxSizeInclusive {

		errs = append(errs, fmt.Errorf(
			"box size has a value of %d, but it is expected to be between %d and %d inclusively",
			sudoku.BoxSize,
			init.Settings.MinimumBoxSizeInclusive,
			init.Settings.MaximumBoxSizeInclusive))
	}

	widthError := init.validateLayoutSizeValue(sudoku.Layout.Width,
		init.Settings.MinimumLayoutSizeInclusive, init.Settings.MaximumLayoutSizeInclusive, "width")
	if widthError != nil {
		errs = append(errs, widthError)
	}

	heightError := init.validateLayoutSizeValue(sudoku.Layout.Height,
		init.Settings.MinimumLayoutSizeInclusive, init.Settings.MaximumLayoutSizeInclusive, "height")
	if heightError != nil {
		errs = append(errs, heightError)
	}

	if len(errs) >= 1 {
		// at this point furher processing is dangerous since we have
		// core settings outside of permitted range
		return errs
	}

	err := init.validateBoxesCount(sudoku)
	if err != nil {
		errs = append(errs, err)
	}

	return errs
}

// validateBoxesCount check if sudoku object contains exactly as many
// sudoku boxes as required by provided sudoku layout settings
func (init *SudokuInit) validateBoxesCount(sudoku *models.Sudoku) error {
	expectedBoxesCount := int(sudoku.Layout.Height * sudoku.Layout.Width)
	actualBoxesCount := len(sudoku.Boxes)
	if expectedBoxesCount != actualBoxesCount {
		return fmt.Errorf(
			"expected %d sudoku boxes, but %d boxes are provided",
			expectedBoxesCount,
			actualBoxesCount)
	}

	return nil
}

// validateBoxesPresence check if sudoku object contains all required sudoku boxes
// with correct indexes
func (init *SudokuInit) validateBoxesPresence(sudoku *models.Sudoku) []error {
	errs := []error{}

	var rowIndex, columnIndex int8
	for rowIndex = 0; rowIndex < sudoku.Layout.Height; rowIndex++ {
		for columnIndex = 0; columnIndex < sudoku.Layout.Width; columnIndex++ {
			box := sudoku.Boxes.FirstOrDefault(nil, func(b *models.SudokuBox) bool {
				return b.IndexRow == rowIndex && b.IndexColumn == columnIndex
			})

			if box == nil {
				errs = append(errs, fmt.Errorf(
					"sudoku box %s is missing",
					helpers.GetCoordinatesString(rowIndex+1, columnIndex+1, true)))

				continue
			}

			if box.Disabled {
				// if box is disabled we do not care for box's cells
				continue
			}

			cellError := init.validateCellsPresence(sudoku, box)
			if cellError != nil {
				errs = append(errs, cellError)
			}
		}
	}

	return errs
}

// validateCellsPresence check if sudoku box contains all required sudoku cells
// with correct indexes
func (init *SudokuInit) validateCellsPresence(sudoku *models.Sudoku,
	box *models.SudokuBox) error {

	expectedCellsCount := int(sudoku.BoxSize * sudoku.BoxSize)
	actualCellsCount := len(box.Cells)

	if actualCellsCount != expectedCellsCount {
		return fmt.Errorf(
			"invalid amount of cells for sudoku box %s. "+
				"Expected %d cells, got %d cells",
			helpers.GetCoordinatesString(box.IndexRow+1, box.IndexColumn+1, true),
			expectedCellsCount,
			actualCellsCount)
	}

	var rowIndex, columnIndex int8

	for rowIndex = 0; rowIndex < sudoku.BoxSize; rowIndex++ {
		for columnIndex = 0; columnIndex < sudoku.BoxSize; columnIndex++ {
			cell := box.Cells.FirstOrDefault(nil, func(c *models.SudokuCell) bool {
				return c.IndexRowInBox == rowIndex && c.IndexColumnInBox == columnIndex
			})

			if cell == nil {
				return fmt.Errorf(
					"sudoku box %s is missing a cell %s",
					helpers.GetCoordinatesString(rowIndex+1, columnIndex+1, true),
					helpers.GetCellCoordinatesString(sudoku, box, cell, true))
			}
		}
	}

	return nil
}

// validateCellsInitialValues check if all sudoku cell that contain a value have
// a value within accepted values range
func (init *SudokuInit) validateCellsInitialValues(sudoku *models.Sudoku) []error {
	errs := []error{}
	minimumValue := 1
	maximumValue := int(sudoku.BoxSize * sudoku.BoxSize)
	var boxRowIndex, boxColumnIndex int8

	for boxRowIndex = 0; boxRowIndex < sudoku.Layout.Height; boxRowIndex++ {
		for boxColumnIndex = 0; boxColumnIndex < sudoku.Layout.Width; boxColumnIndex++ {
			box := sudoku.Boxes.FirstOrDefault(nil, func(b *models.SudokuBox) bool {
				return b.IndexRow == boxRowIndex && b.IndexColumn == boxColumnIndex
			})

			for _, cell := range box.Cells {
				if cell.Value == nil {
					continue
				}

				if *cell.Value < minimumValue || *cell.Value > maximumValue {
					errs = append(errs, fmt.Errorf(
						"box %s has a cell %s with value outside of "+
							"allowed range. Got %d, acceptable values are from range "+
							"from %d to %d incllusive",
						helpers.GetBoxCoordinatesString(box, true),
						helpers.GetCellCoordinatesString(sudoku, box, cell, true),
						*cell.Value,
						minimumValue,
						maximumValue))
				}
			}
		}
	}

	return errs
}

// validateLayoutSizeValue check sudoku layout size requirements
func (init *SudokuInit) validateLayoutSizeValue(actualSize int8,
	minSize int8, maxSize int8, direction string) error {

	if actualSize < minSize || actualSize > maxSize {
		return fmt.Errorf(
			"the sudoku layout %s has a value of %d, but it must be between %d and %d inclusively",
			direction,
			actualSize,
			minSize,
			maxSize)
	}

	return nil
}
