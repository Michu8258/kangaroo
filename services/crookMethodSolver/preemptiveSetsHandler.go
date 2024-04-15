package crookMethodSolver

import (
	"fmt"
	"slices"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/types"
	guid "github.com/nu7hatch/gouuid"
)

type preemptiveSet struct {
	CellsInSet           types.GenericSlice[*models.SudokuCell]
	WholeCollectionCells types.GenericSlice[*models.SudokuCell]
	Values               []int
}

// executePreemptiveSetsLogic searches for preemptive sets and if finds any, it is also
// managed - rest of the cells within the cells collection (box, row, column) will be
// managed in the sence of modifying (truncating) slice of potential values.
//
// # Returns (bool, bool, error), where values means the following
//
// - set management success -  true if and only if there was at least
// one preemptive set found and managed without any error
//
// - empty possible values collection in any cell - after managing the
// preemptive set, and excluding values from sibling cells, if any of
// sibling cell would have empty potential values slice this will be
// true which indicates that sudoku puzzle is unsolvable
//
// - error if any occures
func executePreemptiveSetsLogic(sudoku *models.Sudoku, settings *models.Settings) (bool, bool, error) {
	anyPreemptiveSetHandled := false
	anyCellWithEmptyPotentialValues := false

	if settings.UseDebugPrints {
		fmt.Println("Starting preemptive sets logic execution.")
	}

	for _, subSudoku := range sudoku.SubSudokus {
		// for every box in the subsudoku we want to take care of preemptive sets
		for _, subSudokuBox := range subSudoku.Boxes {
			// box itself
			boxSet := findShortestPreemptiveSet(subSudokuBox.Cells, settings, "box")
			if boxSet != nil {
				siblingWithNoPotentialValues, didModify := processPreemptiveSet(boxSet, settings)
				anyPreemptiveSetHandled = anyPreemptiveSetHandled || didModify
				anyCellWithEmptyPotentialValues = anyCellWithEmptyPotentialValues || siblingWithNoPotentialValues
			}

			// rows
			handleSuccess, missingPotentialValues, err := iterateBoxLines(sudoku,
				subSudoku, subSudokuBox, models.SudokuLineTypeRow, settings,
				func(cell *models.SudokuCell, lineIndex int8) bool {
					return cell.IndexRowInBox == lineIndex && cell.IndexColumnInBox == 0
				})
			if err != nil {
				return false, anyCellWithEmptyPotentialValues, err
			}

			anyPreemptiveSetHandled = anyPreemptiveSetHandled || handleSuccess
			anyCellWithEmptyPotentialValues = anyCellWithEmptyPotentialValues || missingPotentialValues

			// columns
			handleSuccess, missingPotentialValues, err = iterateBoxLines(sudoku,
				subSudoku, subSudokuBox, models.SudokuLineTypeColumn, settings,
				func(cell *models.SudokuCell, lineIndex int8) bool {
					return cell.IndexRowInBox == 0 && cell.IndexColumnInBox == lineIndex
				})
			if err != nil {
				return false, anyCellWithEmptyPotentialValues, err
			}

			anyPreemptiveSetHandled = anyPreemptiveSetHandled || handleSuccess
			anyCellWithEmptyPotentialValues = anyCellWithEmptyPotentialValues || missingPotentialValues
		}
	}

	if settings.UseDebugPrints {
		fmt.Printf("Finished preemptive sets logic execution. Any set processed: %v, "+
			"any cell with zero potential values: %v.\n",
			anyPreemptiveSetHandled, anyCellWithEmptyPotentialValues)
	}

	return anyPreemptiveSetHandled, anyCellWithEmptyPotentialValues, nil
}

// iterateBoxLines iterates through rows and columns of a box within a subsudoku. Finds and processes
// the preemptive sets. Two bolean flags and an error. FIRST flag indicates if the set was found and
// processed successfully. SECOND flag indicates emptiness of at least one sibling cell of
// cells slice containing the preemptive set. ERROR indicates an error occurence.
func iterateBoxLines(sudoku *models.Sudoku, subSudoku *models.SubSudoku, subSudokuBox *models.SudokuBox,
	lineType string, settings *models.Settings, cellFilter func(cell *models.SudokuCell, lineIndex int8) bool) (bool, bool, error) {

	anyPreemptiveSetHandled := false
	anyCellWithEmptyPotentialValues := false
	var lineIndex int8
	for lineIndex = 0; lineIndex < sudoku.BoxSize; lineIndex++ {
		firstCellInLine := subSudokuBox.Cells.FirstOrDefault(nil, func(c *models.SudokuCell) bool {
			return cellFilter(c, lineIndex)
		})

		if firstCellInLine == nil {
			return false, anyCellWithEmptyPotentialValues, fmt.Errorf(
				"could not find cell in sudoku %s. Sudoku box row index %d and column index %d",
				lineType, subSudokuBox.IndexRow, subSudokuBox.IndexColumn)
		}

		theLine := firstCellInLine.MemberOfLines.FirstOrDefault(nil, func(line *models.SudokuLine) bool {
			return line.SubsudokuId == subSudoku.Id && line.LineType == lineType
		})

		if theLine == nil {
			return false, anyCellWithEmptyPotentialValues, fmt.Errorf(
				"could not find sudoku %s. Sudoku box row index %d and column index %d",
				lineType, subSudokuBox.IndexRow, subSudokuBox.IndexColumn)
		}

		theSet := findShortestPreemptiveSet(theLine.Cells, settings, lineType)
		if theSet != nil {
			siblingWithNoPotentialValues, didModify := processPreemptiveSet(theSet, settings)
			anyPreemptiveSetHandled = anyPreemptiveSetHandled || didModify
			anyCellWithEmptyPotentialValues = anyCellWithEmptyPotentialValues || siblingWithNoPotentialValues
		}
	}

	return anyPreemptiveSetHandled, anyCellWithEmptyPotentialValues, nil
}

// findShortestPreemptiveSet finds the preemptive set in the cells colection - the
// shortest one. Returns preemptiveSet data if set was founc.
func findShortestPreemptiveSet(cells types.GenericSlice[*models.SudokuCell], settings *models.Settings,
	collectionType string) *preemptiveSet {

	preemptiveSetCells := types.GenericSlice[*models.SudokuCell]{}
	for _, cell := range cells {
		if cell.PotentialValues == nil {
			continue
		}

		siblingCellsWithPotentialValues := cells.Where(func(c *models.SudokuCell) bool {
			return c.Id != cell.Id && c.PotentialValues != nil && len(*c.PotentialValues) >= 1
		})

		if len(siblingCellsWithPotentialValues) < 1 {
			continue
		}

		for _, siblingCell := range siblingCellsWithPotentialValues {
			// checking if we have mathing potential values
			if cell.PotentialValues.EqualContent(*siblingCell.PotentialValues) {
				// in case we found the preemptive set, we also have to make sure that
				// ther is no other sibling cell with less amount of possible values.
				// this is because possible values of the sibling cell may be a subset
				// of possible values of cell that is consired part of a preemptive set
				if siblingCellsWithPotentialValues.Any(func(sibCell *models.SudokuCell) bool {
					return len(*sibCell.PotentialValues) < len(*cell.PotentialValues)
				}) {
					continue
				}

				// we also need to ommit cases where all cells (the current one and all
				// siblings with possible values) has exactly the same possible values.
				// that case is simply inconclusive
				if siblingCellsWithPotentialValues.All(func(sibCell *models.SudokuCell) bool {
					return sibCell.PotentialValues.EqualContent(*cell.PotentialValues)
				}) {
					continue
				}

				// if count of potential values is less that already found one, we want the shorter one (new match)
				if len(preemptiveSetCells) > 0 && len(*cell.PotentialValues) < len(*preemptiveSetCells[0].PotentialValues) {
					clear(preemptiveSetCells)
				}

				// add cells if not already in the preemptive set
				if !slices.Contains(preemptiveSetCells, cell) {
					preemptiveSetCells = append(preemptiveSetCells, cell)
				}

				if !slices.Contains(preemptiveSetCells, siblingCell) {
					preemptiveSetCells = append(preemptiveSetCells, siblingCell)
				}
			}
		}
	}

	if len(preemptiveSetCells) >= 1 {
		result := &preemptiveSet{
			CellsInSet:           preemptiveSetCells,
			WholeCollectionCells: cells,
			Values:               *preemptiveSetCells[0].PotentialValues,
		}

		if settings.UseDebugPrints {
			fmt.Printf(
				"Found the preemptive set in %s with values %v. Box absolute index (row: %d, column %d), "+
					"cell in the box index (row: %d, column %d). Cells Total: %d, cells in set: %d. "+
					"Collection type: '%s'.\n",
				collectionType, result.Values,
				result.CellsInSet[0].Box.IndexRow, result.CellsInSet[0].Box.IndexColumn,
				result.CellsInSet[0].IndexRowInBox, result.CellsInSet[0].IndexColumnInBox,
				len(result.WholeCollectionCells),
				len(result.CellsInSet),
				collectionType)
		}

		return result
	}

	return nil
}

// processPreemptiveSet processes preemptive set data - that is removes possible values
// appearing in preemptive set from slices of potential values of sibling sudoku cells.
// Returns pair of bools where FIRST is indicating if any of the sibling cell is left
// without any potential value, SECOND indicates if any cell's potential values was
// modified.
func processPreemptiveSet(preemptiveSet *preemptiveSet, settings *models.Settings) (bool, bool) {
	appliedAnyPotentialValuesChange := false
	anyCellWithEmptyPotentialValues := false
	preemptiveSetCellsIds := []guid.UUID{}
	for _, cell := range preemptiveSet.CellsInSet {
		preemptiveSetCellsIds = append(preemptiveSetCellsIds, cell.Id)
	}

	for _, cell := range preemptiveSet.WholeCollectionCells {
		if cell.PotentialValues != nil && !slices.Contains(preemptiveSetCellsIds, cell.Id) {
			truncatedPotentialValues := cell.PotentialValues.Where(func(potentialValue int) bool {
				return !slices.Contains(preemptiveSet.Values, potentialValue)
			})

			// in case there is not change in potential values in the cell
			// we may skip assignment
			if cell.PotentialValues.EqualContent(truncatedPotentialValues) {
				continue
			}

			if len(truncatedPotentialValues) < 1 {
				anyCellWithEmptyPotentialValues = true
				if settings.UseDebugPrints {
					fmt.Println(cell.PotentialValues, preemptiveSet.Values, truncatedPotentialValues)
					fmt.Println("removing potential values from sibling cell of preemptive " +
						"cells leads to leaving no potential values for the cell")
				}
			}

			if settings.UseDebugPrints {
				fmt.Printf("Replacing existing potential values %v, with truncagted slice %v. "+
					"Box absolute indexes (row: %d, column: %d), cell in box indexes (row: %d, column: %d).\n",
					*cell.PotentialValues, truncatedPotentialValues,
					cell.Box.IndexRow, cell.Box.IndexColumn,
					cell.IndexRowInBox, cell.IndexColumnInBox)
			}

			cell.PotentialValues = &truncatedPotentialValues
			appliedAnyPotentialValuesChange = true

			if settings.UseDebugPrints {
				fmt.Printf("Potential values of cell after replacement: %v.\n", *cell.PotentialValues)
			}
		}
	}

	return anyCellWithEmptyPotentialValues, appliedAnyPotentialValuesChange
}
