package dataWriter

import (
	"os"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/services/printer"
)

type DataWriter struct {
	Settings           *models.Settings
	DataPrinter        dataPrinters.IDataPrinter
	TxtPrinterProvider func(file *os.File) printer.IPrinter
}

type IDataWriter interface {
	SaveSudokuToJson(sudoku *models.Sudoku, path string, overwrite bool) (bool, error)
	SaveSudokuDtoToJson(sudokuDto *models.SudokuDTO, path string, overwrite bool) (bool, error)
	SaveSudokuToTxt(sudoku *models.Sudoku, path string, overwrite bool) (bool, error)
}

func GetNewDataWriter(settings *models.Settings,
	dataPrinter dataPrinters.IDataPrinter,
	txtPrinterProvider func(file *os.File) printer.IPrinter) IDataWriter {
	return &DataWriter{
		Settings:           settings,
		DataPrinter:        dataPrinter,
		TxtPrinterProvider: txtPrinterProvider,
	}
}
