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

<<<<<<< HEAD
	for i := 0; i < len(content); i++ {
		char := content[i]
		tType, tText := classifyToken(content, &i)
		if tType == token.UNSPECIFIED {
			printError(1, char)
			hadError = true
			continue
		}
		tokens = append(tokens, newToken(tType, tText, nil))
	}

	tokens = append(tokens, newToken(token.EOF, "", nil))

	return tokens, hadError
}

func classifyToken(content []byte, i *int) (token.TokenType, string) {
	var tt token.TokenType
	var ttStr string
	switch content[*i] {
	case '(':
		tt = token.LEFT_PAREN
		ttStr = "("
	case ')':
		tt = token.RIGHT_PAREN
		ttStr = ")"
	case '{':
		tt = token.LEFT_BRACE
		ttStr = "{"
	case '}':
		tt = token.RIGHT_BRACE
		ttStr = "}"
	case '.':
		tt = token.DOT
		ttStr = "."
	case ',':
		tt = token.COMMA
		ttStr = ","
	case '-':
		tt = token.MINUS
		ttStr = "-"
	case '+':
		tt = token.PLUS
		ttStr = "+"
	case ';':
		tt = token.SEMICOLON
		ttStr = ";"
	case '*':
		tt = token.STAR
		ttStr = "*"
	case '=':
		if *i < len(content)-1 && content[*i+1] == '=' {
			tt = token.EQUAL_EQUAL
			ttStr = "=="
			*i++
		} else {
			tt = token.EQUAL
			ttStr = "="
		}
	default:
		tt = token.UNSPECIFIED
	}
	return tt, ttStr
}

func newToken(tokenType token.TokenType, tokenStr string, value interface{}) token.Token {
	return token.Token{
		TokenType: tokenType,
		Text:      tokenStr,
		Value:     value}
=======
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
>>>>>>> ad204e3 (refactor codebase)
}

func printError(line int, char byte) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", line, string(char))
}
