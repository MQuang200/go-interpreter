package main

import (
	"fmt"
	"os"

	"github.com/MQuang200/go-interpreter/src/scanner"
	t "github.com/MQuang200/go-interpreter/src/token"
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
    tokens, hadError := scanner.Scan(fileContents)
    for _, token := range(tokens) {
      if token.TokenType == t.EOF {
        fmt.Println(token.EOFString())
      } else {
        fmt.Println(token.String())
      }
    }

    if hadError {
      os.Exit(65)
    }
	}
}
