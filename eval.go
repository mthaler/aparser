package aparser

import (
	"aparser/ast"
	"aparser/expr"
)

func Eval(s string) (interface{}, error) {
	b := expr.CreateBuffer(s)
	e := expr.ArithmeticExpression()
	expr.Parse(e, b)
	a := ast.CreateAST(b.Code.Code)
	return a.Evaluate()
}
