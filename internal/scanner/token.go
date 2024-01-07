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
		return "<EOF>"
	default:
		return "N/A"
	}
}

type Token struct {
	kind TokenKind
	line int
}

func NewToken(kind TokenKind, line int) *Token {
	return &Token{
		kind: kind,
		line: line,
	}
}

func (t *Token) String() string {
	return kindToString(t.kind) + "#" + fmt.Sprint(t.line)
}
