package parse

import (
	"Okra/okra/interpreter/scan"
	"testing"
)

func TestParseUnaryExpr(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"-1", "ast.UnaryExpr"},
		{"!true", "ast.UnaryExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			expr := parser.Expression()

			if expr.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseBinaryExpr(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"1 + 2", "ast.BinaryExpr"},
		{"3 - 4", "ast.BinaryExpr"},
		{"5 * 6", "ast.BinaryExpr"},
		{"7 / 8", "ast.BinaryExpr"},
		{"9 > 10", "ast.BinaryExpr"},
		{"11 >= 12", "ast.BinaryExpr"},
		{"13 < 14", "ast.BinaryExpr"},
		{"14 <= 15", "ast.BinaryExpr"},
		{"16 == 17", "ast.BinaryExpr"},
		{"18 != 19", "ast.BinaryExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			expr := parser.Expression()

			if expr.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseGroupingExpr(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"(1 + 2)", "ast.GroupingExpr"},
		{"((1 + 2) * (3 / 4))", "ast.GroupingExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			expr := parser.Expression()

			if expr.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseLiteralExpr(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"true", "ast.LiteralExpr"},
		{"false", "ast.LiteralExpr"},
		{"null", "ast.LiteralExpr"},
		{"3.1415", "ast.LiteralExpr"},
		{"\"abc\"", "ast.LiteralExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			expr := parser.Expression()

			if expr.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseVariableExpr(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"foo", "ast.VariableExpr"},
		{"bar", "ast.VariableExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			expr := parser.Expression()

			if expr.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

// TODO: Fill out tests!
func TestParseAssignmentExpr(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		// {},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := scan.NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			expr := parser.Expression()

			if expr.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}
