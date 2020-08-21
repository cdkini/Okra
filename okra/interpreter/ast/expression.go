package ast

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
	Operator Token
	Operand  Expr
}

func NewUnaryExpr(operator Token, operand Expr) *UnaryExpr {
	return &UnaryExpr{operator, operand}
}

func (u UnaryExpr) GetType() string {
	return fmt.Sprintf("%T", u)
}

// A BinaryExpr expression is one that applies a single operator to a multiple operands.
// BinaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type BinaryExpr struct {
	LeftOperand  Expr
	Operator     Token
	RightOperand Expr
}

func NewBinaryExpr(leftOperand Expr, operator Token, rightOperand Expr) *BinaryExpr {
	return &BinaryExpr{leftOperand, operator, rightOperand}
}

func (b BinaryExpr) GetType() string {
	return fmt.Sprintf("%T", b)
}

// A GroupingExpr sets a higher level of precedence for another expression within its bounds.
// GroupingExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type GroupingExpr struct {
	Expression Expr
}

func NewGroupingExpr(expression Expr) *GroupingExpr {
	return &GroupingExpr{expression}
}

func (g GroupingExpr) GetType() string {
	return fmt.Sprintf("%T", g)
}

// A LiteralExpr is the most basic expression type and represents a fully evaluated value.
// LiteralExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type LiteralExpr struct {
	Val interface{}
}

func NewLiteralExpr(val interface{}) *LiteralExpr {
	return &LiteralExpr{val}
}

func (l LiteralExpr) GetType() string {
	return fmt.Sprintf("%T", l)
}

// TODO: Add docstring
type VariableExpr struct {
	Identifier Token
}

func NewVariableExpr(identifier Token) *VariableExpr {
	return &VariableExpr{identifier}
}

func (v VariableExpr) GetType() string {
	return fmt.Sprintf("%T", v)
}

// TODO: Add docstring
type AssignmentExpr struct {
	Identifier Token
	Val        Expr
}

func NewAssignmentExpr(identifier Token, val Expr) *AssignmentExpr {
	return &AssignmentExpr{identifier, val}
}

func (a AssignmentExpr) GetType() string {
	return fmt.Sprintf("%T", a)
}

type LogicalExpr struct {
	LeftOperand  Expr
	Operator     Token
	RightOperand Expr
}

func NewLogicalExpr(leftOperand Expr, operator Token, rightOperand Expr) *LogicalExpr {
	return &LogicalExpr{leftOperand, operator, rightOperand}
}

func (l LogicalExpr) GetType() string {
	return fmt.Sprintf("%T", l)
}

type CallExpr struct {
	Callee Expr
	Paren  Token
	Args   []Expr
}

func NewCallExpr(callee Expr, paren Token, args []Expr) *CallExpr {
	return &CallExpr{callee, paren, args}
}

func (c CallExpr) GetType() string {
	return fmt.Sprintf("%T", c)
}
