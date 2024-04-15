package commands

import (
	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	crook "github.com/Michu8258/kangaroo/services/crookMethodSolver"
	"github.com/Michu8258/kangaroo/services/dataInputs"
	"github.com/Michu8258/kangaroo/services/initialization"
	"github.com/Michu8258/kangaroo/services/printers"
	"github.com/Michu8258/kangaroo/types"
	"github.com/urfave/cli/v2"
)

// SolveCommand provides solve sudoku command configuration
func SolveCommand(settings *models.Settings) *cli.Command {
	return &cli.Command{
		Name:    "solve",
		Aliases: []string{"s"},
		Usage:   "Solves a provided sudoku puzzle",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "box-size",
				Aliases:     []string{"s"},
				DefaultText: "0",
				Usage:       "How many rows and columns single sudoku box has - in case of classic sudoku it is 3",
			},
			&cli.IntFlag{
				Name:        "layout-width",
				Aliases:     []string{"lw"},
				DefaultText: "0",
				Usage:       "How many boxes there are in the row - in case of classic sudoku it is 3",
			},
			&cli.IntFlag{
				Name:        "layout-height",
				Aliases:     []string{"lh"},
				DefaultText: "0",
				Usage:       "How many boxes there are in the column - in case of classic sudoku it is 3",
			},
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
func solveCommandHandler(request models.SolveCommandRequest, settings *models.Settings) error {
	consolePrinter := types.NewConsolePrinter(settings.SilentConsolePrints)
	rawSudoku, err := getSudokuInputRawData(request, settings)
	if err != nil {
		printers.PrintErrors("Invalid sudoku input", consolePrinter, err)
		return nil
	}

	sudoku := rawSudoku.ToSudoku()
	errs := initialization.InitializeSudoku(sudoku, settings)
	if len(errs) >= 1 {
		printers.PrintErrors("Invalid sudoku configuration", consolePrinter, err)
		return nil
	}

	// TODO printout of important settings
	printSudoku("Provided sudoku input:", sudoku, settings, consolePrinter)

	// TODO - add spinner here
	// https://github.com/charmbracelet/bubbletea/blob/master/examples/spinners/main.go
	solved, errs := crook.SolveWithCrookMethod(sudoku, settings)
	if !solved {
		consolePrinter.PrintError("Failed to solve the sudoku.")
		consolePrinter.PrintNewLine()
		return nil
	}

	if settings.UseDebugPrints && len(errs) >= 1 {
		printers.PrintErrors("Sudoku solution failure reasons:", consolePrinter, err)
		return nil
	}

	printSudoku("Sudoku puzzle solution:", sudoku, settings, consolePrinter)

	return nil
}

// getSudokuInputRawData retrieves sudoku raw data by analyzing the
// request object and executing one of the data sources logic.
func getSudokuInputRawData(request models.SolveCommandRequest, settings *models.Settings) (*models.SudokuDTO, error) {
	if request.InputJsonFile != nil {
		return dataInputs.ReadFromJsonFile(*request.InputJsonFile)
	}

	return dataInputs.ReadFromConsole(request, settings)
}

// printSudoku prints sudoku to standard out
func printSudoku(description string, sudoku *models.Sudoku, settings *models.Settings, printer types.Printer) {
	printer.PrintDefault(description)
	printer.PrintNewLine()
	printers.PrintSudoku(sudoku, settings, printer)
}

// buildSolveCommandRequest retrieves options settings from the command
// and constructs request object.
func buildSolveCommandRequest(context *cli.Context) models.SolveCommandRequest {
	inputJsonFile := context.String("input-file-json")
	outputJsonFile := context.String("output-file-json")
	outputTxtFile := context.String("output-file-txt")
	boxSize := context.Int("box-size")
	layoutWidth := context.Int("layout-width")
	layoutHeight := context.Int("layout-height")

	request := models.SolveCommandRequest{}

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
