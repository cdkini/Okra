package interpret

import "testing"

func TestParsing(t *testing.T) {
	table := []struct {
		input  string
		output Expr
	}{
		// TODO: Fill in parser tests!"}
		{"!x", Unary{Token{}, Expr{Identifier}}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.tokens)
			expr := parser.evaluate()

			if expr != test.output {
				t.Errorf("Expected %q, received %q", test.output.String(), expr.String())
			}
		})
	}
}
