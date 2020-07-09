package interpret

import (
	"fmt"
)

// An Interpreter takes in a given expression and evaluates it into its most basic literal form.
// Interpreter inherits from the Visitor interface, allowing it interact with all Expr types.
type Interpreter struct {
	stmts       []Stmt
	globalScope Environment
	localScope  Environment
}

func NewInterpreter(stmts []Stmt) *Interpreter {
	return &Interpreter{stmts, *NewEnvironment(), *NewEnvironment()}
}

// TODO: Update docstring after changes from stmt
func (i *Interpreter) Interpret() {
	for _, stmt := range i.stmts {
		stmt.accept(i)
	}
}

func (i *Interpreter) visitVariableStmt(stmt VariableStmt) {
	var val interface{}
	if stmt.expr != nil {
		val = stmt.expr.accept(i)
	}
	i.globalScope.putVar(stmt.identifier.lexeme, val)
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

	i.globalScope.putVar(a.identifier.lexeme, value)
	return value
}

func (i *Interpreter) visitVariableExpr(v VariableExpr) interface{} {
	return i.globalScope.getVar(v.identifier)
}

func (i *Interpreter) visitLiteralExpr(l LiteralExpr) interface{} {
	return l.val
}

func (i *Interpreter) visitGroupingExpr(g GroupingExpr) interface{} {
	return g.expression.accept(i)
}

func (i *Interpreter) visitUnaryExpr(u UnaryExpr) interface{} {
	operand := u.operand.accept(i)

	switch u.operator.tokenType {
	case Minus:
		checkNumericValidity("Runtime Error => \"-\" used on non-numeric operand", operand)
		return -evaluateNumeric(operand)
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
		checkNumericValidity("Runtime Error => \"-\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) - evaluateNumeric(rightOperand)
	case Plus:
		checkNumericValidity("Runtime Error => \"+\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) + evaluateNumeric(rightOperand)
	case Slash:
		checkNumericValidity("Runtime Error => \"/\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) / evaluateNumeric(rightOperand)
	case Star:
		checkNumericValidity("Runtime Error => \"*\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) * evaluateNumeric(rightOperand)
	case Greater:
		checkNumericValidity("Runtime Error => \">\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) > evaluateNumeric(rightOperand)
	case Less:
		checkNumericValidity("Runtime Error => \"<\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) < evaluateNumeric(rightOperand)
	case GreaterEqual:
		checkNumericValidity("Runtime Error => \">=\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) >= evaluateNumeric(rightOperand)
	case LessEqual:
		checkNumericValidity("Runtime Error => \"<=\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) <= evaluateNumeric(rightOperand)
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

func evaluateNumeric(i interface{}) float64 {
	t, ok := i.(float64)
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Placeholder"))
	}
	return t
}

func evaluateString(i interface{}) string {
	t, ok := i.(string)
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Placeholder"))
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
		output = string(val) // FIXME: Printing int32 vals rather than Unicode characters
	case float64:
		output = fmt.Sprintf("%v", val)
	}
	return output
}
