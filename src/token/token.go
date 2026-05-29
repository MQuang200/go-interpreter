package token

type TokenType string

const (
	// unspecified
	UNSPECIFIED TokenType = ""

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

	// single or double characters
	EQUAL       TokenType = "EQUAL"
	EQUAL_EQUAL TokenType = "EQUAL_EQUAL"
	BANG        TokenType = "BANG"
	BANG_EQUAL  TokenType = "BANG_EQUAL"

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
