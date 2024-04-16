package printers

import (
	"strings"

	"github.com/Michu8258/kangaroo/types"
)

// TODO - replace raw fmt prints with printer usage - add special log method
// TODO - inject printer into commands (solve, create)

// PrintErrors prints errors list
func PrintErrors(errorsHeader string, printer types.Printer, errors ...error) {
	printer.PrintError(errorsHeader)
	printer.PrintNewLine()

	for _, err := range errors {
		errorString := err.Error()
		if len(errorString) >= 1 {
			firstLetter := errorString[:1]
			restOfError := ""
			if len(errorString) >= 2 {
				restOfError = errorString[1:]
			}

			printer.PrintError("✗ ")
			printer.PrintError(strings.ToUpper(firstLetter))
			printer.PrintError(restOfError)
			printer.PrintError(".")
			printer.PrintNewLine()
		}
	}
}
