package crookMethodSolver

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// printPotentialValues prints debug information to the console when called
func printPotentialValues(sudoku *models.Sudoku, title string) {
	fmt.Printf("==================== POTENTIAL VALUES | %s | ====================\n", title)

	cellValuePrinter := func(v *int) string {
		if v == nil {
			return "-"
		}

		return fmt.Sprintf("%v", *v)
	}

	potentialValuesPrinter := func(potentialValues *models.GenericSlice[int]) string {
		if potentialValues == nil {
			return "-"
		}

		return fmt.Sprintf("%v", *potentialValues)
	}

	var boxRowIndex int8 = 0
	var cellRowIndex int8 = 0
	var boxColumnIndex int8 = 0
	var cellColumnIndex int8 = 0

	for boxRowIndex = 0; boxRowIndex < sudoku.Layout.Height; boxRowIndex++ {
		for cellRowIndex = 0; cellRowIndex < sudoku.BoxSize; cellRowIndex++ {
			for boxColumnIndex = 0; boxColumnIndex < sudoku.Layout.Width; boxColumnIndex++ {
				sudokuBox := sudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
					return box.IndexColumn == int8(boxColumnIndex) && box.IndexRow == boxRowIndex
				})

				for cellColumnIndex = 0; cellColumnIndex < sudoku.BoxSize; cellColumnIndex++ {
					sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCell) bool {
						return cell.IndexColumnInBox == int8(cellColumnIndex) && cell.IndexRowInBox == cellRowIndex
					})

					representation := fmt.Sprintf("%s %v",
						cellValuePrinter(sudokuCell.Value),
						potentialValuesPrinter(sudokuCell.PotentialValues))

					if cellColumnIndex >= sudoku.BoxSize-1 {
						fmt.Printf("%-25s", representation)
					} else {
						fmt.Printf("%-20s", representation)
					}
				}
			}

			fmt.Println()
		}

		if boxRowIndex < sudoku.Layout.Height-1 {
			fmt.Println()
			// fmt.Println()
		}
	}

	fmt.Printf("================== POTENTIAL VALUES END | %s | ====================\n", title)
}
