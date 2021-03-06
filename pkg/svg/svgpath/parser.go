package svgpath

import (
	"errors"
	"strconv"
)

var ErrExpectedCommandToken = errors.New("ErrExpectedCommandToken")
var ErrExpectedFloatToken = errors.New("ErrExpectedFloatToken")

type parser struct {
	current Token
	next    Token
	lexer   *Lexer
}

func newParser(l *Lexer) parser {
	return parser{
		current: l.getNextToken(),
		next:    l.getNextToken(),
		lexer:   l,
	}
}

func (p *parser) parse() ([]Command, error) {
	cmds := []Command{}

	for !p.isCurrentTokenA(EOF) {
		cmd, err := p.parseCommand()
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}

func (p *parser) parseCommand() (Command, error) {
	cmd := Command{}

	if !p.isCurrentTokenA(COMMAND) {
		return cmd, ErrExpectedCommandToken
	}

	cmd.Kind = p.current.Literal[0]

	p.goToNextToken()

	for p.isCurrentTokenA(NUMBER) {
		point, err := p.parsePoint()
		if err != nil {
			if len(cmd.Points) != 0 {
				return cmd, err
			}
			if !errors.Is(err, ErrExpectedFloatToken) {
				return cmd, err
			}
		}
		cmd.Points = append(cmd.Points, point)
	}
	return cmd, nil
}

func (p *parser) parsePoint() (Point, error) {
	point := Point{}

	point.X, _ = strconv.ParseFloat(p.current.Literal, 64)
	p.goToNextToken()

	if p.isCurrentTokenA(COMA) {
		p.goToNextToken()
	}

	if !p.isCurrentTokenA(NUMBER) {
		return point, ErrExpectedFloatToken
	}
	point.Y, _ = strconv.ParseFloat(p.current.Literal, 64)

	p.goToNextToken()

	return point, nil
}

func (p *parser) goToNextToken() {
	p.current = p.next
	p.next = p.lexer.getNextToken()
}

func (p *parser) isCurrentTokenA(t TokenType) bool {
	return p.current.Type == t
}
