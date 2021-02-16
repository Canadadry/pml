package svgnode

import (
	"errors"
	"fmt"
	"github.com/canadadry/pml/pkg/matrix"
	"github.com/canadadry/pml/pkg/svg/svgdrawer"
	"github.com/canadadry/pml/pkg/svg/svgparser"
	"image/color"
	"strconv"
	"strings"
)

func parseStyleAttribute(element *svgparser.Element, transform matrix.Matrix) svgdrawer.Style {
	s := svgdrawer.Style{
		FillColor:   color.RGBA{0, 0, 0, 0},
		BorderSize:  0,
		BorderColor: color.RGBA{0, 0, 0, 0},
		PathStyle:   svgdrawer.None,
	}

	style, ok := element.Attributes["style"]
	if !ok {
		return s
	}

	part := strings.Split(style, ";")

	for _, p := range part {
		arg := strings.Split(p, ":")
		if len(arg) != 2 {
			continue
		}
		switch arg[0] {
		case "fill":
			c, err := parseColorParam(arg[1])
			if err != nil {
				continue
			}
			s.FillColor = c
			switch s.PathStyle {
			case svgdrawer.None:
				s.PathStyle = svgdrawer.Fill
			case svgdrawer.Fill:
				s.PathStyle = svgdrawer.Fill
			default:
				s.PathStyle = svgdrawer.FillAndStroke
			}

		case "stroke":
			c, err := parseColorParam(arg[1])
			if err != nil {
				continue
			}
			s.BorderColor = c
			s.BorderSize = 0.1
			switch s.PathStyle {
			case svgdrawer.None:
				s.PathStyle = svgdrawer.Stroke
			case svgdrawer.Stroke:
				s.PathStyle = svgdrawer.Stroke
			default:
				s.PathStyle = svgdrawer.FillAndStroke
			}

		case "stroke-width":
			if arg[1][len(arg[1])-2:] == "px" {
				width, err := strconv.ParseFloat(arg[1][:len(arg[1])-2], 64)
				if err != nil {
					continue
				}
				newWidth, _ := transform.ProjectPoint(width, 0)
				s.BorderSize = newWidth
			}
		case "fill-rule":
			s.EvenOddRule = arg[1] != "nonzero"
		}
	}

	return s
}

func parseColorParam(attribute string) (color.RGBA, error) {

	if len(attribute) > 4 && attribute[:4] == "rgb(" {
		return parseFuncRgbColor(attribute)
	}

	if (len(attribute) == 7 || len(attribute) == 4) && attribute[0] == '#' {
		return parseHexColor(attribute)
	}

	rgb, ok := colorDict[attribute]
	if !ok {
		return color.RGBA{}, errors.New("ColorNotFoundInDictionnary")
	}
	return rgb, nil
}

func parseHexColor(s string) (c color.RGBA, err error) {
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func parseFuncRgbColor(attribute string) (color.RGBA, error) {
	colorAttr := attribute[4 : len(attribute)-1]
	colorPart := strings.Split(colorAttr, ",")
	if len(colorPart) != 3 {
		return color.RGBA{}, errors.New("NotEnoughPartToParseRGB")
	}
	r, err := strconv.ParseUint(colorPart[0], 10, 8)
	if err != nil {
		return color.RGBA{}, errors.New("ParseUintFailed")
	}
	g, err := strconv.ParseUint(colorPart[1], 10, 8)
	if err != nil {
		return color.RGBA{}, errors.New("ParseUintFailed")
	}
	b, err := strconv.ParseUint(colorPart[2], 10, 8)
	if err != nil {
		return color.RGBA{}, errors.New("ParseUintFailed")
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), 0}, nil
}
