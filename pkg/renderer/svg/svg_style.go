package svg

import (
	"errors"
	"image/color"
	"pml/pkg/renderer/svg/svgparser"
	"strconv"
	"strings"
)

func parseStyleAttribute(element *svgparser.Element) Style {
	s := Style{
		Fill:        false,
		FillColor:   color.RGBA{0, 0, 0, 0},
		BorderSize:  0,
		BorderColor: color.RGBA{0, 0, 0, 0},
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
			s.Fill = true
		case "stroke":
			c, err := parseColorParam(arg[1])
			if err != nil {
				continue
			}
			s.BorderColor = c
			s.BorderSize = 0.1
		case "stroke-width":
			s.BorderSize = 0.1
		}
	}

	return s
}

func parseColorParam(attribute string) (color.RGBA, error) {

	if len(attribute) <= 4 || attribute[:4] != "rgb(" {
		rgb, ok := colorDict[attribute]
		if !ok {
			return color.RGBA{}, errors.New("ColorNotFoundInDictionnary")
		}
		return rgb, nil
	}

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
