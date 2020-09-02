package parse

import (
	"Okra/src/interpreter/scan"
	"testing"
)

func TestParseExpressionStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		{"1 + 1 = 2;", []string{"ast.ExpressionStmt"}},
		{"a + b >= c;", []string{"ast.ExpressionStmt"}},
		{"!d;", []string{"ast.ExpressionStmt"}},
		{"true != false;", []string{"ast.ExpressionStmt"}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

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
		{"print 1;", []string{"ast.PrintStmt"}},
		{"print \"Hello, World!\";", []string{"ast.PrintStmt"}},
		{"print abc;", []string{"ast.PrintStmt"}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

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
		{"var a: 1;", []string{"ast.VariableStmt"}},
		{"var name: \"Bob\";", []string{"ast.VariableStmt"}},
		{"var abc;", []string{"ast.VariableStmt"}},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

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
		// {},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}

// TODO: Fill out tests!
func TestParseIfStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		// {},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}

// TODO: Fill out tests!
func TestParseForStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		// {},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}

// TODO: Fill out tests!
func TestParseFuncStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		// {},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}

// TODO: Fill out tests!
func TestParseReturnStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		// {},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}

// TODO: Fill out tests!
func TestParseStructStmt(t *testing.T) {
	table := []struct {
		input  string
		output []string
	}{
		// {},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmts := parser.Parse()

			for i, stmt := range stmts {
				if stmt.GetType() != test.output[i] {
					t.Errorf("Expected *%v, received %T", test.output[i], stmt)
				}
			}
		})
	}
}
