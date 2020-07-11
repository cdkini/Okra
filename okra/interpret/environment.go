package interpret

// TODO: Add docstring
type Environment struct {
	localScope  map[Token]interface{}
	globalScope map[Token]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{make(map[Token]interface{}), make(map[Token]interface{})}
}

func (e *Environment) putVar(token Token, value interface{}, isLocal bool) {
	if isLocal {
		e.localScope[token] = value
	} else {
		e.globalScope[token] = value
	}
}

func (e *Environment) getVar(token Token) interface{} {
	val1, ok1 := e.localScope[token]
	if !ok1 {
		val2, ok2 := e.globalScope[token]
		if !ok2 {
			ReportErr(-1, NewOkraError(token.col, token.line, "Variable not declared prior to usage"))
		}
		return val2
	}
	return val1
}

func (e *Environment) assignExistingVar(token Token, value interface{}) {
	if _, ok := e.localScope[token]; ok {
		e.localScope[token] = value
		return
	}
	if _, ok2 := e.globalScope[token]; ok2 {
		e.globalScope[token] = value
		return
	}
	ReportErr(-1, NewOkraError(token.col, token.line, "Variable not declared prior to usage"))
}

func (e *Environment) clearLocalScope() {
	for k := range e.localScope {
		delete(e.localScope, k)
	}
}
