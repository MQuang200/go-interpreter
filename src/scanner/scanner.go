package scanner

import (
	"bytes"
	"fmt"
	"os"

	"github.com/MQuang200/go-interpreter/src/token"
)

func Scan(content []byte) ([]token.Token, bool) {
	content = bytes.TrimSpace(content)

	tokens := []token.Token{}
	hadError := false

	for _, char := range content {
		switch char {
		case '(':
			tokens = append(tokens, token.NewToken(token.LEFT_PAREN, char, nil))
		case ')':
			tokens = append(tokens, token.NewToken(token.RIGHT_PAREN, char, nil))
		case '{':
			tokens = append(tokens, token.NewToken(token.LEFT_BRACE, char, nil))
		case '}':
			tokens = append(tokens, token.NewToken(token.RIGHT_BRACE, char, nil))
		case '.':
			tokens = append(tokens, token.NewToken(token.DOT, char, nil))
		case ',':
			tokens = append(tokens, token.NewToken(token.COMMA, char, nil))
		case '-':
			tokens = append(tokens, token.NewToken(token.MINUS, char, nil))
		case '+':
			tokens = append(tokens, token.NewToken(token.PLUS, char, nil))
		case ';':
			tokens = append(tokens, token.NewToken(token.SEMICOLON, char, nil))
		case '*':
			tokens = append(tokens, token.NewToken(token.STAR, char, nil))
		default:
			printError(1, char)
			hadError = true
		}
	}

	tokens = append(tokens, token.NewToken(token.EOF, byte(0), nil))

	return tokens, hadError
}

func printError(line int, char byte) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", line, string(char))
}
