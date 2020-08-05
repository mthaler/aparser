package expr

import (
	"aparser"
	"testing"
)

func Test_charLiteralExpression_Read(t *testing.T) {
	e := charLiteral('s')
	b := aparser.CreateBuffer("test")
	checkNotRead(t, b, e)
	b.IncrementCurrentPosition()
	checkNotRead(t, b, e)
	b.IncrementCurrentPosition()
	checkRead(t, b, e)
}

func TestDoubleLiteralExpression_Read(t *testing.T) {
	e := doubleLiteral()
	//e.SetCreateNode(ast.CreateDoubleOperand)
	b := aparser.CreateBuffer("foo")
	checkNotRead(t, b, e)
	b = aparser.CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	//ast := ast.CreateAST(b)
	//checkEvaluate(t, ast, 3.14)
}

func Test_stringLiteral(t *testing.T) {
	a := stringLiteral("&&")
	if a.str != "&&" {
		t.Error("str should be &&")
	}
}

func Test_stringLiteralExpression_Read(t *testing.T) {
	e := stringLiteral("&&")
	b := aparser.CreateBuffer("foo")
	checkNotRead(t, b, e)
}

func Test_caseInsensitiveStringLiteralExpression_Read(t *testing.T) {
	e := caseInsensitiveStringLiteral("Test")
	b := aparser.CreateBuffer("test")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = aparser.CreateBuffer("Test")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = aparser.CreateBuffer("TEST")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = aparser.CreateBuffer("foo")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
}
