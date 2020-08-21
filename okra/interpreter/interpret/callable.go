package interpret

type Callable interface {
	arity() int
	call(i *Interpreter, args []interface{}) interface{}
}
