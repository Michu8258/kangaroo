package types

import (
	"fmt"
	"os"
)

type Printer interface {
	PrintDefault(text string)
	PrintPrimary(text string)
	PrintSuccess(text string)
	PrintError(text string)
	PrintBorder(text string)
	PrintNewLine()
}

type ConsolePrinter struct {
	isSilent bool
}

func NewConsolePrinter(isSilent bool) ConsolePrinter {
	return ConsolePrinter{
		isSilent: isSilent,
	}
}

func (cp ConsolePrinter) PrintDefault(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.DefaultStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintPrimary(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.PrimaryStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintSuccess(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.SuccessStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintError(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.ErrorStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintBorder(text string) {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.BorderStyle.Render(text))
	}
}

func (cp ConsolePrinter) PrintNewLine() {
	if !cp.isSilent {
		fmt.Printf("%s", OutputStyles.DefaultStyle.Render("\n"))
	}
}

type TxtFilePrinter struct {
	file *os.File
}

func NewTxtFilePrinter(file *os.File) TxtFilePrinter {
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
