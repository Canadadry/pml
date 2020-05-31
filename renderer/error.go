package renderer

import (
	"errors"
)

var (
	rootMustBeDocumentItem      = errors.New("rootMustBeDocumentItem")
	renderingItemNotImplemented = errors.New("renderingItemNotImplemented")
)
