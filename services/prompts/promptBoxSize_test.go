package prompts

import (
	"testing"

	"github.com/Michu8258/kangaroo/testHelpers"
	tea "github.com/charmbracelet/bubbletea"
)

func TestPromptGetBoxSize_AlreadySelected(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	var initialBoxSize int8 = 3

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return model, nil
		})

	result, err := prompter.PromptGetBoxSize(&initialBoxSize)

	if result != 3 {
		t.Errorf("Expected box size to be %d, but is %d.", 3, result)
	}

	if err != nil {
		t.Errorf("Did not expect any error, but got: '%s'", err)
	}
}

func TestPromptGetBoxSize_PreselectOutOfRange(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	var initialBoxSize int8 = settings.MaximumBoxSizeInclusive + 5

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return model, nil
		})

	result, err := prompter.PromptGetBoxSize(&initialBoxSize)

	if result != 3 {
		t.Errorf("Expected box size to be %d, but is %d.", 3, result)
	}

	if err != nil {
		t.Errorf("Did not expect any error, but got: '%s'", err)
	}
}

func TestPromptGetBoxSize_NoPreselection(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return model, nil
		})

	result, err := prompter.PromptGetBoxSize(nil)

	if result != 3 {
		t.Errorf("Expected box size to be %d, but is %d.", 3, result)
	}

	if err != nil {
		t.Errorf("Did not expect any error, but got: '%s'", err)
	}
}
