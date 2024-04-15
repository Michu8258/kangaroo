package models

type Settings struct {
	MinimumLayoutSizeInclusive       int8
	MaximumLayoutSizeInclusive       int8
	MinimumBoxSizeInclusive          int8
	MaximumBoxSizeInclusive          int8
	UseDebugPrints                   bool
	SudokuPrintoutValuePaddingLength int8
	SilentConsolePrints              bool
}
