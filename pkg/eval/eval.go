package eval

func Eval(str string) (float64, error) {
	n, err := NewParser(NewLexer(str)).ParseExpression(LOWEST)
	if err != nil {
		return 0.0, err
	}
	return n.Eval(), nil
}
