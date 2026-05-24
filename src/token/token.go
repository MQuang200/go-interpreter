package token

type TokenType string

const (
	// single characters

	LEFT_PAREN  TokenType = "LEFT_PAREN"
	RIGHT_PAREN TokenType = "RIGHT_PAREN"
	LEFT_BRACE  TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"
	COMMA       TokenType = "COMMA"
	DOT         TokenType = "DOT"
	MINUS       TokenType = "MINUS"
	PLUS        TokenType = "PLUS"
	SEMICOLON   TokenType = "SEMICOLON"
	SLASH       TokenType = "SLASH"
	STAR        TokenType = "STAR"
	EQUAL       TokenType = "EQUAL"
	EQUAL_EQUAL TokenType = "EQUAL_EQUAL"

	// EOF
	EOF TokenType = "EOF"
)

type Token struct {
	TokenType TokenType
	Text      string
	Value     interface{}
}

func (t Token) String() string {
	var value string
	if t.Value == nil {
		value = "null"
	} else {
		value = t.Value.(string)
	}

	return string(t.TokenType) + " " + t.Text + " " + value
}

func (t Token) EOFString() string {
	return string(EOF) + " null"
}
