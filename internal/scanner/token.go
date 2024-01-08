package scanner

import "fmt"

type TokenKind string

const (
	// Literal
	TokenIdentifier = "identifier"
	TokenString     = "string"
	TokenNumber     = "number"

	// Symbols
	TokenEqual       = "equal"
	TokenPlus        = "plus"
	TokenMinus       = "minus"
	TokenStar        = "star"
	TokenDivide      = "divide"
	TokenModulo      = "modulo"
	TokenDot         = "dot"
	TokenSemicolon   = "semicolon"
	TokenComma       = "comma"
	TokenLess        = "less"
	TokenLessEqual   = "less_equal"
	TokenGrater      = "grater"
	TokenGraterEqual = "grater_equal"

	// Keywords
	TokenVar = "var"

	// Other
	TokenEOF = "[EOF]"
)

type Token struct {
	kind     TokenKind
	value    string
	position *Position
}

func NewToken(kind TokenKind, value string, position *Position) *Token {
	return &Token{
		kind:     kind,
		value:    value,
		position: position,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf(
		"<%s `%s` %d:%d>",
		t.kind,
		t.value,
		t.position.line,
		t.position.column,
	)
}
