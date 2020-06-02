package svg

import (
	"errors"
)

var (
	errCannotParseMainTransformAttr = errors.New("errCannotParseMainTransformAttr")
	errMissingAttr                  = errors.New("errMissingAttr")
)
