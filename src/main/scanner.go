package main

import (
	"errors"
	"strconv"
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
	case '\t', '\v', '\f', '\r', ' ':
		break
	case '\n':
		s.line++
		break

	// Single character tokens
	case '(':
		s.addToken(LeftParen)
		break
	case ')':
		s.addToken(RightParen)
		break
	case '{':
		s.addToken(LeftBracket)
		break
	case '}':
		s.addToken(RightBracket)
		break
	case '[':
		s.addToken(LeftBrace)
		break
	case ']':
		s.addToken(RightBrace)
		break
	case ',':
		s.addToken(Comma)
		break
	case '.':
		s.addToken(Dot)
		break
	case '-':
		s.addToken(Minus)
		break
	case '+':
		s.addToken(Plus)
		break
	case ';':
		s.addToken(Semicolon)
		break
	case '*':
		s.addToken(Star)
		break

	// Single or double character tokens
	case '!':
		s.addToken(s.ternaryMatch('=', BangEqual, Bang))
		break
	case '=':
		s.addToken(s.ternaryMatch('=', EqualEqual, Equal))
		break
	case '<':
		s.addToken(s.ternaryMatch('=', LessEqual, Less))
		break
	case '>':
		s.addToken(s.ternaryMatch('=', GreaterEqual, Greater))
		break
	case '&':
		if s.match('&') {
			s.addToken(And)
		}
		break
	case '|':
		if s.match('|') {
			s.addToken(Or)
		}
		break

	// Comments (single and multiline)
	case '/':
		if s.match('/') {
			for s.peek(0) != '\n' && s.curr < len(s.source) {
				s.advance()
			}
		} else if s.match('*') {
			for s.peek(0) != '*' && s.peek(1) != '/' && s.curr < len(s.source) {
				if s.peek(0) == '\n' {
					s.line++
				}
				s.advance()
				s.advance()
			}
		} else {
			s.addToken(Slash)
		}
		break

	case '"':
		err := s.addStringToken()
		checkErr(-1, err)

	default:
		if unicode.IsDigit(c) {
			s.addNumericToken()
		} else if unicode.IsLetter(c) {
			s.addIdentifierToken()
		} else {
			return errors.New("Invalid character")
		}
	}
	return nil
}

func (s *Scanner) advance() rune {
	s.curr++
	return s.source[s.curr-1]
}

func (s *Scanner) addToken(tokenType TokenType, literal ...interface{}) {
	if len(literal) == 1 {
		s.tokens = append(s.tokens, &Token{tokenType, string(s.source[s.start:s.curr]), literal[0], s.line})
	} else {
		s.tokens = append(s.tokens, &Token{tokenType, string(s.source[s.start:s.curr]), nil, s.line})
	}
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

func (s *Scanner) ternaryMatch(expectedChar rune, ifTrue TokenType, ifFalse TokenType) TokenType {
	if s.curr >= len(s.source) {
		return ifFalse
	}
	if s.source[s.curr+1] == expectedChar {
		s.curr++
		return ifTrue
	}
	return ifFalse
}

func (s *Scanner) addStringToken() error {
	for s.peek(0) != '\\' && s.curr < len(s.source) {
		if s.peek(0) == '\n' {
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

func (s *Scanner) addNumericToken() error {
	for unicode.IsDigit(s.peek(0)) {
		s.advance()
	}

	// Consumes floating point values by ignoring DOT TokenType
	if s.peek(0) == '.' && unicode.IsDigit(s.peek(1)) {
		s.advance()
		for unicode.IsDigit(s.peek(0)) {
			s.advance()
		}
	}

	num, err := strconv.ParseFloat(string(s.source[s.start:s.curr]), 64)
	checkErr(-1, err)
	s.addToken(Numeric, num)

	return nil
}

func (s *Scanner) addIdentifierToken() {
	for unicode.IsLetter(s.source[s.curr]) || unicode.IsDigit(s.source[s.curr]) {
		s.advance()
	}
	text := string(s.source[s.start:s.curr])
	id := keywordDict[text]
	if id == 0 {
		s.addToken(Identifier)
	} else {
		s.addToken(id)
	}
}

func (s *Scanner) peek(n int) rune {
	if s.curr+n >= len(s.source) {
		return '\u0000' // Null terminator
	}
	return s.source[s.curr]
}