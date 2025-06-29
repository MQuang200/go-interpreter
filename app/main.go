package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/MQuang200/go-interpreter/src/scanner"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
  fileContents = bytes.TrimSpace(fileContents)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
    tokens := scanner.Scan(fileContents)
    for _, token := range(tokens) {
      fmt.Println(token.String())
    }
	}
}
