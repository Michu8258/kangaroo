package initialization

import (
	"errors"
	"fmt"

	"github.com/Michu8258/kangaroo/models"
	"github.com/beevik/guid"
)

// InitializeSubSudokus sets sub-sudokus data in the main sudoku
// data structure - it finds all existing and settings-matching
// sub-sudocus.
func initializeSubSudokus(sudoku *models.Sudoku) []error {
	errs := []error{}

	// This is amount of boxes that need to appear next to each other
	// in the puzzle. So sub-sudoku will need to be a square of sudoku
	// boxes with this size.
	expectedSize := sudoku.BoxSize

	// Since sudoku puzzle may contain many sub-sudokus (and every sub-sudoku)
	// is a square of sudoku boxes, we are marking minimum and maximum box
	// absolute indexes for the top left box of a subsudoku. This will help us
	// check if rest of required boxes for potential sub sudoku are in place.
	var minimumRowIndex int8 = 0
	maximumRowIndex := sudoku.Layout.Height - expectedSize
	var minimumColumnIndex int8 = 0
	maximumColumnIndex := sudoku.Layout.Width - expectedSize

	if maximumRowIndex < 0 || maximumColumnIndex < 0 {
		errs = append(errs, fmt.Errorf(
			"no possibility to designate any sub-sudoku. Sub-sudoku cannot "+
				"be designated when box size is set to %d and sudoku layout "+
				"width is %d and height is %d",
			sudoku.BoxSize,
			sudoku.Layout.Width,
			sudoku.Layout.Height))

		return errs
	}

	for minimumRowIndex = 0; minimumRowIndex <= maximumRowIndex; minimumRowIndex++ {
		for minimumColumnIndex = 0; minimumColumnIndex <= maximumColumnIndex; minimumColumnIndex++ {
			err := addSubSudoku(sudoku, minimumRowIndex, minimumColumnIndex)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) >= 1 {
		return errs
	}

	err := validateBoxesAssignments(sudoku)
	if err != nil {
		errs = append(errs, err)
	}

	return errs
}

func addSubSudoku(sudoku *models.Sudoku, startRowIndex, startColumnIndex int8) error {
	topLeftSubSudokuBox := sudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
		return box.IndexRow == startRowIndex && box.IndexColumn == startColumnIndex
	})

	if topLeftSubSudokuBox == nil {
		return fmt.Errorf(
			"cannot locate sudoku box with row index %d and column index %d when attempting to build "+
				"sub-sudoku. Those are coordinates of top left sudoku box of potential sub-sudoku",
			startRowIndex, startColumnIndex)
	}

	if topLeftSubSudokuBox.Disabled {
		return nil
	}

	// now we have starting and ending index of sudoku boxes that should be a part
	// of considered sub-sudoku
	endRowIndex := startRowIndex + sudoku.BoxSize - 1
	endColumnIndex := startColumnIndex + sudoku.BoxSize - 1

	subSudokuBoxes := []*models.SudokuBox{}

	for boxRowIndex := startRowIndex; boxRowIndex <= endRowIndex; boxRowIndex++ {
		for boxColumnIndex := startColumnIndex; boxColumnIndex <= endColumnIndex; boxColumnIndex++ {
			potentialSubSudokuBox := sudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
				return box.IndexRow == boxRowIndex && box.IndexColumn == boxColumnIndex
			})

			if potentialSubSudokuBox == nil {
				return fmt.Errorf(
					"cannot locate sudoku box with row index %d and column index %d when attempting to find "+
						"potential sub-sudoku box",
					startRowIndex, startColumnIndex)
			}

			if potentialSubSudokuBox.Disabled {
				// if any of required boxes is disabled, it means this is not a possible
				// sub sudoku area, but this does not mean an error
				return nil
			}

			subSudokuBoxes = append(subSudokuBoxes, potentialSubSudokuBox)
		}
	}

	// we found all required boxes to build a sub-sudoku and all of them are enabled (NOT disabled)
	sudoku.SubSudokus = append(sudoku.SubSudokus, &models.SubSudoku{
		Id:                    *guid.New(),
		Boxes:                 subSudokuBoxes,
		TopLeftBoxRowIndex:    startRowIndex,
		TopLeftBoxColumnIndex: startColumnIndex,
		ChildLines:            []*models.SudokuLine{},
	})
	return nil
}

func validateBoxesAssignments(sudoku *models.Sudoku) error {
	// every not disabled box must appear in at least one sub-sudoku
	notDisabledBoxes := sudoku.Boxes.Where(func(box *models.SudokuBox) bool {
		return !box.Disabled
	})

	if len(notDisabledBoxes) < 1 {
		return errors.New("no not disabled box exists")
	}

	for _, box := range notDisabledBoxes {
		// we are searching for first sub-sudoku which has specific box in its boxes collection
		subSudoku := sudoku.SubSudokus.FirstOrDefault(nil, func(subSudoku *models.SubSudoku) bool {
			boxIsAMember := subSudoku.Boxes.Any(func(subSudokuBox *models.SudokuBox) bool {
				return subSudokuBox.Id == box.Id
			})

			return boxIsAMember
		})

		if subSudoku == nil {
			return fmt.Errorf(
				"found a sudoku box that is not a part of any sub-sudoku. "+
					"Box row index %d, box column index %d",
				box.IndexRow,
				box.IndexColumn)
		}
	}

	return nil
}
