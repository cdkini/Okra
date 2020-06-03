package interpret

import (
	"fmt"
	"os"
)

type OkraError struct {
	line int
	col  int
	msg  string
}

func reportError(line int, loc string, message string) {
	fmt.Println("[Line", line, "] Error", loc, ":", message)
}

func checkErr(code int, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(code)
	}
}
