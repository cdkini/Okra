package interpret

import (
	"fmt"
)

// TODO: Explain Visitor design pattern better!
type Stmt interface {
	accept(StmtVisitor) // TODO: Explain Visitor design pattern
	GetType() string
}

// TODO: Explain Visitor design pattern
type StmtVisitor interface {
	visitExpressionStmt(ExpressionStmt)
	visitPrintStmt(PrintStmt)
	visitVariableStmt(VariableStmt)
}

// TODO: Add docstring
type ExpressionStmt struct {
	expr Expr
}

func (e ExpressionStmt) accept(vst StmtVisitor) {
	vst.visitExpressionStmt(e)
}

func (e ExpressionStmt) GetType() string {
	return fmt.Sprintf("%T", e)
}

// TODO: Add docstring
type PrintStmt struct {
	expr Expr
}

func (p PrintStmt) accept(vst StmtVisitor) {
	vst.visitPrintStmt(p)
}

func (p PrintStmt) GetType() string {
	return fmt.Sprintf("%T", p)
}

// TODO: Add docstring
type VariableStmt struct {
	identifier Token
	expr       Expr
}

func (v VariableStmt) accept(vst StmtVisitor) {
	vst.visitVariableStmt(v)
}

func (v VariableStmt) GetType() string {
	return fmt.Sprintf("%T", v)
}
