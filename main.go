package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"pml/lexer"
	"pml/parser"
	"pml/renderer"
	"pml/template"
	"pml/token"
)

func main() {

	mode := flag.String("mode", "full", "mode : lexer|parser|renderer|template|full")
	filename := flag.String("in", "", "entry pml filename")
	paramfile := flag.String("param", "", "param for pml filename")
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
	case "template":
		if len(*paramfile) == 0 {
			flag.PrintDefaults()
			return
		}

		var dat interface{}
		param, err := ioutil.ReadFile(*paramfile)
		if err != nil {
			fmt.Println("Cannot find file " + *paramfile)
			return
		}
		if err := json.Unmarshal(param, &dat); err != nil {
			fmt.Println("Cannot unmarshall json file " + *paramfile)
			return
		}
		out, err := template.Apply(string(file), dat)
		if err != nil {
			fmt.Printf("failed to transform template : %v\n", err)
			return
		}
		fmt.Println(out)
	case "full":
		if len(*paramfile) == 0 {
			flag.PrintDefaults()
			return
		}

		var dat interface{}
		param, err := ioutil.ReadFile(*paramfile)
		if err != nil {
			fmt.Println("Cannot find file " + *paramfile)
			return
		}
		if err := json.Unmarshal(param, &dat); err != nil {
			fmt.Println("Cannot unmarshall json file " + *paramfile)
			return
		}
		out, err := template.Apply(string(file), dat)
		if err != nil {
			fmt.Printf("failed to transform template : %v\n", err)
			return
		}
		l := lexer.New(out)
		p := parser.New(l)
		r := renderer.New(*output)

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
