package commands

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/services/initialization"
	"github.com/Michu8258/kangaroo/services/printer"
)

// executeSudokuInitialization executes sudoku initialization (validation included)
// based of dto input object. If everything is OK, sudoku data will be printer.
// Returns mapped sudoku object and boolean flag indicating if everything is
// correct up to this point
func executeSudokuInitialization(sudokuDto *models.SudokuDTO, settings *models.Settings,
	printer printer.Printer) (*models.Sudoku, bool) {

	sudoku := sudokuDto.ToSudoku()
	isSudokuPrintable, errs := initialization.InitializeSudoku(sudoku, settings)

	if len(errs) >= 1 {
		dataPrinters.PrintErrors("Invalid sudoku configuration:", printer, errs...)
		printer.PrintNewLine()
		if isSudokuPrintable {
			printSudoku("Invalid sudoku values", sudoku, settings, printer)
		}
		return sudoku, false
	}

	printSudokuConfig(sudoku, printer)
	printSudoku("Provided sudoku input:", sudoku, settings, printer)
	printer.PrintNewLine()

	return sudoku, true
}

// printSudokuConfig prints sudoku configuration with provided printer
func printSudokuConfig(sudoku *models.Sudoku, printer printer.Printer) {
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
func printSudoku(description string, sudoku *models.Sudoku, settings *models.Settings, printer printer.Printer) {
	printer.PrintPrimary(description)
	printer.PrintNewLine()
	dataPrinters.PrintSudoku(sudoku, settings, printer)
}
