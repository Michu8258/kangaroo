package dataPrinters

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/printer"
)

type sudokuPrintoutConfig struct {
	ValueCharactersLength int
	MaxIndex              int
	CharactersPerLine     int
	BoxSize               int
	Padding               int
}

// PrintSudoku prints entire sudoku puzzle pseudo-graphical representation to the console
func (dp *DataPrinter) PrintSudoku(sudoku *models.Sudoku, printer printer.IPrinter) {
	defer func() {
		if err := recover(); err != nil {
			printer.PrintNewLine()
			printer.PrintError("Failed to render a sudoku puzzle.")
			printer.PrintNewLine()
		}
	}()

	printoutConfig := dp.buildSudokuPrintoutConfig(sudoku)
	dp.printTopBorderLine(sudoku, printoutConfig, printer)

	var boxRowIndex int8 = 0
	var cellRowIndex int8 = 0

	for boxRowIndex = 0; boxRowIndex < sudoku.Layout.Height; boxRowIndex++ {
		for cellRowIndex = 0; cellRowIndex < sudoku.BoxSize; cellRowIndex++ {
			dp.printValuesLine(sudoku, printer, printoutConfig, int8(boxRowIndex), int8(cellRowIndex))
			if cellRowIndex < sudoku.BoxSize-1 {
				dp.printMidCellsLine(sudoku, printer, printoutConfig, boxRowIndex)
			}
		}

		if boxRowIndex < sudoku.Layout.Height-1 {
			dp.printMidBoxesLine(sudoku, printoutConfig, printer)
		}
	}

	dp.printBottomBorderLine(sudoku, printoutConfig, printer)
}

// printTopBorderLine prints top border line of a sudoku puzzle
func (dp *DataPrinter) printTopBorderLine(sudoku *models.Sudoku,
	printoutConfig sudokuPrintoutConfig, printer printer.IPrinter) {

	dp.printHorizontalBorderLine(sudoku, printoutConfig, printer, "╔", "═", "╗", "╦")
}

// printMidBoxesLine prints line of a sudoku puzzle that appears between boxes
func (dp *DataPrinter) printMidBoxesLine(sudoku *models.Sudoku,
	printoutConfig sudokuPrintoutConfig, printer printer.IPrinter) {

	dp.printHorizontalBorderLine(sudoku, printoutConfig, printer, "║", "═", "║", "╬")
}

// printBottomBorderLine prints bottom border line of a sudoku puzzle
func (dp *DataPrinter) printBottomBorderLine(sudoku *models.Sudoku,
	printoutConfig sudokuPrintoutConfig, printer printer.IPrinter) {

	dp.printHorizontalBorderLine(sudoku, printoutConfig, printer, "╚", "═", "╝", "╩")
}

// printMidCellsLine prints line of a sudoku puzzle that appears between cells
func (dp *DataPrinter) printMidCellsLine(sudoku *models.Sudoku, printer printer.IPrinter,
	printoutConfig sudokuPrintoutConfig, boxRowIndex int8) {

	printer.PrintBorder("║")

	var sudokuBoxIndex int8 = 0
	var boxColumnIndex int8 = 0

	for sudokuBoxIndex = 0; sudokuBoxIndex < sudoku.Layout.Width; sudokuBoxIndex++ {
		for boxColumnIndex = 0; boxColumnIndex < int8(printoutConfig.BoxSize); boxColumnIndex++ {
			middleSign := "─"
			box := sudoku.Boxes.FirstOrDefault(nil, func(b *models.SudokuBox) bool {
				return b.IndexRow == boxRowIndex && b.IndexColumn == sudokuBoxIndex
			})

			if box != nil && box.Disabled {
				middleSign = " "
			}

			if boxColumnIndex > 0 {
				printer.PrintBorder(middleSign)
			}
			for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength+printoutConfig.Padding*2; characterIndex++ {
				printer.PrintBorder(middleSign)
			}
		}

		if sudokuBoxIndex < sudoku.Layout.Width-1 {
			printer.PrintBorder("║")
		}
	}

	printer.PrintBorder("║")
	printer.PrintNewLine()
}

// printValuesLine prinst single horizontal line with values of a sudoku puzzle
// with respect to padding
func (dp *DataPrinter) printValuesLine(sudoku *models.Sudoku, printer printer.IPrinter,
	printoutConfig sudokuPrintoutConfig, boxRowIndex int8, cellRowIndex int8) {

	printer.PrintBorder("║")
	for boxColumnIndex := 0; boxColumnIndex < int(sudoku.Layout.Width); boxColumnIndex++ {
		sudokuBox := sudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
			return box.IndexColumn == int8(boxColumnIndex) && box.IndexRow == boxRowIndex
		})

		for cellColumnIndex := 0; cellColumnIndex < printoutConfig.BoxSize; cellColumnIndex++ {
			if cellColumnIndex > 0 {
				if sudokuBox.Disabled {
					printer.PrintDefault(" ")
				} else {
					printer.PrintBorder("│")
				}
			}

			if sudokuBox.Disabled {
				for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength; characterIndex++ {
					dp.printNoValuePlaceholder(printoutConfig, printer)
				}
			} else {
				sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCell) bool {
					return cell.IndexColumnInBox == int8(cellColumnIndex) && cell.IndexRowInBox == cellRowIndex
				})

				if sudokuCell.Value == nil {
					dp.printNoValuePlaceholder(printoutConfig, printer)
				} else {
					dp.printValuePadding(printoutConfig, printer)
					dp.printSudokuValue(sudokuCell, printoutConfig, printer)
					dp.printValuePadding(printoutConfig, printer)
				}
			}
		}

		if boxColumnIndex < int(sudoku.Layout.Width)-1 {
			printer.PrintBorder("║")
		}
	}

	printer.PrintBorder("║")
	printer.PrintNewLine()
}

// printSudokuValue prints out correctly formatter sudoku value
func (dp *DataPrinter) printSudokuValue(sudokuCell *models.SudokuCell,
	printoutConfig sudokuPrintoutConfig, printer printer.IPrinter) {

	var printFunc func(text string)

	isCellViolatingSudokuRule := sudokuCell.HasViolationError()

	if sudokuCell.IsInputValue && isCellViolatingSudokuRule {
		printFunc = printer.PrintError
	} else if sudokuCell.IsInputValue {
		printFunc = printer.PrintPrimary
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

// printNoValuePlaceholder prints empty spaces with amount adjusted with characters
// per value and padding
func (dp *DataPrinter) printNoValuePlaceholder(printoutConfig sudokuPrintoutConfig,
	printer printer.IPrinter) {

	dp.printValuePadding(printoutConfig, printer)
	for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength; characterIndex++ {
		printer.PrintDefault(" ")
	}

	dp.printValuePadding(printoutConfig, printer)
}

// printValuePadding prints padding before and after a sudoku value (horizontal padding)
func (dp *DataPrinter) printValuePadding(printoutConfig sudokuPrintoutConfig,
	printer printer.IPrinter) {

	if printoutConfig.Padding >= 1 {
		for paddingIndex := 0; paddingIndex < printoutConfig.Padding; paddingIndex++ {
			printer.PrintDefault(" ")
		}
	}
}

// printHorizontalBorderLine is a function that will print horizontal sudoku line
// (between horizontal boxes or cells) with provided characters
func (dp *DataPrinter) printHorizontalBorderLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig,
	printer printer.IPrinter, startSign string, middleSign string, endSign string, columnCrossSign string) {

	printer.PrintBorder(startSign)

	for sudokuBoxIndex := 0; sudokuBoxIndex < int(sudoku.Layout.Width); sudokuBoxIndex++ {
		for boxColumnIndex := 0; boxColumnIndex < printoutConfig.BoxSize; boxColumnIndex++ {
			if boxColumnIndex > 0 {
				printer.PrintBorder(middleSign)
			}
			for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength+printoutConfig.Padding*2; characterIndex++ {
				printer.PrintBorder(middleSign)
			}
		}

		if sudokuBoxIndex < int(sudoku.Layout.Width)-1 {
			printer.PrintBorder(columnCrossSign)
		}
	}

	printer.PrintBorder(endSign)
	printer.PrintNewLine()
}

// buildSudokuPrintoutConfig creates rintout configuration that is used to
// actually make a printout of a sudoku
func (dp *DataPrinter) buildSudokuPrintoutConfig(sudoku *models.Sudoku) sudokuPrintoutConfig {
	valueCharactersLength := len(strconv.Itoa(int(sudoku.BoxSize * sudoku.BoxSize)))

	valuesPerLine := int(sudoku.Layout.Width * sudoku.BoxSize)
	valuesCharactersCountPerLine := valuesPerLine * valueCharactersLength
	separatorsCount := valuesPerLine + 1

	return sudokuPrintoutConfig{
		ValueCharactersLength: valueCharactersLength,
		MaxIndex:              int(math.Max(0, float64(sudoku.BoxSize-1))),
		CharactersPerLine:     valuesCharactersCountPerLine + separatorsCount,
		BoxSize:               int(sudoku.BoxSize),
		Padding:               int(dp.Settings.SudokuPrintoutValuePaddingLength),
	}
}
