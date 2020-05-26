package main

import (
	"testing"
)

func TestScanTokens(t *testing.T) {
	table := []struct {
		input string
		output []TokenType
	}{
		{"()[]{}", []TokenType{LeftParen, RightParen, LeftBracket, RightBracket, LeftBrace, RightBrace, EOF}},
	}

	for _, test := range table {
		t.Run(test.input, test.output) {
			s := NewScanner(test.input)
		}
	}