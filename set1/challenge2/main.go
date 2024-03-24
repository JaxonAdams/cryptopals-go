package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	// CLI args
	var first_buffer string
	var second_buffer string

	// Parse CLI args
	flag.StringVar(&first_buffer, "first-buffer", "", "Specify a hex value.")
	flag.StringVar(&second_buffer, "second-buffer", "", "Specify a hex value.")
	flag.Parse()

	// Print defaults and exit if input not provided
	if first_buffer == "" || second_buffer == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Decode both buffers
	first_bytes, err := hex.DecodeString(first_buffer)
	if err != nil {
		fmt.Println("Could not parse hex value", first_buffer)
	}

	second_bytes, err := hex.DecodeString(second_buffer)
	if err != nil {
		fmt.Println("Could not parse hex value", second_buffer)
	}

	// Check that both byte arrays are the same length
	if len(first_bytes) != len(second_bytes) {
		fmt.Printf("Buffers must be equal lengths.")
		os.Exit(1)
	}

	// Create a new set of bytes by XOR'ing the two provided buffers
	new_bytes := make([]byte, len(first_bytes))
	for i := range new_bytes {
		new_bytes[i] = first_bytes[i] ^ second_bytes[i]
	}

	// Encode and display the new bytes
	new_hex := hex.EncodeToString(new_bytes)
	fmt.Println(new_hex)
}
