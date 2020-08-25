package interpret

import (
	"Okra/okra/interpreter/ast"
	"Okra/okra/interpreter/env"
	"Okra/okra/okraerr"
	"fmt"
	"strconv"
)

// An Interpreter takes in a given expression and evaluates it into its most basic literal form.
// Interpreter inherits from the Visitor interface, allowing it interact with all Expr types.
type Interpreter struct {
	stmts  []ast.Stmt
	global *env.Environment
	env    *env.Environment
}

func NewInterpreter(stmts []ast.Stmt) *Interpreter {
	// TODO: Open to add standard library methods as part of global
	return &Interpreter{stmts, env.NewEnvironment(nil), env.NewEnvironment(nil)}
}

func (i *Interpreter) LoadStdlib(stdlib map[string]Callable) {
	for k, v := range stdlib {
		i.global.Define(k, v)
	}
}

// TODO: Update docstring after changes from stmt
func (i *Interpreter) Interpret() {
	for _, stmt := range i.stmts {
		i.evalStmt(stmt)
	}
}

func (i *Interpreter) evalStmt(stmt ast.Stmt) {
	switch t := stmt.(type) {
	case *ast.ExpressionStmt:
		i.interpretExpressionStmt(t)
	case *ast.FuncStmt:
		i.interpretFuncStmt(t)
	case *ast.BlockStmt:
		i.interpretBlockStmt(t)
	case *ast.PrintStmt:
		i.interpretPrintStmt(t)
	case *ast.VariableStmt:
		i.interpretVariableStmt(t)
	case *ast.IfStmt:
		i.interpretIfStmt(t)
	case *ast.ForStmt:
		i.interpretForStmt(t)

	default:
		fmt.Println(t.GetType())
	}
}

func (i *Interpreter) evalExpr(expr ast.Expr) interface{} {
	switch t := expr.(type) {
	case *ast.AssignmentExpr:
		return i.interpretAssignmentExpr(t)
	case *ast.BinaryExpr:
		return i.interpretBinaryExpr(t)
	case *ast.GroupingExpr:
		return i.interpretGroupingExpr(t)
	case *ast.LiteralExpr:
		return i.interpretLiteralExpr(t)
	case *ast.UnaryExpr:
		return i.interpretUnaryExpr(t)
	case *ast.VariableExpr:
		return i.interpretVariableExpr(t)
	case *ast.LogicalExpr:
		return i.interpretLogicalExpr(t)
	case *ast.CallExpr:
		return i.interpretCallExpr(t)
	default:
		return nil
	}
}

func (i *Interpreter) interpretIfStmt(stmt *ast.IfStmt) {
	if isTruthy(i.evalExpr(stmt.Condition)) {
		i.evalStmt(stmt.ThenBranch)
	} else if stmt.ElseBranch != nil {
		i.evalStmt(stmt.ElseBranch)
	}
}

func (i *Interpreter) interpretForStmt(stmt *ast.ForStmt) {
	for isTruthy(i.evalExpr(stmt.Condition)) {
		i.evalStmt(stmt.Body)
	}
}

func (i *Interpreter) interpretBlockStmt(stmt *ast.BlockStmt) {
	i.executeBlock(stmt.Stmts, env.NewEnvironment(i.env))
}

func (i *Interpreter) executeBlock(stmts []ast.Stmt, env *env.Environment) {
	prevEnv := i.env

	i.env = env
	defer func() { i.env = prevEnv }()

	for _, s := range stmts {
		i.evalStmt(s)
	}
}

func (i *Interpreter) interpretVariableStmt(stmt *ast.VariableStmt) {
	var val interface{}
	if stmt.Expr != nil {
		val = i.evalExpr(stmt.Expr)
	}

	i.env.Define(stmt.Identifier.Lexeme, val)
}

func (i *Interpreter) interpretExpressionStmt(stmt *ast.ExpressionStmt) {
	i.evalExpr(stmt.Expr)
}

func (i *Interpreter) interpretFuncStmt(stmt *ast.FuncStmt) {
	function := NewFunction(*stmt)
	i.env.Define(stmt.Identifier.Lexeme, function)
}

func (i *Interpreter) interpretPrintStmt(stmt *ast.PrintStmt) {
	value := i.evalExpr(stmt.Expr)
	fmt.Println(value)
}

func (i *Interpreter) interpretLogicalExpr(l *ast.LogicalExpr) interface{} {
	left := i.evalExpr(l.LeftOperand)

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
	return i.evalExpr(l.RightOperand)
}

func (i *Interpreter) interpretAssignmentExpr(a *ast.AssignmentExpr) interface{} {
	value := i.evalExpr(a.Val)
	i.env.Assign(a.Identifier, value)
	return value
}

func (i *Interpreter) interpretVariableExpr(v *ast.VariableExpr) interface{} {
	return i.env.Get(v.Identifier)
}

func (i *Interpreter) interpretLiteralExpr(l *ast.LiteralExpr) interface{} {
	if str, ok := l.Val.(string); ok {
		return str
	} else if num, ok := l.Val.(float64); ok {
		return num
	}
	return nil
}

func (i *Interpreter) interpretGroupingExpr(g *ast.GroupingExpr) interface{} {
	return i.evalExpr(g.Expression)
}

func (i *Interpreter) interpretUnaryExpr(u *ast.UnaryExpr) interface{} {
	operand := i.evalExpr(u.Operand)

	switch u.Operator.Type {
	case ast.Minus:
		checkNumericValidity("Invalid usage of \"-\" on non-numeric operand", operand)
		return -evalNumeric(operand)
	case ast.Bang:
		return !isTruthy(operand)
	}
	return nil
}

func (i *Interpreter) interpretBinaryExpr(b *ast.BinaryExpr) interface{} {
	leftOperand := i.evalExpr(b.LeftOperand)
	rightOperand := i.evalExpr(b.RightOperand)

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
	case ast.EqualEqual:
		return leftOperand == rightOperand
	case ast.BangEqual:
		return leftOperand != rightOperand
	}
	return nil
}

func (i *Interpreter) interpretCallExpr(c *ast.CallExpr) interface{} {
	callee := i.evalExpr(c.Callee)

	var args []interface{}
	for _, arg := range c.Args {
		args = append(args, i.evalExpr(arg))
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
		okraerr.ReportErr(0, 0, "Expect numeric")
	}
	return t
}

func evalString(i interface{}) string {
	t, ok := i.(string)
	if !ok {
		okraerr.ReportErr(0, 0, "Expect string")
	}
	return t
}

func checkNumericValidity(msg string, i ...interface{}) {
	for _, n := range i {
		switch n.(type) {
		case float64:
			continue
		default:
			okraerr.ReportErr(0, 0, msg)
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
