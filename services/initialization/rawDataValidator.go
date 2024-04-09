package initialization

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// ValidateRawData validates raw data - that is before constructing full Sudoku
// object with SudokuLines and references assignment.
func validateRawData(sudoku *models.Sudoku, settings *models.Settings) []error {
	errs := validateLayout(sudoku, settings)
	if len(errs) >= 1 {
		return errs
	}

	errs = validateBoxesPresence(sudoku)
	if len(errs) >= 1 {
		return errs
	}

	errs = validateCellsInitialValues(sudoku)
	if len(errs) >= 1 {
		return errs
	}

	errs = initializeSubSudokus(sudoku)
	if len(errs) >= 1 {
		return errs
	}

	return errs
}

func validateLayout(sudoku *models.Sudoku, settings *models.Settings) []error {
	errs := []error{}

	if sudoku.BoxSize < settings.MinimumBoxSizeInclusive || sudoku.BoxSize > settings.MaximumBoxSizeInclusive {
		errs = append(errs, fmt.Errorf(
			"box size has a value of %d, but it is expected to be between %d and %d inclusively",
			sudoku.BoxSize,
			settings.MinimumBoxSizeInclusive,
			settings.MaximumBoxSizeInclusive))
	}

	widthError := validateLayoutSizeValue(sudoku.Layout.Width,
		settings.MinimumLayoutSizeInclusive, settings.MaximumLayoutSizeInclusive, "width")
	if widthError != nil {
		errs = append(errs, widthError)
	}

	heightError := validateLayoutSizeValue(sudoku.Layout.Height,
		settings.MinimumLayoutSizeInclusive, settings.MaximumLayoutSizeInclusive, "height")
	if heightError != nil {
		errs = append(errs, heightError)
	}

	if len(errs) >= 1 {
		// at this point furher processing is dangerous since we have
		// core settings outside of permitted range
		return errs
	}

	err := validateBoxesCount(sudoku)
	if err != nil {
		errs = append(errs, err)
	}

	return errs
}

func validateBoxesCount(sudoku *models.Sudoku) error {
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

func validateBoxesPresence(sudoku *models.Sudoku) []error {
	errs := []error{}

	var rowIndex, columnIndex int8
	for rowIndex = 0; rowIndex < sudoku.Layout.Height; rowIndex++ {
		for columnIndex = 0; columnIndex < sudoku.Layout.Width; columnIndex++ {
			box := sudoku.Boxes.FirstOrDefault(nil, func(b *models.SudokuBox) bool {
				return b.IndexRow == rowIndex && b.IndexColumn == columnIndex
			})

			if box == nil {
				errs = append(errs, fmt.Errorf(
					"sudoku box with row index %d and column index %d is missing",
					rowIndex,
					columnIndex))

				continue
			}

			if box.Disabled {
				// if box is disabled we do not care for box's cells
				continue
			}

			cellError := validateCellsPresence(sudoku, box)
			if cellError != nil {
				errs = append(errs, cellError)
			}
		}
	}

	return errs
}

func validateCellsPresence(sudoku *models.Sudoku, box *models.SudokuBox) error {
	expectedCellsCount := int(sudoku.BoxSize * sudoku.BoxSize)
	actualCellsCount := len(box.Cells)

	if actualCellsCount != expectedCellsCount {
		return fmt.Errorf(
			"invalid amount of cells for sudoku box with row index %d and column index %d. expected %d cells, got %d cells",
			box.IndexRow,
			box.IndexColumn,
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
					"sudoku box with row index %d and column index %d is missing a cell with row index %d and column index %d",
					box.IndexRow,
					box.IndexColumn,
					rowIndex,
					columnIndex)
			}
		}
	}

	return nil
}

func validateCellsInitialValues(sudoku *models.Sudoku) []error {
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
						"box with row index %d and column index %d has a cell "+
							"(cell row index %d and column index %d) with value outside of "+
							"allowed range. got %d, acceptable values are from range "+
							"from %d to %d inclisive",
						boxRowIndex,
						boxColumnIndex,
						cell.IndexRowInBox,
						cell.IndexColumnInBox,
						*cell.Value,
						minimumValue,
						maximumValue))
				}
			}
		}
	}

	return errs
}

func validateLayoutSizeValue(actualSize int8, minSize int8, maxSize int8, direction string) error {
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
