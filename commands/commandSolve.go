package commands

import (
	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	crook "github.com/Michu8258/kangaroo/services/crookMethodSolver"
	"github.com/Michu8258/kangaroo/services/dataInputs"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/urfave/cli/v2"
)

// TODO - test algorithm against unsolvable sudoku,
// TODO - add saving result to file - one flag, not two - like in create command
// TODO - unify flags and commands descriptions

// SolveCommand provides solve sudoku command configuration
func SolveCommand(commandConfig *CommandConfig) *cli.Command {
	return &cli.Command{
		Name:    "solve",
		Aliases: []string{"s"},
		Usage:   "Solves a provided sudoku puzzle",
		Flags: []cli.Flag{
			&boxSizeFlag,
			&layoutWidthFlag,
			&layoutHeightFlag,
			&cli.StringFlag{Name: "input-file-json",
				Aliases:     []string{"i"},
				DefaultText: "",
				Usage:       "Specify relative path to sudoku JSON configuration file",
			},
			&cli.StringFlag{
				Name:        "output-file-json",
				Aliases:     []string{"oj"},
				DefaultText: "",
				Usage:       "Specify relative path to JSON file where you want to save solution of the sudoku",
			},
			&cli.StringFlag{
				Name:        "output-file-txt",
				Aliases:     []string{"ot"},
				DefaultText: "",
				Usage:       "Specify relative path to TXT file where you want to save solution of the sudoku",
			},
		},
		Action: func(context *cli.Context) error {
			request := buildSolveCommandRequest(context)
			return solveCommandHandler(request, commandConfig)
		},
	}
}

// solveCommandHandler is an entry point function for solve sudoku command
func solveCommandHandler(request *models.SolveCommandRequest, commandConfig *CommandConfig) error {
	rawSudoku, err := getSudokuInputRawData(request, commandConfig)
	if err != nil {
		dataPrinters.PrintErrors("Invalid sudoku input", commandConfig.TerminalPrinter, err)
		return nil
	}

	sudoku, ok := executeSudokuInitialization(rawSudoku,
		commandConfig.Settings, commandConfig.TerminalPrinter)
	if !ok {
		return nil
	}

	solved, errs := crook.SolveWithCrookMethod(sudoku, commandConfig.Settings,
		commandConfig.DebugPrinter)
	if !solved {
		commandConfig.TerminalPrinter.PrintError("Failed to solve the sudoku.\n")
		return nil
	}

	if commandConfig.Settings.UseDebugPrints && len(errs) >= 1 {
		dataPrinters.PrintErrors("Sudoku solution failure reasons:",
			commandConfig.DebugPrinter, err)
		return nil
	}

	printSudoku("Sudoku puzzle solution:", sudoku,
		commandConfig.Settings, commandConfig.TerminalPrinter)

	return nil
}

// getSudokuInputRawData retrieves sudoku raw data by analyzing the
// request object and executing one of the data sources logic.
func getSudokuInputRawData(request *models.SolveCommandRequest, commandConfig *CommandConfig) (*models.SudokuDTO, error) {
	if request.InputJsonFile != nil {
		return dataInputs.ReadFromJsonFile(*request.InputJsonFile)
	}

	return dataInputs.ReadFromConsole(request.AsConfigRequest(),
		commandConfig.Settings, commandConfig.TerminalPrinter, commandConfig.DebugPrinter)
}

// buildSolveCommandRequest retrieves options settings from the command
// and constructs request object.
func buildSolveCommandRequest(context *cli.Context) *models.SolveCommandRequest {
	inputJsonFile := context.String("input-file-json")
	outputJsonFile := context.String("output-file-json")
	outputTxtFile := context.String("output-file-txt")
	boxSize := context.Int(boxSizeFlag.Name)
	layoutWidth := context.Int(layoutWidthFlag.Name)
	layoutHeight := context.Int(layoutHeightFlag.Name)

	request := &models.SolveCommandRequest{}

	if len(inputJsonFile) > 0 {
		request.InputJsonFile = &inputJsonFile
	}

	if len(outputJsonFile) > 0 {
		request.OutputJsonFile = &outputJsonFile
	}

	if len(outputTxtFile) > 0 {
		request.OutputTxtFile = &outputTxtFile
	}

	if boxSize > 0 {
		request.BoxSize = helpers.IntToInt8Pointer(boxSize)
	}

	if layoutWidth > 0 {
		request.LayoutWidth = helpers.IntToInt8Pointer(layoutWidth)
	}

	if layoutHeight > 0 {
		request.LayoutHeight = helpers.IntToInt8Pointer(layoutHeight)
	}

	return request
}
