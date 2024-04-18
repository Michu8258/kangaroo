package testHelpers

type TestPrinter struct {
	PrintedData string
}

func NewTestPrinter() *TestPrinter {
	return &TestPrinter{
		PrintedData: "",
	}
}

func (printer *TestPrinter) PrintDefault(text string) {
	printer.PrintedData += text
}

func (printer *TestPrinter) PrintPrimary(text string) {
	printer.PrintedData += text
}

func (printer *TestPrinter) PrintSuccess(text string) {
	printer.PrintedData += text
}

func (printer *TestPrinter) PrintError(text string) {
	printer.PrintedData += text
}

func (printer *TestPrinter) PrintBorder(text string) {
	printer.PrintedData += text
}

func (printer *TestPrinter) PrintNewLine() {
	printer.PrintedData += "\n"
}
