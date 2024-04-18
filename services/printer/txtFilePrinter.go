package printer

import "os"

type TxtFilePrinter struct {
	file *os.File
}

func NewTxtFilePrinter(file *os.File) IPrinter {
	return TxtFilePrinter{
		file: file,
	}
}

func (fp TxtFilePrinter) PrintDefault(text string) {
	fp.file.WriteString(text)
}

func (fp TxtFilePrinter) PrintPrimary(text string) {
	fp.file.WriteString(text)
}

func (fp TxtFilePrinter) PrintSuccess(text string) {
	fp.file.WriteString(text)
}

func (fp TxtFilePrinter) PrintError(text string) {
	fp.file.WriteString(text)
}

func (fp TxtFilePrinter) PrintBorder(text string) {
	fp.file.WriteString(text)
}

func (fp TxtFilePrinter) PrintNewLine() {
	fp.file.WriteString("\n")
}
