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
	name string
}

func NewStruct(name string) *Struct {
	return &Struct{name}
}

func (s *Struct) Arity() int {
	return 0
}

func (s *Struct) Call(i *Interpreter, args []interface{}) interface{} {
	return NewInstance(*s)

}

type Instance struct {
	class Struct
}

func NewInstance(class Struct) *Instance {
	return &Instance{class}
}
