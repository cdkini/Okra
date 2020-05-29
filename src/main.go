package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cdkini/Okra/interpret"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".okr") {
		fmt.Println("Error: Must use \"okra [script]\" to run a .okr file")
		os.Exit(-1)
	}
	runFile(os.Args[1])
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	interpret.CheckErr(-1, err)
	scanner := interpret.NewScanner(string(bytes))
	tokens := scanner.scanTokens()
	fmt.Println(tokens)
}