package interpret

import (
	"fmt"
	"os"
)

func reportError(line int, loc string, message string) {
	fmt.Println("[Line", line, "] Error", loc, ":", message)
}

func checkErr(code int, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(code)
	}
}
