package crookMethodSolver

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/types"
)

// assignCellsPotentialValues assigns potential sudoku cell values.
// Potential cell values are also known and referred tu under the name
// of sudoku cell mark up.
func assignCellsPotentialValues(sudoku *models.Sudoku, settings *models.Settings) []error {
	errs := []error{}
	minimumValue := 1
	maximumValue := int(sudoku.BoxSize * sudoku.BoxSize)

	for _, subSudoku := range sudoku.SubSudokus {
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				// if cell value is a input one, we can skip checking
				if subSudokuBoxCell.Value != nil && subSudokuBoxCell.IsInputValue {
					continue
				}

				// looking in box containing given cell
				err := findPotentialValuesForCell(
					subSudokuBoxCell,
					subSudokuBoxCell.Box.Cells,
					minimumValue,
					maximumValue)

				if err != nil {
					errs = append(errs, err)
				}

				// then looking for every row/column (line) containing given cell
				for _, subSudokuLine := range subSudokuBoxCell.MemberOfLines {
					err = findPotentialValuesForCell(
						subSudokuBoxCell,
						subSudokuLine.Cells,
						minimumValue,
						maximumValue)

					if err != nil {
						errs = append(errs, err)
					}
				}
			}
		}

	}

	if settings.UseDebugPrints {
		printPotentialValues(sudoku)
	}

	return errs
}

func findPotentialValuesForCell(cell *models.SudokuCell, cellsCollection types.GenericSlice[*models.SudokuCell],
	minimumCellValue int, maximumCellValue int) (errorResult error) {
	// in cas something went wrong
	defer func() {
		if err := recover(); err != nil {
			errorResult = errors.New("fatal error, failed to find potential values for a cell")
		}
	}()

	siblingCellsWithValues := cellsCollection.Where(func(c *models.SudokuCell) bool {
		return c.Id != cell.Id && c.Value != nil
	})

	takenValues := make(types.GenericSlice[int], 0, len(siblingCellsWithValues))
	for _, siblingCell := range siblingCellsWithValues {
		takenValues = append(takenValues, *siblingCell.Value)
	}

	potentialValues := make(types.GenericSlice[int], 0, maximumCellValue+1-minimumCellValue-len(takenValues))
	for i := minimumCellValue; i <= maximumCellValue; i++ {
		if !slices.Contains(takenValues, i) {
			potentialValues = append(potentialValues, i)
		}
	}

	if cell.PotentialValues == nil {
		// this is first iteration for this cell
		cell.PotentialValues = &potentialValues
		return nil
	}

	// in case of another iteration for same cell, we need to merge potential values
	// by taking a common items in both slices
	intersection := cell.PotentialValues.Intersect(potentialValues)
	cell.PotentialValues = &intersection

	return nil
}

func printPotentialValues(sudoku *models.Sudoku) {
	fmt.Println("==================== POTENTIAL VALUES ====================")
	for subSudokuIndex, subSudoku := range sudoku.SubSudokus {
		cellCount := 0
		for _, subSudokuBox := range subSudoku.Boxes {
			for _, subSudokuBoxCell := range subSudokuBox.Cells {
				cellCount += 1
				fmt.Printf("Sub sudoku index: %d, cell count %d, box indices (row: %d, column: %d), "+
					"cell indices (row: %d, column: %d), potential values %v.\n",
					subSudokuIndex,
					cellCount,
					subSudokuBoxCell.Box.IndexRow, subSudokuBoxCell.Box.IndexColumn,
					subSudokuBoxCell.IndexRowInBox, subSudokuBoxCell.IndexColumnInBox,
					subSudokuBoxCell.PotentialValues)
			}
		}
	}
	fmt.Println("================== POTENTIAL VALUES END ==================")
}