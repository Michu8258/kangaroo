package printers

import (
	"github.com/Michu8258/kangaroo/types"
)

// PrintErrors prints errors list
func PrintErrors(errorsHeader string, printer types.Printer, errors ...error) {
	printer.PrintError(errorsHeader)
	printer.PrintNewLine()

	for _, err := range errors {
		printer.PrintError("- ")
		printer.PrintError(err.Error())
		printer.PrintError(".")
		printer.PrintNewLine()
	}
}
