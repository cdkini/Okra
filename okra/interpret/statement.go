package interpret

// TODO: Add docstring
// TODO: Explain Visitor design pattern better!
type Stmt interface {
	accept(StmtVisitor) interface{} // TODO: Explain Visitor design pattern
}

// TODO: Explain Visitor design pattern
type StmtVisitor interface {
	visitExpressionStmt(Expression) interface{}
}

type Expression struct {
	expr Expr
}

func (e Expression) accept(v StmtVisitor) interface{} {
	return v.visitExpressionStmt(e)
}
