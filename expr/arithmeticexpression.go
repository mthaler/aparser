package expr

import "aparser/ast"

/*
Creates an arithmetic expression parser that can parse common arithmetic expressions

Supported operations:
- unary: -
- binary: +, -, *, /, ^ (power), % (modulo)
- ternary: cond ? a : b, if (cond) a else b
- logic: &&, ||, !, ^ (exclusive or)
- relational: ==, !=, <, <=, >, >=
- functions: abs, acos, asin, atan, cos, cosh, exp, log, log10, sign, sin, sinh, sqrt, tan, tanh, round
 */
func ArithmeticExpression() recursiveExpression {

	ws := optionalWhiteSpaces()

	t := stringLiteral("true")
	f := stringLiteral("false")

	num := and(ws, doubleLiteral(), ws)
	boolean := and(ws, or(t, f), ws)
	_string := quotedString("\"", "")

	funcname := and(ws, identifier(), ws)
	_if := stringLiteral("if")

	neg := charLiteral('-')

	add := charLiteral('+')
	sub := charLiteral('-')
	mul := charLiteral('*')
	div := charLiteral('/')
	exp := charLiteral('^')
	mod := charLiteral('%')

	a := stringLiteral("&&")
	o := stringLiteral("||")
	not := charLiteral('!')
	xor := charLiteral('^')

	lt := charLiteral('<')
	le := stringLiteral("<=")
	ge := stringLiteral(">=")
	gt := charLiteral('>')

	eq := stringLiteral("==")
	ne := stringLiteral("!=")

	o7 := and(ws, or(neg, not), ws)
	o6 := and(ws, exp, ws)
	o5 := and(ws, or(mul, div, mod), ws)
	o4 := and(ws, or(add, sub), ws)
	o3 := and(ws, or(le, ge, lt, gt), ws)
	o2 := and(ws, or(eq, ne), ws)
	o1 := and(ws, or(a, o, xor), ws)

	cond := recursive()
	function := and(funcname, charLiteral('('), ws, &cond, ws, zeroOrMore(and(ws, charLiteral(','), &cond)), charLiteral(')'))
	p100 := or(boolean, num, _string, function, and(charLiteral('('), &cond, charLiteral(')')))
	p7 := recursive()
	p7.value = or(and(o7, &p7, ws), p100)
	p6 := and(ws, p7, ws, zeroOrMore(and(o6, p7, ws)))
	p5 := and(ws, p6, ws, zeroOrMore(and(o5, p6, ws)))
	p4 := and(ws, p5, ws, zeroOrMore(and(o4, p5, ws)))
	p3 := and(ws, p4, ws, zeroOrMore(and(o3, p4, ws)))
	p2 := and(ws, p3, ws, zeroOrMore(and(o2, p3, ws)))
	p1 := and(ws, p2, ws, zeroOrMore(and(o1, p2, ws)))
	cond1 := and(ws, _if, ws, charLiteral('('), ws, p1, ws, charLiteral(')'), ws, p1, ws, stringLiteral("else"), ws, p1, ws)
	cond2 := and(p1, optional(and(charLiteral('?'), ws, p1, ws, charLiteral(':'), ws, p1, ws)))
	cond.value = or(cond1, cond2)

	num.SetCreateNode(ast.CreateDoubleOperand)
	boolean.SetCreateNode(ast.CreateBooleanOperand)
	_string.SetCreateNode(ast.CreateStringOperand)

	funcname.SetCreateNode(ast.CreateIdentifier)
	_if.SetCreateNode(ast.CreateIdentifier)

	neg.SetCreateNode(ast.CreateUnaryOperation)

	add.SetCreateNode(ast.CreateBinaryOperation)
	sub.SetCreateNode(ast.CreateBinaryOperation)
	mul.SetCreateNode(ast.CreateBinaryOperation)
	div.SetCreateNode(ast.CreateBinaryOperation)
	mod.SetCreateNode(ast.CreateBinaryOperation)
	exp.SetCreateNode(ast.CreateBinaryOperation)

	a.SetCreateNode(ast.CreateBinaryOperation)
	o.SetCreateNode(ast.CreateBinaryOperation)
	xor.SetCreateNode(ast.CreateBinaryOperation)
	not.SetCreateNode(ast.CreateUnaryOperation)

	lt.SetCreateNode(ast.CreateBinaryOperation)
	le.SetCreateNode(ast.CreateBinaryOperation)
	gt.SetCreateNode(ast.CreateBinaryOperation)
	ge.SetCreateNode(ast.CreateBinaryOperation)

	eq.SetCreateNode(ast.CreateBinaryOperation)
	ne.SetCreateNode(ast.CreateBinaryOperation)

	p7.SetCreateNode(ast.CreateUnaryPrefix)
	p6.SetCreateNode(ast.CreateBinaryRightAssoc)
	p5.SetCreateNode(ast.CreateBinaryLeftAssoc)
	p4.SetCreateNode(ast.CreateBinaryLeftAssoc)
	p3.SetCreateNode(ast.CreateBinaryLeftAssoc)
	p2.SetCreateNode(ast.CreateBinaryLeftAssoc)
	p1.SetCreateNode(ast.CreateBinaryLeftAssoc)

	function.SetCreateNode(ast.CreateFunction)
	cond1.SetCreateNode(ast.CreateFunction)
	cond2.SetCreateNode(ast.CreateTernaryConditional)

	return cond
}
