package commands

import (
	"fmt"
	"path/filepath"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/dataInputs"
	"github.com/Michu8258/kangaroo/services/dataOutputs"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/services/printer"
	"github.com/urfave/cli/v2"
)

// CreateCommand provides create command configuration
func CreateCommand(settings *models.Settings) *cli.Command {
	return &cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage: "Creates sudoku puzzle data and saves to provided " +
			"file paths (JSON and TXT files supported, default is JSON)",
		Flags: []cli.Flag{
			&boxSizeFlag,
			&layoutWidthFlag,
			&layoutHeightFlag,
			&cli.BoolFlag{
				Name:        "overwrite",
				Aliases:     []string{"o"},
				DefaultText: "false",
				Usage:       "Overwrite provided files if exist",
			},
		},
		Action: func(context *cli.Context) error {
			request := buildCreateCommandRequest(context)
			filePaths := context.Args().Slice()
			return createCommandHandler(request, settings, filePaths)
		},
	}
}

// createCommandHandler is an entry point function to create sudoku data file
func createCommandHandler(request *models.CreateCommandRequest, settings *models.Settings,
	destinationFilePaths []string) error {
	consolePrinter := printer.NewTerminalPrinter(settings.SilentConsolePrints)

	if len(destinationFilePaths) < 1 {
		consolePrinter.PrintError("Please provide at least one argument for output file location.")
		consolePrinter.PrintNewLine()
		return nil
	}

	validPaths, errorPaths := validateDestinationFilePaths(destinationFilePaths)
	if len(errorPaths) > 0 {
		dataPrinters.PrintErrors("Optput files listed below are not supported", consolePrinter, errorPaths...)
	}

	if len(validPaths) < 1 {
		consolePrinter.PrintError("No supported file path to save sudoku data to.")
		consolePrinter.PrintNewLine()
		return nil
	}

	sudokuDto, err := dataInputs.ReadFromConsole(request.GetConfigRequest(), settings)
	if err != nil {
		dataPrinters.PrintErrors("Invalid sudoku input", consolePrinter, err)
		return nil
	}

	sudoku, ok := executeSudokuInitialization(sudokuDto, settings, consolePrinter)
	if !ok {
		return nil
	}

	consolePrinter.PrintPrimary("Saving results:")
	consolePrinter.PrintNewLine()
	for _, path := range validPaths {
		saveToFile(sudoku, request, path, consolePrinter, settings)
	}

	return nil
}

// executes save to file logic
func saveToFile(sudoku *models.Sudoku, request *models.CreateCommandRequest,
	path string, printer printer.Printer, settings *models.Settings) {

	extension := filepath.Ext(path)

	var written bool
	var err error

	if extension == ".txt" {
		written, err = dataOutputs.SaveSudokuToTxt(sudoku, settings, path, request.Overwrite)
	} else {
		written, err = dataOutputs.SaveSudokuToJson(sudoku, path, request.Overwrite)
	}

	if err != nil {
		printer.PrintError(fmt.Sprintf("- %s", err))
		printer.PrintNewLine()
		return
	}

	if written {
		printer.PrintSuccess(fmt.Sprintf("- '%s' written successfully", path))
		printer.PrintNewLine()
		return
	}

	printer.PrintDefault(fmt.Sprintf("- '%s' already exists (ommited)", path))
	printer.PrintNewLine()
}

// validateDestinationFilePaths checks if all provided file names have no extension
// or have json or txt extension. Returns slice ov valid names and slice of errors
// for those not matching the criteria.
func validateDestinationFilePaths(destinationFilePaths []string) ([]string, []error) {
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
func buildCreateCommandRequest(context *cli.Context) *models.CreateCommandRequest {
	boxSize := context.Int(boxSizeFlag.Name)
	layoutWidth := context.Int(layoutWidthFlag.Name)
	layoutHeight := context.Int(layoutHeightFlag.Name)
	overwrite := context.Bool("overwrite")

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
