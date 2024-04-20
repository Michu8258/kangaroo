package printer

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestDebugPrinter(t *testing.T) {
	testPrinter(t, func(settings *models.Settings, writer io.Writer) IPrinter {
		return NewDebugPrinter(settings, writer)
	})
}

func TestTerminalPrinter(t *testing.T) {
	testPrinter(t, func(settings *models.Settings, writer io.Writer) IPrinter {
		return NewTerminalPrinter(settings, writer)
	})
}

func TestTxtFilePrinter(t *testing.T) {
	testPrinter(t, func(settings *models.Settings, writer io.Writer) IPrinter {
		return NewTxtFilePrinter(writer)
	})
}

func testPrinter(t *testing.T,
	printerProvider func(settings *models.Settings, writer io.Writer) IPrinter) {

	settings := testHelpers.GetTestSettings()
	settings.UseDebugPrints = true

	strs := []string{"default", "primary", "success", "error", "border", "\n"}
	buffer := &bytes.Buffer{}

	printer := printerProvider(settings, buffer)

	printer.PrintDefault(strs[0])
	printer.PrintPrimary(strs[1])
	printer.PrintSuccess(strs[2])
	printer.PrintError(strs[3])
	printer.PrintBorder(strs[4])
	printer.PrintNewLine()

	printed := buffer.String()

	for _, str := range strs {
		if !strings.Contains(printed, str) {
			t.Errorf("Printed data '%s' does not contain required string: '%s'.",
				printed, str)
		}
	}
}
