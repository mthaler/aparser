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
	operand := or(boolean, group)

	term.value = and(operand, zeroOrMore(and(or(a, o, xor), operand)))

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

//func logicalExpression() recursiveExpression {
//	term := recursive()
//
//	t := stringLiteral("true")
//	f := stringLiteral("false")
//	ws := optionalWhiteSpaces()
//	boolean := and(ws, and(or(t, f), ws))
//
//	a := and(ws, stringLiteral("&&"), ws)
//	o := and(ws, stringLiteral("||"), ws)
//	xor := and(ws, charLiteral('^'), ws)
//
//	op := and(or(a, o, xor))
//
//	term.value = or(op, boolean)
//
//	boolean.SetCreateNode(ast.CreateBooleanOperand)
//	a.SetCreateNode(ast.CreateBinaryOperation)
//	o.SetCreateNode(ast.CreateBinaryOperation)
//	xor.SetCreateNode(ast.CreateBinaryOperation)
//	term.SetCreateNode(ast.CreateBinaryLeftAssoc)
//
//	boolean.SetId("bool")
//
//	return term
//}