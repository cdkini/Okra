package interpret

import (
	"fmt"
)

// An Interpreter takes in a given expression and evaluates it into its most basic literal form.
// Interpreter inherits from the Visitor interface, allowing it interact with all Expr types.
type Interpreter struct {
	stmts  []Stmt
	global *Environment
	env    *Environment
}

func NewInterpreter(stmts []Stmt) *Interpreter {
	// TODO: Open to add standard library methods as part of global
	return &Interpreter{stmts, NewEnvironment(nil), NewEnvironment(nil)}
}

// TODO: Update docstring after changes from stmt
func (i *Interpreter) Interpret() {
	for _, stmt := range i.stmts {
		i.evalStmt(stmt)
	}
}

func (i *Interpreter) evalStmt(stmt Stmt) {
	switch t := stmt.(type) {
	case *ExpressionStmt:
		i.interpretExpressionStmt(t)
	case *BlockStmt:
		i.interpretBlockStmt(t)
	case *PrintStmt:
		i.interpretPrintStmt(t)
	case *VariableStmt:
		i.interpretVariableStmt(t)
	case *IfStmt:
		i.interpretIfStmt(t)
	case *ForStmt:
		i.interpretForStmt(t)
	default:
		fmt.Println(t.GetType())
	}
}

func (i *Interpreter) evalExpr(expr Expr) interface{} {
	switch t := expr.(type) {
	case *AssignmentExpr:
		return i.interpretAssignmentExpr(t)
	case *BinaryExpr:
		return i.interpretBinaryExpr(t)
	case *GroupingExpr:
		return i.interpretGroupingExpr(t)
	case *LiteralExpr:
		return i.interpretLiteralExpr(t)
	case *UnaryExpr:
		return i.interpretUnaryExpr(t)
	case *VariableExpr:
		return i.interpretVariableExpr(t)
	case *LogicalExpr:
		return i.interpretLogicalExpr(t)
	default:
		return nil
	}
}

func (i *Interpreter) interpretIfStmt(stmt *IfStmt) {
	if isTruthy(i.evalExpr(stmt.condition)) {
		i.evalStmt(stmt.thenBranch)
	} else if stmt.elseBranch != nil {
		i.evalStmt(stmt.elseBranch)
	}
}

func (i *Interpreter) interpretForStmt(stmt *ForStmt) {
	for isTruthy(i.evalExpr(stmt.condition)) {
		i.evalStmt(stmt.body)
	}
}

func (i *Interpreter) interpretBlockStmt(stmt *BlockStmt) {
	prevEnv := i.env

	i.env = NewEnvironment(prevEnv)
	defer func() { i.env = prevEnv }()

	for _, s := range stmt.stmts {
		i.evalStmt(s)
	}
}

func (i *Interpreter) interpretVariableStmt(stmt *VariableStmt) {
	var val interface{}
	if stmt.expr != nil {
		val = i.evalExpr(stmt.expr)
	}

	i.env.define(stmt.identifier.lexeme, val)
}

func (i *Interpreter) interpretExpressionStmt(stmt *ExpressionStmt) {
	i.evalExpr(stmt.expr)
}

func (i *Interpreter) interpretPrintStmt(stmt *PrintStmt) {
	value := i.evalExpr(stmt.expr)
	fmt.Println(value)
}

func (i *Interpreter) interpretLogicalExpr(l *LogicalExpr) interface{} {
	left := i.evalExpr(l.leftOperand)

	switch l.operator.tokenType {
	case Or:
		if isTruthy(left) {
			return left
		}
	case And:
		if !isTruthy(left) {
			return left
		}
	}
	return i.evalExpr(l.rightOperand)
}

func (i *Interpreter) interpretAssignmentExpr(a *AssignmentExpr) interface{} {
	value := i.evalExpr(a.val)
	i.env.assign(a.identifier, value)
	return value
}

func (i *Interpreter) interpretVariableExpr(v *VariableExpr) interface{} {
	return i.env.get(v.identifier)
}

func (i *Interpreter) interpretLiteralExpr(l *LiteralExpr) interface{} {
	if str, ok := l.val.(string); ok {
		return str
	} else if num, ok := l.val.(float64); ok {
		return num
	}
	return nil
}

func (i *Interpreter) interpretGroupingExpr(g *GroupingExpr) interface{} {
	return i.evalExpr(g.expression)
}

func (i *Interpreter) interpretUnaryExpr(u *UnaryExpr) interface{} {
	operand := i.evalExpr(u.operand)

	switch u.operator.tokenType {
	case Minus:
		checkNumericValidity("Invalid usage of \"-\" on non-numeric operand", operand)
		return -evalNumeric(operand)
	case Bang:
		return !isTruthy(operand)
	}
	return nil
}

func (i *Interpreter) interpretBinaryExpr(b *BinaryExpr) interface{} {
	leftOperand := i.evalExpr(b.leftOperand)
	rightOperand := i.evalExpr(b.rightOperand)

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

func (i *Interpreter) interpretCallExpr(c *CallExpr) interface{} {
	callee := i.evalExpr(c.callee)

	var args []interface{}
	for _, arg := range c.args {
		args = append(args, i.evalExpr(arg))
	}

	if function, ok := callee.(Callable); !ok {
		if len(args) != function.arity() {
			ReportErr(0, 0, "Expected "+string(function.arity())+" args but got "+string(len(args))+".")
		}
		return function.call(i, args)
	}

	ReportErr(0, 0, "Can only call funcs and structs.")
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
		ReportErr(0, 0, "Expect numeric")
	}
	return t
}

func evalString(i interface{}) string {
	t, ok := i.(string)
	if !ok {
		ReportErr(0, 0, "Expect string")
	}
	return t
}

func checkNumericValidity(msg string, i ...interface{}) {
	for _, n := range i {
		switch n.(type) {
		case float64:
			continue
		default:
			ReportErr(0, 0, msg)
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
