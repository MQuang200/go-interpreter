package token

type TokenType string

const (
  EOF TokenType = "EOF"
  LEFT_PAREN = "LEFT_PAREN"
  RIGHT_PAREN = "RIGHT_PAREN"
  LEFT_BRACE = "LEFT_BRACE"
  RIGHT_BRACE = "RIGHT_BRACE"
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
