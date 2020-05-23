package main

type Scanner struct {
	source  string
	tokens  []*Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{source, make([]*Token, 0), 0, 0, 0}
}

func (s *Scanner) scanTokens() ([]Token, error) {
	for s.current < len(s.source) {
		s.start = s.current
		s.scan()
	}
	s.tokens = append(s.tokens, &Token{EOF, "", nil, s.line})
}

func (s *Scanner) scan() {
	c := s.advance()
	switch c {

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
	case ".":
		s.addToken(DOT, nil)
		break
	case "-":
		s.addToken(MINUS, nil)
		break
	case "+":
		s.addToken(PLUS, nil)
		break
	case ";":
		s.addToken(SEMICOLON, nil)
		break
	case "*":
		s.addToken(STAR, nil)
		break

	case "!":
		s.addToken(s.nextCharMatch("="), BANGEQUAL, BANG)
		break
	case "=":
		s.addToken(s.nextCharMatch("="), EQUALEQUAL, EQUAL)
		break
	case "<":
		s.addToken(s.nextCharMatch("="), LESSEQUAL, LESS)
		break
	case ">":
		s.addToken(s.nextCharMatch("="), GREATEREQUAL, GREATER)
		break
	case "&":
		s.addToken(s.nextCharMatch("&"), AND, INVALID)
		break
	case "|":
		s.addToken(s.nextCharMatch("|"), OR, INVALID)
		break

}

func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current-1])
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	s.tokens = append(s.tokens, &Token{tokenType, s.source[s.start:s.current], literal, s.line})
}

func (s *Scanner) nextCharMatch(expectedChar string, ifTrue TokenType, ifFalse TokenType) bool {
	if s.current >= len(s.source) {
		return false
	}
	if s.source[s.current + 1] == expectedChar {
		s.current++
		return ifTrue
	}
	return ifFalse
}
