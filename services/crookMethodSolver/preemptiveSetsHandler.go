package crookMethodSolver

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/types"
	"github.com/beevik/guid"
)

// executePreemptiveSetsLogic searches for preemptive sets and if finds any, it is also
// managed - rest of the cells within the cells collection (box, row, column) will be
// managed in the sence of modifying (truncating) slice of potential values.
// Returns boolean and error. Bool will be true if and only if there was at least
// one preemptive set found and managed without error.
func executePreemptiveSetsLogic(sudoku *models.Sudoku) (bool, error) {
	anyPreemptiveSetHandled := false

	preemptiveSetResultHandler := func(preemptiveSetFound bool, err error) error {
		if err != nil {
			return err
		}

		if preemptiveSetFound {
			anyPreemptiveSetHandled = true
		}

		return nil
	}

	boxHandler := func(box *models.SudokuBox) error {
		return preemptiveSetResultHandler(managePreemptiveSetsForCellsCollection(box.Cells))
	}

	lineHandler := func(firstCellInLine *models.SudokuLine) error {
		return preemptiveSetResultHandler(managePreemptiveSetsForCellsCollection(firstCellInLine.Cells))
	}

	err := helpers.IterateSubSudokusBoxesRowsCells(
		sudoku,
		true,
		&boxHandler,
		&lineHandler,
		&lineHandler)

	if err != nil {
		return false, err
	}

	return anyPreemptiveSetHandled, nil
}

// managePreemptiveSetsForCellsCollection searches preemptive set inside cells collection (first
// one if many possible) and removes potential values contained in the preemptive set from potential
// values of cells that are not part of the preemtive set. Returns bool and error, the bool indicates
// wheather the preemptive set was found and successfully managed.
func managePreemptiveSetsForCellsCollection(cells types.GenericSlice[*models.SudokuCell]) (bool, error) {
	preemptiveSetCells := types.GenericSlice[*models.SudokuCell]{}
	for _, cell := range cells {
		for _, siblingCell := range cells {
			// inside this first if we know that cell and sibling cell are references to different
			// cells within cells collection and both of them have at least one potential value assigned
			if cell.Id != siblingCell.Id && cell.PotentialValues != nil && siblingCell.PotentialValues != nil {
				// now we are comparing if the sibling cell has the same potantial values
				if cell.PotentialValues.EqualContent(*siblingCell.PotentialValues) {
					// the preemptive set must be the smallest one, even if first find
					// has more potential values, we must choose the one with least
					// amount of potantial vlaues.
					if len(preemptiveSetCells) > 0 {
						fmt.Println("----", len(preemptiveSetCells), len(*cell.PotentialValues), len(*preemptiveSetCells[0].PotentialValues))
						if len(*cell.PotentialValues) < len(*preemptiveSetCells[0].PotentialValues) {
							fmt.Println("rrrrrrrrrrrrrrrrr")
						}
					}

					if len(preemptiveSetCells) > 0 && len(*cell.PotentialValues) < len(*preemptiveSetCells[0].PotentialValues) {
						fmt.Println("fwefwefw")
						preemptiveSetCells = types.GenericSlice[*models.SudokuCell]{}
					}

					// first match, we also add the cell
					if len(preemptiveSetCells) == 0 {
						preemptiveSetCells = append(preemptiveSetCells, cell)
					}

					// add sibling cell
					preemptiveSetCells = append(preemptiveSetCells, siblingCell)
				}
			}
		}

		// this means there is at least one cells pair with same potential values found
		// and we stop iterating because we only want to find one preemptive set (unique)
		// per cells collection
		// if len(preemptiveSetCells) >= 1 {
		// 	break
		// }
	}

	// no match found - no preemptive set
	if len(preemptiveSetCells) < 1 {
		return false, nil
	}

	// remove potential values included in preemptive set from cells which
	// are not a part of the given preemptive set
	preemptiveSet := *preemptiveSetCells[0].PotentialValues
	preemptiveSetCellsIds := []guid.Guid{}
	for _, cell := range preemptiveSetCells {
		preemptiveSetCellsIds = append(preemptiveSetCellsIds, cell.Id)
	}

	for _, cell := range cells {
		if cell.PotentialValues != nil && !slices.Contains(preemptiveSetCellsIds, cell.Id) {
			fmt.Println("aaaaa")
			truncatedPotentialValues := cell.PotentialValues.Where(func(potentialValue int) bool {
				return !slices.Contains(preemptiveSet, potentialValue)
			})

			if len(truncatedPotentialValues) < 1 {
				fmt.Println(cell.PotentialValues, preemptiveSet, truncatedPotentialValues)
				return false, errors.New("removing potential values from sibling cell " +
					"of preemptive cells leads to leaving no potential values for the cell")
			}

			cell.PotentialValues = &truncatedPotentialValues
		}
	}

	return true, nil
}
