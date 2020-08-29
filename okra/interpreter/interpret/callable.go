package interpret

import (
	"Okra/okra/interpreter/ast"
)

type Callable interface {
	Arity() int
	Call(i *Interpreter, args []interface{}) interface{}
}

type Function struct {
	declaration ast.FuncStmt
}

func NewFunction(declaration ast.FuncStmt) *Function {
	return &Function{declaration}
}

func (f *Function) Arity() int {
	return len(f.declaration.Params)
}

func (f *Function) Call(i *Interpreter, args []interface{}) interface{} {
	env := NewEnvironment(i.globalEnv)
	for i, token := range f.declaration.Params {
		env.Define(token.Lexeme, args[i])
	}

	block := i.executeBlock(f.declaration.Body, env)
	if r, ok := block.(*ReturnValue); ok {
		return r.LiteralExpr.Val
	}
	return nil
}

type Struct struct {
	// TODO: Open to update struct attributes and methods
}

func NewStruct() *Struct {
	return &Struct{} // TODO: Open to update struct attributes and methods
}

func (s *Struct) Arity() int {
	// TODO: Open to update struct attributes and methods
	return 0
}

func (s *Struct) Call() interface{} {
	// TODO: Open to update struct attributes and methods
	return nil
}
