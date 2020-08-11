package interpret

// TODO: Add docstring
type Environment struct {
	enclosing *Environment
	vars      map[string]interface{}
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing, make(map[string]interface{})}
}

func (e *Environment) defineVar(identifier string, value interface{}) {
	e.vars[identifier] = value
}

func (e *Environment) assignVar(token Token, value interface{}) {
	if e.enclosing != nil {
		e.enclosing.assignVar(token, value)
		return
	}
	if _, ok := e.vars[token.lexeme]; !ok {
		ReportErr(token.col, token.line, "Variable not declared prior to usage")
	}
	e.vars[token.lexeme] = value
}

func (e *Environment) getVar(token Token) interface{} {
	if e.enclosing != nil {
		return e.enclosing.getVar(token)
	}
	val, ok := e.vars[token.lexeme]
	if !ok {
		ReportErr(token.line, token.col, "Undefined variable '"+token.lexeme+"'")
	}
	return val
}
