package interpret

import (
	"strconv"
	"unicode"
)

// A Scanner takes in some stream of characters and tokenizes them based on Okra's syntax
type Scanner struct {
	source []rune
	tokens []Token // Populated as result of ScanTokens()
	start  int     // Where the current token begins
	curr   int     // Where the scanner is within the source
	line   int
	col    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{[]rune(source), make([]Token, 0), 0, 0, 1, 1}
}

func (s *Scanner) Tokens() []Token {
	return s.tokens
}

func (s *Scanner) Line() int {
	return s.line
}

func (s *Scanner) Col() int {
	return s.col
}

// ScanTokens iterates through the source text and generates tokens based on Okra's defined syntax rules
//   Args: nil
//   Returns: Slice of token pointers
func (s *Scanner) ScanTokens() []Token {
	for s.curr < len(s.source) {
		s.start = s.curr
		s.scan()
	}
	s.tokens = append(s.tokens, Token{EOF, "EOF", nil, s.line, s.col})
	return s.tokens
}

func (s *Scanner) scan() {
	c := s.advance()
	switch c {

	// Ignore all whitespace but record line break
	case '\t', '\v', '\f', '\r', ' ':
		break
	case '\n':
		s.lineBreak()
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
	case '*':
		s.addToken(Star, nil)
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
			break
		}
		ReportErr(-1, NewOkraError(s.line, s.col, "Invalid character"))
	case '|':
		if s.match('|') {
			s.addToken(Or, nil)
			break
		}
		ReportErr(-1, NewOkraError(s.line, s.col, "Invalid character"))

	// Comments
	case '/':
		if s.match('/') {
			for s.peek(0) != '\n' && s.curr < len(s.source) {
				s.advance()
			}
		} else {
			s.addToken(Slash, nil)
		}
		break

	case '"':
		s.addStringToken()

	default:
		if unicode.IsDigit(c) {
			s.addNumericToken()
		} else if unicode.IsLetter(c) {
			s.addIdentifierToken()
		} else {
			ReportErr(-1, NewOkraError(s.line, s.col, "Invalid character"))
		}
	}
}

func (s *Scanner) advance() rune {
	s.curr++
	s.col++
	return s.source[s.curr-1]
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	s.tokens = append(s.tokens, Token{tokenType, string(s.source[s.start:s.curr]), literal, s.line, s.col})
}

func (s *Scanner) match(expectedChar rune) bool {
	if s.curr >= len(s.source) {
		return false
	}
	if s.source[s.curr] == expectedChar {
		s.curr++
		s.col++
		return true
	}
	return false
}

func (s *Scanner) ternaryMatch(expectedChar rune, ifTrue TokenType, ifFalse TokenType) TokenType {
	if s.curr >= len(s.source) {
		return ifFalse
	}
	if s.source[s.curr] == expectedChar {
		s.curr++
		s.col++
		return ifTrue
	}
	return ifFalse
}

func (s *Scanner) addStringToken() {
	for s.peek(0) != '"' && s.curr < len(s.source) {
		if s.peek(0) == '\n' {
			s.lineBreak()
		}
		s.advance()
	}

	if s.curr >= len(s.source) {
		ReportErr(-1, NewOkraError(s.line, s.col, "Unterminated character"))
	}

	s.advance()
	str := s.source[s.start+1 : s.curr-1]
	s.addToken(String, str)
}

func (s *Scanner) addNumericToken() {
	for unicode.IsDigit(s.peek(0)) {
		s.advance()
	}

	// Consumes floating point values by ignoring DOT TokenType
	if s.peek(0) == '.' {
		s.advance()
		for unicode.IsDigit(s.peek(0)) {
			s.advance()
		}
	}

	num, err := strconv.ParseFloat(string(s.source[s.start:s.curr]), 64)
	CheckErr(-1, err, NewOkraError(s.line, s.col, "Could not scan numeric"))
	s.addToken(Numeric, num)
}

func (s *Scanner) addIdentifierToken() {
	for s.curr < len(s.source) && (unicode.IsLetter(s.source[s.curr]) || unicode.IsDigit(s.source[s.curr])) {
		s.advance()
	}
	text := string(s.source[s.start:s.curr])
	s.getKeyword(text)
}

func (s *Scanner) getKeyword(text string) {
	if i, ok := keywordDict[text]; !ok {
		s.addToken(Identifier, nil)
	} else {
		s.addToken(i, nil)
	}
}

func (s *Scanner) peek(n int) rune {
	if s.curr+n >= len(s.source) {
		return '\u0000' // Null terminator
	}
	return s.source[s.curr]
}

func (s *Scanner) lineBreak() {
	s.line++
	s.col = 0
}
