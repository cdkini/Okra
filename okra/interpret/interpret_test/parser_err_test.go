package interpret_test

import (
	. "Okra/okra/interpret"
	"testing"
)

func TestParseError(t *testing.T) {
	table := []struct {
		input  string
		output bool
	}{
		{"1 + 1 == 2;", false},
		{"print \"Hello, World!\";", false},
		{"print a;", false},
		// {"print a;", true},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			_, err := parser.Parse()

			if err == true && test.output == false {
				t.Errorf("Parse error raised when no error is present: %v", test.input)
			}
			if err == false && test.output == true {
				t.Errorf("Parse error not raised when error is present: %v", test.input)
			}
		})
	}
}
