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
	class string
	line  int
	col   int
	msg   string
}

func (e OkraError) Error() string {
	return fmt.Sprintf("%s [%d:%d]: %s", e.class, e.line, e.col, e.msg)
}

func CheckErr(code int, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(code)
	}
}
