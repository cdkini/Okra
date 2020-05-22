package main

type Scanner struct {
	source   string
	tokens   []Token
	start    int
	current  int
	line     int
	hasError bool
}

func NewScanner(source string) *Scanner {
	return &Scanner{source, make([]Token, 0), 0, 0, 0, false}
}

func (s *Scanner) scanTokens() ([]Token, error) {
	for s.current < len(s.source) {
		s.start = s.current
		s.scan()
	}

	s.tokens = append(s.tokens, Token{EOF, "", nil, s.line})
}

func (s *Scanner) scan() {

}
