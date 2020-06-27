package interpret

import (
	"fmt"
	"strings"
)

type Expr interface {
	accept(Visitor) error
	String() string
}

type Visitor interface {
	visitUnary(Unary) error
	visitBinary(Binary) error
	visitGrouping(Grouping) error
	visitLiteral(Literal) error
}

type Unary struct {
	operator Token
	operand  Expr
}

func (u Unary) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(u.operator.lexeme)
	sb.WriteString(" ")
	sb.WriteString(u.operand.String())
	sb.WriteString(")")
	return sb.String()
}

func (u Unary) accept(v Visitor) error {
	return v.visitUnary(u)
}

type Binary struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (b Binary) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(b.operator.lexeme)
	sb.WriteString(" ")
	sb.WriteString(b.leftOperand.String())
	sb.WriteString(" ")
	sb.WriteString(b.rightOperand.String())
	sb.WriteString(")")
	return sb.String()
}

func (b Binary) accept(v Visitor) error {
	return v.visitBinary(b)
}

type Grouping struct {
	expression Expr
}

func (g Grouping) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(g.expression.String())
	sb.WriteString(")")
	return sb.String()
}

func (g Grouping) accept(v Visitor) error {
	return v.visitGrouping(g)
}

type Literal struct {
	val interface{}
}

func (l Literal) String() string {
	return fmt.Sprintf("%v", l.val)
}

func (l Literal) accept(v Visitor) error {
	return v.visitLiteral(l)
}
