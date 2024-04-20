package printer

import "io"

type TxtFilePrinter struct {
	writer io.Writer
}

func NewTxtFilePrinter(writer io.Writer) IPrinter {
	return TxtFilePrinter{
		writer: writer,
	}
}

func (fp TxtFilePrinter) PrintDefault(text string) {
	fp.Print(text)
}

func (fp TxtFilePrinter) PrintPrimary(text string) {
	fp.Print(text)
}

func (fp TxtFilePrinter) PrintSuccess(text string) {
	fp.Print(text)
}

func (fp TxtFilePrinter) PrintError(text string) {
	fp.Print(text)
}

func (fp TxtFilePrinter) PrintBorder(text string) {
	fp.Print(text)
}

func (fp TxtFilePrinter) PrintNewLine() {
	fp.Print("\n")
}

func (fp TxtFilePrinter) Print(text string) {
	fp.writer.Write([]byte(text))
}
