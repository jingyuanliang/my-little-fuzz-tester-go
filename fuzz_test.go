package main

import "testing"

// TestProcessString tests the processString function with the hardcoded input
func TestProcessString(t *testing.T) {
	input := "HelloFuzz"

	result, err := processString(input)
	if err != nil {
		t.Errorf("Unexpected error for input '%s': %v", input, err)
	}

	expected := "Processed: HelloFuzz"
	if result != expected {
		t.Errorf("For input '%s', expected '%s', but got '%s'", input, expected, result)
	}
}
