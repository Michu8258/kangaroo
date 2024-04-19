package main

import "testing"

func TestMain(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("Running main should not cause panic.")
		}
	}()

	run([]string{""})
}
