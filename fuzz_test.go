package main

import "testing"

// TestFuzzer tests the processString function using the fuzzer
func TestFuzzer(t *testing.T) {
	// Initialize the fuzzer
	fuzzer := NewFuzzer(processString)

	// Run the fuzzer multiple times and assert the results
	for i := 0; i < 10; i++ {
		result, err := fuzzer.Fuzz()

		// Check if the error message is as expected for long strings
		if err != nil && err.Error() != "input string is too long" {
			t.Errorf("Unexpected error for input: %v", err)
		}

		// Check that non-error results are non-empty
		if err == nil && result == "" {
			t.Errorf("Expected non-empty result, but got empty string")
		}
	}
}
