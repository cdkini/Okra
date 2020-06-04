package interpret

import "testing"

func TestParsing(t *testing.T) {
	table := []struct {
		input  []Token
		output []Expr
	}{
		// TODO: Fill in parser tests
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
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
