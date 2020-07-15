package interpret_test

import (
	. "Okra/okra/interpret"
	"testing"
)

func TestParseExpressionStmt(t *testing.T) {
	table := []struct {
		id     string
		input  []Token
		output []Stmt
	}{
		{"ExpressionStmt", []Token{}, []Stmt{}},
	}

	for _, test := range table {
		t.Run(test.id, func(t *testing.T) {
			parser := NewParser(test.input)
			stmts := parser.Parse()

			if len(stmts) != len(test.output) {
				t.Errorf("Expected %d stmt, received %d", len(test.output), len(stmts))
			} else {
				for i := range stmts {
					if stmts[i] != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], stmts[i])
					}
				}
			}
		})
	}
}

func TestParsePrintStmt(t *testing.T) {
	table := []struct {
		id     string
		input  []Token
		output []Stmt
	}{
		{"PrintStmt", []Token{}, []Stmt{}},
	}

	for _, test := range table {
		t.Run(test.id, func(t *testing.T) {
			parser := NewParser(test.input)
			stmts := parser.Parse()

			if len(stmts) != len(test.output) {
				t.Errorf("Expected %d stmt, received %d", len(test.output), len(stmts))
			} else {
				for i := range stmts {
					if stmts[i] != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], stmts[i])
					}
				}
			}
		})
	}
}

func TestParseVariableStmt(t *testing.T) {
	table := []struct {
		id     string
		input  []Token
		output []Stmt
	}{
		{"VariableStmt", []Token{}, []Stmt{}},
	}

	for _, test := range table {
		t.Run(test.id, func(t *testing.T) {
			parser := NewParser(test.input)
			stmts := parser.Parse()

			if len(stmts) != len(test.output) {
				t.Errorf("Expected %d stmt, received %d", len(test.output), len(stmts))
			} else {
				for i := range stmts {
					if stmts[i] != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], stmts[i])
					}
				}
			}
		})
	}
}

func TestParseErr(t *testing.T) {
	table := []struct {
		id     string
		input  []Token
		output []Stmt
	}{
		{"Error", []Token{}, []Stmt{}},
	}

	for _, test := range table {
		t.Run(test.id, func(t *testing.T) {
			parser := NewParser(test.input)
			stmts := parser.Parse()

			if len(stmts) != len(test.output) {
				t.Errorf("Expected %d stmt, received %d", len(test.output), len(stmts))
			} else {
				for i := range stmts {
					if stmts[i] != test.output[i] {
						t.Errorf("Expected %v, received %v", test.output[i], stmts[i])
					}
				}
			}
		})
	}
}
