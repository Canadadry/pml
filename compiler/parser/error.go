package parser

import (
	"errors"
)

var (
	errNextTokenIsNotTheExpectedOne = errors.New("errNextTokenIsNotTheExpectedOne")
	errPropertyDefinedTwice         = errors.New("errPropertyDefinedTwice")
	errNotAValueType                = errors.New("errNotAValueType")
	errInvalidIdentifier            = errors.New("errInvalidIdentifier")
)
