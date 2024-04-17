package commands

import (
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/services/dataReader"
	"github.com/Michu8258/kangaroo/services/dataWriter"
	"github.com/Michu8258/kangaroo/services/printer"
	"github.com/Michu8258/kangaroo/services/sudokuInit"
)

type CommandConfig struct {
	Settings        *models.Settings
	TerminalPrinter printer.IPrinter
	DebugPrinter    printer.IPrinter
	DataReader      dataReader.IDataReader
	DataWriter      dataWriter.IDataWriter
	DataPrinter     dataPrinters.IDataPrinter
	SudokuInit      sudokuInit.ISudokuInit
}
