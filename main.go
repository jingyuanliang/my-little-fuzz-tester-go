package main

import (
	"fmt"
	"log"

	"github.com/google/go-tpm/tpm2"
	"github.com/google/go-tpm/tpm2/transport/simulator"
)

func main() {
	// Open the TPM simulator
	tpm, err := simulator.OpenSimulator()
	if err != nil {
		log.Fatalf("Could not connect to TPM simulator: %v", err)
	}
	defer tpm.Close()

	// Request 16 random bytes from the TPM
	grc := tpm2.GetRandom{
		BytesRequested: 16,
	}

	// Execute the command and retrieve the bytes
	randomBytes, err := grc.Execute(tpm)
	if err != nil {
		log.Fatalf("GetRandom failed: %v", err)
	}

	// Print the random bytes as a hex string
	fmt.Printf("Random Bytes: %x\n", randomBytes)
}
