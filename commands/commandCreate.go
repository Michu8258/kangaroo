package commands

import (
	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	"github.com/urfave/cli/v2"
)

// CreateCommand provides create command configuration
func (commandConfig *CommandContext) CreateCommand() *cli.Command {
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
func (commandConfig *CommandContext) createCommandHandler(request *models.CreateCommandRequest,
	destinationFilePaths []string) error {

	if len(destinationFilePaths) < 1 {
		commandConfig.ServiceCollection.TerminalPrinter.PrintError(
			"Please provide at least one argument for output file location.")
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		return nil
	}

	validPaths := commandConfig.validateDestinationFilePaths(destinationFilePaths...)
	if len(validPaths) < 1 {
		return nil
	}

	sudokuDto, err := commandConfig.ServiceCollection.DataReader.
		ReadSudokuFromConsole(request.AsConfigRequest())
	if err != nil {
		commandConfig.ServiceCollection.DataPrinter.
			PrintErrors("Invalid sudoku input", err)
		return nil
	}

	sudoku, ok := commandConfig.executeSudokuInitialization(sudokuDto)
	if !ok {
		return nil
	}

	commandConfig.executeSudokuFilesSave(sudoku, request.AsConfigRequest(), validPaths)

	return nil
}

// buildCreateCommandRequest retrieves options settings from the command
// and constructs request object.
func (commandConfig *CommandContext) buildCreateCommandRequest(
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
