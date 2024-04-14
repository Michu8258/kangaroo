package printers

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/Michu8258/kangaroo/models"
)

type sudokuPrintoutConfig struct {
	ValueCharactersLength int
	MaxIndex              int
	CharactersPerLine     int
	BoxSize               int
	Padding               int
}

// PrintSudoku prints entire sudoku puzzle pseudo-graphical representation to the console
func PrintSudoku(sudoku *models.Sudoku, settings *models.Settings) {
	printoutConfig := buildSudokuPrintoutConfig(sudoku, settings)
	printTopBorderLine(sudoku, printoutConfig)

	for boxRowIndex := 0; boxRowIndex < int(sudoku.Layout.Height); boxRowIndex++ {
		for cellRowIndex := 0; cellRowIndex < int(sudoku.BoxSize); cellRowIndex++ {
			printValuesLine(sudoku, printoutConfig, int8(boxRowIndex), int8(cellRowIndex))
			if cellRowIndex < int(sudoku.BoxSize)-1 {
				printMidCellsLine(sudoku, printoutConfig)
			}
		}

		if boxRowIndex < int(sudoku.Layout.Height)-1 {
			printMidBoxesLine(sudoku, printoutConfig)
		}
	}

	printBottomBorderLine(sudoku, printoutConfig)
}

// printTopBorderLine prints top border line of a sudoku puzzle
func printTopBorderLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig) {
	printHorizontalBorderLine(sudoku, printoutConfig, "╔", "═", "╗", "╦")
}

// printMidCellsLine prints line of a sudoku puzzle that appears between cells
func printMidCellsLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig) {
	printHorizontalBorderLine(sudoku, printoutConfig, "║", "─", "║", "║")
}

// printMidBoxesLine prints line of a sudoku puzzle that appears between boxes
func printMidBoxesLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig) {
	printHorizontalBorderLine(sudoku, printoutConfig, "║", "═", "║", "╬")
}

// printBottomBorderLine prints bottom border line of a sudoku puzzle
func printBottomBorderLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig) {
	printHorizontalBorderLine(sudoku, printoutConfig, "╚", "═", "╝", "╩")
}

// printValuesLine prinst single horizontal line with values of a sudoku puzzle
// with respect to padding
func printValuesLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig,
	boxRowIndex int8, cellRowIndex int8) {
	var buffer bytes.Buffer

	printPadding := func() {
		if printoutConfig.Padding >= 1 {
			for paddingIndex := 0; paddingIndex < printoutConfig.Padding; paddingIndex++ {
				buffer.WriteString(" ")
			}
		}
	}

	buffer.WriteString("║")
	for boxColumnIndex := 0; boxColumnIndex < int(sudoku.Layout.Width); boxColumnIndex++ {
		sudokuBox := sudoku.Boxes.FirstOrDefault(nil, func(box *models.SudokuBox) bool {
			return box.IndexColumn == int8(boxColumnIndex) && box.IndexRow == boxRowIndex
		})

		for cellColumnIndex := 0; cellColumnIndex < printoutConfig.BoxSize; cellColumnIndex++ {
			if cellColumnIndex > 0 {
				if sudokuBox.Disabled {
					buffer.WriteString(" ")
				} else {
					buffer.WriteString("│")
				}
			}

			if sudokuBox.Disabled {
				for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength; characterIndex++ {
					buffer.WriteString(" ")
				}
			} else {
				sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCell) bool {
					return cell.IndexColumnInBox == int8(cellColumnIndex) && cell.IndexRowInBox == cellRowIndex
				})

				if sudokuCell.Value == nil {
					for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength+printoutConfig.Padding*2; characterIndex++ {
						buffer.WriteString(" ")
					}
				} else {
					printPadding()

					switch printoutConfig.ValueCharactersLength {
					case 1:
						buffer.WriteString(strconv.Itoa(*sudokuCell.Value))
					case 2:
						buffer.WriteString(fmt.Sprintf("%-2s", strconv.Itoa(*sudokuCell.Value)))
					case 3:
						buffer.WriteString(fmt.Sprintf("%-3s", strconv.Itoa(*sudokuCell.Value)))
					default:
						buffer.WriteString("x")
					}

					printPadding()
				}
			}
		}

		if boxColumnIndex < int(sudoku.Layout.Width)-1 {
			buffer.WriteString("║")
		}
	}
	buffer.WriteString("║")

	fmt.Println(buffer.String())
}

// printHorizontalBorderLine is a function that will print horizontal sudoku line
// (between horizontal boxes or cells) with provided characters
func printHorizontalBorderLine(sudoku *models.Sudoku, printoutConfig sudokuPrintoutConfig,
	startSign string, middleSign string, endSign string, columnCrossSign string) {
	var buffer bytes.Buffer

	buffer.WriteString(startSign)

	for sudokuBoxIndex := 0; sudokuBoxIndex < int(sudoku.Layout.Width); sudokuBoxIndex++ {
		for boxColumnIndex := 0; boxColumnIndex < printoutConfig.BoxSize; boxColumnIndex++ {
			if boxColumnIndex > 0 {
				buffer.WriteString(middleSign)
			}
			for characterIndex := 0; characterIndex < printoutConfig.ValueCharactersLength+printoutConfig.Padding*2; characterIndex++ {
				buffer.WriteString(middleSign)
			}
		}

		if sudokuBoxIndex < int(sudoku.Layout.Width)-1 {
			buffer.WriteString(columnCrossSign)
		}
	}

	buffer.WriteString(endSign)

	fmt.Println(buffer.String())
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
