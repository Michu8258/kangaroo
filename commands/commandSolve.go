package commands

import (
	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	crook "github.com/Michu8258/kangaroo/services/crookMethodSolver"
	"github.com/Michu8258/kangaroo/services/dataInputs"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/services/printer"
	"github.com/urfave/cli/v2"
)

// TODO - test algorithm against unsolvable sudoku,
// TODO - add saving result to file - one flag, not two - like in create command

// SolveCommand provides solve sudoku command configuration
func SolveCommand(settings *models.Settings) *cli.Command {
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
			return solveCommandHandler(request, settings)
		},
	}
}

// solveCommandHandler is an entry point function for solve sudoku command
func solveCommandHandler(request *models.SolveCommandRequest, settings *models.Settings) error {
	consolePrinter := printer.NewTerminalPrinter(settings.SilentConsolePrints)
	rawSudoku, err := getSudokuInputRawData(request, settings)
	if err != nil {
		dataPrinters.PrintErrors("Invalid sudoku input", consolePrinter, err)
		return nil
	}

	sudoku, ok := executeSudokuInitialization(rawSudoku, settings, consolePrinter)
	if !ok {
		return nil
	}

	solved, errs := crook.SolveWithCrookMethod(sudoku, settings)
	if !solved {
		consolePrinter.PrintError("Failed to solve the sudoku.")
		consolePrinter.PrintNewLine()
		return nil
	}

	if settings.UseDebugPrints && len(errs) >= 1 {
		dataPrinters.PrintErrors("Sudoku solution failure reasons:", consolePrinter, err)
		return nil
	}

	printSudoku("Sudoku puzzle solution:", sudoku, settings, consolePrinter)

	return nil
}

// getSudokuInputRawData retrieves sudoku raw data by analyzing the
// request object and executing one of the data sources logic.
func getSudokuInputRawData(request *models.SolveCommandRequest, settings *models.Settings) (*models.SudokuDTO, error) {
	if request.InputJsonFile != nil {
		return dataInputs.ReadFromJsonFile(*request.InputJsonFile)
	}

	return dataInputs.ReadFromConsole(request.GetConfigRequest(), settings)
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