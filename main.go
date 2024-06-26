package main

import (
	"log"
	"os"

	"github.com/Michu8258/kangaroo/commands"
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services"
	"github.com/urfave/cli/v2"
)

func main() {
	run(os.Args)
}

func run(arguments []string) {
	settings := createSettings()

	commandConfig := &commands.CommandContext{
		Settings:          settings,
		ServiceCollection: services.Build(settings),
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
			commandConfig.CreateCommand(),
			commandConfig.SolveCommand(),
			commandConfig.ExecuteCommand(),
		},
	}

	if err := app.Run(arguments); err != nil {
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
		SudokuBinaryEncoderVersion:       1,
	}
}
