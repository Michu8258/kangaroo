package main

import (
	"log"
	"os"

	"github.com/Michu8258/kangaroo/commands"
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/services/dataReader"
	"github.com/Michu8258/kangaroo/services/dataWriter"
	"github.com/Michu8258/kangaroo/services/printer"
	"github.com/Michu8258/kangaroo/services/sudokuInit"
	"github.com/urfave/cli/v2"
)

// TODO presets commands family
// TODO update documentation markdown file

func main() {
	settings := createSettings()
	terminalPrinter := printer.NewTerminalPrinter(settings)
	debugPrinter := printer.NewDebugPrinter(settings)
	dataPrinter := dataPrinters.GetNewDataPrinter(settings, terminalPrinter)

	commandConfig := &commands.CommandConfig{
		Settings:        settings,
		TerminalPrinter: terminalPrinter,
		DebugPrinter:    debugPrinter,
		DataReader:      dataReader.GetNewDataReader(settings, terminalPrinter, debugPrinter),
		DataPrinter:     dataPrinter,
		DataWriter: dataWriter.GetNewDataWriter(settings, dataPrinter,
			func(file *os.File) printer.IPrinter {
				return printer.NewTxtFilePrinter(file)
			}),
		SudokuInit: sudokuInit.GetNewSudokuInit(settings),
	}

	app := &cli.App{
		Name:           "Kangaroo",
		Usage:          "sudoku puzzle solution",
		Version:        "0.0.1",
		DefaultCommand: "help",
		Authors: []*cli.Author{
			{
				Name:  "The author",
				Email: "the.author@example.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "silent",
				Value:       false,
				Usage:       "Supresses any standard output printing (prompts for inputs will still be printed)",
				Destination: &commandConfig.Settings.SilentConsolePrints,
			},
		},
		Commands: []*cli.Command{
			commandConfig.SolveCommand(),
			commandConfig.CreateCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createSettings() *models.Settings {
	return &models.Settings{
		MinimumLayoutSizeInclusive:       2,
		MaximumLayoutSizeInclusive:       5,
		DefaultLayoutSize:                3,
		MinimumBoxSizeInclusive:          2,
		MaximumBoxSizeInclusive:          5,
		DefaultBoxSize:                   3,
		SudokuPrintoutValuePaddingLength: 1,
		UseDebugPrints:                   false,
		SilentConsolePrints:              false,
	}
}
