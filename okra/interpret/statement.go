package interpret

// TODO: Add docstring
// TODO: Explain Visitor design pattern better!
type Stmt interface {
	accept(StmtVisitor) interface{} // TODO: Explain Visitor design pattern
}

// TODO: Explain Visitor design pattern
type StmtVisitor interface {
	visitExpression(s Stmt)
}
