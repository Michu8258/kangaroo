package testHelpers

import "github.com/Michu8258/kangaroo/models"

func GetTestSudokuDto() *models.SudokuDTO {
	return &models.SudokuDTO{
		BoxSize: 3,
		Layout: models.SudokuLayoutDTO{
			Width:  3,
			Height: 3,
		},
		Boxes: models.GenericSlice[*models.SudokuBoxDTO]{
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    0,
				IndexColumn: 0,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    0,
				IndexColumn: 1,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    0,
				IndexColumn: 2,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    1,
				IndexColumn: 0,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    1,
				IndexColumn: 1,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    1,
				IndexColumn: 2,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    2,
				IndexColumn: 0,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    2,
				IndexColumn: 1,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
			&models.SudokuBoxDTO{
				Disabled:    false,
				IndexRow:    2,
				IndexColumn: 2,
				Cells: models.GenericSlice[*models.SudokuCellDTO]{
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    0,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    1,
						IndexColumnInBox: 2,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 0,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 1,
					},
					&models.SudokuCellDTO{
						Value:            nil,
						IndexRowInBox:    2,
						IndexColumnInBox: 2,
					},
				},
			},
		},
	}
}
