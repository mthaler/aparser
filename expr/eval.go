package expr

import (
	"aparser/ast"
)

func Eval(s string) (interface{}, error) {
	b := CreateBuffer(s)
	e := ArithmeticExpression()
	Parse(e, b)
	a := ast.CreateAST(&b.Code)
	return a.Evaluate()
}
