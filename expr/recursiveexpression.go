package expr

type recursiveExpression struct {
	*abstractExpression
	value Expression
}

func recursive() recursiveExpression {
	a := abstractExpression{}
	return recursiveExpression{abstractExpression: &a}
}

func (r recursiveExpression) parse(buffer *buffer) bool {
	if r.value == nil {
		panic("Value must be set to an expression before reading an expression!")
	}
	return Parse(r.value, buffer)
}
