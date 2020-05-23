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
		s.addToken(LeftParen, nil)
		break
	case ")":
		s.addToken(RightParen, nil)
		break
	case "{":
		s.addToken(LeftBracket, nil)
		break
	case "}":
		s.addToken(RightBracket, nil)
		break
	case "[":
		s.addToken(LeftBrace, nil)
		break
	case "]":
		s.addToken(RightBrace, nil)
		break
	case ",":
		s.addToken(RightBrace, nil)
		break
	case ".":
		s.addToken(RightBrace, nil)
		break
	case "-":
		s.addToken(RightBrace, nil)
		break
	case "+":
		s.addToken(RightBrace, nil)
		break
	case ";":
		s.addToken(Semicolon, nil)
		break
	case "*":
		s.addToken(Star, nil)
		break

}

func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current-1])
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	s.tokens = append(s.tokens, &Token{tokenType, s.source[s.start:s.current], literal, s.line})
}
