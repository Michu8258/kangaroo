package dataWriter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Michu8258/kangaroo/helpers"
	"github.com/Michu8258/kangaroo/models"
)

// SaveSudokuToJson executes sudoku object JSON dump to selected file.
// Returns flag if indicating if file was written and potential error
func (writer *DataWriter) SaveSudokuToJson(sudoku *models.Sudoku, path string, overwrite bool) (bool, error) {
	sudokuDto := sudoku.ToSudokuDto()
	return writer.SaveSudokuDtoToJson(sudokuDto, path, overwrite)
}

// SaveSudokuDtoToJson executes sudoku DTO object JSON dump to selected file.
// Returns flag if indicating if file was written and potential error
func (writer *DataWriter) SaveSudokuDtoToJson(sudokuDto *models.SudokuDTO, path string, overwrite bool) (bool, error) {
	saveConfig := writer.prepareSaveConfig(path, overwrite)
	if saveConfig.shortCircuit {
		return false, saveConfig.err
	}

	jsonBytes, err := json.MarshalIndent(sudokuDto, "", "  ")
	if sudokuDto == nil || err != nil {
		return false, fmt.Errorf("failed to generate sudoku json string for file '%s'",
			saveConfig.absoluteFilePath)
	}

	err = os.WriteFile(saveConfig.absoluteFilePath, jsonBytes, 0644)
	if err != nil {
		return false, fmt.Errorf("failed to save sudoku json data file '%s'",
			saveConfig.absoluteFilePath)
	}

	return true, nil
}

// SaveSudokuToJson executes sudoku object TXT dump to selected file.
// Returns flag if indicating if file was written and potential error
func (writer *DataWriter) SaveSudokuToTxt(sudoku *models.Sudoku,
	path string, overwrite bool) (bool, error) {

	saveConfig := writer.prepareSaveConfig(path, overwrite)
	if saveConfig.shortCircuit {
		return false, saveConfig.err
	}

	file, err := os.Create(saveConfig.absoluteFilePath)
	if err != nil {
		return false, fmt.Errorf("failed to create file '%s'", saveConfig.absoluteFilePath)
	}

	defer file.Close()

	txtPrinter := writer.TxtPrinterProvider(file)
	writer.DataPrinter.PrintSudoku(sudoku, txtPrinter)
	file.Sync()

	return true, nil
}

type saveConfig struct {
	shortCircuit     bool
	absoluteFilePath string
	err              error
}

// prepareSaveConfig executes initial checks and validates logic agains
// file overwrite flag. Returns common file save configuration
func (writer *DataWriter) prepareSaveConfig(path string, overwrite bool) saveConfig {
	absolutePath, err := helpers.MakeFilePathAbsolute(path)
	if err != nil {
		return saveConfig{
			shortCircuit:     true,
			absoluteFilePath: absolutePath,
			err:              err,
		}
	}

	fileExists := false

	if _, err := os.Stat(absolutePath); err == nil {
		fileExists = true
	}

	if fileExists && !overwrite {
		return saveConfig{
			shortCircuit:     true,
			absoluteFilePath: absolutePath,
			err:              nil,
		}
	}

	return saveConfig{
		shortCircuit:     false,
		absoluteFilePath: absolutePath,
		err:              nil,
	}
}
