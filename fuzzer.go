package main

import (
	"fmt"
	"math/rand"
)

// Fuzzer struct that holds a function to be fuzzed
type Fuzzer struct {
	testFunction func(string) (string, error) // The function to be fuzzed
}

// NewFuzzer creates a new fuzzer instance
func NewFuzzer(testFunction func(string) (string, error)) *Fuzzer {
	return &Fuzzer{
		testFunction: testFunction,
	}
}

// randomString generates a random string of a given length
func (f *Fuzzer) randomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	s := make([]byte, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// Fuzz generates random inputs, tests the function, and returns the result and any error
func (f *Fuzzer) Fuzz() (string, error) {
	input := f.randomString(rand.Intn(20) + 1) // Generate a random string of length 1 to 20
	result, err := f.testFunction(input)       // Call the function being fuzzed
	if err != nil {
		fmt.Printf("Error with input '%s': %v\n", input, err) // Print the error if one occurs
		return "", err
	}
	fmt.Printf("Success with input '%s': %s\n", input, result) // Print success
	return result, nil
}
