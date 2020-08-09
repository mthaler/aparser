package expr

type optionalExpression struct {
	*abstractExpression
	expression Expression
}

func optional(e Expression) optionalExpression {
	a := abstractExpression{}
	return optionalExpression{abstractExpression: &a, expression: e}
}

func (o optionalExpression) parse(buffer *buffer) bool {
	Parse(o.expression, buffer)
	return true
}
