package interpret

// TODO: Add docstring
// TODO: Explain Visitor design pattern better!
type Stmt interface {
	accept(StmtVisitor) // TODO: Explain Visitor design pattern
}

// TODO: Explain Visitor design pattern
type StmtVisitor interface {
	visitExpressionStmt(Expression)
}

type Expression struct {
	expr Expr
}

func (e Expression) accept(v StmtVisitor) {
	v.visitExpressionStmt(e)
}
