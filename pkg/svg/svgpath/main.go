package svgpath

import (
	"errors"
	"fmt"
)

var ErrInvalideCommandFound = errors.New("InvalideCommandFound")

func Parse(path string) ([]Command, error) {
	l := newLexer(path)
	parser := newParser(l)
	commands, err := parser.parse()

	if err != nil {
		return nil, err
	}

	for _, cmd := range commands {
		if validateCommand(cmd) == false {
			return nil, fmt.Errorf("%w : %s", ErrInvalideCommandFound, string(cmd.Kind))
		}
	}
	return commands, nil
}
