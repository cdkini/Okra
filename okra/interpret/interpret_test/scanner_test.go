package interpret_test

import (
	. "Okra/okra/interpret"
	"io/ioutil"
	"testing"
)

func TestScanWhitespace(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
		line   int
	}{
		{" ", []TokenType{EOF}, 1},
		{"\t", []TokenType{EOF}, 1},
		{"\r", []TokenType{EOF}, 1},
		{"\v", []TokenType{EOF}, 1},
		{"\f", []TokenType{EOF}, 1},
		{"\n", []TokenType{EOF}, 2},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			tokens := scanner.Tokens()

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].TokenType() != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].TokenType())
					}
					if scanner.Line() != test.line {
						t.Errorf("Expected to be on line %d, actually on %d", test.line, scanner.Line())
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
		{"(", []TokenType{LeftParen, EOF}},
		{"{", []TokenType{LeftBrace, EOF}},
		{"[", []TokenType{LeftBracket, EOF}},
		{")", []TokenType{RightParen, EOF}},
		{"}", []TokenType{RightBrace, EOF}},
		{"]", []TokenType{RightBracket, EOF}},
		{",", []TokenType{Comma, EOF}},
		{".", []TokenType{Dot, EOF}},
		{"-", []TokenType{Minus, EOF}},
		{"+", []TokenType{Plus, EOF}},
		{";", []TokenType{Semicolon, EOF}},
		{"*", []TokenType{Star, EOF}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			tokens := scanner.Tokens()

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].TokenType() != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].TokenType())
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
			scanner.ScanTokens()
			tokens := scanner.Tokens()

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].TokenType() != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].TokenType())
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
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			tokens := scanner.Tokens()

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].TokenType() != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].TokenType())
					}
					if scanner.Line() != test.line {
						t.Errorf("Expected to be on line %d, actually on %d", test.line, scanner.Line())
					}
				}
			}
		})
	}
}

func TestScanMiscellaneous(t *testing.T) {
	table := []struct {
		input  string
		output []TokenType
	}{
		// String
		{"\"abc\"", []TokenType{String, EOF}},
		{"\"okra's tests\"", []TokenType{String, EOF}},

		// Numeric
		{"7", []TokenType{Numeric, EOF}},
		{"3.1415", []TokenType{Numeric, EOF}},

		// Id
		{"name", []TokenType{Identifier, EOF}},
		{"x", []TokenType{Identifier, EOF}},

		// Keyword
		{"class", []TokenType{Class, EOF}},
		{"else", []TokenType{Else, EOF}},
		{"false", []TokenType{False, EOF}},
		{"for", []TokenType{For, EOF}},
		{"func", []TokenType{Func, EOF}},
		{"if", []TokenType{If, EOF}},
		{"null", []TokenType{Null, EOF}},
		{"print", []TokenType{Print, EOF}},
		{"return", []TokenType{Return, EOF}},
		{"super", []TokenType{Super, EOF}},
		{"this", []TokenType{This, EOF}},
		{"true", []TokenType{True, EOF}},
		{"var", []TokenType{Variable, EOF}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			tokens := scanner.Tokens()

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i].TokenType() != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], tokens[i].TokenType())
					}
				}
			}
		})
	}
}

func TestScanFiles(t *testing.T) {
	table := []struct {
		path   string
		output []TokenType
	}{
		{"test_files/test1.okr", []TokenType{
			Print,
			String,
			Semicolon,
			EOF,
		}},
		{"test_files/test2.okr", []TokenType{
			Variable,
			Identifier,
			Equal,
			Numeric,
			Semicolon,
			Variable,
			Identifier,
			Equal,
			Numeric,
			Semicolon,
			Print,
			Variable,
			Plus,
			Variable,
			Semicolon,
			EOF,
		}},
		{"test_files/test3.okr", []TokenType{
			Func,
			Identifier,
			LeftParen,
			Identifier,
			RightParen,
			LeftBrace,
			Return,
			Identifier,
			Star,
			Identifier,
			RightBrace,
			Semicolon,
			EOF,
		}},
		{"test_files/test4.okr", []TokenType{
			Struct,
			Identifier,
			LeftBrace,
			Identifier,
			LeftParen,
			RightParen,
			LeftBrace,
			RightBrace,
			RightBrace,
			EOF,
		}},
	}

	for _, test := range table {
		t.Run(test.path, func(t *testing.T) {
			bytes, _ := ioutil.ReadFile(test.path)
			scanner := NewScanner(string(bytes))
			tokens := scanner.ScanTokens()

			if len(tokens) != len(test.output) {
				t.Errorf("Expected %d tokens, received %d", len(test.output), len(tokens))
			}
		})
	}
}
