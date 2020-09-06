package interpret

import (
	"testing"

	. "github.com/cdkini/Okra/src/interpreter/ast"
)

func TestInterpretBlockStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretBlockStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}

func TestInterpretExpressionStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretExpressionStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}

func TestInterpretForStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretForStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}

func TestInterpretFuncStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretFuncStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}

func TestInterpretIfStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretIfStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}

func TestInterpretPrintStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretPrintStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}

func TestInterpretReturnStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretReturnStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}

func TestInterpretStructStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretStructStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}

func TestInterpretVariableStmt(t *testing.T) {
	table := []struct {
		input  Stmt
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretVariableStmt", func(t *testing.T) {
			interpreter := NewInterpreter()
			stmt := interpreter.interpretStmt(test.input)

			if stmt != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, stmt)
			}
		})
	}
}
