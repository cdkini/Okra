package interpret

// TODO: Add docstring
type Environment struct {
	enclosing *Environment
	values    map[string]interface{}
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing, make(map[string]interface{})}
}

func (e *Environment) defineVar(identifier string, value interface{}) {
	e.values[identifier] = value
}

func (e *Environment) assignVar(token Token, value interface{}) {
	if e.enclosing != nil {
		e.enclosing.assignVar(token, value)
		return
	}
	if _, ok := e.values[token.lexeme]; !ok {
		ReportErr(NewOkraError(token.col, token.line, "Variable not declared prior to usage"))
	}
	e.values[token.lexeme] = value
}

func (e *Environment) getVar(token Token) interface{} {
	if e.enclosing != nil {
		return e.enclosing.getVar(token)
	}
	val, ok := e.values[token.lexeme]
	if !ok {
		ReportErr(NewOkraError(token.line, token.col, "Undefined variable '"+token.lexeme+"'"))
	}
	return val
}
