package expr

import (
	"aparser/ast"
	"testing"
)

func TestArithmeticExpressionLiteral(t *testing.T) {
	e := ArithmeticExpression()
	b := CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	a := ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 3.14)
	b = CreateBuffer("false")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, false)
	b = CreateBuffer("true")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
}

func TestArithmeticExpressionSimple(t *testing.T) {
	e := ArithmeticExpression()
	b := CreateBuffer("3 + 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a := ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 7.0)
	b = CreateBuffer("3 - 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, -1.0)
	b = CreateBuffer("3 * 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 12.0)
	b = CreateBuffer("3 / 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 0.75)
	b = CreateBuffer("false || true")
	checkRead(t, b, e)
	checkPosition(t, b, 13)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
	b = CreateBuffer("3 + 4 + 5")
	checkRead(t, b, e)
	checkPosition(t, b, 9)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 12.0)
	b = CreateBuffer("3 - 4 - 5")
	checkRead(t, b, e)
	checkPosition(t, b, 9)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, -6.0)
}

func TestArithmeticExpressionGroups(t *testing.T) {
	e := ArithmeticExpression()
	b := CreateBuffer("(3 + 4)")
	checkRead(t, b, e)
	checkPosition(t, b, 7)
	a := ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 7.0)
	b = CreateBuffer("3 * (4 + 5)")
	checkRead(t, b, e)
	checkPosition(t, b, 11)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 27.0)
	b = CreateBuffer("false || (true && false)")
	checkRead(t, b, e)
	checkPosition(t, b, 24)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, false)
	b = CreateBuffer("(true ||false) || (true && false)")
	checkRead(t, b, e)
	checkPosition(t, b, 33)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
}

func TestArithmeticExpressionFunction(t *testing.T) {
	e := ArithmeticExpression()
	b := CreateBuffer("abs(-5)")
	checkRead(t, b, e)
	checkPosition(t, b, 7)
	a := ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 5.0)
}

func TestArithmeticExpressionRelational(t *testing.T) {
	e := ArithmeticExpression()
	b := CreateBuffer("3 < 3")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a := ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, false)
	b = CreateBuffer("3 < 4")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
	b = CreateBuffer("3 > 3")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, false)
	b = CreateBuffer("4 > 3")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
	b = CreateBuffer("3 <= 3")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
	b = CreateBuffer("4 <= 3")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, false)
	b = CreateBuffer("3 >= 3")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
	b = CreateBuffer("3 >= 4")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, false)
}

func TestArithmeticExpressionTernaryOperation(t *testing.T) {
	e := ArithmeticExpression()
	b := CreateBuffer("true ? 4 : 5")
	checkRead(t, b, e)
	checkPosition(t, b, 12)
	a := ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 4.0)
	b = CreateBuffer("false ? 4 : 5")
	checkRead(t, b, e)
	checkPosition(t, b, 13)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 5.0)
	b = CreateBuffer("if (true) 4 else 5")
	checkRead(t, b, e)
	checkPosition(t, b, 18)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 4.0)
	b = CreateBuffer("if (false) 4 else 5")
	checkRead(t, b, e)
	checkPosition(t, b, 19)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 5.0)
}

func TestArithmeticExpressionEquality(t *testing.T) {
	e := ArithmeticExpression()
	b := CreateBuffer("3 == 3")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
	a := ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
	b = CreateBuffer("3 == 4")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, false)
	b = CreateBuffer("3 != 3")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, false)
	b = CreateBuffer("3 != 4")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
	a = ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, true)
}

