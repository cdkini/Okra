package interpret

import (
	"fmt"
	"os"
)

/*
   TODO:
   SyntaxError
   RuntimeError
*/

type OkraError struct {
	line int
	col  int
	msg  string
}

func NewOkraError(line int, col int, msg string) OkraError {
	return OkraError{line, col, msg}
}

func (e OkraError) Error() string {
	return fmt.Sprintf("OkraError [%d:%d]: %s", e.line, e.col, e.msg)
}

func CheckErr(code int, err error, oe OkraError) {
	if err != nil {
		ThrowErr(code, oe)
	}
}

func ThrowErr(code int, oe OkraError) {
	fmt.Println(oe.Error())
	os.Exit(code)
}
