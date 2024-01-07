package scanner

import (
	"fmt"
	"unicode"
)

type Scanner struct {
	source          []rune
	tokens          []*Token
	line            int
	column          int
	start           int
	currentPosition int
}

func NewScanner(source string) *Scanner {
	s := Scanner{
		source:          []rune(source),
		tokens:          []*Token{},
		line:            1,
		column:          0,
		start:           0,
		currentPosition: 0,
	}

	return &s
}

func (s *Scanner) ScanTokens() ([]*Token, error) {
	for !s.isAtEnd() {
		if err := s.scanToken(); err != nil {
			return nil, err
		}
	}

	s.addToken(TokenEOF, "EOF")

	return s.tokens, nil
}

func (s *Scanner) isAtEnd() bool {
	return s.currentPosition+1 >= len(s.source)
}

func (s *Scanner) scanToken() error {
	s.start = s.currentPosition

	var c = s.advance()

	switch {
	case c == '=':
		s.addToken(TokenEqual, "=")
	case c == ' ' || c == '\t' || c == '\r':
		// skip
	case c == '\n':
		s.line++
		s.column = 1
	case s.isAlpha(c):
		s.identifier()
	default:
		return s.error("Unexpected token %s", string(c))
	}

	return nil
}

func (s *Scanner) addToken(kind TokenKind, value string) {
	token := NewToken(kind, value, s.getPosition())
	s.tokens = append(s.tokens, token)
}

func (s *Scanner) advance() rune {
	var c = s.source[s.currentPosition]
	s.currentPosition++
	s.column++
	return c
}

func (s *Scanner) peek() rune {
	return s.source[s.currentPosition]
}
func (s *Scanner) peekNext() (rune, error) {
	var nextPos = s.currentPosition + 1

	if nextPos >= len(s.source) {
		return 0, fmt.Errorf("reached the end of the source")
	}

	return s.source[nextPos], nil
}

func (s *Scanner) getPosition() *Position {
	size := s.currentPosition - s.start - 1
	return NewPosition(s.line, s.column-size)
}

func (s *Scanner) identifier() {

	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	word := string(s.source[s.start:s.currentPosition])

	if word == "var" {
		s.addToken(TokenVar, word)
		return
	}

	s.addToken(TokenIdentifier, word)
}

func (s *Scanner) isAlphaNumeric(c rune) bool {
	return s.isAlpha(c) || unicode.IsDigit(c)
}

func (s *Scanner) isAlpha(c rune) bool {
	return unicode.IsLetter(c) || c == '_'
}

func (s *Scanner) error(msg string, args ...any) error {
	msg += fmt.Sprintf(", at position %d:%d", s.line, s.column)
	return fmt.Errorf(msg, args...)
}
