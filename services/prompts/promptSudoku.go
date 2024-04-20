package prompts

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Michu8258/kangaroo/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sudokuValuesPrompt struct {
	sudokuDTO         *models.SudokuDTO
	settings          *models.Settings
	quit              bool
	charactersPerCell int
	currentBox        *models.SudokuBoxDTO
	currentCell       *models.SudokuCellDTO
}

// PromptSudokuValues wraps logic for prompting user for sudoku values
// input with respect to prior configuration (box size, layout sizes)
func (prompter *Prompter) PromptSudokuValues(sudokuDto *models.SudokuDTO) error {
	failError := fmt.Errorf("failed to get sudoku values from manual input")
	initialModel, err := buildSudokuValuesPromptModel(sudokuDto, prompter.Settings)
	if err != nil {
		prompter.TerminalPrinter.PrintError(err.Error())
		prompter.TerminalPrinter.PrintNewLine()
		return failError
	}

	model, err := prompter.TeaProgramRunner(initialModel)
	if err != nil {
		prompter.TerminalPrinter.PrintError(err.Error())
		prompter.TerminalPrinter.PrintNewLine()
		return failError
	}

	if _, ok := model.(sudokuValuesPrompt); ok {
		return nil
	}

	return failError
}

// Init iniitalizes tea model state
func (m sudokuValuesPrompt) Init() tea.Cmd {
	return nil
}

// Update updates tea model state
func (m sudokuValuesPrompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch messageType := msg.(type) {
	case tea.KeyMsg:
		switch messageType.String() {
		case "ctrl+c", "esc":
			m.quit = true
			return m, tea.Quit

		case "enter":
			m.quit = true
			return m, tea.Quit

		case "up", "k":
			goUpSudokuCell(&m)

		case "down", "j":
			goDownSudokuCell(&m)

		case "left", "h":
			goLeftSudokuCell(&m)

		case "right", "l":
			goRightSudokuCell(&m)

		case "0":
			appendValue(&m, 0)

		case "1":
			appendValue(&m, 1)

		case "2":
			appendValue(&m, 2)

		case "3":
			appendValue(&m, 3)

		case "4":
			appendValue(&m, 4)

		case "5":
			appendValue(&m, 5)

		case "6":
			appendValue(&m, 6)

		case "7":
			appendValue(&m, 7)

		case "8":
			appendValue(&m, 8)

		case "9":
			appendValue(&m, 9)

		case "delete":
			clearCurrentCellValue(&m)

		case "backspace":
			backspaceCurrentCellValue(&m)

		case "e":
			changeDisableStateOfCurrentBox(&m, true)

		case "d":
			changeDisableStateOfCurrentBox(&m, false)
		}

	}

	return m, nil
}

// view renders output based on model state
func (m sudokuValuesPrompt) View() string {
	if m.quit {
		return ""
	}

	builder := strings.Builder{}

	// iterate through all rows (boxes and cells)
	var boxRowIndex int8 = 0
	var maxBoxRowIndex int8 = m.sudokuDTO.Layout.Height - 1
	var cellRowIndex int8 = 0
	var maxCellRowIndex int8 = m.sudokuDTO.BoxSize - 1

	for boxRowIndex = 0; boxRowIndex <= maxBoxRowIndex; boxRowIndex++ {
		for cellRowIndex = 0; cellRowIndex <= maxCellRowIndex; cellRowIndex++ {
			if boxRowIndex == 0 && cellRowIndex == 0 {
				printTopBorderLine(&builder, &m)
			}

			printSudokuValuesLine(&builder, &m, boxRowIndex, cellRowIndex)

			if cellRowIndex < m.sudokuDTO.BoxSize-1 {
				printMidCellsLine(&builder, &m)
			}
		}

		if boxRowIndex < m.sudokuDTO.Layout.Height-1 {
			printMidBoxesLine(&builder, &m)
		}
	}

	printBottomBorderLine(&builder, &m)
	builder.WriteString("\n")

	printSudokuControls(&builder)

	return builder.String()
}

// printTopBorderLine prints top border line of a sudoku puzzle
func printTopBorderLine(builder *strings.Builder, model *sudokuValuesPrompt) {
	printSudokuHorizontalBoxLine(builder, model, "╔", "═", "╗", "╦")
}

// printMidCellsLine prints line of a sudoku puzzle that appears between cells
func printMidCellsLine(builder *strings.Builder, model *sudokuValuesPrompt) {
	printSudokuHorizontalBoxLine(builder, model, "║", "─", "║", "║")
}

// printMidBoxesLine prints line of a sudoku puzzle that appears between boxes
func printMidBoxesLine(builder *strings.Builder, model *sudokuValuesPrompt) {
	printSudokuHorizontalBoxLine(builder, model, "║", "═", "║", "╬")
}

// printBottomBorderLine prints bottom border line of a sudoku puzzle
func printBottomBorderLine(builder *strings.Builder, model *sudokuValuesPrompt) {
	printSudokuHorizontalBoxLine(builder, model, "╚", "═", "╝", "╩")
}

// printSudokuHorizontalBoxLine prints sudoku vertival line - between boxes
func printSudokuHorizontalBoxLine(builder *strings.Builder, model *sudokuValuesPrompt,
	startSign string, middleSign string, endSign string, columnCrossSign string) {

	paddingLength := int(model.settings.SudokuPrintoutValuePaddingLength)
	charsPerCell := model.charactersPerCell + 2*paddingLength

	builder.WriteString(models.TerminalStyles.BorderStyle.Render(startSign))

	var boxColumnIndex int8 = 0
	var cellColumnIndex int8 = 0
	for boxColumnIndex = 0; boxColumnIndex < model.sudokuDTO.Layout.Width; boxColumnIndex++ {
		for cellColumnIndex = 0; cellColumnIndex < model.sudokuDTO.BoxSize; cellColumnIndex++ {
			if cellColumnIndex > 0 {
				builder.WriteString(models.TerminalStyles.BorderStyle.Render(middleSign))
			}
			for characterIndex := 0; characterIndex < charsPerCell; characterIndex++ {
				builder.WriteString(models.TerminalStyles.BorderStyle.Render(middleSign))
			}
		}

		if boxColumnIndex < model.sudokuDTO.Layout.Width-1 {
			builder.WriteString(models.TerminalStyles.BorderStyle.Render(columnCrossSign))
		}
	}

	builder.WriteString(models.TerminalStyles.BorderStyle.Render(endSign))
	builder.WriteString(models.TerminalStyles.BorderStyle.Render("\n"))
}

// printSudokuValuesLine prints single row of sudoku values
func printSudokuValuesLine(builder *strings.Builder, model *sudokuValuesPrompt,
	boxRowIndex int8, cellRowIndex int8) {

	builder.WriteString(models.TerminalStyles.BorderStyle.Render("║"))

	var boxColumnIndex int8 = 0
	var cellColumnIndex int8 = 0

	for boxColumnIndex = 0; boxColumnIndex < model.sudokuDTO.Layout.Width; boxColumnIndex++ {
		sudokuBox := model.sudokuDTO.Boxes.FirstOrDefault(nil, func(box *models.SudokuBoxDTO) bool {
			return box.IndexColumn == boxColumnIndex && box.IndexRow == boxRowIndex
		})

		for cellColumnIndex = 0; cellColumnIndex < model.sudokuDTO.BoxSize; cellColumnIndex++ {
			if cellColumnIndex > 0 {
				builder.WriteString(models.TerminalStyles.BorderStyle.Render("│"))
			}

			sudokuCell := sudokuBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
				return cell.IndexColumnInBox == cellColumnIndex && cell.IndexRowInBox == cellRowIndex
			})

			printSudokuCell(builder, model, sudokuBox, sudokuCell)
		}

		if boxColumnIndex < model.sudokuDTO.Layout.Width-1 {
			builder.WriteString(models.TerminalStyles.BorderStyle.Render("║"))
		}
	}

	builder.WriteString(models.TerminalStyles.BorderStyle.Render("║"))
	builder.WriteString("\n")
}

// printSudokuCell prints single sudoku cell balue
func printSudokuCell(builder *strings.Builder, model *sudokuValuesPrompt,
	box *models.SudokuBoxDTO, cell *models.SudokuCellDTO) {

	var style lipgloss.Style

	isActiveCell := model.currentBox != nil &&
		model.currentCell != nil &&
		model.currentBox.IndexRow == box.IndexRow &&
		model.currentBox.IndexColumn == box.IndexColumn &&
		model.currentCell.IndexRowInBox == cell.IndexRowInBox &&
		model.currentCell.IndexColumnInBox == cell.IndexColumnInBox

	if isActiveCell {
		style = models.TerminalStyles.SuccessStyle
	} else if box.Disabled {
		style = models.TerminalStyles.BorderStyle
	} else if cell.Value == nil {
		style = models.TerminalStyles.BorderStyle
	} else {
		style = models.TerminalStyles.PrimaryStyle
	}

	printValuePadding(builder, model, style)
	if cell.Value == nil {
		for characterIndex := 0; characterIndex < model.charactersPerCell; characterIndex++ {
			builder.WriteString(style.Render("_"))
		}
	} else {
		stringValue := strconv.Itoa(*cell.Value)
		if len(stringValue) > model.charactersPerCell {
			stringValue = stringValue[:model.charactersPerCell]
		}

		switch model.charactersPerCell {
		case 1:
			builder.WriteString(style.Render(stringValue))
		case 2:
			builder.WriteString(style.Render(fmt.Sprintf("%-2s", stringValue)))
		case 3:
			builder.WriteString(style.Render(fmt.Sprintf("%-3s", stringValue)))
		default:
			builder.WriteString(style.Render("x"))
		}
	}
	printValuePadding(builder, model, style)
}

// printValuePadding prints padding before and after a sudoku value (horizontal padding)
func printValuePadding(builder *strings.Builder, model *sudokuValuesPrompt, style lipgloss.Style) {
	var paddingIndex int8 = 0
	if model.settings.SudokuPrintoutValuePaddingLength >= 1 {
		for paddingIndex = 0; paddingIndex < model.settings.SudokuPrintoutValuePaddingLength; paddingIndex++ {
			builder.WriteString(style.Render(" "))
		}
	}
}

// printSudokuControls print controls of sudoku editor
func printSudokuControls(builder *strings.Builder) {
	builder.WriteString(models.TerminalStyles.PrimaryStyle.Render("Controls:"))
	builder.WriteString("\n")

	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("Move up: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("↑/k"))
	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("\tMove down: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("↓/j"))
	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("\tMove left: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("←/h"))
	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("\tMove right: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("→/l"))
	builder.WriteString("\n")

	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("InsertValue: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("numbers 0-9"))
	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("\tDelete value: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("delete"))
	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("\tBackspace: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("backspace"))
	builder.WriteString("\n")

	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("Finish and confirm: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("enter"))
	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("\tCancel: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("esc/ctrl+c"))
	builder.WriteString(models.TerminalStyles.DefaultStyle.Render("\tEnable/disable box: "))
	builder.WriteString(models.TerminalStyles.SuccessStyle.Render("e/d"))
	builder.WriteString("\n")
}

// goUpSudokuCell navigates to the cell on the top from current one
func goUpSudokuCell(model *sudokuValuesPrompt) {
	maxCellRowIndex := model.sudokuDTO.BoxSize - 1

	// first cell in column - cannot go further top
	if model.currentBox.IndexRow <= 0 &&
		model.currentCell.IndexRowInBox <= 0 {
		return
	}

	// not first cell in the box
	if model.currentCell.IndexRowInBox > 0 {
		newCell := model.currentBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
			return cell.IndexColumnInBox == model.currentCell.IndexColumnInBox &&
				cell.IndexRowInBox == model.currentCell.IndexRowInBox-1
		})

		if newCell == nil {
			return
		}

		model.currentCell = newCell
		return
	}

	// first cell in the box
	newBox := model.sudokuDTO.Boxes.FirstOrDefault(nil, func(box *models.SudokuBoxDTO) bool {
		return box.IndexColumn == model.currentBox.IndexColumn && box.IndexRow == model.currentBox.IndexRow-1
	})

	if newBox == nil {
		return
	}

	newCell := newBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
		return cell.IndexColumnInBox == model.currentCell.IndexColumnInBox && cell.IndexRowInBox == maxCellRowIndex
	})

	if newCell == nil {
		return
	}

	model.currentBox = newBox
	model.currentCell = newCell
}

// goDownSudokuCell navigates to the cell on the bottom from current one
func goDownSudokuCell(model *sudokuValuesPrompt) {
	maxBoxRowIndex := model.sudokuDTO.Layout.Width - 1
	maxCellRowIndex := model.sudokuDTO.BoxSize - 1

	// last cell in column - cannot go further down
	if model.currentBox.IndexRow >= maxBoxRowIndex &&
		model.currentCell.IndexRowInBox >= maxCellRowIndex {
		return
	}

	// not last cell in the box
	if model.currentCell.IndexRowInBox < maxCellRowIndex {
		newCell := model.currentBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
			return cell.IndexColumnInBox == model.currentCell.IndexColumnInBox &&
				cell.IndexRowInBox == model.currentCell.IndexRowInBox+1
		})

		if newCell == nil {
			return
		}

		model.currentCell = newCell
		return
	}

	//last cell in the box
	newBox := model.sudokuDTO.Boxes.FirstOrDefault(nil, func(box *models.SudokuBoxDTO) bool {
		return box.IndexColumn == model.currentBox.IndexColumn && box.IndexRow == model.currentBox.IndexRow+1
	})

	if newBox == nil {
		return
	}

	newCell := newBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
		return cell.IndexColumnInBox == model.currentCell.IndexColumnInBox && cell.IndexRowInBox == 0
	})

	if newCell == nil {
		return
	}

	model.currentBox = newBox
	model.currentCell = newCell
}

// goLeftSudokuCell navigates to the cell on the left from current one
func goLeftSudokuCell(model *sudokuValuesPrompt) {
	maxCellColumnIndex := model.sudokuDTO.BoxSize - 1

	// first cell in row - cannot go further left
	if model.currentBox.IndexColumn <= 0 &&
		model.currentCell.IndexColumnInBox <= 0 {
		return
	}

	// not first cell in the box
	if model.currentCell.IndexColumnInBox > 0 {
		newCell := model.currentBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
			return cell.IndexRowInBox == model.currentCell.IndexRowInBox &&
				cell.IndexColumnInBox == model.currentCell.IndexColumnInBox-1
		})

		if newCell == nil {
			return
		}

		model.currentCell = newCell
		return
	}

	// first cell in the box
	newBox := model.sudokuDTO.Boxes.FirstOrDefault(nil, func(box *models.SudokuBoxDTO) bool {
		return box.IndexRow == model.currentBox.IndexRow && box.IndexColumn == model.currentBox.IndexColumn-1
	})

	if newBox == nil {
		return
	}

	newCell := newBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
		return cell.IndexRowInBox == model.currentCell.IndexRowInBox && cell.IndexColumnInBox == maxCellColumnIndex
	})

	if newCell == nil {
		return
	}

	model.currentBox = newBox
	model.currentCell = newCell
}

// goRightSudokuCell navigates to the cell on the right from current one
func goRightSudokuCell(model *sudokuValuesPrompt) {
	maxBoxColumnIndex := model.sudokuDTO.Layout.Width - 1
	maxCellColumnIndex := model.sudokuDTO.BoxSize - 1

	// last cell in row - cannot go further right
	if model.currentBox.IndexColumn >= maxBoxColumnIndex &&
		model.currentCell.IndexColumnInBox >= maxCellColumnIndex {
		return
	}

	// not last cell in the box
	if model.currentCell.IndexColumnInBox < maxCellColumnIndex {
		newCell := model.currentBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
			return cell.IndexRowInBox == model.currentCell.IndexRowInBox &&
				cell.IndexColumnInBox == model.currentCell.IndexColumnInBox+1
		})

		if newCell == nil {
			return
		}

		model.currentCell = newCell
		return
	}

	//last cell in the box
	newBox := model.sudokuDTO.Boxes.FirstOrDefault(nil, func(box *models.SudokuBoxDTO) bool {
		return box.IndexRow == model.currentBox.IndexRow && box.IndexColumn == model.currentBox.IndexColumn+1
	})

	if newBox == nil {
		return
	}

	newCell := newBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
		return cell.IndexRowInBox == model.currentCell.IndexRowInBox && cell.IndexColumnInBox == 0
	})

	if newCell == nil {
		return
	}

	model.currentBox = newBox
	model.currentCell = newCell
}

// appendValue appends another char to value of the cell (string based)
func appendValue(model *sudokuValuesPrompt, value int) {
	if model.currentCell.Value == nil && value < 1 {
		return
	}

	if model.currentCell.Value == nil {
		model.currentCell.Value = &value
		return
	}

	currentLength := len(strconv.Itoa(*model.currentCell.Value))
	if currentLength >= model.charactersPerCell {
		return
	}

	newValue, err := strconv.Atoi(strconv.Itoa(*model.currentCell.Value) + strconv.Itoa(value))
	if err != nil {
		return
	}

	model.currentCell.Value = &newValue
}

// clearCurrentCellValue removes value from current cell
func clearCurrentCellValue(model *sudokuValuesPrompt) {
	model.currentCell.Value = nil
}

// backspaceCurrentCellValue removes last digit from current cell
func backspaceCurrentCellValue(model *sudokuValuesPrompt) {
	if model.currentCell.Value == nil {
		return
	}

	stringValue := strconv.Itoa(*model.currentCell.Value)
	currentLength := len(stringValue)
	if currentLength <= 1 {
		model.currentCell.Value = nil
		return
	}

	newStringValue := stringValue[:currentLength-1]
	newValue, err := strconv.Atoi(newStringValue)
	if err != nil {
		return
	}

	model.currentCell.Value = &newValue
}

// changeDisableStateOfCurrentBox enables or disables current box
func changeDisableStateOfCurrentBox(model *sudokuValuesPrompt, newEnabled bool) {
	if model.currentBox == nil {
		return
	}

	newDisabled := !newEnabled
	if model.currentBox.Disabled != newDisabled {
		model.currentBox.Disabled = newDisabled
	}
}

// buildSudokuValuesPromptModel builds a model for sudoku values input (prompt)
func buildSudokuValuesPromptModel(sudokuDto *models.SudokuDTO, settings *models.Settings) (
	*sudokuValuesPrompt, error) {

	firstBox := sudokuDto.Boxes.FirstOrDefault(nil, func(box *models.SudokuBoxDTO) bool {
		return box.IndexRow == 0 && box.IndexColumn == 0
	})

	if firstBox == nil {
		return nil, fmt.Errorf("failed to find sudoku box with indexex (row: 0, column: 0)")
	}

	firstCell := firstBox.Cells.FirstOrDefault(nil, func(cell *models.SudokuCellDTO) bool {
		return cell.IndexRowInBox == 0 && cell.IndexColumnInBox == 0
	})

	if firstCell == nil {
		return nil, fmt.Errorf(
			"failed to find sudoku cell with indexex (row: 0, column: 0) within first box")
	}

	return &sudokuValuesPrompt{
		sudokuDTO:         sudokuDto,
		settings:          settings,
		quit:              false,
		charactersPerCell: len(strconv.Itoa(int(sudokuDto.BoxSize * sudokuDto.BoxSize))),
		currentBox:        firstBox,
		currentCell:       firstCell,
	}, nil
}
