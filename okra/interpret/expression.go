package interpret

import (
	"fmt"
)

// An Expr groups together source code that can be reduced to a value. In order to allow different
// structs that inherit from Expr to interact with one another, the ExprVisitor design pattern is used.
// TODO: Explain ExprVisitor design pattern better!
type Expr interface {
	GetType() string // TODO: Implement type for Expr and Stmts
}

// A UnaryExpr expression is one that applies a single operator to a single operand.
// UnaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type UnaryExpr struct {
	operator Token
	operand  Expr
}

func (u UnaryExpr) GetType() string {
	return fmt.Sprintf("%T", u)
}

// A BinaryExpr expression is one that applies a single operator to a multiple operands.
// BinaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type BinaryExpr struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (b BinaryExpr) GetType() string {
	return fmt.Sprintf("%T", b)
}

// A GroupingExpr sets a higher level of precedence for another expression within its bounds.
// GroupingExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type GroupingExpr struct {
	expression Expr
}

func (g GroupingExpr) GetType() string {
	return fmt.Sprintf("%T", g)
}

// A LiteralExpr is the most basic expression type and represents a fully evaluated value.
// LiteralExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type LiteralExpr struct {
	val interface{}
}

func (l LiteralExpr) GetType() string {
	return fmt.Sprintf("%T", l)
}

// TODO: Add docstring
type VariableExpr struct {
	identifier Token
}

func (v VariableExpr) GetType() string {
	return fmt.Sprintf("%T", v)
}

// TODO: Add docstring
type AssignmentExpr struct {
	identifier Token
	val        Expr
}

func (a AssignmentExpr) GetType() string {
	return fmt.Sprintf("%T", a)
}

type LogicalExpr struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (l LogicalExpr) GetType() string {
	return fmt.Sprintf("%T", l)
}

type CallExpr struct {
	callee Expr
	paren  Token
	args   []Expr
}

func (c CallExpr) GetType() string {
	return fmt.Sprintf("%T", c)
}
