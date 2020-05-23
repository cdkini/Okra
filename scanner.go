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

	// Ignore all whitespace but record line break
	case " ":
	case "\\r":
	case "\\t":
		break
	case "\\n":
		s.line++
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

	// Single or double character tokens
	case "!":
		s.addToken(s.ternaryMatch("=", BANGEQUAL, BANG), nil)
		break
	case "=":
		s.addToken(s.ternaryMatch("=", EQUALEQUAL, EQUAL), nil)
		break
	case "<":
		s.addToken(s.ternaryMatch("=", LESSEQUAL, LESS), nil)
		break
	case ">":
		s.addToken(s.ternaryMatch("=", GREATEREQUAL, GREATER), nil)
		break
	case "&":
		if s.match("&") {
			s.addToken(AND, nil)
		}
		break
	case "|":
		if s.match("|") {
			s.addToken(OR, nil)
		}
		break

	// Comments (single and multiline)
	case "/":
		if match("/") {
			// TODO: Fix comment implementation
		} else {
			s.addToken(SLASH, nil)
		}
		break
	case "*":
		if match("/") {
			// TODO: Fix comment implementation
		} else {
			s.addToken(STAR, nil)
		}
		break

	case "\"":
		tokenType, err := addStringToken()




}

func (s *Scanner) advance() string {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	s.tokens = append(s.tokens, &Token{tokenType, s.source[s.start:s.current], literal, s.line})
}

func (s *Scanner) match(expectedChar string) bool {
	if s.current >= len(s.source) {
		return false
	}
	if s.source[s.current+1] == expectedChar {
		s.current++
		return true
	}
	return false
}

func (s *Scanner) addStringToken() (TokenType, error) {
	for s.peek() != "\"" && s.current < len(s.source) {
		if (speek() == '\n') {
			s.line++;                           
		}
		advance();  
	}
}

func (s *Scanner) ternaryMatch(expectedChar string, ifTrue TokenType, ifFalse TokenType) TokenType {
	if s.current >= len(s.source) {
		return false
	}
	if s.source[s.current+1] == expectedChar {
		s.current++
		return ifTrue
	}
	return ifFalse
}

func (s *Scanner) peek() string {
	if s.current >= len(s.source) {
		return "\\0"
	}
	return s.source[s.current]
}
