package printers

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/types"
)

type sudokuPrintoutConfig struct {
	ValueCharactersLength int
	MaxIndex              int
	CharactersPerLine     int
	BoxSize               int
	Padding               int
}

// PrintSudoku prints entire sudoku puzzle pseudo-graphical representation to the console
func PrintSudoku(sudoku *models.Sudoku, settings *models.Settings, printer types.Printer) {
	printoutConfig := buildSudokuPrintoutConfig(sudoku, settings)
	printTopBorderLine(sudoku, printoutConfig, printer)

	for boxRowIndex := 0; boxRowIndex < int(sudoku.Layout.Height); boxRowIndex++ {
		for cellRowIndex := 0; cellRowIndex < int(sudoku.BoxSize); cellRowIndex++ {
			printValuesLine(sudoku, printoutConfig, int8(boxRowIndex), int8(cellRowIndex), printer)
			if cellRowIndex < int(sudoku.BoxSize)-1 {
				printMidCellsLine(sudoku, printoutConfig, printer)
			}
		}

		if boxRowIndex < int(sudoku.Layout.Height)-1 {
			printMidBoxesLine(sudoku, printoutConfig, printer)
		}
	}

	printBottomBorderLine(sudoku, printoutConfig, printer)
}

// printTopBorderLine prints top border line of a sudoku puzzle
func printTopBorderLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig, printer types.Printer) {
	printHorizontalBorderLine(sudoku, printoutConfig, "╔", "═", "╗", "╦", printer)
}

// printMidCellsLine prints line of a sudoku puzzle that appears between cells
func printMidCellsLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig, printer types.Printer) {
	printHorizontalBorderLine(sudoku, printoutConfig, "║", "─", "║", "║", printer)
}

// printMidBoxesLine prints line of a sudoku puzzle that appears between boxes
func printMidBoxesLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig, printer types.Printer) {
	printHorizontalBorderLine(sudoku, printoutConfig, "║", "═", "║", "╬", printer)
}

// printBottomBorderLine prints bottom border line of a sudoku puzzle
func printBottomBorderLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig, printer types.Printer) {
	printHorizontalBorderLine(sudoku, printoutConfig, "╚", "═", "╝", "╩", printer)
}

// printValuesLine prinst single horizontal line with values of a sudoku puzzle
// with respect to padding
func printValuesLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig,
	boxRowIndex int8, cellRowIndex int8, printer types.Printer) {

	printer.PrintDefault("║")
	for boxColumnIndex := 0; boxColumnIndex < int(sudoku.Layout.Width); boxColumnIndex++ {
		sudokuBox := sudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
			return box.IndexColumn == int8(boxColumnIndex) && box.IndexRow == boxRowIndex
		})

		for cellColumnIndex := 0; cellColumnIndex < printoutConfig.BoxSize; cellColumnIndex++ {
			if cellColumnIndex > 0 {
				if sudokuBox.Disabled {
					printer.PrintDefault(" ")
				} else {
					printer.PrintDefault("│")
				}
			}

			if sudokuBox.Disabled {
				for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength; characterIndex++ {
					printer.PrintDefault(" ")
				}
			} else {
				sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCell) bool {
					return cell.IndexColumnInBox == int8(cellColumnIndex) && cell.IndexRowInBox == cellRowIndex
				})

				if sudokuCell.Value == nil {
					for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength+printoutConfig.Padding*2; characterIndex++ {
						printer.PrintDefault(" ")
					}
				} else {
					printValuePadding(printoutConfig, printer)
					printSudokuValue(sudokuCell, printoutConfig, printer)
					printValuePadding(printoutConfig, printer)
				}
			}
		}

		if boxColumnIndex < int(sudoku.Layout.Width)-1 {
			printer.PrintDefault("║")
		}
	}

	printer.PrintDefault("║")
	printer.PrintNewLine()
}

// printSudokuValue prints out correctly formatter sudoku value
func printSudokuValue(sudokuCell *models.SudokuCell, printoutConfig sudokuPrintoutConfig, printer types.Printer) {
	var printFunc func(text string)

	isCellViolatingSudokuRule := sudokuCell.HasViolationError()

	if sudokuCell.IsInputValue && isCellViolatingSudokuRule {
		printFunc = printer.PrintBoldError
	} else if sudokuCell.IsInputValue {
		printFunc = printer.PrintBold
	} else if isCellViolatingSudokuRule {
		printFunc = printer.PrintError
	} else {
		printFunc = printer.PrintDefault
	}

	switch printoutConfig.ValueCharactersLength {
	case 1:
		printFunc(strconv.Itoa(*sudokuCell.Value))
	case 2:
		printFunc(fmt.Sprintf("%-2s", strconv.Itoa(*sudokuCell.Value)))
	case 3:
		printFunc(fmt.Sprintf("%-3s", strconv.Itoa(*sudokuCell.Value)))
	default:
		printFunc("x")
	}
}

// printValuePadding prints padding before and after a sudoku value (horizontal padding)
func printValuePadding(printoutConfig sudokuPrintoutConfig, printer types.Printer) {
	if printoutConfig.Padding >= 1 {
		for paddingIndex := 0; paddingIndex < printoutConfig.Padding; paddingIndex++ {
			printer.PrintDefault(" ")
		}
	}
}

// printHorizontalBorderLine is a function that will print horizontal sudoku line
// (between horizontal boxes or cells) with provided characters
func printHorizontalBorderLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig,
	startSign string, middleSign string, endSign string, columnCrossSign string, printer types.Printer) {

	printer.PrintDefault(startSign)

	for sudokuBoxIndex := 0; sudokuBoxIndex < int(sudoku.Layout.Width); sudokuBoxIndex++ {
		for boxColumnIndex := 0; boxColumnIndex < printoutConfig.BoxSize; boxColumnIndex++ {
			if boxColumnIndex > 0 {
				printer.PrintDefault(middleSign)
			}
			for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength+printoutConfig.Padding*2; characterIndex++ {
				printer.PrintDefault(middleSign)
			}
		}

		if sudokuBoxIndex < int(sudoku.Layout.Width)-1 {
			printer.PrintDefault(columnCrossSign)
		}
	}

	printer.PrintDefault(endSign)
	printer.PrintNewLine()
}

// buildSudokuPrintoutConfig creates rintout configuration that is used to
// actually make a printout of a sudoku
func buildSudokuPrintoutConfig(sudoku *models.Sudoku, settings *models.Settings) sudokuPrintoutConfig {
	valueCharactersLength := len(strconv.Itoa(int(sudoku.BoxSize * sudoku.BoxSize)))

	valuesPerLine := int(sudoku.Layout.Width * sudoku.BoxSize)
	valuesCharactersCountPerLine := valuesPerLine * valueCharactersLength
	separatorsCount := valuesPerLine + 1

	return sudokuPrintoutConfig{
		ValueCharactersLength: valueCharactersLength,
		MaxIndex:              int(math.Max(0, float64(sudoku.BoxSize-1))),
		CharactersPerLine:     valuesCharactersCountPerLine + separatorsCount,
		BoxSize:               int(sudoku.BoxSize),
		Padding:               int(settings.SudokuPrintoutValuePaddingLength),
	}
}
