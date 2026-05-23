package token

type TokenType string

const (
<<<<<<< HEAD
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
	EQUAL       TokenType = "EQUAL"
	EQUAL_EQUAL TokenType = "EQUAL_EQUAL"
=======
	// single characters

	LEFT_PAREN  = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
	LEFT_BRACE  = "LEFT_BRACE"
	RIGHT_BRACE = "RIGHT_BRACE"
	COMMA       = "COMMA"
	DOT         = "DOT"
	MINUS       = "MINUS"
	PLUS        = "PLUS"
	SEMICOLON   = "SEMICOLON"
	SLASH       = "SLASH"
	STAR        = "STAR"
>>>>>>> ad204e3 (refactor codebase)

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
<<<<<<< HEAD

	return string(t.TokenType) + " " + t.Text + " " + value
}

func (t Token) EOFString() string {
	return string(EOF) + " null"
=======
	return string(t.TokenType) + " " + t.Text + " " + value
}

func NewToken(tokenType TokenType, char byte, value interface{}) Token {
	return Token{
		TokenType: tokenType,
		Text:      string(char),
		Value:     value}
}

func (t Token) EOFString() string {
	return string(EOF) + "  null"
>>>>>>> ad204e3 (refactor codebase)
}
