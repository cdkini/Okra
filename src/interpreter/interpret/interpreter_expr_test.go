package interpret

import (
	"testing"

	. "github.com/cdkini/Okra/src/interpreter/ast"
)

func TestInterpretAssignmentExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretAssignmentExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretBinaryExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretBinaryExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretCallExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretCallExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretGetExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretGetExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretGroupingExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretGroupingExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretLiteralExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretLiteralExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretLogicalExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretLogicalExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretSetExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretSetExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretThisExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretThisExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretUnaryExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretUnaryExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}

func TestInterpretVariableExpr(t *testing.T) {
	table := []struct {
		input  Expr
		output interface{}
	}{
		{},
	}

	for _, test := range table {
		t.Run("TestInterpretVariableExpr", func(t *testing.T) {
			interpreter := NewInterpreter()
			expr := interpreter.interpretExpr(test.input)

			if expr != test.output {
				t.Errorf("Expected %v, recieved %v", test.output, expr)
			}
		})
	}
}
