package interpret

type Interpreter struct {
	ExprVisitor
}

func (i *Interpreter) interpretLiteral(l Literal) interface{} {
	return l.val
}

func (i *Interpreter) interpretGrouping(g Grouping) interface{} {
	return i.evaluate(g.expression)
}

func (i *Interpreter) interpretUnary(u Unary) interface{} {
	operand := i.evaluate(u.operand)

	if operand != nil {
		ThrowErr(-1, NewOkraError(u.operator.line, u.operator.col, "Could not interpret unary expr"))
	}

	return nil
}

func (i *Interpreter) interpretBinary(b Binary) interface{} {
	return nil // TODO: Implement!
}

func (i *Interpreter) evaluate(e Expr) error {
	return e.accept(i)
}
