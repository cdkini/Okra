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
		if c == '\n' {
			s.line++
		}
		break

	// Single character tokens
	case '(':
		s.addToken(LeftParen, nil)
		break
	case ')':
		s.addToken(RightParen, nil)
		break
	case '{':
		s.addToken(LeftBracket, nil)
		break
	case '}':
		s.addToken(RightBracket, nil)
		break
	case '[':
		s.addToken(LeftBrace, nil)
		break
	case ']':
		s.addToken(RightBrace, nil)
		break
	case ',':
		s.addToken(Comma, nil)
		break
	case '.':
		s.addToken(Dot, nil)
		break
	case '-':
		s.addToken(Minus, nil)
		break
	case '+':
		s.addToken(Plus, nil)
		break
	case ';':
		s.addToken(Semicolon, nil)
		break

	// Single or double character tokens
	case '!':
		s.addToken(s.ternaryMatch('=', BangEqual, Bang), nil)
		break
	case '=':
		s.addToken(s.ternaryMatch('=', EqualEqual, Equal), nil)
		break
	case '<':
		s.addToken(s.ternaryMatch('=', LessEqual, Less), nil)
		break
	case '>':
		s.addToken(s.ternaryMatch('=', GreaterEqual, Greater), nil)
		break
	case '&':
		if s.match('&') {
			s.addToken(And, nil)
		}
		break
	case '|':
		if s.match('|') {
			s.addToken(Or, nil)
		}
		break

	// Comments (single and multiline)
	case '/':
		if s.match('/') {
			// TODO: Fix comment implementation
		} else {
			s.addToken(Slash, nil)
		}
		break
	case '*':
		if s.match('/') {
			// TODO: Fix comment implementation
		} else {
			s.addToken(Star, nil)
		}
		break

	case '"':
		err := s.addStringToken()
		checkErr(-1, err)

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

func (s *Scanner) addStringToken() error {
	for s.peek() != '\\' && s.curr < len(s.source) {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.curr >= len(s.source) {
		return errors.New("Unterminated string")
	}

	s.advance()
	str := s.source[s.start+1 : s.curr-1]
	s.addToken(String, str)

	return nil
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
