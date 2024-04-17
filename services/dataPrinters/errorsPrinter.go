package dataPrinters

import (
	"strings"
)

// PrintErrors prints errors list
func (dp *DataPrinter) PrintErrors(errorsHeader string, errors ...error) {
	dp.TerminalPrinter.PrintError(errorsHeader)
	dp.TerminalPrinter.PrintNewLine()

	for _, err := range errors {
		errorString := err.Error()
		if len(errorString) >= 1 {
			firstLetter := errorString[:1]
			restOfError := ""
			if len(errorString) >= 2 {
				restOfError = errorString[1:]
			}

			dp.TerminalPrinter.PrintError("✗ ")
			dp.TerminalPrinter.PrintError(strings.ToUpper(firstLetter))
			dp.TerminalPrinter.PrintError(restOfError)
			dp.TerminalPrinter.PrintError(".")
			dp.TerminalPrinter.PrintNewLine()
		}
	}
}
