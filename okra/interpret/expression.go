package interpret

import (
	"fmt"
	"strings"
)

// An Expr groups together source code that can be reduced to a value. In order to allow different
// structs that inherit from Expr to interact with one another, the ExprVisitor design pattern is used.
// TODO: Explain ExprVisitor design pattern better!
type Expr interface {
	accept(ExprVisitor) interface{} // TODO: Explain ExprVisitor design pattern
	String() string                 // Used for parser debugging
}

// TODO: Explain ExprVisitor design pattern
type ExprVisitor interface {
	visitUnaryExpr(UnaryExpr) interface{}
	visitBinaryExpr(BinaryExpr) interface{}
	visitGroupingExpr(GroupingExpr) interface{}
	visitLiteralExpr(LiteralExpr) interface{}
	visitVariableExpr(VariableExpr) interface{}
}

// A UnaryExpr expression is one that applies a single operator to a single operand.
// UnaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type UnaryExpr struct {
	operator Token
	operand  Expr
}

func (u UnaryExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(" + u.operator.lexeme + " " + u.operand.String() + ")")
	return sb.String()
}

func (u UnaryExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitUnaryExpr(u)
}

// A BinaryExpr expression is one that applies a single operator to a multiple operands.
// BinaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type BinaryExpr struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (b BinaryExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(" + b.operator.lexeme + " " + b.leftOperand.String() + " " + b.rightOperand.String() + ")")
	return sb.String()
}

func (b BinaryExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitBinaryExpr(b)
}

// A GroupingExpr sets a higher level of precedence for another expression within its bounds.
// GroupingExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type GroupingExpr struct {
	expression Expr
}

func (g GroupingExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(" + g.expression.String() + ")")
	return sb.String()
}

func (g GroupingExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitGroupingExpr(g)
}

// A LiteralExpr is the most basic expression type and represents a fully evaluated value.
// LiteralExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type LiteralExpr struct {
	val interface{}
}

func (l LiteralExpr) String() string {
	return fmt.Sprintf("%v", l.val)
}

func (l LiteralExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitLiteralExpr(l)
}

type VariableExpr struct {
	identifier Token
}

func (v VariableExpr) String() string {
	return v.identifier.lexeme
}

func (v VariableExpr) accept(visitor ExprVisitor) interface{} {
	return visitor.visitVariableExpr(v)
}
