package binarySudokuManager

import "github.com/Michu8258/kangaroo/models"

type BinarySudokuManager struct {
	Settings *models.Settings
}

type IBinarySudokuManager interface {
	ReadFromBase64(base64Data string) (*models.SudokuDTO, error)
	ReadFromBytes(sudokuData []byte) (*models.SudokuDTO, error)
	ToBase64(sudokuDto *models.SudokuDTO) (string, error)
	ToBytes(sudokuDto *models.SudokuDTO) ([]byte, error)
}

func GetNewBinarySudokuManager(settings *models.Settings) IBinarySudokuManager {
	return &BinarySudokuManager{
		Settings: settings,
	}
}
