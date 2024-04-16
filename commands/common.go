package commands

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printers"
	"github.com/Michu8258/kangaroo/types"
)

// printSudokuConfig prints sudoku configuration with provided printer
func printSudokuConfig(sudoku *models.Sudoku, printer types.Printer) {
	printer.PrintPrimary("Selected sudoku puzzle configuration:")
	printer.PrintNewLine()
	printer.PrintDefault(fmt.Sprintf("- sudoku box size %d", sudoku.BoxSize))
	printer.PrintNewLine()
	printer.PrintDefault(fmt.Sprintf("- sudoku layout width %d", sudoku.Layout.Width))
	printer.PrintNewLine()
	printer.PrintDefault(fmt.Sprintf("- sudoku layout height %d", sudoku.Layout.Width))
	printer.PrintNewLine()
	printer.PrintNewLine()
}

// printSudoku prints sudoku to standard out
func printSudoku(description string, sudoku *models.Sudoku, settings *models.Settings, printer types.Printer) {
	printer.PrintPrimary(description)
	printer.PrintNewLine()
	printers.PrintSudoku(sudoku, settings, printer)
}
