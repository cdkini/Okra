package ast

import (
	"fmt"
)

// An Expr groups together source code that can be reduced to a value. We use this generic interface in the method
// signatures of our parser and interpreter due to not knowing the specific type until runtime.
type Expr interface {
	GetType() string
}

// A UnaryExpr is an expression that applies a single operator to a single operand.
// UnaryExpr successfully fulfills all of the Expr interface's methods.
type UnaryExpr struct {
	Operator Token // Either '-' or '!'
	Operand  Expr
}

func NewUnaryExpr(operator Token, operand Expr) *UnaryExpr {
	return &UnaryExpr{operator, operand}
}

func (u UnaryExpr) GetType() string {
	return fmt.Sprintf("%T", u)
}

// A BinaryExpr expression is one that applies a single operator to a multiple operands.
// BinaryExpr successfully fulfills all of the Expr interface's methods.
type BinaryExpr struct {
	LeftOperand  Expr
	Operator     Token // Either '+', '-', '*', or '/'
	RightOperand Expr
}

func NewBinaryExpr(leftOperand Expr, operator Token, rightOperand Expr) *BinaryExpr {
	return &BinaryExpr{leftOperand, operator, rightOperand}
}

func (b BinaryExpr) GetType() string {
	return fmt.Sprintf("%T", b)
}

// A GroupingExpr sets a higher level of precedence for another expression within its bounds.
// GroupingExpr successuflly fulfills all of the Expr interface's methods.
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
// LiteralExpr successuflly fulfills all of the Expr interface's methods.
type LiteralExpr struct {
	Val interface{}
}

func NewLiteralExpr(val interface{}) *LiteralExpr {
	return &LiteralExpr{val}
}

func (l LiteralExpr) GetType() string {
	return fmt.Sprintf("%T", l)
}

// A VariableExpr represents the declaration a user-defined variable.
// VariableExpr successuflly fulfills all of the Expr interface's methods.
type VariableExpr struct {
	Identifier Token
}

func NewVariableExpr(identifier Token) *VariableExpr {
	return &VariableExpr{identifier}
}

func (v VariableExpr) GetType() string {
	return fmt.Sprintf("%T", v)
}

// An AssignmentExpr represents either the assignment of a previously defined variable.
// AssignmentExpr successuflly fulfills all of the Expr interface's methods.
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

// A LogicalExpr encapsulates the control flow constructs of 'AND' and 'OR'.
// LogicalExpr successfully fulfills all of the Expr interface's methods.
type LogicalExpr struct {
	LeftOperand  Expr
	Operator     Token // Either '&&' or '||'
	RightOperand Expr
}

func NewLogicalExpr(leftOperand Expr, operator Token, rightOperand Expr) *LogicalExpr {
	return &LogicalExpr{leftOperand, operator, rightOperand}
}

func (l LogicalExpr) GetType() string {
	return fmt.Sprintf("%T", l)
}

// A CallExpr represents the invokation of a user-defined function or structure.
// CallExpr successfully fulfills all of the Expr interface's methods.
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

// A GetExpr represents the retrieval of an instance variable or attribute of a structure.
// GetExpr successfully fulfills all of the Expr interface's methods.
type GetExpr struct {
	Object   Expr
	Property Token // As represented by this.property
}

func NewGetExpr(object Expr, property Token) *GetExpr {
	return &GetExpr{object, property}
}

func (g GetExpr) GetType() string {
	return fmt.Sprintf("%T", g)
}

// A SetExpr represents the alteration of an instance variable or attribute of a structure.
// SetExpr successfully fulfills all of the Expr interface's methods.
type SetExpr struct {
	Object   Expr
	Property Token
	Val      Expr
}

func NewSetExpr(object Expr, property Token, val Expr) *SetExpr {
	return &SetExpr{object, property, val}
}

func (s SetExpr) GetType() string {
	return fmt.Sprintf("%T", s)
}

// A ThisExpr represents the declaration and assignment of an instance variable
// or attribute within the constructor of a user-defined structure.
// ThisExpr successfully fulfills all of the Expr interface's methods.
type ThisExpr struct {
	Keyword Token // Will always be 'this'
}

func NewThisExpr(keyword Token) *ThisExpr {
	return &ThisExpr{keyword}
}

func (t ThisExpr) GetType() string {
	return fmt.Sprintf("%T", t)
}
