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
	visitVariableStmt(VariableStmt)
}

type ExpressionStmt struct {
	expr Expr
}

func (e ExpressionStmt) accept(vst StmtVisitor) {
	vst.visitExpressionStmt(e)
}

type PrintStmt struct {
	expr Expr
}

func (p PrintStmt) accept(vst StmtVisitor) {
	vst.visitPrintStmt(p)
}

type VariableStmt struct {
	identifier Token
	expr       Expr
}

func (v VariableStmt) accept(vst StmtVisitor) {
	vst.visitVariableStmt(v)
}
