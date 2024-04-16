package commands

import "github.com/urfave/cli/v2"

var boxSizeFlag cli.IntFlag = cli.IntFlag{
	Name:        "box-size",
	Aliases:     []string{"s"},
	DefaultText: "0",
	Usage:       "How many rows and columns single sudoku box has - in case of classic sudoku it is 3",
}

var layoutWidthFlag cli.IntFlag = cli.IntFlag{
	Name:        "layout-width",
	Aliases:     []string{"lw"},
	DefaultText: "0",
	Usage:       "How many boxes there are in the row - in case of classic sudoku it is 3",
}

var layoutHeightFlag cli.IntFlag = cli.IntFlag{
	Name:        "layout-height",
	Aliases:     []string{"lh"},
	DefaultText: "0",
	Usage:       "How many boxes there are in the column - in case of classic sudoku it is 3",
}
