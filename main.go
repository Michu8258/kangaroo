package main

import (
	"fmt"
	"log"

	"github.com/Michu8258/kangaroo/models"
	crook "github.com/Michu8258/kangaroo/services/crookMethodSolver"
	"github.com/Michu8258/kangaroo/services/dataInputs"
	"github.com/Michu8258/kangaroo/services/initialization"
)

func main() {
	fmt.Println("KANGAROO")

	settings := createSettings()
	rawSudokuData, err := dataInputs.ReadFromFile("./textConfigs/simple01.json")
	if err != nil {
		log.Fatalln(err)
	}

	sudoku := rawSudokuData.ToSudoku()
	errs := initialization.InitializeSudoku(sudoku, settings)
	if len(errs) >= 1 {
		for _, err := range errs {
			log.Println(err)
		}
		return
	}

	fmt.Println("Amount of subSudokus", len(sudoku.SubSudokus))

	errs = crook.SolveWithCrookMethod(sudoku, settings)
	if len(errs) >= 1 {
		for _, err := range errs {
			log.Println(err)
		}
		return
	}
}

func createSettings() *models.Settings {
	return &models.Settings{
		MinimumLayoutSizeInclusive: 2,
		MaximumLayoutSizeInclusive: 10,
		MinimumBoxSizeInclusive:    2,
		MaximumBoxSizeInclusive:    5,
		UseDebugPrints:             true,
	}
}
