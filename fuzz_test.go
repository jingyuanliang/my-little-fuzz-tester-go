package main

import (
	"fmt"
	"os"
	"testing"
)

// WriteTestReport writes the result of fuzz tests to a report file
func WriteTestReport(report string) error {
	// Create a report file
	file, err := os.Create("fuzz_test_report.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the report content to the file
	_, err = file.WriteString(report)
	return err
}

// TestFuzzer tests the processString function using the fuzzer
func TestFuzzer(t *testing.T) {
	// Initialize the fuzzer and prepare to log results
	fuzzer := NewFuzzer(processString)
	report := "Fuzz Test Report:\n"

	// Run the fuzzer multiple times and assert the results
	for i := 0; i < 10; i++ {
		result, err := fuzzer.Fuzz()

		if err != nil && err.Error() == "input string is too long" {
			// Log expected error without failing the test
			t.Logf("Test %d: Expected error - %v", i+1, err)
			report += fmt.Sprintf("Test %d: Error - %v\n", i+1, err)
		} else if err != nil {
			// Log unexpected error and fail the test
			t.Errorf("Test %d: Unexpected error - %v", i+1, err)
			report += fmt.Sprintf("Test %d: Unexpected error - %v\n", i+1, err)
		} else {
			// Log success
			t.Logf("Test %d: Success - %s", i+1, result)
			report += fmt.Sprintf("Test %d: Success - %s\n", i+1, result)
		}
	}

	// Write the report to a file
	err := WriteTestReport(report)
	if err != nil {
		t.Fatalf("Failed to write test report: %v", err)
	}
}
