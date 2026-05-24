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

	for i := 0; i < len(content); i++ {
		switch char := content[i]; char {
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
			if i < len(content)-1 && content[i+1] == '=' {
				tokens = append(tokens, newToken(token.EQUAL_EQUAL, "==", nil))
				i++
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
