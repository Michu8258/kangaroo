package models

import (
	"github.com/Michu8258/kangaroo/types"
	guid "github.com/nu7hatch/gouuid"
)

type SudokuCellDTO struct {
	Value            *int `json:"value"`
	IndexRowInBox    int8 `json:"indexRowInBox"`
	IndexColumnInBox int8 `json:"indexColumnInBox"`
}

type SudokuBoxDTO struct {
	Disabled    bool                               `json:"disabled"`
	IndexRow    int8                               `json:"indexRow"`
	IndexColumn int8                               `json:"indexColumn"`
	Cells       types.GenericSlice[*SudokuCellDTO] `json:"cells"`
}

type SudokuLayoutDTO struct {
	Width  int8 `json:"width"`
	Height int8 `json:"height"`
}

type SudokuDTO struct {
	BoxSize int8                              `json:"boxSize"`
	Layout  SudokuLayoutDTO                   `json:"layout"`
	Boxes   types.GenericSlice[*SudokuBoxDTO] `json:"boxes"`
}

// ToSudoku converts raw sudoku DTO object to internally managed object
// representing sudoku with all dependencies and computed data.
func (sudokuDto *SudokuDTO) ToSudoku() *Sudoku {
	sudoku := &Sudoku{
		BoxSize: sudokuDto.BoxSize,
		Layout: SudokuLayout{
			Height: sudokuDto.Layout.Height,
			Width:  sudokuDto.Layout.Width,
		},
		Boxes:      types.GenericSlice[*SudokuBox]{},
		SubSudokus: []*SubSudoku{},
		Result:     Unspecified,
	}

	for _, sudokuBoxDto := range sudokuDto.Boxes {
		boxId, _ := guid.NewV4()
		sudokuBox := &SudokuBox{
			Id:           *boxId,
			Disabled:     sudokuBoxDto.Disabled,
			IndexRow:     sudokuBoxDto.IndexRow,
			IndexColumn:  sudokuBoxDto.IndexColumn,
			Cells:        types.GenericSlice[*SudokuCell]{},
			ViolatesRule: false,
		}

		for _, sudokuCellDto := range sudokuBoxDto.Cells {
			cellId, _ := guid.NewV4()
			sudokuBox.Cells = append(sudokuBox.Cells, &SudokuCell{
				Id:               *cellId,
				Value:            sudokuCellDto.Value,
				IsInputValue:     sudokuCellDto.Value != nil,
				PotentialValues:  nil,
				IndexRowInBox:    sudokuCellDto.IndexRowInBox,
				IndexColumnInBox: sudokuCellDto.IndexColumnInBox,
				Box:              nil,
				MemberOfLines:    types.GenericSlice[*SudokuLine]{},
			})
		}

		sudoku.Boxes = append(sudoku.Boxes, sudokuBox)
	}

	return sudoku
}
