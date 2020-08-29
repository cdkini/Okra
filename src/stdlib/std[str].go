package stdlib

import (
	"Okra/src/interpreter/interpret"
	"fmt"
)

type StdStr struct{}

func (s *StdStr) Call(i *interpret.Interpreter, args []interface{}) interface{} {
	fmt.Println("Hello, World!")
	return nil
}

func (s *StdStr) Arity() int {
	return 0
}
