package interpret

import (
	"fmt"
)

// An Interpreter takes in a given expression and evaluates it into its most basic literal form.
// Interpreter inherits from the Visitor interface, allowing it interact with all Expr types.
type Interpreter struct {
	stmts        []Stmt
	env          Environment
	inLocalScope bool
}

func NewInterpreter(stmts []Stmt) *Interpreter {
	return &Interpreter{stmts, *NewEnvironment(), false}
}

// TODO: Update docstring after changes from stmt
func (i *Interpreter) Interpret() {
	for _, stmt := range i.stmts {
		stmt.accept(i)
	}
}

func (i *Interpreter) visitVariableStmt(stmt VariableStmt) {
	if stmt.expr == nil {
		i.env.putVar(stmt.identifier.lexeme, nil, i.inLocalScope)
	}
	i.env.putVar(stmt.identifier.lexeme, stmt.expr.accept(i), i.inLocalScope)
}

func (i *Interpreter) visitExpressionStmt(stmt ExpressionStmt) {
	stmt.expr.accept(i)
}

func (i *Interpreter) visitPrintStmt(stmt PrintStmt) {
	value := stmt.expr.accept(i)
	fmt.Println(value)
}

func (i *Interpreter) visitAssignmentExpr(a AssignmentExpr) interface{} {
	value := a.val.accept(i)
	i.env.putVar(a.identifier.lexeme, value, i.inLocalScope)
	return value
}

func (i *Interpreter) visitVariableExpr(v VariableExpr) interface{} {
	return i.env.getVar(v.identifier)
}

func (i *Interpreter) visitLiteralExpr(l LiteralExpr) interface{} {
	if str, ok := l.val.(string); ok {
		return str
	} else if num, ok := l.val.(float64); ok {
		return num
	}
	return nil
}

func (i *Interpreter) visitGroupingExpr(g GroupingExpr) interface{} {
	return g.expression.accept(i)
}

func (i *Interpreter) visitUnaryExpr(u UnaryExpr) interface{} {
	operand := u.operand.accept(i)

	switch u.operator.tokenType {
	case Minus:
		checkNumericValidity("Invalid usage of \"-\" on non-numeric operand", operand)
		return -evalNumeric(operand)
	case Bang:
		return !isTruthy(operand)
	}
	return nil
}

func (i *Interpreter) visitBinaryExpr(b BinaryExpr) interface{} {
	leftOperand := b.leftOperand.accept(i)
	rightOperand := b.rightOperand.accept(i)

	switch b.operator.tokenType {
	case Minus:
		checkNumericValidity("Invalid usage of \"-\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) - evalNumeric(rightOperand)
	case Plus:
		checkNumericValidity("Invalid usage of \"+\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) + evalNumeric(rightOperand)
	case Slash:
		checkNumericValidity("Invalid usage of \"/\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) / evalNumeric(rightOperand)
	case Star:
		checkNumericValidity("Invalid usage of \"*\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) * evalNumeric(rightOperand)
	case Greater:
		checkNumericValidity("Invalid usage of \">\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) > evalNumeric(rightOperand)
	case Less:
		checkNumericValidity("Invalid usage of \"<\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) < evalNumeric(rightOperand)
	case GreaterEqual:
		checkNumericValidity("Invalid usage of \">=\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) >= evalNumeric(rightOperand)
	case LessEqual:
		checkNumericValidity("Invalid usage of \"<=\" on non-numeric operands", leftOperand, rightOperand)
		return evalNumeric(leftOperand) <= evalNumeric(rightOperand)
	case EqualEqual:
		return leftOperand == rightOperand
	case BangEqual:
		return leftOperand != rightOperand
	}
	return nil
}

func isTruthy(i interface{}) bool {
	switch val := i.(type) {
	case nil:
		return false
	case bool:
		return val
	default:
		return true
	}
}

func evalNumeric(i interface{}) float64 {
	t, ok := i.(float64)
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Expect numeric"))
	}
	return t
}

func evalString(i interface{}) string {
	t, ok := i.(string)
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Expect string"))
	}
	return t
}

func checkNumericValidity(msg string, i ...interface{}) {
	for _, n := range i {
		switch n.(type) {
		case float64:
			continue
		default:
			ReportErr(-1, NewOkraError(0, 0, msg))
		}
	}
}

func cleanPrintOutput(input interface{}) string {
	var output string
	switch val := input.(type) {
	case []int32:
		output = string(val)
	case float64:
		output = fmt.Sprintf("%v", val)
	}
	return output
}
