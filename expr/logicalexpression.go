package expr

import "aparser/ast"

func logicalExpression() recursiveExpression {
	term := recursive()

	t := stringLiteral("true")
	f := stringLiteral("false")
	ws := optionalWhiteSpaces()
	boolean := and(ws, and(or(t, f), ws))
	a := and(ws, stringLiteral("&&"), ws)
	o := and(ws, stringLiteral("||"), ws)
	xor := and(ws, charLiteral('^'), ws)

	group := and(ws, charLiteral('('), ws, &term, ws, charLiteral(')'), ws)
	op := or(boolean, group)

	term.value = and(op, zeroOrMore(and(or(a, o, xor), op)))

	boolean.SetCreateNode(ast.CreateBooleanOperand)
	a.SetCreateNode(ast.CreateBinaryOperation)
	o.SetCreateNode(ast.CreateBinaryOperation)
	xor.SetCreateNode(ast.CreateBinaryOperation)
	term.SetCreateNode(ast.CreateBinaryLeftAssoc)

	boolean.SetId("bool")
	a.SetId("and")
	o.SetId("or")
	xor.SetId("xor")

	return term
}