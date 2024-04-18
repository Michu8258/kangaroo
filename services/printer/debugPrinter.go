package printer

import (
	"fmt"

	"github.com/Michu8258/kangaroo/models"
)

type DebugPrinter struct {
	settings *models.Settings
}

func NewDebugPrinter(settings *models.Settings) IPrinter {
	return DebugPrinter{
		settings: settings,
	}
}

func (dp DebugPrinter) PrintDefault(text string) {
	dp.Print(text)
}

func (dp DebugPrinter) PrintPrimary(text string) {
	dp.Print(text)
}

func (dp DebugPrinter) PrintSuccess(text string) {
	dp.Print(text)
}

func (dp DebugPrinter) PrintError(text string) {
	dp.Print(text)
}

func (dp DebugPrinter) PrintBorder(text string) {
	dp.Print(text)
}

func (dp DebugPrinter) PrintNewLine() {
	if dp.settings.UseDebugPrints {
		fmt.Println()
	}
}

func (dp DebugPrinter) Print(text string) {
	if dp.settings.UseDebugPrints {
		fmt.Printf("%s", models.TerminalStyles.DebugStyle.Render(text))
	}
}
