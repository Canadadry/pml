package main

import (
	"flag"
	"fmt"
	"github.com/canadadry/pml/compiler"
	"github.com/canadadry/pml/pkg/pdf"
	"github.com/canadadry/pml/pkg/svg"
	"io"
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
		return fmt.Errorf("Cannot find file '%s' : %w", *filename, err)
	}
	defer fIn.Close()

	var fParam io.ReadCloser
	fParam, err = os.Open(*paramfile)
	if err != nil {
		fParam = nil
	} else {
		defer fParam.Close()
	}

	fOut, err := os.Create(*output)
	if err != nil {
		return fmt.Errorf("Cannot create output file " + *output)
	}
	defer fOut.Close()

	return compiler.Run(fIn, fOut, fParam, pdf.New(svg.Draw))
}
