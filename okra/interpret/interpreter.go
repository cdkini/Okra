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

	// FIXME: Fixed type switch to convert to negative float
	switch t := operand.(type) {
	case float64:
		return -float64(operand)
	}
}

func (i *Interpreter) interpretBinary(b Binary) interface{} {
	return nil // TODO: Implement!
}

func (i *Interpreter) evaluate(e Expr) interface{} {
	return e.accept(e)
}
