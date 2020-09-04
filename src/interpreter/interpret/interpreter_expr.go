package interpret

import (
	"Okra/src/interpreter/ast"
	"Okra/src/okraerr"
	"strconv"
)

// interpretExpr is a helper function used to interpret Expr attributes of Stmt instances and evaluating them
// into returnable values. The method determines which interpret method to use at runtime based on Expr type.
func (i *Interpreter) interpretExpr(expr ast.Expr) interface{} {
	switch t := expr.(type) {

	case *ast.AssignmentExpr:
		return i.interpretAssignmentExpr(t)

	case *ast.BinaryExpr:
		return i.interpretBinaryExpr(t)

	case *ast.CallExpr:
		return i.interpretCallExpr(t)

	case *ast.GetExpr:
		return i.interpretGetExpr(t)

	case *ast.GroupingExpr:
		return i.interpretGroupingExpr(t)

	case *ast.LiteralExpr:
		return i.interpretLiteralExpr(t)

	case *ast.LogicalExpr:
		return i.interpretLogicalExpr(t)

	case *ast.SetExpr:
		return i.interpretSetExpr(t)

	case *ast.ThisExpr:
		return i.interpretThisExpr(t)

	case *ast.UnaryExpr:
		return i.interpretUnaryExpr(t)

	case *ast.VariableExpr:
		return i.interpretVariableExpr(t)

	default:
		return nil
	}
}

func (i *Interpreter) interpretAssignmentExpr(a *ast.AssignmentExpr) interface{} {
	value := i.interpretExpr(a.Val)
	i.env.Assign(a.Identifier, value)
	return value
}

func (i *Interpreter) interpretBinaryExpr(b *ast.BinaryExpr) interface{} {
	leftOperand := i.interpretExpr(b.LeftOperand)
	rightOperand := i.interpretExpr(b.RightOperand)

	switch b.Operator.Type {
	case ast.Minus:
		checkNumericValidity("Invalid usage of \"-\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) - evalNumeric(rightOperand)
	case ast.Plus:
		checkNumericValidity("Invalid usage of \"+\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) + evalNumeric(rightOperand)
	case ast.Slash:
		checkNumericValidity("Invalid usage of \"/\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) / evalNumeric(rightOperand)
	case ast.Star:
		checkNumericValidity("Invalid usage of \"*\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) * evalNumeric(rightOperand)
	case ast.Greater:
		checkNumericValidity("Invalid usage of \">\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) > evalNumeric(rightOperand)
	case ast.Less:
		checkNumericValidity("Invalid usage of \"<\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) < evalNumeric(rightOperand)
	case ast.GreaterEqual:
		checkNumericValidity("Invalid usage of \">=\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) >= evalNumeric(rightOperand)
	case ast.LessEqual:
		checkNumericValidity("Invalid usage of \"<=\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) <= evalNumeric(rightOperand)
	case ast.Equal:
		return leftOperand == rightOperand
	case ast.BangEqual:
		return leftOperand != rightOperand
	}
	return nil
}

func (i *Interpreter) interpretCallExpr(c *ast.CallExpr) interface{} {
	callee := i.interpretExpr(c.Callee)

	var args []interface{}
	for _, arg := range c.Args {
		args = append(args, i.interpretExpr(arg))
	}

	if function, ok := callee.(Callable); ok {
		if len(args) != function.Arity() {
			okraerr.ReportErr(0, 0, "Expected "+strconv.Itoa(function.Arity())+" args but got "+strconv.Itoa(len(args))+".")
		}
		return function.Call(i, args)
	}

	okraerr.ReportErr(0, 0, "Can only call funcs and structs.")
	return nil
}

func (i *Interpreter) interpretLiteralExpr(l *ast.LiteralExpr) interface{} {
	if str, ok := l.Val.(string); ok {
		return str
	} else if num, ok := l.Val.(float64); ok {
		return num
	} else if boolean, ok := l.Val.(bool); ok {
		return boolean
	}
	return l.Val
}

func (i *Interpreter) interpretLogicalExpr(l *ast.LogicalExpr) interface{} {
	left := i.interpretExpr(l.LeftOperand)

	switch l.Operator.Type {
	case ast.Or:
		if isTruthy(left) {
			return left
		}
	case ast.And:
		if !isTruthy(left) {
			return left
		}
	}
	return i.interpretExpr(l.RightOperand)
}

func (i *Interpreter) interpretGetExpr(g *ast.GetExpr) interface{} {
	object := i.interpretExpr(g.Object)
	if instance, ok := object.(*Instance); ok {
		return instance.get(g.Property)
	}
	okraerr.ReportErr(0, 0, "Only struct instances have properties.")
	return nil
}

func (i *Interpreter) interpretGroupingExpr(g *ast.GroupingExpr) interface{} {
	return i.interpretExpr(g.Expression)
}

func (i *Interpreter) interpretSetExpr(s *ast.SetExpr) interface{} {
	object := i.interpretExpr(s.Object)
	if instance, ok := object.(*Instance); !ok {
		okraerr.ReportErr(0, 0, "Only struct instances have properties.")
	} else {
		val := i.interpretStmt(s.Val)
		instance.set(s.Property, val)
		return val
	}
	return nil
}

func (i *Interpreter) interpretThisExpr(t *ast.ThisExpr) interface{} {
	// FIXME: Not currently working; open to check and write tests
	return i.env.Get(t.Keyword)
}

func (i *Interpreter) interpretUnaryExpr(u *ast.UnaryExpr) interface{} {
	operand := i.interpretExpr(u.Operand)

	switch u.Operator.Type {
	case ast.Minus:
		checkNumericValidity("Invalid usage of \"-\" on non-numeric operand", operand)
		return -evalNumeric(operand)
	case ast.Bang:
		return !isTruthy(operand)
	}
	return nil
}

func (i *Interpreter) interpretVariableExpr(v *ast.VariableExpr) interface{} {
	return i.env.Get(v.Identifier)
}
