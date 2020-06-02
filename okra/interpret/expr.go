package interpret

type Expr interface {
	accept(Visitor) error
}

type Visitor interface {
	visitUnary(Unary) error
	visitBinary(Binary) error
	visitGrouping(Grouping) error
	visitLiteral(Literal) error
}

type Unary struct {
	operator Token
	right    Expr
}

func (u Unary) accept(v Visitor) error {
	return v.visitUnary(u)
}

type Binary struct {
	left     Expr
	operator Token
	right    Expr
}

func (b Binary) accept(v Visitor) error {
	return v.visitBinary(b)
}

type Grouping struct {
	expression Expr
}

func (g Grouping) accept(v Visitor) error {
	return v.visitGrouping(g)
}

type Literal struct {
	val interface{}
}

func (l Literal) accept(v Visitor) error {
	return v.visitLiteral(l)
}
