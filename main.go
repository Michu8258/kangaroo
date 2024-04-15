package main

import (
	"log"
	"os"

	"github.com/Michu8258/kangaroo/commands"
	"github.com/Michu8258/kangaroo/models"
	"github.com/urfave/cli/v2"
)

func main() {
	settings := createSettings()

	app := &cli.App{
		Name:           "Kangaroo",
		Usage:          "sudoku puzzle solution",
		Version:        "0.0.1",
		DefaultCommand: "solve",
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
				Destination: &settings.SilentConsolePrints,
			},
		},
		Commands: []*cli.Command{
			commands.SolveCommand(settings),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createSettings() *models.Settings {
	return &models.Settings{
		MinimumLayoutSizeInclusive:       2,
		MaximumLayoutSizeInclusive:       6,
		DefaultLayoutSize:                3,
		MinimumBoxSizeInclusive:          2,
		MaximumBoxSizeInclusive:          5,
		DefaultBoxSize:                   3,
		SudokuPrintoutValuePaddingLength: 1,
		UseDebugPrints:                   false,
		SilentConsolePrints:              false,
	}
}
