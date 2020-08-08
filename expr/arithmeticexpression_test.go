package expr

import (
	"aparser"
	"aparser/ast"
	"testing"
)

func TestArithmeticExpressionLiteral(t *testing.T) {
	e := arithmeticExpression()
	b := aparser.CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, 3.14)
	b = aparser.CreateBuffer("false")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, false)
	b = aparser.CreateBuffer("true")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, true)
}

func TestArithmeticExpressionSimple(t *testing.T) {
	e := arithmeticExpression()
	b := aparser.CreateBuffer("3 + 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, 7.0)
	b = aparser.CreateBuffer("3 - 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, -1.0)
	b = aparser.CreateBuffer("3 * 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, 12.0)
	b = aparser.CreateBuffer("3 / 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, 0.75)
	b = aparser.CreateBuffer("false || true")
	checkRead(t, b, e)
	checkPosition(t, b, 13)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, true)
}

func TestArithmeticExpressionGroups(t *testing.T) {
	e := arithmeticExpression()
	b := aparser.CreateBuffer("(3 + 4)")
	checkRead(t, b, e)
	checkPosition(t, b, 7)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, 7.0)
	b = aparser.CreateBuffer("3 * (4 + 5)")
	checkRead(t, b, e)
	checkPosition(t, b, 11)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, 27.0)
	b = aparser.CreateBuffer("false || (true && false)")
	checkRead(t, b, e)
	checkPosition(t, b, 24)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, false)
	b = aparser.CreateBuffer("(true ||false) || (true && false)")
	checkRead(t, b, e)
	checkPosition(t, b, 33)
	a = ast.CreateAST(b)
	checkEvaluate(t, a, true)
}

func TestArithmeticExpressionFunction(t *testing.T) {
	e := arithmeticExpression()
	b := aparser.CreateBuffer("abs(-5)")
	checkRead(t, b, e)
	checkPosition(t, b, 7)
	a := ast.CreateAST(b)
	checkEvaluate(t, a, 5.0)
}
