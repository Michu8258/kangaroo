package prompts

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
	tea "github.com/charmbracelet/bubbletea"
)

func TestPromptSudokuValues_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	sudokuDto := testHelpers.GetTestSudokuDto()

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			mdl := model.(*sudokuValuesPrompt)
			return *mdl, nil
		})

	err := prompter.PromptSudokuValues(sudokuDto)

	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
	}
}

func TestPromptSudokuValues_Error(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	sudokuDto := testHelpers.GetTestSudokuDto()

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return nil, errors.New("some error")
		})

	err := prompter.PromptSudokuValues(sudokuDto)

	if err == nil {
		t.Errorf("An error was expected but none was returned from Sudoku prompt")
	}
}

func TestInit_SudokuPrompt(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	model, _ := buildSudokuValuesPromptModel(testHelpers.GetTestSudokuDto(), settings)
	result := model.Init()
	if result != nil {
		t.Error("Select prompt init results in non nil command.")
	}
}

func TestUpdate_SudokuPrompt(t *testing.T) {
	testCases := []struct {
		name                 string
		teaKeyMessage        tea.KeyMsg
		expectsResultCommand bool
		modelStateValidator  func(model sudokuValuesPrompt) bool
	}{
		{
			name:                 "escape",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyEsc},
			expectsResultCommand: true,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == true
			},
		},
		{
			name:                 "escape 2",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyEscape},
			expectsResultCommand: true,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == true
			},
		},
		{
			name:                 "enter",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyEnter},
			expectsResultCommand: true,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == true
			},
		},
		{
			name:                 "up - k",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "up",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyUp},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "up - k",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "down",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyDown},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "down - j",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "left",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyLeft},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "left - h",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'h'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "right",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRight},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "right - l",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'l'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false
			},
		},
		{
			name:                 "zero",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'0'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && model.currentCell.Value == nil
			},
		},
		{
			name:                 "one",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 1
			},
		},
		{
			name:                 "two",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'2'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 2
			},
		},
		{
			name:                 "three",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'3'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 3
			},
		},
		{
			name:                 "four",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'4'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 4
			},
		},
		{
			name:                 "five",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'5'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 5
			},
		},
		{
			name:                 "six",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'6'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 6
			},
		},
		{
			name:                 "seven",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'7'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 7
			},
		},
		{
			name:                 "eight",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'8'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 8
			},
		},
		{
			name:                 "nine",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'9'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && *model.currentCell.Value == 9
			},
		},
		{
			name:                 "delete",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyDelete},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && model.currentCell.Value == nil
			},
		},
		{
			name:                 "backspace",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyBackspace},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && model.currentCell.Value == nil
			},
		},
		{
			name:                 "enable box",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'e'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && model.currentBox.Disabled == false
			},
		},
		{
			name:                 "disable box",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'d'}},
			expectsResultCommand: false,
			modelStateValidator: func(model sudokuValuesPrompt) bool {
				return model.quit == false && model.currentBox.Disabled == true
			},
		},
	}

	for _, testCase := range testCases {
		settings := testHelpers.GetTestSettings()
		model, _ := buildSudokuValuesPromptModel(testHelpers.GetTestSudokuDto(), settings)
		resultModel, command := model.Update(testCase.teaKeyMessage)
		hasCommand := command != nil

		if testCase.expectsResultCommand != hasCommand {
			t.Errorf("%s: invalid returned command state", testCase.name)
		}

		if resultModel == nil {
			t.Errorf("%s: no model returned from update", testCase.name)
		}

		if !testCase.modelStateValidator(resultModel.(sudokuValuesPrompt)) {
			t.Errorf("%s: invalid model state", testCase.name)
		}
	}
}

func TestView_SudokuPrompt_Render(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	sudokuDto := getTestSudokDto(t)
	model, _ := buildSudokuValuesPromptModel(sudokuDto, settings)

	expectedSubstrings := []string{
		"╔═══════════╦═══════════╦═══════════╗",
		"║───────────║───────────║───────────║",
		"║═══════════╬═══════════╬═══════════║",
		"╚═══════════╩═══════════╩═══════════╝",
		"║ _ │ _ │ _ ║ _ │ _ │ 4 ║ 3 │ 1 │ _ ║",
		"║ 3 │ _ │ 9 ║ 2 │ 7 │ _ ║ 5 │ 6 │ _ ║",
		"║ 4 │ 1 │ _ ║ _ │ 5 │ _ ║ _ │ _ │ _ ║",
		"║ 9 │ _ │ _ ║ 5 │ 2 │ 7 ║ _ │ _ │ 1 ║",
		"║ 5 │ _ │ _ ║ _ │ _ │ 6 ║ _ │ _ │ _ ║",
		"║ 1 │ 7 │ _ ║ _ │ _ │ _ ║ 8 │ 5 │ 2 ║",
		"║ _ │ 3 │ 2 ║ 7 │ _ │ _ ║ 1 │ 4 │ 6 ║",
		"║ 6 │ _ │ _ ║ 3 │ 4 │ 8 ║ _ │ _ │ _ ║",
		"║ _ │ _ │ 4 ║ _ │ _ │ _ ║ _ │ _ │ 3 ║",
		"Controls:",
		"Move up: ↑/k",
		"Move down: ↓/j",
		"Move left: ←/h",
		"Move right: →/l",
		"InsertValue: numbers 0-9",
		"Delete value: delete",
		"Backspace: backspace",
		"Finish and confirm: enter",
		"Cancel: esc/ctrl+c",
		"Enable/disable box: e/d",
	}

	viewString := model.View()

	for _, expectedSubsting := range expectedSubstrings {
		if !strings.Contains(viewString, expectedSubsting) {
			t.Errorf("Sudoku prompt view string does not contain '%s' substring.",
				expectedSubsting)
		}
	}
}

func TestView_SudokuPrompt_Quit(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	model, _ := buildSudokuValuesPromptModel(testHelpers.GetTestSudokuDto(), settings)
	model.quit = true

	viewString := model.View()

	if viewString != "" {
		t.Errorf("Sudoku prompt view string should be empty, but it is: '%s'", viewString)
	}
}

func TestGoUpSudokuCell(t *testing.T) {
	// these indexes are indexes in single dimension array
	// it goes left to right, then top to bottom
	testCases := []sudokuMovementTestCaseData{
		{
			name:                    "At the top",
			initialActiveBoxIndex:   0,
			initialActiveCellIndex:  0,
			expectedActualBoxIndex:  0,
			expectedActualCellIndex: 0,
		},
		{
			name:                    "Middle of the top box",
			initialActiveBoxIndex:   0,
			initialActiveCellIndex:  3,
			expectedActualBoxIndex:  0,
			expectedActualCellIndex: 0,
		},
		{
			name:                    "Top of second box (in column)",
			initialActiveBoxIndex:   3,
			initialActiveCellIndex:  0,
			expectedActualBoxIndex:  0,
			expectedActualCellIndex: 6,
		},
	}

	testMovements(t, "Movement UP", testCases, goUpSudokuCell)
}

func TestGoDownSudokuCell(t *testing.T) {
	// these indexes are indexes in single dimension array
	// it goes left to right, then top to bottom
	testCases := []sudokuMovementTestCaseData{
		{
			name:                    "At the bottom",
			initialActiveBoxIndex:   6,
			initialActiveCellIndex:  6,
			expectedActualBoxIndex:  6,
			expectedActualCellIndex: 6,
		},
		{
			name:                    "Middle of the bottom box",
			initialActiveBoxIndex:   6,
			initialActiveCellIndex:  3,
			expectedActualBoxIndex:  6,
			expectedActualCellIndex: 6,
		},
		{
			name:                    "Bottom of second box (in column)",
			initialActiveBoxIndex:   3,
			initialActiveCellIndex:  6,
			expectedActualBoxIndex:  6,
			expectedActualCellIndex: 0,
		},
	}

	testMovements(t, "Movement Bottom", testCases, goDownSudokuCell)
}

func TestGoLeftSudokuCell(t *testing.T) {
	// these indexes are indexes in single dimension array
	// it goes left to right, then top to bottom
	testCases := []sudokuMovementTestCaseData{
		{
			name:                    "At the left",
			initialActiveBoxIndex:   0,
			initialActiveCellIndex:  0,
			expectedActualBoxIndex:  0,
			expectedActualCellIndex: 0,
		},
		{
			name:                    "Middle of the left box",
			initialActiveBoxIndex:   0,
			initialActiveCellIndex:  1,
			expectedActualBoxIndex:  0,
			expectedActualCellIndex: 0,
		},
		{
			name:                    "left of second box (in row)",
			initialActiveBoxIndex:   1,
			initialActiveCellIndex:  0,
			expectedActualBoxIndex:  0,
			expectedActualCellIndex: 2,
		},
	}

	testMovements(t, "Movement Left", testCases, goLeftSudokuCell)
}

func TestGoRightSudokuCell(t *testing.T) {
	// these indexes are indexes in single dimension array
	// it goes left to right, then top to bottom
	testCases := []sudokuMovementTestCaseData{
		{
			name:                    "At the right",
			initialActiveBoxIndex:   2,
			initialActiveCellIndex:  2,
			expectedActualBoxIndex:  2,
			expectedActualCellIndex: 2,
		},
		{
			name:                    "Middle of the right box",
			initialActiveBoxIndex:   2,
			initialActiveCellIndex:  1,
			expectedActualBoxIndex:  2,
			expectedActualCellIndex: 2,
		},
		{
			name:                    "left of second box (in row)",
			initialActiveBoxIndex:   1,
			initialActiveCellIndex:  2,
			expectedActualBoxIndex:  2,
			expectedActualCellIndex: 0,
		},
	}

	testMovements(t, "Movement Right", testCases, goRightSudokuCell)
}

func TestAppendValue(t *testing.T) {
	testCases := []struct {
		initialCellValue  *int
		inputValue        int
		expectedCellValue *int
	}{
		{
			initialCellValue:  nil,
			inputValue:        5,
			expectedCellValue: getIntPointer(5),
		},
		{
			initialCellValue:  nil,
			inputValue:        0,
			expectedCellValue: nil,
		},
		{
			initialCellValue:  getIntPointer(1),
			inputValue:        0,
			expectedCellValue: getIntPointer(10),
		},
		{
			initialCellValue:  getIntPointer(11),
			inputValue:        1,
			expectedCellValue: getIntPointer(11),
		},
	}

	for testIndex, testCase := range testCases {
		sudokuDto := getTestSudokDto(t)
		settings := testHelpers.GetTestSettings()
		model, _ := buildSudokuValuesPromptModel(sudokuDto, settings)
		model.charactersPerCell = 2
		model.currentCell = &models.SudokuCellDTO{
			Value: testCase.initialCellValue,
		}

		appendValue(model, testCase.inputValue)

		if model.currentCell.Value == nil && testCase.expectedCellValue == nil {
			continue
		}

		if model.currentCell.Value == nil {
			t.Errorf("%d: cell calue is nil, expected value: %d",
				testIndex, *testCase.expectedCellValue)
		}

		if *model.currentCell.Value != *testCase.expectedCellValue {
			t.Errorf("%d: wrong cell value, got %d, expected value: %d",
				testIndex, *model.currentCell.Value, *testCase.expectedCellValue)
		}
	}
}

func TestBackspaceCurrentCellValue(t *testing.T) {
	testCases := []struct {
		initialCellValue  *int
		inputValue        int
		expectedCellValue *int
	}{
		{
			initialCellValue:  getIntPointer(5),
			expectedCellValue: nil,
		},
		{
			initialCellValue:  nil,
			expectedCellValue: nil,
		},
		{
			initialCellValue:  getIntPointer(11),
			expectedCellValue: getIntPointer(1),
		},
	}

	for testIndex, testCase := range testCases {
		sudokuDto := getTestSudokDto(t)
		settings := testHelpers.GetTestSettings()
		model, _ := buildSudokuValuesPromptModel(sudokuDto, settings)
		model.charactersPerCell = 10
		model.currentCell = &models.SudokuCellDTO{
			Value: testCase.initialCellValue,
		}

		backspaceCurrentCellValue(model)

		if model.currentCell.Value == nil && testCase.expectedCellValue == nil {
			continue
		}

		if model.currentCell.Value == nil {
			t.Errorf("%d: cell calue is nil, expected value: %d",
				testIndex, *testCase.expectedCellValue)
		}

		if *model.currentCell.Value != *testCase.expectedCellValue {
			t.Errorf("%d: wrong cell value, got %d, expected value: %d",
				testIndex, *model.currentCell.Value, *testCase.expectedCellValue)
		}
	}
}

type sudokuMovementTestCaseData struct {
	name                    string
	initialActiveBoxIndex   int
	initialActiveCellIndex  int
	expectedActualBoxIndex  int
	expectedActualCellIndex int
}

func testMovements(t *testing.T, testsSuiteName string, testCases []sudokuMovementTestCaseData,
	movementFunc func(*sudokuValuesPrompt)) {
	for _, testCase := range testCases {
		sudokuDto := getTestSudokDto(t)
		settings := testHelpers.GetTestSettings()
		model, _ := buildSudokuValuesPromptModel(sudokuDto, settings)
		model.currentBox = model.sudokuDTO.Boxes[testCase.initialActiveBoxIndex]
		model.currentCell = model.currentBox.Cells[testCase.initialActiveCellIndex]

		movementFunc(model)

		if model.currentBox != model.sudokuDTO.Boxes[testCase.expectedActualBoxIndex] {
			t.Errorf("%s - %s: Unexpected box is set as current when going up in Sudoku",
				testsSuiteName, testCase.name)
		}

		if model.currentCell != model.currentBox.Cells[testCase.expectedActualCellIndex] {
			t.Errorf("%s - %s: Unexpected cell is set as current when going up in Sudoku",
				testsSuiteName, testCase.name)
		}
	}
}

func getTestSudokDto(t *testing.T) *models.SudokuDTO {
	testFilePath := "../../testConfigs/simple1.json"
	bytes, err := os.ReadFile(testFilePath)
	if err != nil {
		t.Errorf("failed to read sudoku test config from file '%s'", testFilePath)
	}
	sudoku := models.Sudoku{}
	json.Unmarshal(bytes, &sudoku)
	return sudoku.ToSudokuDto()
}

func getIntPointer(value int) *int {
	return &value
}
