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
	isSilent         bool
}

func NewConsolePrinter(isSilent bool) ConsolePrinter {
	return ConsolePrinter{
		defaultPrinter:   color.New(),
		boldPrinter:      color.New(color.Bold, color.FgHiBlue),
		errorPrinter:     color.New(color.FgRed),
		boldErrorPrinter: color.New(color.Bold, color.FgRed),
		isSilent:         isSilent,
	}
}

func (cp ConsolePrinter) PrintDefault(text string) {
	if !cp.isSilent {
		cp.defaultPrinter.Printf(text)
	}
}

func (cp ConsolePrinter) PrintBold(text string) {
	if !cp.isSilent {
		cp.boldPrinter.Printf(text)
	}
}

func (cp ConsolePrinter) PrintError(text string) {
	if !cp.isSilent {
		cp.errorPrinter.Printf(text)
	}
}

func (cp ConsolePrinter) PrintBoldError(text string) {
	if !cp.isSilent {
		cp.boldErrorPrinter.Printf(text)
	}
}

func (cp ConsolePrinter) PrintNewLine() {
	if !cp.isSilent {
		cp.defaultPrinter.Printf("\n")
	}
}
