package commands

import (
	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	"github.com/urfave/cli/v2"
)

// SolveCommand provides solve sudoku command configuration
func (commandConfig *CommandContext) SolveCommand() *cli.Command {
	return &cli.Command{
		Name:    "solve",
		Aliases: []string{"s"},
		Usage: "Solves a provided sudoku puzzle. There are few formats supported\n" +
			"for this command. You can pass an input data json file path using -i flag.\n" +
			"This option will have precedence over all other because the file contains all\n" +
			"the data required to build sudoku object. In case no -i flag is passed, then\n" +
			"cli works in manual mode - it will ask for box size, and sudoku layout and all\n" +
			"sudoku values. Box size, and sudoku layout prompts may be ommited by using -s,\n" +
			"--lw and --lh flags. You can save result of sulution to a file with a -o flag.",
		Flags: []cli.Flag{
			&boxSizeFlag,
			&layoutWidthFlag,
			&layoutHeightFlag,
			&overwriteFileFlag,
			&cli.StringFlag{Name: "input-file",
				Aliases:     []string{"i"},
				DefaultText: "",
				Usage:       "Specify path to sudoku JSON configuration file",
			},
			&cli.StringFlag{
				Name:        "output-file",
				Aliases:     []string{"o"},
				DefaultText: "",
				Usage:       "Specify path to file where you want to save solution of the sudoku (JSON or TXT, JSON is default)",
			},
		},
		Action: func(context *cli.Context) error {
			request := commandConfig.buildSolveCommandRequest(context)
			return commandConfig.solveCommandHandler(request)
		},
	}
}

// solveCommandHandler is an entry point function for solve sudoku command
func (commandConfig *CommandContext) solveCommandHandler(request *models.SolveCommandRequest) error {
	rawSudoku, err := commandConfig.getSudokuInputRawData(request)
	if err != nil {
		commandConfig.ServiceCollection.DataPrinter.
			PrintErrors("Invalid sudoku input", err)
		return nil
	}

	sudoku, ok := commandConfig.executeSudokuInitialization(rawSudoku, true)
	if !ok {
		return nil
	}

	solved, errs := commandConfig.ServiceCollection.Solver.Solve(sudoku)
	if !solved {
		commandConfig.ServiceCollection.TerminalPrinter.
			PrintError("Failed to solve the sudoku.")
		commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
		return nil
	}

	if commandConfig.Settings.UseDebugPrints && len(errs) >= 1 {
		commandConfig.ServiceCollection.DataPrinter.PrintErrors(
			"Sudoku solution failure reasons:", err)
		return nil
	}

	commandConfig.printSudoku("Sudoku puzzle solution:", sudoku)

	if request.OutputFile != nil {
		validPaths := commandConfig.validateDestinationFilePaths(*request.OutputFile)
		if len(validPaths) >= 1 {
			commandConfig.ServiceCollection.TerminalPrinter.PrintNewLine()
			commandConfig.executeSudokuFilesSave(sudoku, request.AsConfigRequest(), validPaths)
		}
	}

	return nil
}

// getSudokuInputRawData retrieves sudoku raw data by analyzing the
// request object and executing one of the data sources logic.
func (commandConfig *CommandContext) getSudokuInputRawData(
	request *models.SolveCommandRequest) (*models.SudokuDTO, error) {

	if request.InputJsonFile != nil {
		return commandConfig.ServiceCollection.DataReader.
			ReadSudokuFromJsonFile(*request.InputJsonFile)
	}

	return commandConfig.ServiceCollection.DataReader.
		ReadSudokuFromConsole(request.AsConfigRequest())
}

// buildSolveCommandRequest retrieves options settings from the command
// and constructs request object.
func (commandConfig *CommandContext) buildSolveCommandRequest(context *cli.Context) *models.SolveCommandRequest {
	boxSize := context.Int(boxSizeFlag.Name)
	layoutWidth := context.Int(layoutWidthFlag.Name)
	layoutHeight := context.Int(layoutHeightFlag.Name)
	inputJsonFile := context.String("input-file")
	outputFile := context.String("output-file")
	overwrite := context.Bool(overwriteFileFlag.Name)

	request := &models.SolveCommandRequest{}

	if boxSize > 0 {
		request.BoxSize = helpers.IntToInt8Pointer(boxSize)
	}

	if layoutWidth > 0 {
		request.LayoutWidth = helpers.IntToInt8Pointer(layoutWidth)
	}

	if layoutHeight > 0 {
		request.LayoutHeight = helpers.IntToInt8Pointer(layoutHeight)
	}

	if len(inputJsonFile) > 0 {
		request.InputJsonFile = &inputJsonFile
	}

	if len(outputFile) > 0 {
		request.OutputFile = &outputFile
	}

	if overwrite {
		request.Overwrite = true
	}

	return request
}
