package interpret

import (
	"Okra/src/interpreter/ast"
	"Okra/src/okraerr"
)

type Callable interface {
	Arity() int
	Call(i *Interpreter, args []interface{}) interface{}
}

type Function struct {
	declaration   ast.FuncStmt
	closure       *Environment
	isConstructor bool
}

func NewFunction(declaration ast.FuncStmt, closure *Environment, isConstructor bool) *Function {
	return &Function{declaration, closure, isConstructor}
}

func (f *Function) Arity() int {
	return len(f.declaration.Params)
}

func (f *Function) Call(i *Interpreter, args []interface{}) interface{} {
	env := NewEnvironment(f.closure)
	for i, token := range f.declaration.Params {
		env.Define(token.Lexeme, args[i])
	}

	block := i.executeBlock(f.declaration.Body, env)
	if r, ok := block.(*ReturnValue); ok {
		return r.LiteralExpr.Val
	}
	if f.isConstructor {
		return f.closure.vals["this"]
	}
	return nil
}

func (f *Function) bind(instance *Instance) *Function {
	env := NewEnvironment(f.closure)
	env.Define("this", instance)
	return NewFunction(f.declaration, env, f.isConstructor)
}

type Struct struct {
	name    string
	methods map[string]*Function
}

func NewStruct(name string, methods map[string]*Function) *Struct {
	return &Struct{name, methods}
}

func (s *Struct) Arity() int {
	if init := s.findMethod("construct"); init != nil {
		return init.Arity()
	}
	return 0
}

func (s *Struct) Call(i *Interpreter, args []interface{}) interface{} {
	instance := NewInstance(*s)
	if init := s.findMethod("construct"); init != nil {
		init.bind(instance).Call(i, args)
	}
	return NewInstance(*s)
}

func (s *Struct) findMethod(method string) *Function {
	if _, ok := s.methods[method]; ok {
		return s.methods[method]
	}
	return nil
}

type Instance struct {
	class  Struct
	fields map[string]interface{}
}

func NewInstance(class Struct) *Instance {
	return &Instance{class, make(map[string]interface{})}
}

func (i *Instance) get(property ast.Token) interface{} {
	if val, ok := i.fields[property.Lexeme]; ok {
		return val
	}
	if method := i.class.findMethod(property.Lexeme); method != nil {
		return method.bind(i)
	}
	okraerr.ReportErr(0, 0, "Undefined property "+property.Lexeme+".")
	return nil
}

func (i *Instance) set(property ast.Token, val interface{}) {
	i.fields[property.Lexeme] = val
}
