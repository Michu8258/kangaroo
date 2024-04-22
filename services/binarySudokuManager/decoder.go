package binarySudokuManager

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/Michu8258/kangaroo/models"
)

// ReadFromBase64 reads sudoku DTO from base64 representation
// of sudoku data
func (manager *BinarySudokuManager) ReadFromBase64(base64Data string) (*models.SudokuDTO, error) {
	sudokuDataBytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, err
	}

	return manager.ReadFromBytes(sudokuDataBytes)
}

// ReadFromBytes reads sudoku DTO from bytes of binary representation
// of sudoku data
func (manager *BinarySudokuManager) ReadFromBytes(sudokuData []byte) (*models.SudokuDTO, error) {
	version, err := getVersionNumber(sudokuData)
	if err != nil {
		return nil, err
	}

	handlers := map[uint16]func(sudokuData []byte) (*models.SudokuDTO, error){
		1: manager.ReadVersion1,
	}

	matchingHandler, ok := handlers[version]
	if !ok {
		return nil, fmt.Errorf(
			"sudoku data binary representation version %d is not supported",
			version)
	}

	return matchingHandler(sudokuData)
}

// ReadVersion1 implements binary data to sudoku DTO parsing for version 1
// of binary representation format
func (manager *BinarySudokuManager) ReadVersion1(sudokuData []byte) (*models.SudokuDTO, error) {
	boxSize, err := getSudokuBoxSize(sudokuData)
	if err != nil {
		return nil, err
	}

	layout, err := getSudokuLayout(sudokuData)
	if err != nil {
		return nil, err
	}

	boxes, err := getBoxesData(sudokuData, boxSize, layout)
	if err != nil {
		return nil, err
	}

	sudokuDto := &models.SudokuDTO{
		BoxSize: boxSize,
		Layout:  *layout,
		Boxes:   boxes,
	}

	return sudokuDto, nil
}

// getBoxesData reads complete ssudoku boxes data out of provided binary
// representation. It includes enabled/disabled state handling and nil
// values
func getBoxesData(sudokuData []byte, boxSize int8, layout *models.SudokuLayoutDTO) (
	[]*models.SudokuBoxDTO, error) {

	enableState, enableStateDataBytesCount, err := getBoxesEnableStateData(sudokuData, layout)
	if err != nil {
		return nil, err
	}

	sudokuBoxesDataStartIndex := 5 + enableStateDataBytesCount
	oneBoxDataBytesCount := int(boxSize * boxSize)

	boxes := []*models.SudokuBoxDTO{}
	for boxIndex, isEnabled := range enableState {
		box := &models.SudokuBoxDTO{}
		box.Disabled = !isEnabled
		box.Cells = []*models.SudokuCellDTO{}
		setBoxIndexes(box, layout, boxIndex)

		err := assignCellValues(box, boxSize, sudokuData,
			sudokuBoxesDataStartIndex, oneBoxDataBytesCount)
		if err != nil {
			return nil, err
		}

		if isEnabled {
			sudokuBoxesDataStartIndex += oneBoxDataBytesCount
		}

		boxes = append(boxes, box)
	}

	return boxes, nil
}

// assignCellValues assigns cells with values to the box according to bytes
// in sudoku binary representation. It takes care of the case if box is
// disabled - cells with no values will be added to the box and no attempt
// of reading bytes from binary representation will be performed.
func assignCellValues(box *models.SudokuBoxDTO, boxSize int8, sudokuData []byte,
	startByteIndex int, oneBoxDataBytesCount int) error {

	cellsCount := int(boxSize * boxSize)

	if box.Disabled {
		for i := 0; i < cellsCount; i++ {
			cell := &models.SudokuCellDTO{}
			setCellIndexes(cell, boxSize, i)
			box.Cells = append(box.Cells, cell)
		}

		return nil
	}

	if len(sudokuData) < startByteIndex+oneBoxDataBytesCount {
		return fmt.Errorf(
			"sudoku data does not have sufficient cells values information")
	}

	for i := 0; i < cellsCount; i++ {
		cell := &models.SudokuCellDTO{}
		setCellIndexes(cell, boxSize, i)

		valueByteIndex := startByteIndex + i
		value := int(sudokuData[valueByteIndex])

		if value > 0 {
			cell.Value = &value
		}

		box.Cells = append(box.Cells, cell)
	}

	return nil
}

// getBoxesEnableStateData creates a slice representing which of boxes encoded in binary data
// are enabled and which are not. It returns a slice where index on a flag is an index of a
// box in the binary data, an int indicating how many bytes of binary data was used to encode
// boxes enable/disable state and an error if occired
func getBoxesEnableStateData(sudokuData []byte, layout *models.SudokuLayoutDTO) ([]bool, int, error) {
	startingByteIndex := 5
	boxesCount := int(layout.Height * layout.Height)
	amountOfBytes := int(math.Ceil(float64(boxesCount) / 8))

	if len(sudokuData) < startingByteIndex+amountOfBytes+1 {
		return []bool{}, 0, fmt.Errorf(
			"sudoku data does not contain boxes enable/disable state information")
	}

	sudokuBoxesEnableDataBytes := sudokuData[startingByteIndex : startingByteIndex+amountOfBytes]

	result := []bool{}
	for i := 0; i < boxesCount; i++ {
		byteIndex := int(float64(i / 8))
		bitIndex := i % 8
		logicalAndOperationValue := byte(math.Pow(2, float64(7-bitIndex)))
		result = append(result, sudokuBoxesEnableDataBytes[byteIndex]&logicalAndOperationValue > 0)
	}

	return result, amountOfBytes, nil
}

// setBoxIndexes calculates and assigns row and column index to the box,
// based of its single dimension index
func setBoxIndexes(box *models.SudokuBoxDTO, layout *models.SudokuLayoutDTO, boxIndex int) {
	rowIndex := int8(math.Floor(float64(boxIndex) / float64(layout.Width)))
	columnIndex := int8(boxIndex) - (rowIndex * layout.Width)

	box.IndexRow = rowIndex
	box.IndexColumn = columnIndex
}

// setCellIndexes calculates and assigns row and column index to the cell,
// based of its single dimension index
func setCellIndexes(cell *models.SudokuCellDTO, boxSize int8, cellIndex int) {
	rowIndex := int8(math.Floor(float64(cellIndex) / float64(boxSize)))
	columnIndex := int8(cellIndex) - (rowIndex * boxSize)

	cell.IndexRowInBox = rowIndex
	cell.IndexColumnInBox = columnIndex
}

// getVersionNumber reads binary representation version from binary representation
func getVersionNumber(sudokuData []byte) (uint16, error) {
	if len(sudokuData) < 2 {
		return 0, errors.New("sudoku data does not contain version information")
	}

	version := binary.BigEndian.Uint16(sudokuData[:2])
	return version, nil
}

// getSudokuBoxSize reads sudoku box size from binary representation
func getSudokuBoxSize(sudokuData []byte) (int8, error) {
	if len(sudokuData) < 3 {
		return 0, errors.New("sudoku data does not contain box size information")
	}

	return int8(sudokuData[2]), nil
}

// getSudokuLayout reads sudoku layout configuration from binary representation
func getSudokuLayout(sudokuData []byte) (*models.SudokuLayoutDTO, error) {
	if len(sudokuData) < 5 {
		return nil, errors.New("sudoku data does not contain layout information")
	}

	layout := &models.SudokuLayoutDTO{
		Width:  int8(sudokuData[3]),
		Height: int8(sudokuData[4]),
	}

	return layout, nil
}
