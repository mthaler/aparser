package expr

import (
	"aparser"
	"aparser/ast"
)

func Eval(s string) (interface{}, error) {
	b := aparser.CreateBuffer(s)
	e := ArithmeticExpression()
	Parse(e, b)
	a := ast.CreateAST(&b.Code)
	return a.Evaluate()
}
