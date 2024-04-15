package types

import "github.com/fatih/color"

type Printer interface {
	PrintDefault(text string)
	PrintBold(text string)
	PrintError(text string)
	PrintBoldError(text string)
	PrintNewLine()
}

type ConsolePrinter struct {
	defaultPrinter   *color.Color
	boldPrinter      *color.Color
	errorPrinter     *color.Color
	boldErrorPrinter *color.Color
}

func NewConsolePrinter() ConsolePrinter {
	return ConsolePrinter{
		defaultPrinter:   color.New(),
		boldPrinter:      color.New(color.Bold, color.FgHiBlue),
		errorPrinter:     color.New(color.FgRed),
		boldErrorPrinter: color.New(color.Bold, color.FgRed),
	}
}

func (cp ConsolePrinter) PrintDefault(text string) {
	cp.defaultPrinter.Printf(text)
}

func (cp ConsolePrinter) PrintBold(text string) {
	cp.boldPrinter.Printf(text)
}

func (cp ConsolePrinter) PrintError(text string) {
	cp.errorPrinter.Printf(text)
}

func (cp ConsolePrinter) PrintBoldError(text string) {
	cp.boldErrorPrinter.Printf(text)
}

func (cp ConsolePrinter) PrintNewLine() {
	cp.defaultPrinter.Printf("\n")
}
