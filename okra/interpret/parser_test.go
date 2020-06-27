package interpret

import "testing"

func TestParsing(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		// TODO: Fill in parser tests!"}
		{"1 + 2", "(+ 1 2)"},
		{"1 * 2 + 1 / 3", "()"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.tokens)
			expr := parser.Parse()
			str := expr.String()

			if str != test.output {
				t.Errorf("Expected %q, received %q", test.output, expr)
			}
		})
	}
}
