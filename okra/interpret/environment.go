package interpret

// TODO: Add docstring
type Environment struct {
	enclosing *Environment
	vars      map[string]interface{}
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing, make(map[string]interface{})}
}

func (e *Environment) define(identifier string, value interface{}) {
	e.vars[identifier] = value
}

func (e *Environment) assign(token Token, value interface{}) {
	if _, ok := e.vars[token.lexeme]; ok {
		e.vars[token.lexeme] = value
	} else if e.enclosing != nil {
		e.enclosing.assign(token, value)
	} else {
		ReportErr(token.col, token.line, "Variable not declared prior to usage")
	}
}

func (e *Environment) get(token Token) interface{} {
	if val, ok := e.vars[token.lexeme]; ok {
		return val
	}
	if e.enclosing != nil {
		return e.enclosing.get(token)
	}
	ReportErr(token.line, token.col, "Undefined variable '"+token.lexeme+"'")
	return nil
}
