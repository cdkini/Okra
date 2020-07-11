package interpret

// TODO: Add docstring
type Environment struct {
	varMap map[Token]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{make(map[Token]interface{})}
}

func (e *Environment) putVar(token Token, value interface{}) {
	e.varMap[token] = value
}

func (e *Environment) getVar(token Token) interface{} {
	val, ok := e.varMap[token]
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Placeholder"))
	}
	return val
}

func (e *Environment) assignExistingVar(token Token, value interface{}) {
	if _, ok := e.varMap[token]; ok {
		e.varMap[token] = value
	} else {
		ReportErr(-1, NewOkraError(0, 0, "Placeholder"))
	}
}

func (e *Environment) clear() {
	for k := range e.varMap {
		delete(e.varMap, k)
	}
}
