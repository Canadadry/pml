package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"pml/lexer"
	"pml/parser"
	"pml/renderer"
	"pml/token"
)

func main() {

	mode := flag.String("mode", "renderer", "mode : lexer|parser|renderer")
	filename := flag.String("in", "", "entry pml filename")
	output := flag.String("out", "out.pdf", "pdf output for renderer mode")

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
	p := parser.New(l)
	r := renderer.New(*output)

	switch *mode {
	case "lexer":
		tok := l.GetNextToken()
		for tok.Type != token.EOF {
			fmt.Println(tok)
			tok = l.GetNextToken()
		}
	case "parser":
		item, err := p.Parse()
		if err != nil {
			fmt.Printf("parsing failed : %v", err)
			return
		}
		fmt.Println(item)
	case "renderer":
		item, err := p.Parse()
		if err != nil {
			fmt.Printf("parsing failed : %v", err)
			return
		}
		err = r.Render(item)
		if err != nil {
			fmt.Printf("rendering failed : %v", err)
			return
		}
	default:
		fmt.Printf("Mode not handle : " + *mode)
		flag.PrintDefaults()
	}
}
