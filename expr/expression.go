/*
Package expr implements a simple library for parsing arithmetic expressions.

There are two types of expressions: simple expressions and complex expressions.

Examples for simple expressions are charLiteralExpression which parses a char literal, e.g. '(' or ')',
stringLiteralExpression which parses a string, e.g. "false" or "true".

An example for a complex expression is orderedChoiceExpression which can be used to e.g. build a boolean
literal expression that parses both "false" or "true".
*/
package expr

import (
	"aparser/ast"
)

type Expression interface {
	parse(buffer *buffer) bool
	CreateNode() func(text string, code *ast.Code) interface{}
	SetCreateNode(m func(text string, code *ast.Code) interface{})
}

func Parse(e Expression, b *buffer) bool {
	currentPosition := b.currentPosition
	matchPosition := b.matchPosition
	b.setMatchPosition(b.currentPosition)
	codeBlockStart := b.CurrentCodeBlockStart()
	codeBlockEnd := b.CurrentCodeBlockEndPosition()
	b.SetCurrentCodeBlockStartPosition(codeBlockEnd)

	match := e.parse(b)
	if match {
		result := onMatch(e, b)
		if result != nil && result != ast.PT {
			b.SetCurrentCodeBlock(result)
		}
	} else {
		b.setCurrentPosition(currentPosition)
		b.setMatchPosition(matchPosition)
		b.SetCurrentCodeBlockEndPosition(codeBlockEnd)
	}
	b.setMatchPosition(matchPosition)
	b.SetCurrentCodeBlockStartPosition(codeBlockStart)
	return match
}

func onMatch(e Expression, b *buffer) interface{} {
	m := e.CreateNode()
	if m != nil {
		return m(b.currentMatch(), &b.Code)
	} else {
		return ast.PT
	}
}
