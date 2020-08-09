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
	"aparser"
)

type Expression interface {
	parse(buffer *aparser.Buffer) bool
	CreateNode() func(text string, code *aparser.Code) interface{}
	SetCreateNode(m func(text string, code *aparser.Code) interface{})
}

func Parse(e Expression, b *aparser.Buffer) bool {
	currentPosition := b.CurrentPosition()
	matchPosition := b.MatchPosition()
	b.SetMatchPosition(b.CurrentPosition())
	codeBlockStart := b.CurrentCodeBlockStart()
	codeBlockEnd := b.CurrentCodeBlockEndPosition()
	b.SetCurrentCodeBlockStartPosition(codeBlockEnd)

	match := e.parse(b)
	if match {
		result := onMatch(e, b)
		if result != nil && result != aparser.PT {
			b.SetCurrentCodeBlock(result)
		}
	} else {
		b.SetCurrentPosition(currentPosition)
		b.SetMatchPosition(matchPosition)
		b.SetCurrentCodeBlockEndPosition(codeBlockEnd)
	}
	b.SetMatchPosition(matchPosition)
	b.SetCurrentCodeBlockStartPosition(codeBlockStart)
	return match
}

func onMatch(e Expression, b *aparser.Buffer) interface{} {
	m := e.CreateNode()
	if m != nil {
		return m(b.CurrentMatch(), &b.Code)
	} else {
		return aparser.PT
	}
}
