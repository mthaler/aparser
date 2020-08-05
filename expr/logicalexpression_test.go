package expr

import (
	"aparser"
	"aparser/ast"
	"testing"
)

func TestLogicalExpressionLiterals(t *testing.T) {
	e := logicalExpression()
	b := aparser.CreateBuffer("false")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, false)
	b = aparser.CreateBuffer("true")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, true)
}

func TestLogicalExpressionSimple(t *testing.T) {
	e := logicalExpression()
	b := aparser.CreateBuffer("false || true")
	checkRead(t, b, e)
	checkPosition(t, b, 13)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, true)
}

func TestLogicalExpressionWithGroup(t *testing.T) {
	e := logicalExpression()
	b := aparser.CreateBuffer("false || (true && false)")
	checkRead(t, b, e)
	checkPosition(t, b, 24)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, false)
	b = aparser.CreateBuffer("(true ||false) || (true && false)")
	checkRead(t, b, e)
	checkPosition(t, b, 33)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, true)
}
