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
			tokens = append(tokens, newToken(token.LEFT_PAREN, string(char), nil))
		case ')':
			tokens = append(tokens, newToken(token.RIGHT_PAREN, string(char), nil))
		case '{':
			tokens = append(tokens, newToken(token.LEFT_BRACE, string(char), nil))
		case '}':
			tokens = append(tokens, newToken(token.RIGHT_BRACE, string(char), nil))
		case '.':
			tokens = append(tokens, newToken(token.DOT, string(char), nil))
		case ',':
			tokens = append(tokens, newToken(token.COMMA, string(char), nil))
		case '-':
			tokens = append(tokens, newToken(token.MINUS, string(char), nil))
		case '+':
			tokens = append(tokens, newToken(token.PLUS, string(char), nil))
		case ';':
			tokens = append(tokens, newToken(token.SEMICOLON, string(char), nil))
		case '*':
			tokens = append(tokens, newToken(token.STAR, string(char), nil))
		case '=':
			s := len(tokens)
			if len(tokens) != 0 && tokens[s-1].TokenType == token.EQUAL {
				tokens = tokens[:s-1]
				tokens = append(tokens, newToken(token.EQUAL_EQUAL, "==", nil))
			} else {
				tokens = append(tokens, newToken(token.EQUAL, string(char), nil))
			}
		default:
			printError(1, char)
			hadError = true
		}
	}

	tokens = append(tokens, newToken(token.EOF, "", nil))

	return tokens, hadError
}

func newToken(tokenType token.TokenType, tokenStr string, value interface{}) token.Token {
	return token.Token{
		TokenType: tokenType,
		Text:      tokenStr,
		Value:     value}
}

func printError(line int, char byte) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", line, string(char))
}
