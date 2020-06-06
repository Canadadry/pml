package svg

import (
	"fmt"
	"pml/pkg/renderer/svg/matrix"
	"pml/pkg/renderer/svg/svgparser"
	"regexp"
	"strconv"
	"strings"
)

func svgPath(element *svgparser.Element, worldToParent matrix.Matrix) (*svgNode, error) {
	sp := &svgNode{
		worldToLocal: worldToParent,
		commands:     []command{},
	}

	d, ok := element.Attributes["d"]
	if ok {
		err := sp.parsePath(d)
		if err != nil {
			return nil, err
		}
	}

	// style, ok := element.Attributes["style"]
	// if ok {
	// 	err := sp.loadStyle(style)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return sp, nil
}

func (sp *svgNode) parsePath(path string) error {
	r, err := regexp.Compile("([MmLlHhVvCcZz])([-0-9,\\. ]*)")
	if err != nil {
		return err
	}
	founds := r.FindAllString(path, -1)

	for _, f := range founds {

		if len(f) == 1 {
			sp.commands = append(sp.commands, command{f[0], 0, 0, 0, 0, 0, 0})
			continue
		}

		kind := f[0]
		simplyfied := strings.ReplaceAll(f[1:], ",", " ")
		params := strings.Split(simplyfied, " ")
		if len(params) != 2 && len(params) != 4 && len(params) != 6 {
			return fmt.Errorf("too many param (got %d) while parsing %s", len(params), f)
		}

		x1, err := strconv.ParseFloat(params[0], 64)
		y1, err := strconv.ParseFloat(params[1], 64)
		x1, y1, _ = sp.worldToLocal.Project(x1, y1, 1.0)

		x2 := float64(0)
		y2 := float64(0)

		if len(params) > 2 {
			x2, err = strconv.ParseFloat(params[2], 64)
			y2, err = strconv.ParseFloat(params[3], 64)
			x2, y2, _ = sp.worldToLocal.Project(x2, y2, 1.0)
		}

		x3 := float64(0)
		y3 := float64(0)
		if len(params) > 4 {
			x3, err = strconv.ParseFloat(params[4], 64)
			y3, err = strconv.ParseFloat(params[5], 64)
			x3, y3, _ = sp.worldToLocal.Project(x3, y3, 1.0)
		}
		if err != nil {
			return err
		}

		sp.commands = append(sp.commands, command{kind, x1, y1, x2, y2, x3, y3})
	}
	return nil
}
