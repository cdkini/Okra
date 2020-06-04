package interpret

import "fmt"

type Expr interface {
	accept(Visitor) error
	String() string // FIXME: Update String methods to work with tests
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

func (u Unary) accept(v Visitor) error {
	return v.visitUnary(u)
}

func (u Unary) String() string {
	return fmt.Sprintf("%v%v", u.operator, u.operand)
}

type Binary struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (b Binary) accept(v Visitor) error {
	return v.visitBinary(b)
}

func (b Binary) String() string {
	return fmt.Sprintf("%v %v %v", b.leftOperand, b.operator, b.rightOperand)
}

type Grouping struct {
	expression Expr
}

func (g Grouping) accept(v Visitor) error {
	return v.visitGrouping(g)
}

func (g Grouping) String() string {
	return fmt.Sprintf("%v", g.expression)
}

type Literal struct {
	val interface{}
}

func (l Literal) accept(v Visitor) error {
	return v.visitLiteral(l)
}

func (l Literal) String() string {
	return fmt.Sprintf("%v", l.val)
}
