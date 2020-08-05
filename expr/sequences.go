package expr

import "aparser"

type sequenceExpression struct {
	*abstractExpression
	expression1 Expression
	expression2 Expression
}

func sequence(e1 Expression, e2 Expression) sequenceExpression {
	a := abstractExpression{}
	return sequenceExpression{abstractExpression: &a, expression1: e1, expression2: e2}
}

func (s sequenceExpression) parse(buffer *aparser.Buffer) bool {
	if Parse(s.expression1, buffer) && Parse(s.expression2, buffer) {
		return true
	}
	return false
}

type zeroOrMoreExpression struct {
	*abstractExpression
	expression Expression
}

func zeroOrMore(e Expression) zeroOrMoreExpression {
	a := abstractExpression{}
	return zeroOrMoreExpression{abstractExpression: &a, expression: e}
}

func (z zeroOrMoreExpression) parse(buffer *aparser.Buffer) bool {
	for Parse(z.expression, buffer) {
	}
	return true
}

type oneOrMoreExpression struct {
	*abstractExpression
	expression Expression
}

func oneOrMore(e Expression) oneOrMoreExpression {
	a := abstractExpression{}
	return oneOrMoreExpression{abstractExpression: &a, expression: e}
}

func (o oneOrMoreExpression) parse(buffer *aparser.Buffer) bool {
	success := false
	for Parse(o.expression, buffer) {
		success = true
	}
	return success
}

