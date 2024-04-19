package dataPrinters

import (
	"errors"
	"strings"
	"testing"

	"github.com/Michu8258/kangaroo/testHelpers"
)

func TestPrintErrors(t *testing.T) {
	settings := testHelpers.GetTestSettings()
	testPrinter := testHelpers.NewTestPrinter()

	header := "Errors title"
	errs := []error{
		errors.New("some error 1"),
		errors.New("some error 2"),
		errors.New("some error 3"),
	}

	expectedPrints := []string{
		"Errors title",
		"Some error 1",
		"Some error 2",
		"Some error 3",
	}

	dataPrinter := GetNewDataPrinter(settings, testPrinter)

	dataPrinter.PrintErrors(header, errs...)

	for _, expectedPrint := range expectedPrints {
		if !strings.Contains(testPrinter.PrintedData, expectedPrint) {
			t.Errorf("Printout does not contain expected string: '%s", expectedPrint)
		}
	}
}
