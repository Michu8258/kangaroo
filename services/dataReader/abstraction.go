package dataReader

import (
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
	"github.com/Michu8258/kangaroo/services/prompts"
)

type DataReader struct {
	Settings        *models.Settings
	TerminalPrinter printer.IPrinter
	DebugPrinter    printer.IPrinter
	Prompter        prompts.IPrompter
}

type IDataReader interface {
	ReadSudokuFromConsole(request *models.SudokuConfigRequest) (*models.SudokuDTO, error)
	ReadSudokuFromJsonFile(path string) (*models.SudokuDTO, error)
}

func GetNewDataReader(settings *models.Settings,
	terminalPrinter printer.IPrinter,
	debugPrinter printer.IPrinter,
	prompter prompts.IPrompter) IDataReader {
	return &DataReader{
		Settings:        settings,
		TerminalPrinter: terminalPrinter,
		DebugPrinter:    debugPrinter,
		Prompter:        prompter,
	}
}
