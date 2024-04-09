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
	errs := initialization.InitializeSudoku(rawSudokuData, settings)
	if len(errs) >= 1 {
		for _, err := range errs {
			log.Println(err)
		}
		log.Fatal("Failed to initialize the sudoku object")
	}

	fmt.Println("Hello kangaroo")
	fmt.Println("Amount of subsudokus", len(rawSudokuData.SubSudokus))
	fmt.Println("Data for cell in the center")
	middleBox := rawSudokuData.SubSudokus[0].Boxes[4]
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
}

func createSettings() *models.Settings {
	return &models.Settings{
		MinimumLayoutSizeInclusive: 2,
		MaximumLayoutSizeInclusive: 10,
		MinimumBoxSizeInclusive:    2,
		MaximumBoxSizeInclusive:    5,
	}
}
