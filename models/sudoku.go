package models

import "github.com/beevik/guid"

type SudokuCell struct {
	Id               guid.Guid      `json:"id"`
	Value            *int           `json:"value"`
	PotentialValues  *[]int         `json:"potentialValues"`
	IndexRowInBox    int8           `json:"indexRowInBox"`
	IndexColumnInBox int8           `json:"indexColumnInBox"`
	Box              *SudokuBox     `json:"box"`
	MemberOfLines    *[]*SudokuLine `json:"memberOfLines"`
}

type SudokuLine struct {
	Cells []*SudokuCell `json:"cells"`
}

type SudokuBox struct {
	Id          guid.Guid     `json:"id"`
	Disabled    bool          `json:"disabled"`
	IndexRow    int8          `json:"indexRow"`
	IndexColumn int8          `json:"indexColumn"`
	Cells       []*SudokuCell `json:"cells"`
}

type SubSudoku struct {
	Boxes []*SudokuBox `json:"boxes"`
}

type SudokuLayout struct {
	Width  int8 `json:"width"`
	Height int8 `json:"height"`
}

type Sudoku struct {
	BoxSize    int8         `json:"boxSize"`
	Layout     SudokuLayout `json:"layout"`
	Boxes      []*SudokuBox `json:"boxes"`
	SubSudokus []*SubSudoku `json:"subSudokus"`
}
