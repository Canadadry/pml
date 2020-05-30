package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"pml/lexer"
	"pml/parser"
	"pml/token"
)

func main() {

	mode := flag.String("mode", "lexer", "mode : lexer|parser")
	filename := flag.String("filename", "", "entry filename")

	flag.Parse()

	if len(*filename) == 0 {
		flag.PrintDefaults()
		return
	}

	file, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println("Cannot find file " + *filename)
		return
	}

	l := lexer.New(string(file))

	switch *mode {
	case "lexer":
		tok := l.GetNextToken()
		for tok.Type != token.EOF {
			fmt.Println(tok)
			tok = l.GetNextToken()
		}
	case "parser":
		p := parser.New(l)
		item, err := p.Parse()
		if err != nil {
			fmt.Printf("failed : %v", err)
			return
		}
		fmt.Println(item)
	default:
		fmt.Printf("Mode not handle : " + *mode)
		flag.PrintDefaults()
	}
}
