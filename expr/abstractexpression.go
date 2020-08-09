package expr

import "aparser"

type abstractExpression struct {
	createNode func(text string, c *aparser.Code) interface{}
}

func (a *abstractExpression) CreateNode() func(text string, c *aparser.Code) interface{} {
	return a.createNode
}

func (a *abstractExpression) SetCreateNode(createNode func(text string, c *aparser.Code) interface{}) {
	a.createNode = createNode
}
