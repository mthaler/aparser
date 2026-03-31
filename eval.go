/*
Package aparser implements a simple library for creating arithmetic expression parsers.
It also provides a pre-build arithmetic expression parser that supports parsing common arithmetic expressions.
*/
package aparser

import (
	"github.com/mthaler/aparser/ast"
	"github.com/mthaler/aparser/expr"
)

// Eval evaluates the given arithmetic expression using the pre-build arithmetic expression parser.
func Eval(s string) (interface{}, error) {
	b := expr.CreateBuffer(s)
	e := expr.ArithmeticExpression()
	expr.Parse(e, b)
	a := ast.CreateAST(b.Code.Code)
	return a.Evaluate()
}
