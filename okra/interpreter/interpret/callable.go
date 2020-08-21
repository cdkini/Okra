package interpret

type Callable interface {
	Arity() int
	Call(i *Interpreter, args []interface{}) interface{}
}
