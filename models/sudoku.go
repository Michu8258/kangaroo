package models

import (
	guid "github.com/nu7hatch/gouuid"
)

const SudokuLineTypeRow = "row"
const SudokuLineTypeColumn = "column"

type SudokuCell struct {
	Id               guid.UUID
	Value            *int
	IsInputValue     bool
	PotentialValues  *GenericSlice[int]
	IndexRowInBox    int8
	IndexColumnInBox int8
	Box              *SudokuBox
	MemberOfLines    GenericSlice[*SudokuLine]
}

func (cell *SudokuCell) HasViolationError() bool {
	if cell.Box != nil && cell.Box.ViolatesRule {
		return true
	}

	if cell.MemberOfLines == nil || len(cell.MemberOfLines) < 1 {
		return false
	}

	return cell.MemberOfLines.Any(func(line *SudokuLine) bool {
		return line.ViolatesRule
	})
}

type SudokuLine struct {
	Cells        GenericSlice[*SudokuCell]
	LineType     string
	ViolatesRule bool
	SubsudokuId  guid.UUID
}

type SudokuBox struct {
	Id           guid.UUID
	Disabled     bool
	IndexRow     int8
	IndexColumn  int8
	Cells        GenericSlice[*SudokuCell]
	ViolatesRule bool
}

type SubSudoku struct {
	Id                    guid.UUID
	Boxes                 GenericSlice[*SudokuBox]
	TopLeftBoxRowIndex    int8
	TopLeftBoxColumnIndex int8
	ChildLines            GenericSlice[*SudokuLine]
}

type SudokuLayout struct {
	Width  int8
	Height int8
}

type Sudoku struct {
	BoxSize    int8
	Layout     SudokuLayout
	Boxes      GenericSlice[*SudokuBox]
	SubSudokus GenericSlice[*SubSudoku]
	Result     SudokuResultType
}

type SudokuValueGuess struct {
	GuessedValue            int
	GuessedCell             *SudokuCell
	SubsudokuId             guid.UUID
	PotentialValuesSnapshot map[guid.UUID]*[]int
}

type SudokuResultType int8

const (
	Unspecified         SudokuResultType = 0
	SuccessfullSolution SudokuResultType = 1
	Failure             SudokuResultType = 2
	InvalidGuess        SudokuResultType = 3
	UnsolvableSudoku    SudokuResultType = 4
)

// ToSudoku converts internal sudoku object to DTO object.
// Suitable for serialization to json
func (sudoku *Sudoku) ToSudokuDto() *SudokuDTO {
	sudokuDto := &SudokuDTO{
		BoxSize: sudoku.BoxSize,
		Layout: SudokuLayoutDTO{
			Height: sudoku.Layout.Height,
			Width:  sudoku.Layout.Width,
		},
		Boxes: GenericSlice[*SudokuBoxDTO]{},
	}

	for _, sudokuBox := range sudoku.Boxes {
		sudokuBoxDto := &SudokuBoxDTO{
			Disabled:    sudokuBox.Disabled,
			IndexRow:    sudokuBox.IndexRow,
			IndexColumn: sudokuBox.IndexColumn,
			Cells:       GenericSlice[*SudokuCellDTO]{},
		}

		for _, sudokuCell := range sudokuBox.Cells {
			sudokuBoxDto.Cells = append(sudokuBoxDto.Cells, &SudokuCellDTO{
				Value:            sudokuCell.Value,
				IndexRowInBox:    sudokuCell.IndexRowInBox,
				IndexColumnInBox: sudokuCell.IndexColumnInBox,
			})
		}

		sudokuDto.Boxes = append(sudokuDto.Boxes, sudokuBoxDto)
	}

	return sudokuDto
}
