package commands

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

// executeSudokuInitialization executes sudoku initialization (validation included)
// based of dto input object. If everything is OK, sudoku data will be printer.
// Returns mapped sudoku object and boolean flag indicating if everything is
// correct up to this point
func (commandConfig *CommandConfig) executeSudokuInitialization(
	sudokuDto *models.SudokuDTO) (*models.Sudoku, bool) {

	sudoku := sudokuDto.ToSudoku()
	isSudokuPrintable, errs := commandConfig.SudokuInit.InitializeSudoku(sudoku)

	if len(errs) >= 1 {
		commandConfig.DataPrinter.PrintErrors("Invalid sudoku configuration:", errs...)
		commandConfig.TerminalPrinter.PrintNewLine()
		if isSudokuPrintable {
			commandConfig.printSudoku("Invalid sudoku values", sudoku)
		}
		return sudoku, false
	}

	commandConfig.printSudokuConfig(sudoku)
	commandConfig.printSudoku("Provided sudoku input:", sudoku)
	commandConfig.TerminalPrinter.PrintNewLine()

	return sudoku, true
}

// printSudokuConfig prints sudoku configuration with provided printer
func (commandConfig *CommandConfig) printSudokuConfig(sudoku *models.Sudoku) {
	commandConfig.TerminalPrinter.PrintPrimary("Selected sudoku puzzle configuration:")
	commandConfig.TerminalPrinter.PrintNewLine()
	commandConfig.TerminalPrinter.PrintDefault(fmt.Sprintf("- sudoku box size %d", sudoku.BoxSize))
	commandConfig.TerminalPrinter.PrintNewLine()
	commandConfig.TerminalPrinter.PrintDefault(fmt.Sprintf("- sudoku layout width %d", sudoku.Layout.Width))
	commandConfig.TerminalPrinter.PrintNewLine()
	commandConfig.TerminalPrinter.PrintDefault(fmt.Sprintf("- sudoku layout height %d", sudoku.Layout.Width))
	commandConfig.TerminalPrinter.PrintNewLine()
	commandConfig.TerminalPrinter.PrintNewLine()
}

// printSudoku prints sudoku to standard out
func (commandConfig *CommandConfig) printSudoku(description string,
	sudoku *models.Sudoku) {

	commandConfig.TerminalPrinter.PrintPrimary(description)
	commandConfig.TerminalPrinter.PrintNewLine()
	commandConfig.DataPrinter.PrintSudoku(sudoku, commandConfig.TerminalPrinter)
}
