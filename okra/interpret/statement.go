package interpret

import (
	"fmt"
)

// TODO: Explain Visitor design pattern better!
type Stmt interface {
	GetType() string
}

// TODO: Add docstring
type ExpressionStmt struct {
	expr Expr
}

func (e ExpressionStmt) GetType() string {
	return fmt.Sprintf("%T", e)
}

// TODO: Add docstring
type PrintStmt struct {
	expr Expr
}

func (p PrintStmt) GetType() string {
	return fmt.Sprintf("%T", p)
}

// TODO: Add docstring
type VariableStmt struct {
	identifier Token
	expr       Expr
}

func (v VariableStmt) GetType() string {
	return fmt.Sprintf("%T", v)
}

type BlockStmt struct {
	stmts []Stmt
}

func (b BlockStmt) GetType() string {
	return fmt.Sprintf("%T", b)
}

type IfStmt struct {
	condition  Expr
	thenBranch Stmt
	elseBranch Stmt
}

func (i IfStmt) GetType() string {
	return fmt.Sprintf("%T", i)
}

type ForStmt struct {
	condition Expr
	body      Stmt
}

func (f ForStmt) GetType() string {
	return fmt.Sprintf("%T", f)
}
