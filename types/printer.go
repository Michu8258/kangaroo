package types

import (
	"fmt"
)

type Printer interface {
	PrintDefault(text string)
	PrintPrimary(text string)
	PrintError(text string)
	PrintBorder(text string)
	PrintNewLine()
}

type ConsolePrinter struct {
	isSilent bool
}

func NewConsolePrinter(isSilent bool) ConsolePrinter {
	return ConsolePrinter{
		isSilent: isSilent,
	}
}

func (cp ConsolePrinter) PrintDefault(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.DefaultStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintPrimary(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.PrimaryStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintError(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.ErrorStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintBorder(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.BorderStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintNewLine() {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.DefaultStyle.Render("\n"))
	}
}
