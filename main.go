package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Must use \"okra [script]\" to run a .okr file")
		os.Exit(-1)
	}
	runFile(os.Args[1])
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	check(err)
	run(string(bytes))
}

func run(source string) {
	// scanner := &Scanner{}
	// tokens := scanner.scanTokens()

	// for i, token := range tokens {
	// 	fmt.Println(i, token)
	// }
}
