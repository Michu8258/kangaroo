package dataPrinters

import (
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
)

type DataPrinter struct {
	Settings        *models.Settings
	TerminalPrinter printer.IPrinter
}

type IDataPrinter interface {
	PrintErrors(errorsHeader string, errors ...error)
	PrintSudoku(sudoku *models.Sudoku, printer printer.IPrinter)
}

func GetNewDataPrinter(settings *models.Settings, terminalPrinter printer.IPrinter) IDataPrinter {
	return &DataPrinter{
		Settings:        settings,
		TerminalPrinter: terminalPrinter,
	}
}
