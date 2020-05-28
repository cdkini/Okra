package interpret

import (
	"testing"
)

func TestScanWhitespace(t *testing.T) {
	table := []struct {
		in   string
		out  []TokenType
		line int
	}{
		{" ", []TokenType{EOF}, 1},
		{"\t\r\v\f", []TokenType{EOF}, 1},
		{"\n", []TokenType{EOF}, 2},
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
					if scanner.line != test.line {
						t.Errorf("Expected to be on line %d, actually on %d", scanner.line, test.line)
					}
				}
			}
		})
	}
}

func TestScanSingleCharTokens(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{"({[", []TokenType{LeftParen, LeftBracket, LeftBrace, EOF}},
		{")}]", []TokenType{RightParen, RightBracket, RightBrace, EOF}},
		{",.-", []TokenType{Comma, Dot, Minus, EOF}},
		{"+;*", []TokenType{Plus, Semicolon, Star, EOF}},
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanDoubleCharTokens(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{"!", []TokenType{Bang, EOF}},
		// {"!=", []TokenType{BangEqual, EOF}},
		{"=", []TokenType{Equal, EOF}},
		// {"==", []TokenType{EqualEqual, EOF}},
		{">", []TokenType{Greater, EOF}},
		// {">=", []TokenType{GreaterEqual, EOF}},
		{"<", []TokenType{Less, EOF}},
		// {"<=", []TokenType{LessEqual, EOF}},
		{"&&", []TokenType{And, EOF}},
		// {"||", []TokenType{Or, EOF}},
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanComments(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{"// This is a comment", []TokenType{EOF}},
		{"// Random text: +-=&&!forclass", []TokenType{EOF}},
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanString(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanNumeric(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}

func TestScanIdentifierAndKeyword(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{}, // TODO: Fill out tests!
	}

	for _, test := range table {
		t.Run(test.in, func(t *testing.T) {
			scanner := NewScanner(test.in)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.out) {
				t.Errorf("Expected %d tokens, received %d", len(tokens), len(test.out))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.out[i] {
						t.Errorf("Expected %v, received %v", tokens[i].tokenType, test.out[i])
					}
				}
			}
		})
	}
}
