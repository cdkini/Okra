package ast

import (
	"fmt"
)

// TODO: Explain Visitor design pattern better!
type Stmt interface {
	GetType() string
}

// TODO: Add docstring
type ExpressionStmt struct {
	Expr Expr
}

func NewExpressionStmt(expr Expr) *ExpressionStmt {
	return &ExpressionStmt{expr}
}

func (e ExpressionStmt) GetType() string {
	return fmt.Sprintf("%T", e)
}

// TODO: Add docstring
type PrintStmt struct {
	Expr Expr
}

func NewPrintStmt(expr Expr) *PrintStmt {
	return &PrintStmt{expr}
}

func (p PrintStmt) GetType() string {
	return fmt.Sprintf("%T", p)
}

// TODO: Add docstring
type VariableStmt struct {
	Identifier Token
	Expr       Expr
}

func NewVariableStmt(identifier Token, expr Expr) *VariableStmt {
	return &VariableStmt{identifier, expr}
}

func (v VariableStmt) GetType() string {
	return fmt.Sprintf("%T", v)
}

type BlockStmt struct {
	Stmts []Stmt
}

func NewBlockStmt(stmts []Stmt) *BlockStmt {
	return &BlockStmt{stmts}
}

func (b BlockStmt) GetType() string {
	return fmt.Sprintf("%T", b)
}

type IfStmt struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

func NewIfStmt(condition Expr, thenBranch Stmt, elseBranch Stmt) *IfStmt {
	return &IfStmt{condition, thenBranch, elseBranch}
}

func (i IfStmt) GetType() string {
	return fmt.Sprintf("%T", i)
}

type ForStmt struct {
	Condition Expr
	Body      Stmt
}

func NewForStmt(condition Expr, body Stmt) *ForStmt {
	return &ForStmt{condition, body}
}

func (f ForStmt) GetType() string {
	return fmt.Sprintf("%T", f)
}

type FuncStmt struct {
	Identifier Token
	Params     []Token
	Body       []Stmt
}

func NewFuncStmt(identifier Token, params []Token, body []Stmt) *FuncStmt {
	return &FuncStmt{identifier, params, body}
}

func (f FuncStmt) GetType() string {
	return fmt.Sprintf("%T", f)
}
