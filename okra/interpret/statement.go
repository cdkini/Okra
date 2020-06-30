package interpret

// TODO: Add docstring
// TODO: Explain Visitor design pattern better!
type Stmt interface {
	accept(StmtVisitor) // TODO: Explain Visitor design pattern
}

// TODO: Explain Visitor design pattern
type StmtVisitor interface {
	visitExpressionStmt(ExpressionStmt)
	visitPrintStmt(PrintStmt)
}

type ExpressionStmt struct {
	expr Expr
}

func (e ExpressionStmt) accept(v StmtVisitor) {
	v.visitExpressionStmt(e)
}

type PrintStmt struct {
	expr Expr
}

func (p PrintStmt) accept(v StmtVisitor) {
	v.visitPrintStmt(p)
}
