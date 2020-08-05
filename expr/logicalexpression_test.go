package expr

import (
	"aparser"
	"aparser/ast"
	"testing"
)

func TestLogicalExpression(t *testing.T) {
	e := logicalExpression()
	b := aparser.CreateBuffer("false || true")
	checkRead(t, b, e)
	checkPosition(t, b, 13)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, true)
	b = aparser.CreateBuffer("false || (true && false)")
	checkRead(t, b, e)
	checkPosition(t, b, 24)
	//a = ast.CreateAST(b)
	//checkEvaluate(t, a, true)
}
