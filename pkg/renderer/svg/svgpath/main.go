package svgpath

import (
	"errors"
)

func Parse(path string) ([]Command, error) {
	l := newLexer(path)
	parser := newParser(l)
	commands, err := parser.parse()

	if err != nil {
		return nil, err
	}

	for _, cmd := range commands {
		if validateCommand(cmd) == false {
			return nil, errors.New("InvalideCommandFound : " + string(cmd.Kind))
		}
	}
	return commands, nil
}
