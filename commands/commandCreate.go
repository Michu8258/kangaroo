package commands

import (
	"fmt"
	"path/filepath"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	"github.com/urfave/cli/v2"
)

// CreateCommand provides create command configuration
func (commandConfig *CommandConfig) CreateCommand() *cli.Command {
	return &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage: "Creates sudoku puzzle data and saves to provided " +
			"file paths (JSON and TXT files supported, default is JSON)",
		Flags: []cli.Flag{
			&boxSizeFlag,
			&layoutWidthFlag,
			&layoutHeightFlag,
			&overwriteFileFlag,
		},
		Action: func(context *cli.Context) error {
			request := commandConfig.buildCreateCommandRequest(context)
			filePaths := context.Args().Slice()
			return commandConfig.createCommandHandler(request, filePaths)
		},
	}
}

// createCommandHandler is an entry point function to create sudoku data file
func (commandConfig *CommandConfig) createCommandHandler(request *models.CreateCommandRequest,
	destinationFilePaths []string) error {

	if len(destinationFilePaths) < 1 {
		commandConfig.TerminalPrinter.PrintError(
			"Please provide at least one argument for output file location.")
		commandConfig.TerminalPrinter.PrintNewLine()
		return nil
	}

	validPaths, errorPaths := commandConfig.validateDestinationFilePaths(destinationFilePaths)
	if len(errorPaths) > 0 {
		commandConfig.DataPrinter.PrintErrors(
			"Optput files listed below are not supported", errorPaths...)
	}

	if len(validPaths) < 1 {
		commandConfig.TerminalPrinter.PrintError(
			"No supported file path to save sudoku data to.")
		commandConfig.TerminalPrinter.PrintNewLine()
		return nil
	}

	sudokuDto, err := commandConfig.DataReader.ReadSudokuFromConsole(request.AsConfigRequest())
	if err != nil {
		commandConfig.DataPrinter.PrintErrors("Invalid sudoku input", err)
		return nil
	}

	sudoku, ok := commandConfig.executeSudokuInitialization(sudokuDto)
	if !ok {
		return nil
	}

	commandConfig.TerminalPrinter.PrintPrimary("Saving results:")
	commandConfig.TerminalPrinter.PrintNewLine()
	for _, path := range validPaths {
		commandConfig.saveToFile(sudoku, request, path)
	}

	return nil
}

// executes save to file logic
func (commandConfig *CommandConfig) saveToFile(sudoku *models.Sudoku,
	request *models.CreateCommandRequest, path string) {

	extension := filepath.Ext(path)

	var written bool
	var err error

	if extension == ".txt" {
		written, err = commandConfig.DataWriter.SaveSudokuToTxt(sudoku, path, request.Overwrite)
	} else {
		written, err = commandConfig.DataWriter.SaveSudokuToJson(sudoku, path, request.Overwrite)
	}

	if err != nil {
		commandConfig.TerminalPrinter.PrintError(fmt.Sprintf("- %s", err))
		commandConfig.TerminalPrinter.PrintNewLine()
		return
	}

	if written {
		commandConfig.TerminalPrinter.PrintSuccess(fmt.Sprintf("- '%s' written successfully", path))
		commandConfig.TerminalPrinter.PrintNewLine()
		return
	}

	commandConfig.TerminalPrinter.PrintDefault(fmt.Sprintf("- '%s' already exists (ommited)", path))
	commandConfig.TerminalPrinter.PrintNewLine()
}

// validateDestinationFilePaths checks if all provided file names have no extension
// or have json or txt extension. Returns slice ov valid names and slice of errors
// for those not matching the criteria.
func (commandConfig *CommandConfig) validateDestinationFilePaths(
	destinationFilePaths []string) ([]string, []error) {

	validPaths := []string{}
	errs := []error{}

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

		errs = append(errs, fmt.Errorf("unsupported file extension for path '%s'", destinationPath))
	}

	return validPaths, errs
}

// buildCreateCommandRequest retrieves options settings from the command
// and constructs request object.
func (commandConfig *CommandConfig) buildCreateCommandRequest(
	context *cli.Context) *models.CreateCommandRequest {

	boxSize := context.Int(boxSizeFlag.Name)
	layoutWidth := context.Int(layoutWidthFlag.Name)
	layoutHeight := context.Int(layoutHeightFlag.Name)
	overwrite := context.Bool(overwriteFileFlag.Name)

	request := &models.CreateCommandRequest{}

	if boxSize > 0 {
		request.BoxSize = helpers.IntToInt8Pointer(boxSize)
	}

	if layoutWidth > 0 {
		request.LayoutWidth = helpers.IntToInt8Pointer(layoutWidth)
	}

	if layoutHeight > 0 {
		request.LayoutHeight = helpers.IntToInt8Pointer(layoutHeight)
	}

	if overwrite {
		request.Overwrite = true
	}

	return request
}
