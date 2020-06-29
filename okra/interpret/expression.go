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
	visitUnary(Unary) interface{}
	visitBinary(Binary) interface{}
	visitGrouping(Grouping) interface{}
	visitLiteral(Literal) interface{}
}

// A Unary expression is one that applies a single operator to a single operand.
// Unary inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type Unary struct {
	operator Token
	operand  Expr
}

func (u Unary) String() string {
	var sb strings.Builder
	sb.WriteString("(" + u.operator.lexeme + " " + u.operand.String() + ")")
	return sb.String()
}

func (u Unary) accept(v ExprVisitor) interface{} {
	return v.visitUnary(u)
}

// A Binary expression is one that applies a single operator to a multiple operands.
// Binary inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type Binary struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (b Binary) String() string {
	var sb strings.Builder
	sb.WriteString("(" + b.operator.lexeme + " " + b.leftOperand.String() + " " + b.rightOperand.String() + ")")
	return sb.String()
}

func (b Binary) accept(v ExprVisitor) interface{} {
	return v.visitBinary(b)
}

// A Grouping sets a higher level of precedence for another expression within its bounds.
// Grouping inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type Grouping struct {
	expression Expr
}

func (g Grouping) String() string {
	var sb strings.Builder
	sb.WriteString("(" + g.expression.String() + ")")
	return sb.String()
}

func (g Grouping) accept(v ExprVisitor) interface{} {
	return v.visitGrouping(g)
}

// A Literal is the most basic expression type and represents a fully evaluated value.
// Literal inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type Literal struct {
	val interface{}
}

func (l Literal) String() string {
	return fmt.Sprintf("%v", l.val)
}

func (l Literal) accept(v ExprVisitor) interface{} {
	return v.visitLiteral(l)
}
