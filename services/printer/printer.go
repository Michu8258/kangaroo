package printer

type Printer interface {
	PrintDefault(text string)
	PrintPrimary(text string)
	PrintSuccess(text string)
	PrintError(text string)
	PrintBorder(text string)
	PrintNewLine()
}
