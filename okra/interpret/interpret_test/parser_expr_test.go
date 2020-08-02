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
		{}, // TODO: Fill out tests!
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
		{}, // TODO: Fill out tests!
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
		{}, // TODO: Fill out tests!
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
		{}, // TODO: Fill out tests!
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

func TestParseAssignmentExpr(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{}, // TODO: Fill out tests!
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
