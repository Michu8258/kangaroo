package printer

import (
	"io"

	"github.com/Michu8258/kangaroo/models"
)

type TerminalPrinter struct {
	settings *models.Settings
	writer   io.Writer
}

func NewTerminalPrinter(settings *models.Settings, writer io.Writer) IPrinter {
	return TerminalPrinter{
		settings: settings,
		writer:   writer,
	}
}

func (cp TerminalPrinter) PrintDefault(text string) {
	if !cp.settings.SilentConsolePrints {
		cp.writer.Write([]byte(models.TerminalStyles.DefaultStyle.Render(text)))
	}
}

func (cp TerminalPrinter) PrintPrimary(text string) {
	if !cp.settings.SilentConsolePrints {
		cp.writer.Write([]byte(models.TerminalStyles.PrimaryStyle.Render(text)))
	}
}

func (cp TerminalPrinter) PrintSuccess(text string) {
	if !cp.settings.SilentConsolePrints {
		cp.writer.Write([]byte(models.TerminalStyles.SuccessStyle.Render(text)))
	}
}

func (cp TerminalPrinter) PrintError(text string) {
	if !cp.settings.SilentConsolePrints {
		cp.writer.Write([]byte(models.TerminalStyles.ErrorStyle.Render(text)))
	}
}

func (cp TerminalPrinter) PrintBorder(text string) {
	if !cp.settings.SilentConsolePrints {
		cp.writer.Write([]byte(models.TerminalStyles.BorderStyle.Render(text)))
	}
}

func (cp TerminalPrinter) PrintNewLine() {
	if !cp.settings.SilentConsolePrints {
		cp.writer.Write([]byte(models.TerminalStyles.DefaultStyle.Render("\n")))
	}
}
