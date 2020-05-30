package parser

import (
	"errors"
)

var (
	errNextTokenIsNotTheExpectedOne = errors.New("errNextTokenIsNotTheExpectedOne")
	errPropertyDefinedTwice         = errors.New("errPropertyDefinedTwice")
)
