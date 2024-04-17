package main

import (
	"log"
	"os"

	"github.com/Michu8258/kangaroo/commands"
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
	"github.com/urfave/cli/v2"
)

// TODO presets commands family
// TODO update documentation markdown file
// TODO dependency inversion - services, inject printer type

func main() {
	settings := createSettings()

	commandConfig := &commands.CommandConfig{
		Settings:        settings,
		TerminalPrinter: printer.NewTerminalPrinter(settings),
		DebugPrinter:    printer.NewDebugPrinter(settings),
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
				Usage:       "supresses any standard output printing",
				Destination: &commandConfig.Settings.SilentConsolePrints,
			},
		},
		Commands: []*cli.Command{
			commands.SolveCommand(commandConfig),
			commands.CreateCommand(commandConfig),
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
