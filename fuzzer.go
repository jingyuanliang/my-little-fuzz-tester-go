package main

// Fuzzer struct that holds a function
type Fuzzer struct {
	testFunction func(string) (string, error) // The function to be fuzzed
}

// NewFuzzer creates a new fuzzer instance
func NewFuzzer(testFunction func(string) (string, error)) *Fuzzer {
	return &Fuzzer{
		testFunction: testFunction,
	}
}
