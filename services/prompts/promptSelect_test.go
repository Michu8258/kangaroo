package prompts

import (
	"errors"
	"strings"
	"testing"

	"github.com/Michu8258/kangaroo/models"
	"github.com/Michu8258/kangaroo/testHelpers"
	tea "github.com/charmbracelet/bubbletea"
)

func TestPromptMakeSelectChoice_Success(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	testOptions := getTestOptions()

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return model, nil
		})

	result, err := prompter.PromptMakeSelectChoice("test title", testOptions, 1)

	if err != nil {
		t.Errorf("Unexpected error: '%s'", err)
	}

	if result.Value != testOptions[1].Value {
		t.Errorf("Unexpected result value. Expected %d, got %d.",
			testOptions[1].Value, result.Value)
	}
}

func TestPromptMakeSelectChoice_Error(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	testOptions := getTestOptions()

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return nil, errors.New("some error")
		})

	_, err := prompter.PromptMakeSelectChoice("test title", testOptions, 1)

	if err == nil {
		t.Errorf("An error was expected but none was returned from Select prompt")
	}
}

func TestInit_SelectPrompt(t *testing.T) {
	model := getPromptSelectModel(1)
	result := model.Init()
	if result != nil {
		t.Error("Select prompt init results in non nil command.")
	}
}

func TestUpdate_SelectPrompt(t *testing.T) {
	testCases := []struct {
		name                 string
		teaKeyMessage        tea.KeyMsg
		initialOptionIndex   int
		expectsResultCommand bool
		modelStateValidator  func(model promptSelect) bool
	}{
		{
			name:                 "escape",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyEsc},
			initialOptionIndex:   1,
			expectsResultCommand: true,
			modelStateValidator: func(model promptSelect) bool {
				return model.quit == true
			},
		},
		{
			name:                 "escape 2",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyEscape},
			initialOptionIndex:   1,
			expectsResultCommand: true,
			modelStateValidator: func(model promptSelect) bool {
				return model.quit == true
			},
		},
		{
			name:                 "enter",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyEnter},
			initialOptionIndex:   1,
			expectsResultCommand: true,
			modelStateValidator: func(model promptSelect) bool {
				return model.quit == true && model.activeChoice.Value == getTestOptions()[1].Value
			},
		},
		{
			name:                 "down",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyDown},
			initialOptionIndex:   1,
			expectsResultCommand: false,
			modelStateValidator: func(model promptSelect) bool {
				return model.quit == false && model.cursor == 2
			},
		},
		{
			name:                 "down with overflow",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyDown},
			initialOptionIndex:   2,
			expectsResultCommand: false,
			modelStateValidator: func(model promptSelect) bool {
				return model.quit == false && model.cursor == 0
			},
		},
		{
			name:                 "up",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyUp},
			initialOptionIndex:   1,
			expectsResultCommand: false,
			modelStateValidator: func(model promptSelect) bool {
				return model.quit == false && model.cursor == 0
			},
		},
		{
			name:                 "up with overflow",
			teaKeyMessage:        tea.KeyMsg{Type: tea.KeyUp},
			initialOptionIndex:   0,
			expectsResultCommand: false,
			modelStateValidator: func(model promptSelect) bool {
				return model.quit == false && model.cursor == 2
			},
		},
	}

	for _, testCase := range testCases {
		model := getPromptSelectModel(testCase.initialOptionIndex)
		resultModel, command := model.Update(testCase.teaKeyMessage)
		hasCommand := command != nil

		if testCase.expectsResultCommand != hasCommand {
			t.Errorf("%s: invalid returned command state", testCase.name)
		}

		if resultModel == nil {
			t.Errorf("%s: no model returned from update", testCase.name)
		}

		if !testCase.modelStateValidator(resultModel.(promptSelect)) {
			t.Errorf("%s: invalid model state", testCase.name)
		}
	}
}

func TestView_SelectPrompt_Render(t *testing.T) {
	model := getPromptSelectModel(1)
	expectedSubstrings := []string{
		"[✓]",
		"[ ]",
		model.title,
	}

	for _, option := range model.choices {
		expectedSubstrings = append(expectedSubstrings, option.Label)
	}

	viewString := model.View()

	for _, expectedSubsting := range expectedSubstrings {
		if !strings.Contains(viewString, expectedSubsting) {
			t.Errorf("Select prompt view string does not contain '%s' substring.",
				expectedSubsting)
		}
	}
}

func TestView_SelectPrompt_Quit(t *testing.T) {
	model := getPromptSelectModel(1)
	model.quit = true

	viewString := model.View()

	if viewString != "" {
		t.Errorf("Select prompt view string should be empty, but it is: '%s'", viewString)
	}
}

func getPromptSelectModel(initialOptionIndex int) promptSelect {
	options := getTestOptions()

	return promptSelect{
		cursor:       initialOptionIndex,
		title:        "test select title",
		activeChoice: options[initialOptionIndex],
		choices:      options,
		quit:         false,
	}
}

func getTestOptions() []models.PromptSelectOption {
	return []models.PromptSelectOption{
		{
			Label: "option 1",
			Value: 1,
		},
		{
			Label: "option 2",
			Value: 2,
		},
		{
			Label: "option 2",
			Value: 2,
		},
	}
}
