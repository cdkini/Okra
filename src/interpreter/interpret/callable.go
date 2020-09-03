package interpret

// A Callable is a data object that has the capacity to be invoked. Per Okra's definition, this is applicable only to
// functions, methods, and structure instances. Invokation requires usage of 'object()' per the parser and interpreter.
type Callable interface {
	Arity() int // Returns the number of arguments or parameters
	Call(i *Interpreter, args []interface{}) interface{}
}
