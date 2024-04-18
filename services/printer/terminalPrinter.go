package printer

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

type TerminalPrinter struct {
	settings *models.Settings
}

func NewTerminalPrinter(settings *models.Settings) IPrinter {
	return TerminalPrinter{
		settings: settings,
	}
}

func (cp TerminalPrinter) CanPrint() bool {
	return !cp.settings.SilentConsolePrints
}

func (cp TerminalPrinter) PrintDefault(text string) {
	if cp.CanPrint() {
		fmt.Printf("%s", models.TerminalStyles.DefaultStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintPrimary(text string) {
	if cp.CanPrint() {
		fmt.Printf("%s", models.TerminalStyles.PrimaryStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintSuccess(text string) {
	if cp.CanPrint() {
		fmt.Printf("%s", models.TerminalStyles.SuccessStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintError(text string) {
	if cp.CanPrint() {
		fmt.Printf("%s", models.TerminalStyles.ErrorStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintBorder(text string) {
	if cp.CanPrint() {
		fmt.Printf("%s", models.TerminalStyles.BorderStyle.Render(text))
	}
}

func (cp TerminalPrinter) PrintNewLine() {
	if cp.CanPrint() {
		fmt.Printf("%s", models.TerminalStyles.DefaultStyle.Render("\n"))
	}
}
