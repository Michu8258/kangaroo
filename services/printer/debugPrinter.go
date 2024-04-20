package printer

import (
	"io"

	"github.com/Michu8258/kangaroo/models"
)

type DebugPrinter struct {
	settings *models.Settings
	writer   io.Writer
}

func NewDebugPrinter(settings *models.Settings, writer io.Writer) IPrinter {
	return DebugPrinter{
		settings: settings,
		writer:   writer,
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
	dp.Print("\n")
}

func (dp DebugPrinter) Print(text string) {
	if dp.settings.UseDebugPrints {
		dp.writer.Write([]byte(models.TerminalStyles.DebugStyle.Render(text)))
	}
}
