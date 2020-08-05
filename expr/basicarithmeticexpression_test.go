package expr

import (
	"aparser"
	"aparser/ast"
	"testing"
)

func Test_basicArithmeticExpression_number(t *testing.T) {
	e := basicArithmeticExpression()
	b := aparser.CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, 3.14)
}

func Test_basicArithmeticExpression_sum(t *testing.T) {
	e := basicArithmeticExpression()
	b := aparser.CreateBuffer("3 + 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, 7.0)
}
