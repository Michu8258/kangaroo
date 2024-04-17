package commands

import (
	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	crook "github.com/Michu8258/kangaroo/services/crookMethodSolver"
	"github.com/urfave/cli/v2"
)

// TODO - test algorithm against unsolvable sudoku,
// TODO - add tests with script

// SolveCommand provides solve sudoku command configuration
func (commandConfig *CommandConfig) SolveCommand() *cli.Command {
	return &cli.Command{
		Name:    "solve",
		Aliases: []string{"s"},
		Usage:   "Solves a provided sudoku puzzle",
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
func (commandConfig *CommandConfig) solveCommandHandler(request *models.SolveCommandRequest) error {
	rawSudoku, err := commandConfig.getSudokuInputRawData(request)
	if err != nil {
		commandConfig.DataPrinter.PrintErrors("Invalid sudoku input", err)
		return nil
	}

	sudoku, ok := commandConfig.executeSudokuInitialization(rawSudoku)
	if !ok {
		return nil
	}

	solver := crook.GetNewSudokuSolver(commandConfig.Settings, commandConfig.DebugPrinter)
	solved, errs := solver.Solve(sudoku)
	if !solved {
		commandConfig.TerminalPrinter.PrintError("Failed to solve the sudoku.")
		commandConfig.TerminalPrinter.PrintNewLine()
		return nil
	}

	if commandConfig.Settings.UseDebugPrints && len(errs) >= 1 {
		commandConfig.DataPrinter.PrintErrors("Sudoku solution failure reasons:", err)
		return nil
	}

	commandConfig.printSudoku("Sudoku puzzle solution:", sudoku)

	if request.OutputFile != nil {
		validPaths := commandConfig.validateDestinationFilePaths(*request.OutputFile)
		if len(validPaths) >= 1 {
			commandConfig.TerminalPrinter.PrintNewLine()
			commandConfig.executeSudokuFilesSave(sudoku, request.AsConfigRequest(), validPaths)
		}
	}

	return nil
}

// getSudokuInputRawData retrieves sudoku raw data by analyzing the
// request object and executing one of the data sources logic.
func (commandConfig *CommandConfig) getSudokuInputRawData(
	request *models.SolveCommandRequest) (*models.SudokuDTO, error) {

	if request.InputJsonFile != nil {
		return commandConfig.DataReader.ReadSudokuFromJsonFile(*request.InputJsonFile)
	}

	return commandConfig.DataReader.ReadSudokuFromConsole(request.AsConfigRequest())
}

// buildSolveCommandRequest retrieves options settings from the command
// and constructs request object.
func (commandConfig *CommandConfig) buildSolveCommandRequest(context *cli.Context) *models.SolveCommandRequest {
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
