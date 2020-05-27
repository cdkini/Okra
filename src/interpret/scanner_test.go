package interpret

import (
	"testing"
)

func TestScanTokens(t *testing.T) {
	table := []struct {
		in  string
		out []TokenType
	}{
		{"+*", []TokenType{Plus, Star, EOF}},
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
