package token

type TokenType string

const (
  // single characters

  LEFT_PAREN = "LEFT_PAREN"
  RIGHT_PAREN = "RIGHT_PAREN"
  LEFT_BRACE = "LEFT_BRACE"
  RIGHT_BRACE = "RIGHT_BRACE"
  COMMA = "COMMA"
  DOT = "DOT"
  MINUS = "MINUS"
  PLUS = "PLUS"
  SEMICOLON = "SEMICOLON"
  SLASH = "SLASH"
  STAR = "STAR"

  // EOF
  EOF TokenType = "EOF"
)

type Token struct {
  TokenType TokenType
  Text string
  Value interface{}
}

func(t Token) String() string{
  var value string
  if t.Value == nil {
    value = "null"
  } else {
    value = t.Value.(string)
  }
  return string(t.TokenType) + " " + t.Text + " " + value
}
