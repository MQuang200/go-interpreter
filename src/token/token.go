package token

type TokenType int

const (
  EOF TokenType = iota
  LEFT_PAREN
  RIGHT_PAREN
)

type Token struct {
  TokenType TokenType
  Text string
  Value interface{}
}

func(t TokenType) String() string {
  switch t {
  case LEFT_PAREN:
    return "LEFT_PAREN"
  case RIGHT_PAREN:
    return "RIGHT_PAREN"
  default:
    return ""
  }
}

func(t Token) String() string{
  var value string
  if t.Value == nil {
    value = "null"
  } else {
    value = t.Value.(string)
  }
  return t.TokenType.String() + " " + t.Text + " " + value
}
