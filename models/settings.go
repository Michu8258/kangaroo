package models

type Settings struct {
	MinimumLayoutSizeInclusive       int8
	MaximumLayoutSizeInclusive       int8
	DefaultLayoutSize                int8
	MinimumBoxSizeInclusive          int8
	MaximumBoxSizeInclusive          int8
	DefaultBoxSize                   int8
	SudokuPrintoutValuePaddingLength int8
	UseDebugPrints                   bool
	SilentConsolePrints              bool
	SudokuBinaryEncoderVersion       uint16
}
