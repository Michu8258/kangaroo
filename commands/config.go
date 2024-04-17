package commands

import (
	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
)

type CommandConfig struct {
	Settings        *models.Settings
	TerminalPrinter printer.Printer
	DebugPrinter    printer.Printer
}
