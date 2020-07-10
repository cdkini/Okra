package interpret

import (
	"fmt"
	"reflect"
)

// An Expr groups together source code that can be reduced to a value. In order to allow different
// structs that inherit from Expr to interact with one another, the ExprVisitor design pattern is used.
// TODO: Explain ExprVisitor design pattern better!
type Expr interface {
	accept(ExprVisitor) interface{} // TODO: Explain ExprVisitor design pattern
	String() string                 // Used for parser debugging
	getType() string                // TODO: Implement type for Expr and Stmts
}

// TODO: Explain ExprVisitor design pattern
type ExprVisitor interface {
	visitUnaryExpr(UnaryExpr) interface{}
	visitBinaryExpr(BinaryExpr) interface{}
	visitGroupingExpr(GroupingExpr) interface{}
	visitLiteralExpr(LiteralExpr) interface{}
	visitVariableExpr(VariableExpr) interface{}
	visitAssignmentExpr(AssignmentExpr) interface{}
}

// A UnaryExpr expression is one that applies a single operator to a single operand.
// UnaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type UnaryExpr struct {
	operator Token
	operand  Expr
}

func (u UnaryExpr) accept(vst ExprVisitor) interface{} {
	return vst.visitUnaryExpr(u)
}

func (u UnaryExpr) String() string {
	return u.getType() + ": (" + u.operator.lexeme + " " + u.operand.String() + ")"
}

func (u UnaryExpr) getType() string {
	return reflect.TypeOf(u).String()
}

// A BinaryExpr expression is one that applies a single operator to a multiple operands.
// BinaryExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type BinaryExpr struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (b BinaryExpr) accept(vst ExprVisitor) interface{} {
	return vst.visitBinaryExpr(b)
}

func (b BinaryExpr) String() string {
	return b.getType() + ": (" + b.operator.lexeme + " " + b.leftOperand.String() + " " + b.rightOperand.String() + ")"
}

func (b BinaryExpr) getType() string {
	return reflect.TypeOf(b).String()
}

// A GroupingExpr sets a higher level of precedence for another expression within its bounds.
// GroupingExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type GroupingExpr struct {
	expression Expr
}

func (g GroupingExpr) accept(vst ExprVisitor) interface{} {
	return vst.visitGroupingExpr(g)
}

func (g GroupingExpr) String() string {
	return g.getType() + ": (" + g.expression.String() + ")"
}

func (g GroupingExpr) getType() string {
	return reflect.TypeOf(g).String()
}

// A LiteralExpr is the most basic expression type and represents a fully evaluated value.
// LiteralExpr inherits from the Expr interface in order to utilize the ExprVisitor design pattern.
type LiteralExpr struct {
	val interface{}
}

func (l LiteralExpr) accept(vst ExprVisitor) interface{} {
	return vst.visitLiteralExpr(l)
}

func (l LiteralExpr) String() string {
	return fmt.Sprintf(l.getType()+": %v", l.val)
}

func (l LiteralExpr) getType() string {
	return reflect.TypeOf(l).String()
}

// TODO: Add docstring
type VariableExpr struct {
	identifier Token
}

func (v VariableExpr) accept(vst ExprVisitor) interface{} {
	return vst.visitVariableExpr(v)
}

func (v VariableExpr) String() string {
	return v.getType() + ": " + v.identifier.lexeme
}

func (v VariableExpr) getType() string {
	return reflect.TypeOf(v).String()
}

// TODO: Add docstring
type AssignmentExpr struct {
	identifier Token
	val        Expr
}

func (a AssignmentExpr) accept(vst ExprVisitor) interface{} {
	return vst.visitAssignmentExpr(a)
}

func (a AssignmentExpr) String() string {
	return a.identifier.lexeme + " = " + a.val.String()
}

func (a AssignmentExpr) getType() string {
	return reflect.TypeOf(a).String()
}
