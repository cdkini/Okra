package interpret_test

import (
	. "Okra/okra/interpret"
	"testing"
)

func TestParseUnaryExpr(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"-1", "interpret.UnaryExpr"},
		{"!true", "interpret.UnaryExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
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
		{"1 + 2", "interpret.BinaryExpr"},
		{"3 - 4", "interpret.BinaryExpr"},
		{"5 * 6", "interpret.BinaryExpr"},
		{"7 / 8", "interpret.BinaryExpr"},
		{"9 > 10", "interpret.BinaryExpr"},
		{"11 >= 12", "interpret.BinaryExpr"},
		{"13 < 14", "interpret.BinaryExpr"},
		{"14 <= 15", "interpret.BinaryExpr"},
		{"16 == 17", "interpret.BinaryExpr"},
		{"18 != 19", "interpret.BinaryExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
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
		{"(1 + 2)", "interpret.GroupingExpr"},
		{"((1 + 2) * (3 / 4))", "interpret.GroupingExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
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
		{"true", "interpret.LiteralExpr"},
		{"false", "interpret.LiteralExpr"},
		{"null", "interpret.LiteralExpr"},
		{"3.1415", "interpret.LiteralExpr"},
		{"\"abc\"", "interpret.LiteralExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
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
		{"foo", "interpret.VariableExpr"},
		{"bar", "interpret.VariableExpr"},
	}

	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			scanner := NewScanner(test.input)
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
			scanner := NewScanner(test.input)
			scanner.ScanTokens()
			parser := NewParser(scanner.Tokens())
			expr := parser.Expression()

			if expr.GetType() != test.output {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}
