package interpret_test

import (
	. "Okra/okra/interpret"
	"testing"
)

func TestParseExpressionStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		{"1 + 1 == 2;", []string{"interpret.ExpressionStmt"}},
		{"a + b >= c;", []string{"interpret.ExpressionStmt"}},
		{"!d;", []string{"interpret.ExpressionStmt"}},
		{"true != false;", []string{"interpret.ExpressionStmt"}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts, err := parser.Parse()

			if err {
				t.Error("Parse error noted")
			}

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}

func TestParsePrintStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		{"print 1;", []string{"interpret.PrintStmt"}},
		{"print \"Hello, World!\";", []string{"interpret.PrintStmt"}},
		{"print abc;", []string{"interpret.PrintStmt"}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts, err := parser.Parse()

			if err {
				t.Error("Parse error noted")
			}

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}

func TestParseVariableStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		{"var a = 1;", []string{"interpret.VariableStmt"}},
		{"var name = \"Bob\";", []string{"interpret.VariableStmt"}},
		{"var abc;", []string{"interpret.VariableStmt"}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts, err := parser.Parse()

			if err {
				t.Error("Parse error noted")
			}

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}

// TODO: Fill out tests!
func TestParseBlockStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		{"", []string{"interpret.BlockStmt"}},
		{"", []string{"interpret.BlockStmt"}},
		{"", []string{"interpret.BlockStmt"}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts, err := parser.Parse()

			if err {
				t.Error("Parse error noted")
			}

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}
