package main

import (
	"flag"
	"fmt"
	"github.com/canadadry/pml/compiler"
	"github.com/canadadry/pml/pkg/pdf"
	"github.com/canadadry/pml/pkg/svg"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {

	var (
		filename  string
		main      string
		paramfile string
		output    string
	)

	flag.StringVar(&filename, "in", "", "entry pml filename or directory (will only load sub .pml files)")
	flag.StringVar(&main, "main", "", "if in is a folder, main should contain the main filename")
	flag.StringVar(&paramfile, "param", "", "param for pml filename")
	flag.StringVar(&output, "out", "out.pdf", "pdf output for renderer mode")

	flag.Parse()

	if len(filename) == 0 {
		flag.PrintDefaults()
		return nil
	}

	files, err := LoadInputFiles(filename)
	if err != nil {
		return err
	}
	defer CloseFiles(files)

	if len(files) == 1 {
		for n := range files {
			main = n
		}
	} else {
		if len(main) == 0 {
			return fmt.Errorf("You must select your main file")
		}
	}

	var fParam io.ReadCloser
	fParam, err = os.Open(paramfile)
	if err != nil {
		fParam = nil
	} else {
		defer fParam.Close()
	}

	fOut, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("Cannot create output file " + output)
	}
	defer fOut.Close()

	return compiler.Run(files, main, fOut, fParam, pdf.New(svg.Draw))
}

func LoadInputFiles(filename string) (map[string]io.Reader, error) {
	files := map[string]io.Reader{}

	f, err := os.Open(filename)
	if err != nil {
		return files, fmt.Errorf("Cannot find file '%s' : %w", filename, err)
	}
	stat, err := f.Stat()
	if err != nil {
		return files, fmt.Errorf("Cannot stat file '%s' : %w", filename, err)
	}

	if !stat.IsDir() {
		files[filename] = f
		return files, nil
	}

	pattern := filepath.Clean(filename + "/*.pml")
	paths, err := filepath.Glob(pattern)
	if err != nil {
		return files, fmt.Errorf("Cannot glob '%s' : %w", pattern, err)
	}

	for _, p := range paths {
		f, err := os.Open(p)
		if err != nil {
			return files, fmt.Errorf("Cannot find file '%s' : %w", p, err)
		}
		stat, err := f.Stat()
		if err != nil {
			return files, fmt.Errorf("Cannot stat file '%s' : %w", p, err)
		}
		if stat.IsDir() {
			continue
		}
		files[p] = f
	}

	return files, nil
}

func CloseFiles(files map[string]io.Reader) {
	for _, f := range files {
		fcloser, ok := f.(io.ReadCloser)
		if !ok {
			continue
		}
		fcloser.Close()
	}
}
