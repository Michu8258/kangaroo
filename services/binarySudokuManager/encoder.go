package binarySudokuManager

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/Michu8258/kangaroo/models"
)

// ToBase64 converts sudoku dto object to its base64 string representation
func (manager *BinarySudokuManager) ToBase64(sudokuDto *models.SudokuDTO) (string, error) {
	dataBytes, err := manager.ToBytes(sudokuDto)
	if err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(dataBytes)
	return base64Str, nil
}

// ToBytes converts sudoku dto object to its binary data representation
func (manager *BinarySudokuManager) ToBytes(sudokuDto *models.SudokuDTO) ([]byte, error) {
	var version uint16 = manager.Settings.SudokuBinaryEncoderVersion
	result := []byte{}

	handlers := map[uint16]func(sudokuDto *models.SudokuDTO, result []byte) ([]byte, error){
		1: manager.WriteVersion1,
	}

	matchingHandler, ok := handlers[version]
	if !ok {
		return result, fmt.Errorf(
			"sudoku data binary representation version %d is not supported",
			version)
	}

	result = binary.BigEndian.AppendUint16(result, version)
	return matchingHandler(sudokuDto, result)
}

// WriteVersion1 implements logic for writing sudoku binary data for version 1
func (manager *BinarySudokuManager) WriteVersion1(sudokuDto *models.SudokuDTO, result []byte) ([]byte, error) {
	//box size
	result = append(result, byte(sudokuDto.BoxSize))

	// layout width and height
	result = append(result, byte(sudokuDto.Layout.Width), byte(sudokuDto.Layout.Height))

	// enabed/disabled boxes data
	enableStateBytes, err := buildEnableStateBytes(sudokuDto)
	if err != nil {
		return result, err
	}
	result = append(result, enableStateBytes...)

	// cells data
	cellsBytes, err := buildCellsData(sudokuDto)
	if err != nil {
		return result, err
	}
	result = append(result, cellsBytes...)

	return result, nil
}

// buildEnableStateBytes creates binary data holding information about
// sudoku bixes enabled/disabled state
func buildEnableStateBytes(sudokuDto *models.SudokuDTO) ([]byte, error) {
	var rowIndex int8 = 0
	var columnIndex int8 = 0

	boxesCount := sudokuDto.Layout.Width * sudokuDto.Layout.Height
	amountOfBytes := int(math.Ceil(float64(boxesCount) / 8))
	result := make([]byte, amountOfBytes)

	for rowIndex = 0; rowIndex < sudokuDto.Layout.Height; rowIndex++ {
		for columnIndex = 0; columnIndex < sudokuDto.Layout.Width; columnIndex++ {
			sudokuBox := sudokuDto.Boxes.FirstOrDefault(nil, func(box *models.SudokuBoxDTO) bool {
				return box.IndexRow == rowIndex && box.IndexColumn == columnIndex
			})

			if sudokuBox == nil {
				return result, errors.New(
					"could not find box during enable/disable sudoku binary data construction")
			}

			boxIndex := rowIndex*sudokuDto.Layout.Width + columnIndex
			byteIndex := int(float64(boxIndex / 8))
			bitIndex := boxIndex % 8
			logicalOrOperationValue := byte(math.Pow(2, float64(7-bitIndex)))

			if !sudokuBox.Disabled {
				result[byteIndex] |= logicalOrOperationValue
			}
		}
	}

	return result, nil
}

// buildCellsData creates binary data holding information about cell values
func buildCellsData(sudokuDto *models.SudokuDTO) ([]byte, error) {
	var rowIndex int8 = 0
	var columnIndex int8 = 0

	result := []byte{}

	// iterate through boxes - we have to do it by searching by indexes because order of data
	// is important in binary representation
	for rowIndex = 0; rowIndex < sudokuDto.Layout.Height; rowIndex++ {
		for columnIndex = 0; columnIndex < sudokuDto.Layout.Width; columnIndex++ {
			sudokuBox := sudokuDto.Boxes.FirstOrDefault(nil, func(box *models.SudokuBoxDTO) bool {
				return box.IndexRow == rowIndex && box.IndexColumn == columnIndex
			})

			if !sudokuBox.Disabled {
				values := make([]byte, sudokuDto.BoxSize*sudokuDto.BoxSize)
				var cellRowIndex int8 = 0
				var cellColumnIndex int8 = 0

				// iterate through cells in box - we have to do it by searching by indexes
				// because order of data is important in binary representation
				for cellRowIndex = 0; cellRowIndex < sudokuDto.BoxSize; cellRowIndex++ {
					for cellColumnIndex = 0; cellColumnIndex < sudokuDto.BoxSize; cellColumnIndex++ {
						sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
							return cell.IndexRowInBox == cellRowIndex && cell.IndexColumnInBox == cellColumnIndex
						})

						if sudokuCell == nil {
							return result, errors.New(
								"could not find cell during values sudoku binary data construction")
						}

						value := 0
						if sudokuCell.Value != nil {
							value = *sudokuCell.Value
						}

						values[cellRowIndex*sudokuDto.BoxSize+cellColumnIndex] = byte(value)
					}
				}

				result = append(result, values...)
			}
		}
	}

	return result, nil
}
