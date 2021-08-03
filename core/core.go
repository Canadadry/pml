package core

import (
	"github.com/canadadry/pml/compiler/renderer"
)

type Core struct {
	Drawer renderer.PdfDrawer
}

func New(pdf renderer.Pdf) Core {
	return Core{pdf.Init()}
}
