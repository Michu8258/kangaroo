package dataPrinters

import (
	"strings"

	"github.com/Michu8258/kangaroo/services/printer"
)

// PrintErrors prints errors list
func PrintErrors(errorsHeader string, printer printer.Printer, errors ...error) {
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
