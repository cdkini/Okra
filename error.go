package main

import (
	"fmt"
	"os"
)

func ReportError(line int, loc string, message string) {
	fmt.Println("[line", line, "] Error", loc, ":", message)
}

func Check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(-1)
	}
}
