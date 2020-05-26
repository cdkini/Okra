package interpret

import (
	"testing"
)

func TestScanTokens(t *testing.T) {
	table := []struct {
		input string
		output []TokenType
	}{
		{"()[]{}", []TokenType{LeftParen, RightParen, LeftBracket, RightBracket, LeftBrace, RightBrace, EOF}}
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			s := NewScanner(test.input)
			tokens := s.scanTokens()

			if len(tokens) != len(test.output) {
				t.Errorf("Error: Expected %d tokens, got %d", len(test.output), len(tokens))
			} else {
				for i := range tokens {
					if tokens[i] != test.out[i] {
						t.Errorf("Error: Expected %v, got %v", test.output[i], tokens[i])
					}
				}
			}
		})
	}
}