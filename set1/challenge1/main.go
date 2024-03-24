package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	// CLI args
	var input string

	// Parse CLI args
	flag.StringVar(&input, "input", "", "Specify a hex value.")
	flag.Parse()

	// Print defaults and exit if input not provided
	if input == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Attempt to decode the given hex string
	as_bytes, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("Invalid hex value provided.")
		os.Exit(1)
	}

	// Encode the decoded bytes with Base64
	as_b64 := base64.StdEncoding.EncodeToString(as_bytes)

	// Print the newly encoded value
	fmt.Println("Encoded value:", as_b64)
}
