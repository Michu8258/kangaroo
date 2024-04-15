package models

import (
	"github.com/Michu8258/kangaroo/types"
	"github.com/beevik/guid"
)

const SudokuLineTypeRow = "row"
const SudokuLineTypeColumn = "column"

type SudokuCell struct {
	Id               guid.Guid
	Value            *int
	IsInputValue     bool
	PotentialValues  *types.GenericSlice[int]
	IndexRowInBox    int8
	IndexColumnInBox int8
	Box              *SudokuBox
	MemberOfLines    types.GenericSlice[*SudokuLine]
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
	Cells        types.GenericSlice[*SudokuCell]
	LineType     string
	ViolatesRule bool
	SubsudokuId  guid.Guid
}

type SudokuBox struct {
	Id           guid.Guid
	Disabled     bool
	IndexRow     int8
	IndexColumn  int8
	Cells        types.GenericSlice[*SudokuCell]
	ViolatesRule bool
}

type SubSudoku struct {
	Id                    guid.Guid
	Boxes                 types.GenericSlice[*SudokuBox]
	TopLeftBoxRowIndex    int8
	TopLeftBoxColumnIndex int8
	ChildLines            types.GenericSlice[*SudokuLine]
}

type SudokuLayout struct {
	Width  int8
	Height int8
}

type Sudoku struct {
	BoxSize    int8
	Layout     SudokuLayout
	Boxes      types.GenericSlice[*SudokuBox]
	SubSudokus types.GenericSlice[*SubSudoku]
	Result     SudokuResultType
}

type SudokuValueGuess struct {
	GuessedValue            int
	GuessedCell             *SudokuCell
	SubsudokuId             guid.Guid
	PotentialValuesSnapshot map[guid.Guid]*[]int
}

type SudokuResultType int8

const (
	Unspecified         SudokuResultType = 0
	SuccessfullSolution SudokuResultType = 1
	Failure             SudokuResultType = 2
	InvalidGuess        SudokuResultType = 3
	UnsolvableSudoku    SudokuResultType = 4
)
