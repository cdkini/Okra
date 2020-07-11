package interpret

import (
	"reflect"
)

// TODO: Explain Visitor design pattern better!
type Stmt interface {
	accept(StmtVisitor) // TODO: Explain Visitor design pattern
	String() string
	getType() string
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

func (e ExpressionStmt) String() string {
	return e.getType() + ": " + e.expr.String()
}

func (e ExpressionStmt) getType() string {
	return reflect.TypeOf(e).String()
}

// TODO: Add docstring
type PrintStmt struct {
	expr Expr
}

func (p PrintStmt) accept(vst StmtVisitor) {
	vst.visitPrintStmt(p)
}

func (p PrintStmt) String() string {
	return p.getType() + p.getType() + ": " + p.expr.String()
}

func (p PrintStmt) getType() string {
	return reflect.TypeOf(p).String()
}

// TODO: Add docstring
type VariableStmt struct {
	identifier Token
	expr       Expr
}

func (v VariableStmt) accept(vst StmtVisitor) {
	vst.visitVariableStmt(v)
}

func (v VariableStmt) String() string {
	return v.getType() + v.getType() + ": " + v.identifier.String() + " = " + v.expr.String()
}

func (v VariableStmt) getType() string {
	return reflect.TypeOf(v).String()
}
