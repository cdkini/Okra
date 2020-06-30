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
}

// A UnaryExpr expression is one that applies a single operator to a single operand.
// UnaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type UnaryExpr struct {
	operator Token
	operand  Expr
}

func (ue UnaryExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(" + ue.operator.lexeme + " " + ue.operand.String() + ")")
	return sb.String()
}

func (ue UnaryExpr) accept(ev ExprVisitor) interface{} {
	return ev.visitUnaryExpr(ue)
}

// A BinaryExpr expression is one that applies a single operator to a multiple operands.
// BinaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type BinaryExpr struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (be BinaryExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(" + be.operator.lexeme + " " + be.leftOperand.String() + " " + be.rightOperand.String() + ")")
	return sb.String()
}

func (be BinaryExpr) accept(ve ExprVisitor) interface{} {
	return ve.visitBinaryExpr(be)
}

// A GroupingExpr sets a higher level of precedence for another expression within its bounds.
// GroupingExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type GroupingExpr struct {
	expression Expr
}

func (ge GroupingExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(" + ge.expression.String() + ")")
	return sb.String()
}

func (ge GroupingExpr) accept(ve ExprVisitor) interface{} {
	return ve.visitGroupingExpr(ge)
}

// A LiteralExpr is the most basic expression type and represents a fully evaluated value.
// LiteralExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type LiteralExpr struct {
	val interface{}
}

func (le LiteralExpr) String() string {
	return fmt.Sprintf("%v", le.val)
}

func (le LiteralExpr) accept(ev ExprVisitor) interface{} {
	return ev.visitLiteralExpr(le)
}

type VariableExpr struct {
	identifier Token
}

func (ve VariableExpr) String() string {
	return ve.identifier.lexeme
}

func (ve VariableExpr) accept(ev ExprVisitor) interface{} {

}
