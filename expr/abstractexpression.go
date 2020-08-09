package expr

import (
	"aparser"
)

type abstractExpression struct {
	createNode func(buffer *aparser.Buffer) interface{}
}

func (a *abstractExpression) CreateNode() func(buffer *aparser.Buffer) interface{} {
	return a.createNode
}

func (a *abstractExpression) SetCreateNode(m func(buffer *aparser.Buffer) interface{}) {
	a.createNode = m
}
