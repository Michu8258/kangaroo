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
	fmt.Println(rawSudokuData.Boxes[0].Cells[0].Id.String())
	fmt.Println(rawSudokuData.Boxes[0].Cells[0].Value)
	fmt.Println(rawSudokuData.Boxes[0].Cells[0].PotentialValues)
}

func createSettings() *models.Settings {
	return &models.Settings{
		MinimumLayoutSizeInclusive: 2,
		MaximumLayoutSizeInclusive: 10,
		MinimumBoxSizeInclusive:    2,
		MaximumBoxSizeInclusive:    5,
	}
}
