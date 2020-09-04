package interpret

import (
	"Okra/src/interpreter/ast"
	"fmt"
)

// interpretStmt is a helper function used in Interpret that does the brunt of the interpretation of the
// Stmts produced by the parser. The method determines which interpret method to use at runtime based on Stmt type.
func (i *Interpreter) interpretStmt(stmt ast.Stmt) interface{} {
	switch t := stmt.(type) {

	case *ast.BlockStmt:
		return i.interpretBlockStmt(t)

	case *ast.ExpressionStmt:
		return i.interpretExpressionStmt(t)

	case *ast.ForStmt:
		return i.interpretForStmt(t)

	case *ast.FuncStmt:
		return i.interpretFuncStmt(t)

	case *ast.IfStmt:
		return i.interpretIfStmt(t)

	case *ast.PrintStmt:
		return i.interpretPrintStmt(t)

	case *ast.ReturnStmt:
		return i.interpretReturnStmt(t)

	case *ast.StructStmt:
		return i.interpretStructStmt(t)

	case *ast.VariableStmt:
		return i.interpretVariableStmt(t)

	default:
		return nil
	}
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

func (i *Interpreter) interpretExpressionStmt(stmt *ast.ExpressionStmt) interface{} {
	return i.interpretExpr(stmt.Expr)
}

func (i *Interpreter) interpretForStmt(stmt *ast.ForStmt) interface{} {
	for isTruthy(i.interpretExpr(stmt.Condition)) {
		i.interpretStmt(stmt.Body)
	}
	return nil
}

func (i *Interpreter) interpretFuncStmt(stmt *ast.FuncStmt) interface{} {
	function := NewFunction(*stmt, i.env, false)
	i.env.Define(stmt.Identifier.Lexeme, function)
	return nil
}

func (i *Interpreter) interpretIfStmt(stmt *ast.IfStmt) interface{} {
	if isTruthy(i.interpretExpr(stmt.Condition)) {
		i.interpretStmt(stmt.ThenBranch)
	} else if stmt.ElseBranch != nil {
		i.interpretStmt(stmt.ElseBranch)
	}
	return nil
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

func (i *Interpreter) interpretReturnStmt(stmt *ast.ReturnStmt) interface{} {
	var val interface{}
	if stmt.Val != nil {
		val = i.interpretExpr(stmt.Val)
	}
	return NewReturnValue(ast.NewLiteralExpr(val))
}

func (i *Interpreter) interpretStructStmt(stmt *ast.StructStmt) interface{} {
	i.env.Define(stmt.Name.Lexeme, nil)
	methods := make(map[string]*Function)
	for _, method := range stmt.Methods {
		methods[method.Identifier.Lexeme] = NewFunction(method, i.env, method.Identifier.Lexeme == "construct")
	}
	i.env.Assign(stmt.Name, NewStructure(stmt.Name.Lexeme, methods))
	return nil
}

func (i *Interpreter) interpretVariableStmt(stmt *ast.VariableStmt) interface{} {
	var val interface{}
	if stmt.Expr != nil {
		val = i.interpretExpr(stmt.Expr)
	}

	i.env.Define(stmt.Identifier.Lexeme, val)
	return nil
}
