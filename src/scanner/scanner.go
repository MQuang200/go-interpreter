package scanner

import (
	"fmt"

	"github.com/MQuang200/go-interpreter/src/token"
)

func Scan(content []byte) []token.Token{
	tokens := []token.Token{}

	for _, char := range content {
		switch char {
		case '(':
      tokens = append(tokens, newToken(token.LEFT_PAREN, char, nil))
    case ')':
      tokens = append(tokens, newToken(token.RIGHT_PAREN, char, nil))
    case '{':
      tokens = append(tokens, newToken(token.LEFT_BRACE, char, nil))
    case '}':
      tokens = append(tokens, newToken(token.RIGHT_BRACE, char, nil))
    case '.':
      tokens = append(tokens, newToken(token.DOT, char, nil))
    case ',':
      tokens = append(tokens, newToken(token.COMMA, char, nil))
    case '-':
      tokens = append(tokens, newToken(token.MINUS, char, nil))
    case '+':
      tokens = append(tokens, newToken(token.PLUS, char, nil))
    case ';':
      tokens = append(tokens, newToken(token.SEMICOLON, char, nil))
    case '*':
      tokens = append(tokens, newToken(token.STAR, char, nil))
    default:
      printError(1, char)
    }
	}

  tokens = append(tokens, newToken(token.EOF, byte(0), nil))

  return tokens
}

func newToken(tokenType token.TokenType, char byte, value interface{}) token.Token{
  return token.Token{
    TokenType: tokenType,
    Text: string(char),
    Value: value,
  }
}

func printError(line int, char byte) {
  fmt.Printf("[line %d] Error: Unexpected character: %s\n", line, string(char))
}
