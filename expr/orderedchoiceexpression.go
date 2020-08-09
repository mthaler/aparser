package expr

type orderedChoiceExpression struct {
	*abstractExpression
	expression1 Expression
	expression2 Expression
}

func orderedChoice(e1 Expression, e2 Expression) orderedChoiceExpression {
	a := abstractExpression{}
	return orderedChoiceExpression{abstractExpression: &a, expression1: e1, expression2: e2}
}

func (o orderedChoiceExpression) parse(buffer *buffer) bool {
	if Parse(o.expression1, buffer) {
		return true
	} else if Parse(o.expression2, buffer) {
		return true
	}
	return false
}
