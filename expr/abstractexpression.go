package expr

import (
	"aparser"
)

type abstractExpression struct {
	id         string
	createNode func(buffer *aparser.Buffer) interface{}
}

func (a *abstractExpression) Id() string {
	return a.id
}

func (a *abstractExpression) SetId(id string) {
	a.id = id
}

func (a *abstractExpression) CreateNode() func(buffer *aparser.Buffer) interface{} {
	return a.createNode
}

func (a *abstractExpression) SetCreateNode(m func(buffer *aparser.Buffer) interface{}) {
	a.createNode = m
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