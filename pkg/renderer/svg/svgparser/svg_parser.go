package svgparser

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type Element struct {
	Name       string
	Attributes map[string]string
	Children   []*Element
	Content    string
}

var ErrMissingAttr = errors.New("errMissingAttr")
var ErrParsingAttr = errors.New("ErrParsingAttr")

func Parse(source io.Reader) (*Element, error) {
	raw, err := ioutil.ReadAll(source)
	if err != nil {
		return nil, err
	}
	decoder := xml.NewDecoder(bytes.NewReader(raw))
	decoder.CharsetReader = charset.NewReaderLabel
	element, err := decodeFirstLine(decoder)
	if err != nil {
		return nil, err
	}
	if err := element.decode(decoder); err != nil && err != io.EOF {
		return nil, err
	}
	return element, nil
}

func decodeFirstLine(decoder *xml.Decoder) (*Element, error) {
	for {
		token, err := decoder.Token()
		if token == nil && err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		switch element := token.(type) {
		case xml.StartElement:
			return newElement(element), nil
		}
	}
	return &Element{}, nil
}

func newElement(token xml.StartElement) *Element {
	element := &Element{}
	attributes := make(map[string]string)
	for _, attr := range token.Attr {
		attributes[attr.Name.Local] = attr.Value
	}
	element.Name = token.Name.Local
	element.Attributes = attributes
	return element
}

func (e *Element) decode(decoder *xml.Decoder) error {
	for {
		token, err := decoder.Token()
		if token == nil && err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch element := token.(type) {
		case xml.StartElement:
			nextElement := newElement(element)
			err := nextElement.decode(decoder)
			if err != nil {
				return err
			}

			e.Children = append(e.Children, nextElement)

		case xml.CharData:
			data := strings.TrimSpace(string(element))
			if data != "" {
				e.Content = string(element)
			}

		case xml.EndElement:
			if element.Name.Local == e.Name {
				return nil
			}
		}
	}
	return nil
}

func (e *Element) ReadAttributeAsFloat(attribute string) (float64, error) {
	attr, ok := e.Attributes[attribute]
	if !ok {
		return 0, fmt.Errorf("%w : cannot found %s", ErrMissingAttr, attribute)
	}

	parsed, err := strconv.ParseFloat(attr, 64)
	if err != nil {
		return 0, fmt.Errorf("%w : error while parsing %s : %v", ErrParsingAttr, attribute, err)
	}

	return parsed, nil
}
