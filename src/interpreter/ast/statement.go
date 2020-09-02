package ast

import (
	"fmt"
)

// A Stmt groups together source code that results in a side-effect or notable change in program state.
// We use this generic interface in the method signatures of our parser
// and interpreter due to not knowing the specific type until runtime.
type Stmt interface {
	GetType() string
}

// An ExpressionStmt is a wrapper around an Expr. One is only instantiated if the
// evaluated Stmt does not meet the criteria of any other Stmt, making it an Expr by default.
// ExpressionStmt successfully fulfills all of the Stmt interface's methods.
type ExpressionStmt struct {
	Expr Expr
}

func NewExpressionStmt(expr Expr) *ExpressionStmt {
	return &ExpressionStmt{expr}
}

func (e ExpressionStmt) GetType() string {
	return fmt.Sprintf("%T", e)
}

// A PrintSmt encapsulates an Expr that's result is to be displayed to the console or terminal.
// PrintStmt successfully fulfills all of the Stmt interface's methods.
type PrintStmt struct {
	Expr Expr
}

func NewPrintStmt(expr Expr) *PrintStmt {
	return &PrintStmt{expr}
}

func (p PrintStmt) GetType() string {
	return fmt.Sprintf("%T", p)
}

// A VariableStmt covers both the declaration and assignment of a user-defined variable.
// VariableStmt successfully fulfills all of the Stmt interface's methods.
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

// A BlockStmt wraps around a slice of additional Stmt and signifies a change in scoping.
// BlockStmt successfully fulfills all of the Stmt interface's methods.
type BlockStmt struct {
	Stmts []Stmt
}

func NewBlockStmt(stmts []Stmt) *BlockStmt {
	return &BlockStmt{stmts}
}

func (b BlockStmt) GetType() string {
	return fmt.Sprintf("%T", b)
}

// A IfStmt represents a control flow construct based on the evaluation of a boolean condition.
// IfStmt successfully fulfills all of the Stmt interface's methods.
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

// A ForStmt represents a control flow construct that loops based on evaluation of a boolean condition.
// ForStmt successfully fulfills all of the Stmt interface's methods.
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

// A FuncStmt represents the definition of a function or method.
// FuncStmt successfully fulfills all of the Stmt interface's methods.
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

// A ReturnStmt is used to cease function execution and return an evaluated Expr back to the user.
// ReturnStmt successfully fulfills all of the Stmt interface's methods.
type ReturnStmt struct {
	Keyword Token // Will always be 'return'
	Val     Expr
}

func NewReturnStmt(keyword Token, val Expr) *ReturnStmt {
	return &ReturnStmt{keyword, val}
}

func (g ReturnStmt) GetType() string {
	return fmt.Sprintf("%T", g)
}

// A StructStmt represents the definition of a struct.
// Struct successfully fulfills all of the Stmt interface's methods.
type StructStmt struct {
	Name    Token
	Methods []FuncStmt
}

func NewStructStmt(name Token, methods []FuncStmt) *StructStmt {
	return &StructStmt{name, methods}
}

func (s StructStmt) GetType() string {
	return fmt.Sprintf("%T", s)
}
