package interpret

type Expr interface {
	accept(ExprVisitor) error
}

type ExprVisitor interface {
	visitUnary(Unary) error
	visitBinary(Binary) error
	visitGrouping(Grouping) error
	visitLiteral(Literal) error
}

type Unary struct {
	operator Token
	operand  Expr
}

func (u Unary) accept(ev ExprVisitor) error {
	return ev.visitUnary(u)
}

type Binary struct {
	leftOperand  Expr
	operator     Token
	rightOperand Expr
}

func (b Binary) accept(ev ExprVisitor) error {
	return ev.visitBinary(b)
}

type Grouping struct {
	expression Expr
}

func (g Grouping) accept(ev ExprVisitor) error {
	return ev.visitGrouping(g)
}

type Literal struct {
	val interface{}
}

func (l Literal) accept(ev ExprVisitor) error {
	return ev.visitLiteral(l)
}
