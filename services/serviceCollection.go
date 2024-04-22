package services

import (
	"os"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/binarySudokuManager"
	crook "github.com/Michu8258/kangaroo/services/crookMethodSolver"
	"github.com/Michu8258/kangaroo/services/dataPrinters"
	"github.com/Michu8258/kangaroo/services/dataReader"
	"github.com/Michu8258/kangaroo/services/dataWriter"
	"github.com/Michu8258/kangaroo/services/printer"
	"github.com/Michu8258/kangaroo/services/prompts"
	"github.com/Michu8258/kangaroo/services/sudokuInit"
	tea "github.com/charmbracelet/bubbletea"
)

type ServiceCollection struct {
	TerminalPrinter printer.IPrinter
	DebugPrinter    printer.IPrinter
	DataReader      dataReader.IDataReader
	DataWriter      dataWriter.IDataWriter
	DataPrinter     dataPrinters.IDataPrinter
	SudokuInit      sudokuInit.ISudokuInit
	Prompter        prompts.IPrompter
	Solver          crook.ISudokuSolver
	SudokuEncoder   binarySudokuManager.IBinarySudokuManager
}

// Build creates a service collection to use in the application
func Build(settings *models.Settings) *ServiceCollection {
	terminalPrinter := printer.NewTerminalPrinter(settings, os.Stdout)
	debugPrinter := printer.NewDebugPrinter(settings, os.Stdout)
	dataPrinter := dataPrinters.GetNewDataPrinter(settings, terminalPrinter)
	prompter := prompts.GetNewPrompter(
		settings,
		terminalPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			program := tea.NewProgram(model, opts...)
			return program.Run()
		})

	return &ServiceCollection{
		TerminalPrinter: terminalPrinter,
		DebugPrinter:    debugPrinter,
		Prompter:        prompter,
		DataPrinter:     dataPrinter,
		SudokuInit:      sudokuInit.GetNewSudokuInit(settings),
		DataReader:      dataReader.GetNewDataReader(settings, terminalPrinter, debugPrinter, prompter),
		DataWriter: dataWriter.GetNewDataWriter(settings, dataPrinter,
			func(file *os.File) printer.IPrinter {
				return printer.NewTxtFilePrinter(file)
			}),
		Solver:        crook.GetNewSudokuSolver(settings, debugPrinter),
		SudokuEncoder: binarySudokuManager.GetNewBinarySudokuManager(settings),
	}
}
