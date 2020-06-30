package interpret

import "testing"

func TestInterpreting(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		// TODO: Fill in interpreter tests!"}
		{"1+1", ""},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			// FIXME: Test broken after addition of statements
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.tokens)
			stmts := parser.Parse()
			interpreter := NewInterpreter()
			interpreter.Interpret(stmts)

			// if res != test.output {
			// 	t.Errorf("Expected %q, received %q", test.output, res)
			// }
		})
	}
}
