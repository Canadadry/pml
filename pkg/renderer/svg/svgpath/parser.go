package svgpath

import (
	"errors"
	"strconv"
)

var ErrExpectedCommandToken = errors.New("ErrExpectedCommandToken")
var ErrExpectedFloatToken = errors.New("ErrExpectedFloatToken")
var ErrExpectedComaToken = errors.New("ErrExpectedComaToken")

type parser struct {
	current Token
	next    Token
	lexer   *Lexer
}

func NewParser(l *Lexer) parser {
	return parser{
		current: l.GetNextToken(),
		next:    l.GetNextToken(),
		lexer:   l,
	}
}

func (p *parser) Parse() ([]Command, error) {
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
		point := Point{}

		if !p.isCurrentTokenA(NUMBER) {
			return cmd, ErrExpectedFloatToken
		}
		var err error
		point.X, err = strconv.ParseFloat(p.current.Literal, 64)
		if err != nil {
			return cmd, err
		}
		p.goToNextToken()

		if !p.isCurrentTokenA(COMA) {
			return cmd, ErrExpectedComaToken
		}
		p.goToNextToken()

		if !p.isCurrentTokenA(NUMBER) {
			return cmd, ErrExpectedFloatToken
		}
		point.Y, err = strconv.ParseFloat(p.current.Literal, 64)
		if err != nil {
			return cmd, err
		}

		p.goToNextToken()

		cmd.Points = append(cmd.Points, point)
	}
	return cmd, nil
}

func (p *parser) goToNextToken() {
	p.current = p.next
	p.next = p.lexer.GetNextToken()
}

func (p *parser) isCurrentTokenA(t TokenType) bool {
	return p.current.Type == t
}
