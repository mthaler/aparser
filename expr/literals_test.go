package expr

import (
	"aparser/ast"
	"testing"
)

func Test_charLiteralExpression_Read(t *testing.T) {
	e := charLiteral('s')
	b := CreateBuffer("test")
	checkNotRead(t, b, e)
	b.incrementCurrentPosition()
	checkNotRead(t, b, e)
	b.incrementCurrentPosition()
	checkRead(t, b, e)
}

func TestDoubleLiteralExpression_Read(t *testing.T) {
	e := doubleLiteral()
	e.SetCreateNode(ast.CreateDoubleOperand)
	b := CreateBuffer("foo")
	checkNotRead(t, b, e)
	b = CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	a := ast.CreateAST(b.Code.Code)
	checkEvaluate(t, a, 3.14)
}

func Test_stringLiteral(t *testing.T) {
	a := stringLiteral("&&")
	if a.str != "&&" {
		t.Error("str should be &&")
	}
}

func Test_stringLiteralExpression_Read(t *testing.T) {
	e := stringLiteral("&&")
	b := CreateBuffer("foo")
	checkNotRead(t, b, e)
}

func Test_caseInsensitiveStringLiteralExpression_Read(t *testing.T) {
	e := caseInsensitiveStringLiteral("Test")
	b := CreateBuffer("test")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = CreateBuffer("Test")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = CreateBuffer("TEST")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = CreateBuffer("foo")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
}
