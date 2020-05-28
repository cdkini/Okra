package interpret

import (
	"testing"
)

func TestScanWhitespace(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
		line   int
	}{
		{" ", []TokenType{EOF}, 1},
		{"\t\r\v\f", []TokenType{EOF}, 1},
		{"\n", []TokenType{EOF}, 2},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].tokenType)
					}
					if scanner.line != test.line {
						t.Errorf("Expected to be on line %d, actually on %d", test.line, scanner.line)
					}
				}
			}
		})
	}
}

func TestScanSingleCharTokens(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
	}{
		{"({[", []TokenType{LeftParen, LeftBracket, LeftBrace, EOF}},
		{")}]", []TokenType{RightParen, RightBracket, RightBrace, EOF}},
		{",.-", []TokenType{Comma, Dot, Minus, EOF}},
		{"+;*", []TokenType{Plus, Semicolon, Star, EOF}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].tokenType)
					}
				}
			}
		})
	}
}

func TestScanDoubleCharTokens(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
	}{
		{"!", []TokenType{Bang, EOF}},
		{"!=", []TokenType{BangEqual, EOF}},
		{"=", []TokenType{Equal, EOF}},
		{"==", []TokenType{EqualEqual, EOF}},
		{">", []TokenType{Greater, EOF}},
		{">=", []TokenType{GreaterEqual, EOF}},
		{"<", []TokenType{Less, EOF}},
		{"<=", []TokenType{LessEqual, EOF}},
		{"&&", []TokenType{And, EOF}},
		{"||", []TokenType{Or, EOF}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].tokenType)
					}
				}
			}
		})
	}
}

func TestScanComments(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
		line   int
	}{
		{"// This is a comment", []TokenType{EOF}, 1},
		{"// Random text: +-=&&!forclassfunc", []TokenType{EOF}, 1},
		{"// Line break ends comment \n", []TokenType{EOF}, 2},
		{"/* Yet another comment */", []TokenType{EOF}, 1},
		{"/* \n * A \n * proper \n * multiline * \n comment */", []TokenType{EOF}, 5},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].tokenType)
					}
					if scanner.line != test.line {
						t.Errorf("Expected to be on line %d, actually on %d", test.line, scanner.line)
					}
				}
			}
		})
	}
}

func TestScanString(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
	}{
		{}, // TODO: Fill output tests!
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].tokenType)
					}
				}
			}
		})
	}
}

func TestScanNumeric(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
	}{
		{}, // TODO: Fill output tests!
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].tokenType)
					}
				}
			}
		})
	}
}

func TestScanIdentifierAndKeyword(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
	}{
		{}, // TODO: Fill output tests!
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.scanTokens()
			tokens := scanner.tokens

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].tokenType != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].tokenType)
					}
				}
			}
		})
	}
}
