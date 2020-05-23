package main

import (
	"errors"
	"unicode"
)

type Scanner struct {
	source []rune
	tokens []*Token
	start  int
	curr   int
	line   int
}

func NewScanner(source string) *Scanner {
	return &Scanner{[]rune(source), make([]*Token, 0), 0, 0, 0}
}

func (s *Scanner) scanTokens() []*Token {
	for s.curr < len(s.source) {
		s.start = s.curr
		err := s.scan()
		checkErr(-1, err)
	}
	s.tokens = append(s.tokens, &Token{EOF, "", nil, s.line})
	return s.tokens
}

func (s *Scanner) scan() error {
	c := s.advance()
	switch c {

	// Ignore all whitespace but record line break
	case unicode.IsSpace(c):
		if (c == '\n') {
			s.line++
		}
		break

	// Single character tokens
	case "(":
		s.addToken(LEFTPAREN, nil)
		break
	case ")":
		s.addToken(RIGHTPAREN, nil)
		break
	case "{":
		s.addToken(LEFTBRACKET, nil)
		break
	case "}":
		s.addToken(RIGHTBRACKET, nil)
		break
	case "[":
		s.addToken(LEFTBRACE, nil)
		break
	case "]":
		s.addToken(RIGHTBRACE, nil)
		break
	case ",":
		s.addToken(COMMA, nil)
		break
	case '.':
		s.addToken(DOT, nil)
		break
	case '-':
		s.addToken(MINUS, nil)
		break
	case '+':
		s.addToken(PLUS, nil)
		break
	case ';':
		s.addToken(SEMICOLON, nil)
		break

	// Single or double character tokens
	case '!':
		s.addToken(s.ternaryMatch('=', BANGEQUAL, BANG), nil)
		break
	case '=':
		s.addToken(s.ternaryMatch('=', EQUALEQUAL, EQUAL), nil)
		break
	case '<':
		s.addToken(s.ternaryMatch('=', LESSEQUAL, LESS), nil)
		break
	case '>':
		s.addToken(s.ternaryMatch('=', GREATEREQUAL, GREATER), nil)
		break
	case '&':
		if s.match('&') {
			s.addToken(AND, nil)
		}
		break
	case '|':
		if s.match('|') {
			s.addToken(OR, nil)
		}
		break

	// Comments (single and multiline)
	case '/':
		if s.match('/') {
			// TODO: Fix comment implementation
		} else {
			s.addToken(SLASH, nil)
		}
		break
	case '*':
		if s.match('/') {
			// TODO: Fix comment implementation
		} else {
			s.addToken(STAR, nil)
		}
		break

	case '"':
		tokenType, err := s.addStringToken()

	default:
		return errors.New("Invalid character")
	}
	return nil
}

func (s *Scanner) advance() rune {
	s.curr++
	return s.source[s.curr-1]
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	s.tokens = append(s.tokens, &Token{tokenType, string(s.source[s.start:s.curr]), literal, s.line})
}

func (s *Scanner) match(expectedChar rune) bool {
	if s.curr >= len(s.source) {
		return false
	}
	if s.source[s.curr+1] == expectedChar {
		s.curr++
		return true
	}
	return false
}

func (s *Scanner) addStringToken() (TokenType, error) {
	for s.peek() != '\' && s.curr < len(s.source) {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
}

func (s *Scanner) addNumericToken() (TokenType, error) {
	// TODO: Open to implement
}

func (s *Scanner) ternaryMatch(expectedChar rune, ifTrue TokenType, ifFalse TokenType) TokenType {
	if s.curr >= len(s.source) {
		return ifFalse
	}
	if string(s.source[s.curr+1]) == expectedChar {
		s.curr++
		return ifTrue
	}
	return ifFalse
}

func (s *Scanner) peek() rune {
	if s.curr < len(s.source) {
		return nil
	}
	return s.source[s.curr]
}
