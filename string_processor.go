package main

import (
	"errors"
	"fmt"
)

// processString processes the input string and returns an error if the string is too long
func processString(input string) (string, error) {
	// Strings longer than 10 characters are invalid
	if len(input) > 10 {
		return "", errors.New("input string is too long")
	}

	// If the string is valid, return the processed result
	return fmt.Sprintf("Processed: %s", input), nil
}
