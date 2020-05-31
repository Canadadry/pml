package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"pml/cmd"
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

	switch *mode {
	case "lexer":
		cmd.Lexer(string(file))
	case "parser":
		err := cmd.Parser(string(file))
		if err != nil {
			fmt.Println(err)
			return
		}
	case "renderer":
		err := cmd.Renderer(string(file), *output)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "template":
		if len(*paramfile) == 0 {
			flag.PrintDefaults()
			return
		}
		param, err := ioutil.ReadFile(*paramfile)
		if err != nil {
			fmt.Println("Cannot find file " + *paramfile)
			return
		}
		err = cmd.Template(string(file), param)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "full":
		if len(*paramfile) == 0 {
			flag.PrintDefaults()
			return
		}
		param, err := ioutil.ReadFile(*paramfile)
		if err != nil {
			fmt.Println("Cannot find file " + *paramfile)
			return
		}
		err = cmd.Full(string(file), *output, param)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Printf("Mode not handle : " + *mode)
		flag.PrintDefaults()
	}
}
