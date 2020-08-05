package expr

import "aparser/ast"

func basicArithmeticExpression() Expression {

	ws := optionalWhiteSpaces()
	number := and(ws, doubleLiteral(), ws)

	add := and(ws, charLiteral('+'), ws)
	sub := and(ws, charLiteral('-'), ws)

	addSubOp := and(ws, or(add, sub), ws)

	term := or(and(number, addSubOp, number), number)

	number.SetCreateNode(ast.CreateDoubleOperand)
	add.SetCreateNode(ast.CreateBinaryOperation)
	sub.SetCreateNode(ast.CreateBinaryOperation)
	term.SetCreateNode(ast.CreateBinaryLeftAssoc)

	return term
}
