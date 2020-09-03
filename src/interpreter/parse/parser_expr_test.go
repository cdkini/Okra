package parse

import (
	. "Okra/src/interpreter/ast"
	"testing"
)

func mock(tokens ...TokenType) []Token {
	var input []Token
	for _, t := range tokens {
		input = append(input, Token{t, "", nil, 0, 0})
	}
	input = append(input, Token{EOF, "", nil, 0, 0})
	return input
}

func TestParseUnaryExpr(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(Minus, Numeric), UnaryExpr{}}, // Test "-1"
		{mock(Bang, True), UnaryExpr{}},     // Test "!true;"
	}

	for _, test := range table {
		t.Run("TestParseUnaryExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseBinaryExpr(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(Numeric, Plus, Numeric), BinaryExpr{}},         // Test "1 + 2"
		{mock(Numeric, Minus, Numeric), BinaryExpr{}},        // Test "3 - 4"
		{mock(Numeric, Star, Numeric), BinaryExpr{}},         // Test "5 * 6"
		{mock(Numeric, Slash, Numeric), BinaryExpr{}},        // Test "7 / 8"
		{mock(Numeric, Greater, Numeric), BinaryExpr{}},      // Test "9 > 10"
		{mock(Numeric, GreaterEqual, Numeric), BinaryExpr{}}, // Test "11 >= 12"
		{mock(Numeric, Less, Numeric), BinaryExpr{}},         // Test "13 < 14"
		{mock(Numeric, LessEqual, Numeric), BinaryExpr{}},    // Test "14 <= 15"
		{mock(Numeric, Equal, Numeric), BinaryExpr{}},        // Test "16 = 17"
		{mock(Numeric, BangEqual, Numeric), BinaryExpr{}},    // Test "18 != 19"
	}

	for _, test := range table {
		t.Run("TestParseBinaryExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseGroupingExpr(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(LeftParen, Numeric, Plus, Numeric, RightParen), GroupingExpr{}}, // Test "(1 + 2)"
		{mock(LeftParen, LeftParen, Numeric, Plus, Numeric, RightParen, Star,
			LeftParen, Numeric, RightParen, RightParen), GroupingExpr{}}, // Test "((1 + 2) * (3))""
	}

	for _, test := range table {
		t.Run("TestParseBinaryExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseLiteralExpr(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(True), LiteralExpr{}},    // Test "true"
		{mock(False), LiteralExpr{}},   // Test "false"
		{mock(Null), LiteralExpr{}},    // Test "null"
		{mock(Numeric), LiteralExpr{}}, // Test "3.1415"
		{mock(String), LiteralExpr{}},  // Test "\"abc\""
	}

	for _, test := range table {
		t.Run("TestParseBinaryExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseVariableExpr(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(Identifier), VariableExpr{}}, // Test "foo"
		{mock(Identifier), VariableExpr{}}, // Test "bar"
	}

	for _, test := range table {
		t.Run("TestParseVariableExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseAssignmentExpr(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(Identifier, Colon, Numeric, Semicolon), AssignmentExpr{}}, // Test "a: 5;"
		{mock(Identifier, Colon, String, Semicolon), AssignmentExpr{}},  // Test "b: \"abc\";"
		{mock(Identifier, Colon, True, Semicolon), AssignmentExpr{}},    // Test "c: true;"
		{mock(Identifier, Colon, Null, Semicolon), AssignmentExpr{}},    // Test "d: null;"
	}

	for _, test := range table {
		t.Run("TestParseVariableExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseLogicalExpr(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(Identifier, And, Identifier, Semicolon), LogicalExpr{}}, // Test "a && b;"
		{mock(Identifier, Or, Identifier, Semicolon), LogicalExpr{}},  // Test "c || d;"
	}

	for _, test := range table {
		t.Run("TestParseVariableExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseCallExpr(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(Identifier, LeftParen, RightParen, Semicolon), CallExpr{}}, // Test "funcName();"
		{mock(Identifier, LeftParen, RightParen, Semicolon), CallExpr{}}, // Test "StructName();"
	}

	for _, test := range table {
		t.Run("TestParseVariableExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}

func TestParseStructExprs(t *testing.T) {
	table := []struct {
		input  []Token
		output Expr
	}{
		{mock(Identifier, Dot, Identifier), GetExpr{}},                    // Test "class.property"
		{mock(This, Dot, Identifier), GetExpr{}},                          // Test "this.value"
		{mock(Identifier, Dot, Identifier, Colon, Identifier), SetExpr{}}, // Test "class.property: newProperty"
		{mock(This, Dot, Identifier, Colon, Identifier), SetExpr{}},       // Test "this.value: value"
		{mock(This), ThisExpr{}},                                          // Test "this"
	}

	for _, test := range table {
		t.Run("TestParseVariableExpr", func(t *testing.T) {
			parser := NewParser(test.input)
			expr := parser.Expression()

			if expr.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, expr)
			}
		})
	}
}
