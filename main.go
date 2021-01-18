package main

import (
	"flag"
	"fmt"
	"github.com/canadadry/pml/pkg/domain"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {

	filename := flag.String("in", "", "entry pml filename")
	paramfile := flag.String("param", "", "param for pml filename")
	output := flag.String("out", "out.pdf", "pdf output for renderer mode")

	flag.Parse()

	if len(*filename) == 0 {
		flag.PrintDefaults()
		return nil
	}

	fIn, err := os.Open(*filename)
	if err != nil {
		return fmt.Errorf("Cannot find file " + *filename)
	}
	defer fIn.Close()

	fParam, _ := os.Open(*paramfile)
	if fParam != nil {
		defer fParam.Close()
	}

	fOut, err := os.Create(*output)
	if err != nil {
		return fmt.Errorf("Cannot create output file " + *output)
	}
	defer fOut.Close()

	return domain.Run(fIn, fOut, fParam)
}
