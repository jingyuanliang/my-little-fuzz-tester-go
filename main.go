package main

import "fmt"

func main() {
	fmt.Println("Starting Go Fuzzer...")

	// Hardcoded input string for testing
	input := "HelloFuzz"

	// Run the string processing function with the input
	result, err := processString(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %v\n", result)
	}
}
