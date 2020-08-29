package scan

import (
	"Okra/okra/interpreter/ast"
	"Okra/okra/okraerr"
	"strconv"
	"unicode"
)

// A Scanner takes in some stream of characters and tokenizes them based on Okra's syntax
type Scanner struct {
	source []rune
	tokens []ast.Token // Populated as result of ScanTokens()
	start  int         // Where the current token begins
	curr   int         // Where the scanner is within the source
	line   int
	col    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{[]rune(source), make([]ast.Token, 0), 0, 0, 1, 1}
}

func (s *Scanner) Tokens() []ast.Token {
	return s.tokens
}

func (s *Scanner) Line() int {
	return s.line
}

// ScanTokens iterates through the source text and generates tokens based on Okra's defined syntax rules
//   Args: nil
//   Returns: Slice of token pointers
func (s *Scanner) ScanTokens() []ast.Token {
	for s.curr < len(s.source) {
		s.start = s.curr
		s.scan()
	}
	s.tokens = append(s.tokens, ast.Token{ast.EOF, "EOF", nil, s.line, s.col})
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
		s.addToken(ast.LeftParen, nil)
		break
	case ')':
		s.addToken(ast.RightParen, nil)
		break
	case '{':
		s.addToken(ast.LeftBrace, nil)
		break
	case '}':
		s.addToken(ast.RightBrace, nil)
		break
	case ',':
		s.addToken(ast.Comma, nil)
		break
	case '.':
		s.addToken(ast.Dot, nil)
		break
	case '-':
		s.addToken(ast.Minus, nil)
		break
	case '+':
		s.addToken(ast.Plus, nil)
		break
	case ';':
		s.addToken(ast.Semicolon, nil)
		break
	case '*':
		s.addToken(ast.Star, nil)
		break
	case '=':
		s.addToken(ast.Equal, nil)
		break
	case ':':
		s.addToken(ast.Colon, nil)
		break

	// Single or double character tokens
	case '!':
		s.addToken(s.ternaryMatch('=', ast.BangEqual, ast.Bang), nil)
		break
	case '<':
		s.addToken(s.ternaryMatch('=', ast.LessEqual, ast.Less), nil)
		break
	case '>':
		s.addToken(s.ternaryMatch('=', ast.GreaterEqual, ast.Greater), nil)
		break
	case '&':
		if s.match('&') {
			s.addToken(ast.And, nil)
			break
		}
		okraerr.ReportErr(s.line, s.col, "Invalid character")
	case '|':
		if s.match('|') {
			s.addToken(ast.Or, nil)
			break
		}
		okraerr.ReportErr(s.line, s.col, "Invalid character")

	// Comments
	case '/':
		if s.match('/') {
			for s.peek(0) != '\n' && s.curr < len(s.source) {
				s.advance()
			}
		} else {
			s.addToken(ast.Slash, nil)
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
			okraerr.ReportErr(s.line, s.col, "Invalid character")
		}
	}
}

func (s *Scanner) advance() rune {
	s.curr++
	s.col++
	return s.source[s.curr-1]
}

func (s *Scanner) addToken(tokenType ast.TokenType, literal interface{}) {
	s.tokens = append(s.tokens, ast.Token{tokenType, string(s.source[s.start:s.curr]), literal, s.line, s.col})
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

func (s *Scanner) ternaryMatch(expectedChar rune, ifTrue ast.TokenType, ifFalse ast.TokenType) ast.TokenType {
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
		okraerr.ReportErr(s.line, s.col, "Unterminated character")
	}

	s.advance()
	str := s.source[s.start+1 : s.curr-1]
	s.addToken(ast.String, str)
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
	okraerr.CheckErr(err, s.line, s.col, "Could not scan numeric")
	s.addToken(ast.Numeric, num)
}

func (s *Scanner) addIdentifierToken() {
	for s.curr < len(s.source) && (unicode.IsLetter(s.source[s.curr]) || unicode.IsDigit(s.source[s.curr])) {
		s.advance()
	}
	text := string(s.source[s.start:s.curr])
	s.getKeyword(text)
}

func (s *Scanner) getKeyword(text string) {
	if i, ok := ast.KeywordDict[text]; !ok {
		s.addToken(ast.Identifier, nil)
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
