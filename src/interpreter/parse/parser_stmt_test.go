package parse

import (
	. "Okra/src/interpreter/ast"
	"testing"
)

func TestParseExpressionStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(Numeric, Plus, Numeric, Equal, Numeric, Semicolon), ExpressionStmt{}},                 // Test "1 + 1 = 2;"
		{mock(Identifier, Plus, Identifier, GreaterEqual, Identifier, Semicolon), ExpressionStmt{}}, // Test "a + b >= c;"
		{mock(Bang, Identifier, Semicolon), ExpressionStmt{}},                                       // Test "!d;
		{mock(True, BangEqual, False, Semicolon), ExpressionStmt{}},                                 // Test "true != false;"
	}

	for _, test := range table {
		t.Run("TestParseExpressionStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParsePrintStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(Print, Numeric, Semicolon), PrintStmt{}},    // Test "print 1;"
		{mock(Print, String, Semicolon), PrintStmt{}},     // Test "print \"Hello, World!\";"
		{mock(Print, Identifier, Semicolon), PrintStmt{}}, // Test "print abc;"
	}

	for _, test := range table {
		t.Run("TestParsePrintStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseVariableStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(Variable, Identifier, Colon, Numeric, Semicolon), VariableStmt{}}, // Test "var a: 1;"
		{mock(Variable, Identifier, Colon, String, Semicolon), VariableStmt{}},  // Test "var name: \"Bob\";"
		{mock(Variable, Identifier, Semicolon), VariableStmt{}},                 // Test "var abc;"
	}

	for _, test := range table {
		t.Run("TestParseVariableStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseBlockStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(LeftBrace, RightBrace), BlockStmt{}},                                                  // Test "{}"
		{mock(LeftBrace, Variable, Identifier, Colon, Numeric, Semicolon, RightBrace), BlockStmt{}}, // Test {var x: 1;}
		{mock(LeftBrace, Numeric, Plus, Numeric, Semicolon, RightBrace), BlockStmt{}},               // Test {1 + 1;}
	}

	for _, test := range table {
		t.Run("TestParseBlockStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseIfStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(If, LeftParen, True, RightParen, LeftBrace, RightBrace), IfStmt{}},                      // Test "if (true) {}"
		{mock(If, LeftParen, Identifier, Less, Numeric, RightParen, LeftBrace, RightBrace), IfStmt{}}, // Test "if (x < 5) {}"
		{mock(If, LeftParen, Bang, Identifier, RightParen, LeftBrace, RightBrace), IfStmt{}},          // Test "if (!y) {}"
	}

	for _, test := range table {
		t.Run("TestParseIfStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseForStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(For, LeftParen, True, RightParen, LeftBrace, RightBrace), ForStmt{}},                      // Test "for (true) {}"
		{mock(For, LeftParen, Identifier, Less, Numeric, RightParen, LeftBrace, RightBrace), ForStmt{}}, // Test "for (x < 5) {}"
		{mock(For, LeftParen, Bang, Identifier, RightParen, LeftBrace, RightBrace), ForStmt{}},          // Test "for (!y) {}"
	}

	for _, test := range table {
		t.Run("TestParseForStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseFuncStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(Func, Identifier, Colon, Colon, LeftBrace, RightBrace), FuncStmt{}},                                                           // Test "func abs :: {}"
		{mock(Func, Identifier, Colon, Identifier, Colon, LeftBrace, Print, Identifier, Semicolon, RightBrace), FuncStmt{}},                 // Test "func hello : name : { print name; }"
		{mock(Func, Identifier, Colon, Identifier, Colon, LeftBrace, Return, Identifier, Plus, Numeric, Semicolon, RightBrace), FuncStmt{}}, // Test "func addOne : x : { return x + 1; }"
	}

	for _, test := range table {
		t.Run("TestParseFuncStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseReturnStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(Return, Numeric, Semicolon), ReturnStmt{}}, // Test "return 1;"
		{mock(Return, Semicolon), ReturnStmt{}},          // Test "return;""
	}

	for _, test := range table {
		t.Run("TestParseReturnStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}

func TestParseStructStmt(t *testing.T) {
	table := []struct {
		input  []Token
		output Stmt
	}{
		{mock(Struct, Identifier, LeftBrace, RightBrace), StructStmt{}},                                                  // Test "struct Dog {}"
		{mock(Struct, Identifier, LeftBrace, Identifier, Colon, Colon, LeftBrace, RightBrace, RightBrace), StructStmt{}}, // Test "struct Person { construct :: {} }"
	}

	for _, test := range table {
		t.Run("TestParseStructStmt", func(t *testing.T) {
			parser := NewParser(test.input)
			stmt := parser.Parse()[0]

			if stmt.GetType() != test.output.GetType() {
				t.Errorf("Expected *%v, received %T", test.output, stmt)
			}
		})
	}
}
