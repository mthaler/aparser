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

func and(expressions ...Expression) Expression {
	l := len(expressions)
	if l == 0 {
		panic("and called with empty expressions")
	} else if l == 1 {
		return expressions[0]
	} else {
		e1 := expressions[0]
		e2 := and(expressions[1:]...)
		return sequence(e1, e2)
	}
}

func or(expressions ...Expression) Expression {
	l := len(expressions)
	if l == 0 {
		panic("or called with empty expressions")
	} else if l == 1 {
		return expressions[0]
	} else {
		e1 := expressions[0]
		e2 := or(expressions[1:]...)
		return orderedChoice(e1, e2)
	}
}
