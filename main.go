package main

import (
	"flag"
	"fmt"
	"github.com/canadadry/pml/cmd"
	"io/ioutil"
	"os"
)

func main() {

	mode := flag.String("mode", "direct", "mode : direct|api")
	filename := flag.String("in", "", "entry pml filename or folder if in api mode")
	paramfile := flag.String("param", "", "param for pml filename (unused in api mode)")
	output := flag.String("out", "out.pdf", "pdf output for renderer mode  (unused in api mode)")

	flag.Parse()

	switch *mode {
	case "direct":
		if len(*filename) == 0 {
			flag.PrintDefaults()
			return
		}

		file, err := ioutil.ReadFile(*filename)
		if err != nil {
			fmt.Println("Cannot find file " + *filename)
			return
		}
		param := []byte("{}")
		if len(*paramfile) > 0 {
			param, err = ioutil.ReadFile(*paramfile)
			if err != nil {
				fmt.Println("Cannot find file " + *paramfile)
				return
			}
		}

		fOut, err := os.Create(*output)
		if err != nil {
			fmt.Println("Cannot create file " + *output)
			return
		}
		defer fOut.Close()
		err = cmd.Full(string(file), fOut, param)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "api":
		err := cmd.Api(*filename)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Printf("Mode not handle : " + *mode)
		flag.PrintDefaults()
	}
}
