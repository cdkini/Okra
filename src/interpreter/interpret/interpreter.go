package interpret

import (
	"Okra/src/interpreter/ast"
	"Okra/src/okraerr"
	"fmt"
	"strconv"
)

// An Interpreter takes in a given expression and evaluates it into its most basic literal form.
// Interpreter inherits from the Visitor interface, allowing it interact with all Expr types.
type Interpreter struct {
	env     *Environment
	globals *Environment
}

func NewInterpreter() *Interpreter {
	// TODO: Open to add standard library methods as part of global
	return &Interpreter{NewEnvironment(nil), NewEnvironment(nil)}
}

// func (i *Interpreter) LoadStdlib(stdlib map[string]Callable) {
// 	for k, v := range stdlib {
// 		i.globalEnv.Define(k, v)
// 	}
// }

// TODO: Update docstring after changes from stmt
func (i *Interpreter) Interpret(stmts []ast.Stmt) {
	for _, stmt := range stmts {
		i.interpretStmt(stmt)
	}
}

func (i *Interpreter) interpretStmt(stmt ast.Stmt) interface{} {
	switch t := stmt.(type) {
	case *ast.ExpressionStmt:
		return i.interpretExpressionStmt(t)
	case *ast.FuncStmt:
		return i.interpretFuncStmt(t)
	case *ast.StructStmt:
		return i.interpretStructStmt(t)
	case *ast.ReturnStmt:
		return i.interpretReturnStmt(t)
	case *ast.BlockStmt:
		return i.interpretBlockStmt(t)
	case *ast.PrintStmt:
		return i.interpretPrintStmt(t)
	case *ast.VariableStmt:
		return i.interpretVariableStmt(t)
	case *ast.IfStmt:
		return i.interpretIfStmt(t)
	case *ast.ForStmt:
		return i.interpretForStmt(t)

	default:
		return nil
	}
}

func (i *Interpreter) interpretExpr(expr ast.Expr) interface{} {
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
	case *ast.GetExpr:
		return i.interpretGetExpr(t)
	case *ast.SetExpr:
		return i.interpretSetExpr(t)
	case *ast.ThisExpr:
		return i.interpretThisExpr(t)

	default:
		return nil
	}
}

func (i *Interpreter) interpretIfStmt(stmt *ast.IfStmt) interface{} {
	if isTruthy(i.interpretExpr(stmt.Condition)) {
		i.interpretStmt(stmt.ThenBranch)
	} else if stmt.ElseBranch != nil {
		i.interpretStmt(stmt.ElseBranch)
	}
	return nil
}

func (i *Interpreter) interpretForStmt(stmt *ast.ForStmt) interface{} {
	for isTruthy(i.interpretExpr(stmt.Condition)) {
		i.interpretStmt(stmt.Body)
	}
	return nil
}

func (i *Interpreter) interpretBlockStmt(stmt *ast.BlockStmt) interface{} {
	i.executeBlock(stmt.Stmts, NewEnvironment(i.env))
	return nil
}

func (i *Interpreter) executeBlock(stmts []ast.Stmt, env *Environment) interface{} {
	prevEnv := i.env
	i.env = env

	defer func() { i.env = prevEnv }()
	var eval interface{}
	for _, stmt := range stmts {
		if eval = i.interpretStmt(stmt); eval != nil {
			break
		}
	}
	return eval
}

func (i *Interpreter) interpretVariableStmt(stmt *ast.VariableStmt) interface{} {
	var val interface{}
	if stmt.Expr != nil {
		val = i.interpretExpr(stmt.Expr)
	}

	i.env.Define(stmt.Identifier.Lexeme, val)
	return nil
}

func (i *Interpreter) interpretExpressionStmt(stmt *ast.ExpressionStmt) interface{} {
	return i.interpretExpr(stmt.Expr)
}

func (i *Interpreter) interpretFuncStmt(stmt *ast.FuncStmt) interface{} {
	function := NewFunction(*stmt, i.env, false)
	i.env.Define(stmt.Identifier.Lexeme, function)
	return nil
}

func (i *Interpreter) interpretStructStmt(stmt *ast.StructStmt) interface{} {
	i.env.Define(stmt.Name.Lexeme, nil)
	methods := make(map[string]*Function)
	for _, method := range stmt.Methods {
		methods[method.Identifier.Lexeme] = NewFunction(method, i.env, method.Identifier.Lexeme == "construct")
	}
	i.env.Assign(stmt.Name, NewStruct(stmt.Name.Lexeme, methods))
	return nil
}

func (i *Interpreter) interpretReturnStmt(stmt *ast.ReturnStmt) interface{} {
	var val interface{}
	if stmt.Val != nil {
		val = i.interpretExpr(stmt.Val)
	}
	return NewReturnValue(ast.NewLiteralExpr(val))
}

func (i *Interpreter) interpretPrintStmt(stmt *ast.PrintStmt) interface{} {
	value := i.interpretExpr(stmt.Expr)
	switch v := value.(type) {
	case nil:
		fmt.Println("null")
	default:
		fmt.Printf("%v\n", v)
	}
	return nil
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

func (i *Interpreter) interpretAssignmentExpr(a *ast.AssignmentExpr) interface{} {
	value := i.interpretExpr(a.Val)
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
	} else if boolean, ok := l.Val.(bool); ok {
		return boolean
	}
	return l.Val
}

func (i *Interpreter) interpretGroupingExpr(g *ast.GroupingExpr) interface{} {
	return i.interpretExpr(g.Expression)
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

func (i *Interpreter) interpretGetExpr(g *ast.GetExpr) interface{} {
	object := i.interpretExpr(g.Object)
	if instance, ok := object.(*Instance); ok {
		return instance.get(g.Property)
	}
	okraerr.ReportErr(0, 0, "Only struct instances have properties.")
	return nil
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
	// TODO: Not currently working; open to check and write tests
	return i.env.Get(t.Keyword)
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
