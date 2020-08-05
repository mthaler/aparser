package expr

import "aparser/ast"

func arithmeticExpression() recursiveExpression {

	expression := recursive()
	operand := recursive()

	ws := optionalWhiteSpaces()
	number := and(ws, doubleLiteral(), ws)
	funcName := identifier()

	add := and(ws, charLiteral('+'), ws)
	sub := and(ws, charLiteral('-'), ws)
	mul := and(ws, charLiteral('*'), ws)
	div := and(ws, charLiteral('/'), ws)
	exp := and(ws, charLiteral('^'), ws)
	neg := and(ws, charLiteral('-'), ws)

	function := and(funcName, charLiteral('('), &expression, charLiteral(')'))
	group := and(ws, charLiteral('('), ws, &expression, charLiteral(')'), ws)
	operand2 := or(number, function, group)
	operand.value = or(and(neg, &operand), operand2)
	powerOp := and(ws, exp, ws, &operand, ws)
	power := and(ws, &operand, ws, zeroOrMore(powerOp))
	mulDivOp := and(ws, or(mul, div), ws, power, ws)
	term := and(ws, power, ws, zeroOrMore(mulDivOp))
	addSubOp := and(ws, or(add, sub), ws, term, ws)
	term2 := and(ws, term, ws, zeroOrMore(addSubOp))
	expression.value = term2

	number.SetCreateNode(ast.CreateDoubleOperand)
	add.SetCreateNode(ast.CreateBinaryOperation)
	sub.SetCreateNode(ast.CreateBinaryOperation)
	mul.SetCreateNode(ast.CreateBinaryOperation)
	div.SetCreateNode(ast.CreateBinaryOperation)
	exp.SetCreateNode(ast.CreateBinaryOperation)
	neg.SetCreateNode(ast.CreateUnaryOperation)
	funcName.SetCreateNode(ast.CreateIdentifier)
	term.SetCreateNode(ast.CreateBinaryLeftAssoc)
	expression.SetCreateNode(ast.CreateBinaryLeftAssoc)
	power.SetCreateNode(ast.CreateBinaryRightAssoc)
	function.SetCreateNode(ast.CreateFunction)
	operand.SetCreateNode(ast.CreateUnaryPrefix)

	group.SetId("group")
	operand.SetId("factor")
	term.SetId("term")
	power.SetId("power")
	number.SetId("number")
	add.SetId("add")
	sub.SetId("sub")
	mul.SetId("mul")
	div.SetId("div")
	exp.SetId("exp")
	neg.SetId("neg")
	expression.SetId("expression")

	return expression
}
