package parse

import (
	"Okra/src/interpreter/scan"
	"testing"
)

func TestParseExpressionStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"1 + 1 = 2;", "ast.ExpressionStmt"},
		{"a + b >= c;", "ast.ExpressionStmt"},
		{"!d;", "ast.ExpressionStmt"},
		{"true != false;", "ast.ExpressionStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParsePrintStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"print 1;", "ast.PrintStmt"},
		{"print \"Hello, World!\";", "ast.PrintStmt"},
		{"print abc;", "ast.PrintStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseVariableStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"var a: 1;", "ast.VariableStmt"},
		{"var name: \"Bob\";", "ast.VariableStmt"},
		{"var abc;", "ast.VariableStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseBlockStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"{}", "ast.BlockStmt"},
		{"{var x: 1;}", "ast.BlockStmt"},
		{"{1 + 1;}", "ast.BlockStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseIfStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"if (true) {}", "ast.IfStmt"},
		{"if (x < 5) {}", "ast.IfStmt"},
		{"if (!y) {}", "ast.IfStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseForStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"for (true) {}", "ast.ForStmt"},
		{"for (x < 5) {}", "ast.ForStmt"},
		{"for (!y) {}", "ast.ForStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseFuncStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"func abs :: {}", "ast.FuncStmt"},
		{"func hello : name : { print \"Hello\" + name; }", "ast.FuncStmt"},
		{"func addOne : x : { return x + 1; }", "ast.FuncStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseReturnStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"return 1;", "ast.ReturnStmt"},
		{"return;", "ast.ReturnStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseStructStmt(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"struct Dog {}", "ast.StructStmt"},
		{"struct Person { construct :: {} }", "ast.StructStmt"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}
