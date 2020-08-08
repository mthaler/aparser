package expr

import "aparser"

type optionalExpression struct {
	*abstractExpression
	expression Expression
}

func optional(e Expression) optionalExpression {
	a := abstractExpression{}
	return optionalExpression{abstractExpression: &a, expression: e}
}

func (o optionalExpression) parse(buffer *aparser.Buffer) bool {
	Parse(o.expression, buffer)
	return true
}
