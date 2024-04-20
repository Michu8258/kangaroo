package prompts

import (
	"testing"

	"github.com/Michu8258/kangaroo/testHelpers"
	tea "github.com/charmbracelet/bubbletea"
)

func TestPromptGetLayoutSize_AlreadySelected(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	var initialLayoutSize int8 = 3

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return model, nil
		})

	result, err := prompter.PromptGetLayoutSize(&initialLayoutSize, "line")

	if result != 3 {
		t.Errorf("Expected layout size to be %d, but is %d.", 3, result)
	}

	if err != nil {
		t.Errorf("Did not expect any error, but got: '%s'", err)
	}
}

func TestPromptGetLayoutSize_PreselectOutOfRange(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()
	var initialLayoutSize int8 = settings.MaximumBoxSizeInclusive + 5

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return model, nil
		})

	result, err := prompter.PromptGetLayoutSize(&initialLayoutSize, "line")

	if result != 3 {
		t.Errorf("Expected layout size to be %d, but is %d.", 3, result)
	}

	if err != nil {
		t.Errorf("Did not expect any error, but got: '%s'", err)
	}
}

func TestPromptGetLayoutSize_NoPreselection(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()

	prompter := GetNewPrompter(settings, testPrinter,
		func(model tea.Model, opts ...tea.ProgramOption) (tea.Model, error) {
			return model, nil
		})

	result, err := prompter.PromptGetLayoutSize(nil, "line")

	if result != 3 {
		t.Errorf("Expected layout size to be %d, but is %d.", 3, result)
	}

	if err != nil {
		t.Errorf("Did not expect any error, but got: '%s'", err)
	}
}
