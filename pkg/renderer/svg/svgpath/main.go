package svgpath

func Parse(path string) ([]Command, error) {
	l := newLexer(path)
	parser := newParser(l)
	commands, err := parser.parse()
	if err != nil {
		return nil, err
	}

	for _, cmd := range commands {
		if validateCommand(cmd) == false {
			return nil, err
		}
	}
	return commands, nil
}
