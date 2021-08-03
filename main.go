package main

import (
	"fmt"
	"github.com/canadadry/pml/core"
	"github.com/canadadry/pml/doc"
	"github.com/canadadry/pml/pkg/pdf"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	for name, doc := range doc.Docs {
		output := "doc/" + name + ".pdf"
		fOut, err := os.Create(output)
		if err != nil {
			return fmt.Errorf("Cannot create output file " + output)
		}
		defer fOut.Close()

		c := core.New(pdf.New(nil))

		err = doc(c)
		if err != nil {
			return fmt.Errorf("while rendering %s : %w", name, err)
		}
		err = c.Drawer.Output(fOut)
		if err != nil {
			return fmt.Errorf("while saving %s : %w", name, err)
		}
	}
	return nil
}
