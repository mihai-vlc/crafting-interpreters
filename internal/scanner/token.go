package scanner

import "fmt"

type TokenKind int32

const (
	// Literal
	TokenIdentifier TokenKind = iota
	TokenString
	TokenNumber

	// Symbols
	TokenEqual

	// Keywords
	TokenVar

	// Other
	TokenEOF
)

func kindToString(kind TokenKind) string {
	switch kind {
	case TokenIdentifier:
		return "identifier"
	case TokenString:
		return "string"
	case TokenNumber:
		return "number"
	case TokenEqual:
		return "="
	case TokenVar:
		return "var"
	case TokenEOF:
		return "[EOF]"
	default:
		return "N/A"
	}
}

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
		kindToString(t.kind),
		t.value,
		t.position.line,
		t.position.column,
	)
}
