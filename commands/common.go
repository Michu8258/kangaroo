package commands

import (
	"fmt"
	"path/filepath"

	"github.com/Michu8258/kangaroo/models"
)

// validateDestinationFilePaths checks if all provided file names have no extension
// or have json or txt extension. Returns slice of valid names
func (commandConfig *CommandContext) validateDestinationFilePaths(
	destinationFilePaths ...string) []string {

	validPaths := []string{}
	errorPaths := []error{}

	for _, destinationPath := range destinationFilePaths {
		extension := filepath.Ext(destinationPath)
		if len(extension) < 1 {
			validPaths = append(validPaths, destinationPath+".json")
			continue
		}

		if extension == ".json" || extension == ".txt" {
			validPaths = append(validPaths, destinationPath)
			continue
		}

		errorPaths = append(errorPaths, fmt.Errorf("unsupported file extension for path '%s'", destinationPath))
	}

	if len(errorPaths) > 0 {
		commandConfig.ServiceCollection.DataPrinter.PrintErrors(
			"Optput files listed below are not supported", errorPaths...)
	}

	if len(validPaths) < 1 {
		commandConfig.ServiceCollection.TerminalPrinter.PrintError(
			"No supported file path to save sudoku data to.")
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
	}

	return validPaths
}

// executeSudokuFilesSave executes iterative sudoku files save with
// results printing uncluded
func (commandConfig *CommandContext) executeSudokuFilesSave(sudoku *models.Sudoku,
	request *models.SudokuConfigRequest, paths []string) {
	commandConfig.ServiceCollection.TerminalPrinter.PrintPrimary("Saving results:")
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
	for _, path := range paths {
		commandConfig.saveSudokuToFile(sudoku, request.AsConfigRequest(), path)
	}
}

// executes save to file logic
func (commandConfig *CommandContext) saveSudokuToFile(sudoku *models.Sudoku,
	request *models.SudokuConfigRequest, path string) {

	extension := filepath.Ext(path)

	var written bool
	var err error

	if extension == ".txt" {
		written, err = commandConfig.ServiceCollection.DataWriter.
			SaveSudokuToTxt(sudoku, path, request.Overwrite)
	} else {
		written, err = commandConfig.ServiceCollection.DataWriter.
			SaveSudokuToJson(sudoku, path, request.Overwrite)
	}

	if err != nil {
		commandConfig.ServiceCollection.TerminalPrinter.PrintError(fmt.Sprintf("- %s", err))
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		return
	}

	if written {
		commandConfig.ServiceCollection.TerminalPrinter.PrintSuccess(
			fmt.Sprintf("- '%s' written successfully", path))
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		return
	}

	commandConfig.ServiceCollection.TerminalPrinter.PrintDefault(
		fmt.Sprintf("- '%s' already exists (ommited)", path))
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
}

// executeSudokuInitialization executes sudoku initialization (validation included)
// based of dto input object. If everything is OK, sudoku data will be printed.
// Returns mapped sudoku object and boolean flag indicating if everything is
// correct up to this point
func (commandConfig *CommandContext) executeSudokuInitialization(
	sudokuDto *models.SudokuDTO) (*models.Sudoku, bool) {

	sudoku := sudokuDto.ToSudoku()
	isSudokuPrintable, errs := commandConfig.ServiceCollection.
		SudokuInit.InitializeSudoku(sudoku)

	if len(errs) >= 1 {
		commandConfig.ServiceCollection.DataPrinter.PrintErrors(
			"Invalid sudoku configuration:", errs...)
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		if isSudokuPrintable {
			commandConfig.printSudoku("Invalid sudoku values", sudoku)
		}
		return sudoku, false
	}

	commandConfig.printSudokuConfig(sudoku)
	commandConfig.printSudoku("Provided sudoku input:", sudoku)
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()

	return sudoku, true
}

// printSudokuConfig prints sudoku configuration with provided printer
func (commandConfig *CommandContext) printSudokuConfig(sudoku *models.Sudoku) {
	commandConfig.ServiceCollection.TerminalPrinter.PrintPrimary(
		"Selected sudoku puzzle configuration:")
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()

	commandConfig.ServiceCollection.TerminalPrinter.PrintDefault(
		fmt.Sprintf("- sudoku box size %d", sudoku.BoxSize))
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()

	commandConfig.ServiceCollection.TerminalPrinter.PrintDefault(
		fmt.Sprintf("- sudoku layout width %d", sudoku.Layout.Width))
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()

	commandConfig.ServiceCollection.TerminalPrinter.PrintDefault(
		fmt.Sprintf("- sudoku layout height %d", sudoku.Layout.Width))
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()

	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
}

// printSudoku prints sudoku to standard out
func (commandConfig *CommandContext) printSudoku(description string,
	sudoku *models.Sudoku) {

	commandConfig.ServiceCollection.TerminalPrinter.PrintPrimary(description)
	commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
	commandConfig.ServiceCollection.DataPrinter.PrintSudoku(
		sudoku, commandConfig.ServiceCollection.TerminalPrinter)
}
