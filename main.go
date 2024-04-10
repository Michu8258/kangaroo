package main

import (
	"fmt"
	"log"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/services/dataInputs"
	"github.com/Michu8258/kangaroo/services/initialization"
)

func main() {
	rawSudokuData, err := dataInputs.ReadFromFile("./textConfigs/simple01.json")
	if err != nil {
		log.Fatalln(err)
	}

	settings := createSettings()
	sudoku := rawSudokuData.ToSudoku()
	errs := initialization.InitializeSudoku(sudoku, settings)
	if len(errs) >= 1 {
		for _, err := range errs {
			log.Println(err)
		}
		return
	}

	fmt.Println("Hello kangaroo")
	fmt.Println("Amount of subsudokus", len(sudoku.SubSudokus))
	fmt.Println("Data for cell in the center")
	middleBox := sudoku.SubSudokus[0].Boxes[4]
	fmt.Println(middleBox)
	middleCell := middleBox.Cells[4]
	fmt.Println(middleCell)
	fmt.Println("lines count", len(middleCell.MemberOfLines))
	for index, line := range middleCell.MemberOfLines {
		fmt.Print("LINE", index, "\t")
		for _, cell := range line.Cells {
			if cell.Value != nil {
				fmt.Print(*cell.Value, " ")
			} else {
				fmt.Print(0, " ")
			}
		}
		fmt.Println()
	}
	fmt.Println("subsudoku lines count", len(sudoku.SubSudokus[0].ChildLines))
}

func createSettings() *models.Settings {
	return &models.Settings{
		MinimumLayoutSizeInclusive: 2,
		MaximumLayoutSizeInclusive: 10,
		MinimumBoxSizeInclusive:    2,
		MaximumBoxSizeInclusive:    5,
	}
}
