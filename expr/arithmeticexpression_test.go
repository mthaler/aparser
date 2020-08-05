package expr

import (
	"aparser"
	"aparser/ast"
	"testing"
)

func Test_arithmeticExpression_Read_doubleLiteral(t *testing.T) {
	e := arithmeticExpression()
	b := aparser.CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, 3.14)
}

func Test_arithmeticExpression(t *testing.T) {
	e := arithmeticExpression()
	b := aparser.CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = aparser.CreateBuffer("2 * (3 + 4)")
	checkRead(t, b, e)
	checkPosition(t, b, 11)
}