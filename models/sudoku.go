package models

import (
	"github.com/Michu8258/kangaroo/types"
	"github.com/beevik/guid"
)

// TODO - introduce serializable sudoku DTO that will be possible to Convert to full Sudoku object

type SudokuCell struct {
	Id               guid.Guid                       `json:"id"`
	Value            *int                            `json:"value"`
	PotentialValues  *types.GenericSlice[int]        `json:"potentialValues"`
	IndexRowInBox    int8                            `json:"indexRowInBox"`
	IndexColumnInBox int8                            `json:"indexColumnInBox"`
	Box              *SudokuBox                      `json:"box"`
	MemberOfLines    types.GenericSlice[*SudokuLine] `json:"memberOfLines"`
}

type SudokuLine struct {
	Cells types.GenericSlice[*SudokuCell] `json:"cells"`
}

type SudokuBox struct {
	Id          guid.Guid                       `json:"id"`
	Disabled    bool                            `json:"disabled"`
	IndexRow    int8                            `json:"indexRow"`
	IndexColumn int8                            `json:"indexColumn"`
	Cells       types.GenericSlice[*SudokuCell] `json:"cells"`
}

type SubSudoku struct {
	Boxes                 types.GenericSlice[*SudokuBox] `json:"boxes"`
	TopLeftBoxRowIndex    int8                           `json:"topLeftBoxRowIndex"`
	TopLeftBoxColumnIndex int8                           `json:"topLeftBoxColumnIndex"`
}

type SudokuLayout struct {
	Width  int8 `json:"width"`
	Height int8 `json:"height"`
}

type Sudoku struct {
	BoxSize    int8                           `json:"boxSize"`
	Layout     SudokuLayout                   `json:"layout"`
	Boxes      types.GenericSlice[*SudokuBox] `json:"boxes"`
	SubSudokus types.GenericSlice[*SubSudoku] `json:"subSudokus"`
}
