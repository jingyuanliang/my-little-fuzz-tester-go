package main

import (
	"fmt"
	"log"

	"github.com/google/go-tpm/tpm2"
	"github.com/google/go-tpm/tpm2/transport" // Import the Windows transport package
)

func main() {
	fmt.Println("This represents the main program.")

	// Open the TPM device using the Windows transport
	tpm, err := transport.OpenTPM() // Use the appropriate OpenTPM function for Windows
	if err != nil {
		log.Fatalf("Failed to open TPM device: %v", err)
	}
	defer tpm.Close() // Ensure the TPM is closed properly when done

	// Define the number of bytes to request
	const bytesRequested = 16

	// Create a GetRandom command
	getRandomCmd := tpm2.GetRandom{BytesRequested: bytesRequested}

	fmt.Printf("GetRandom command: %+v\n", getRandomCmd)
}
