package expr

import (
	"aparser/ast"
)

type abstractExpression struct {
	createNode func(text string, c *ast.Code) interface{}
}

func (a *abstractExpression) CreateNode() func(text string, c *ast.Code) interface{} {
	return a.createNode
}

func (a *abstractExpression) SetCreateNode(createNode func(text string, c *ast.Code) interface{}) {
	a.createNode = createNode
}
