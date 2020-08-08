/*
Package ast implements a simple library for building ab abstract syntax tree for arithmetic expressions.

There are two types of nodes: operands and operations. An operand represents values, e.g. a double, a boolean or a string.
Operations represent arithmetic operations, e.g. addition, subtraction etc.
*/
package ast

type Node interface {
	Evaluate() (interface{}, error)
}
