package svg

import (
	"fmt"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"pml/pkg/renderer/svg/svgpath"
)

func svgPath(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {
	sp := &svgNode{
		worldToLocal: worldToParent,
	}

	d, ok := element.Attributes["d"]
	if !ok {
		return nil, fmt.Errorf("No attribute 'd' found")
	}

	l := svgpath.NewLexer(d)
	p := svgpath.NewParser(l)

	cmd, err := p.Parse()
	if err != nil {
		return nil, err
	}

	sp.commands = cmd

	return sp, nil
}
