package printer

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

type TerminalPrinter struct {
	isSilent bool
}

func NewTerminalPrinter(isSilent bool) TerminalPrinter {
	return TerminalPrinter{
		isSilent: isSilent,
	}
}

func (cp TerminalPrinter) PrintDefault(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", models.TerminalStyles.DefaultStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintPrimary(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", models.TerminalStyles.PrimaryStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintSuccess(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", models.TerminalStyles.SuccessStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintError(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", models.TerminalStyles.ErrorStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintBorder(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", models.TerminalStyles.BorderStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintNewLine() {
	if !cp.isSilent {
		fmt.Printf("%s", models.TerminalStyles.DefaultStyle.Render("\n"))
	}
}
