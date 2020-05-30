package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pml/lexer"
	"pml/token"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("You need to provide an input file")
		return
	}

	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Cannot find file " + os.Args[1])
		return
	}

	l := lexer.New(string(file))

	tok := l.GetNextToken()
	for tok.Type != token.EOF {
		fmt.Println(tok)
		tok = l.GetNextToken()
	}
}
