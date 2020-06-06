package svg

import (
	"errors"
)

var (
	errCannotParseMainTransformAttr = errors.New("errCannotParseMainTransformAttr")
	errCannotParseElement           = errors.New("errCannotParseElement")
)
