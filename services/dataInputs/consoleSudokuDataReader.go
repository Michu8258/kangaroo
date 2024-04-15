package dataInputs

import "github.com/Michu8258/kangaroo/models"

func ReadFromConsole(request models.SolveCommandRequest, settings *models.Settings) (*models.SudokuDTO, error) {

}

func readBoxSize(request models.SolveCommandRequest, settings *models.Settings) int8 {
	if request.BoxSize == nil {

	}

	if *request.BoxSize < settings.MinimumBoxSizeInclusive || *request.BoxSize > settings.MaximumBoxSizeInclusive {

	}

	return *request.BoxSize
}
