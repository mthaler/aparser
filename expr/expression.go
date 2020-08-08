/*
Package expr implements a simple library for parsing expressions.

There are two types of expressions: simple expressions and complex expressions.

Examples for simple expressions are charLiteralExpression which parses a char literal, e.g. '(' or ')',
stringLiteralExpression which parses a string, e.g. "false" or "true".

An example for a complex expression is orderedChoiceExpresssion which can be used to e.g. build a boolean
literal expression that parses both "false" or "true".
*/
package expr

import (
	"aparser"
)

type Expression interface {
	Id() string
	SetId(id string)
	CreateNode() func(buffer *aparser.Buffer) interface{}
	SetCreateNode(m func(buffer *aparser.Buffer) interface{})
	parse(buffer *aparser.Buffer) bool
}

func Parse(e Expression, buffer *aparser.Buffer) bool {
	currentPosition := buffer.CurrentPosition()
	matchPosition := buffer.MatchPosition()
	buffer.SetMatchPosition(buffer.CurrentPosition())
	codeBlockStart := buffer.CurrentCodeBlockStart()
	codeBlockEnd := buffer.CurrentCodeBlockEndPosition()
	buffer.SetCurrentCodeBlockStartPosition(codeBlockEnd)

	match := e.parse(buffer)
	if match {
		result := onMatch(e, buffer)
		if result != nil && result != aparser.PT {
			buffer.SetCurrentCodeBlock(result)
		}
	} else {
		buffer.SetCurrentPosition(currentPosition)
		buffer.SetMatchPosition(matchPosition)
		buffer.SetCurrentCodeBlockEndPosition(codeBlockEnd)
	}
	buffer.SetMatchPosition(matchPosition)
	buffer. SetCurrentCodeBlockStartPosition(codeBlockStart)
	return match
}

func onMatch(e Expression, buffer *aparser.Buffer) interface{} {
	m := e.CreateNode()
	if m != nil {
		return m(buffer)
	} else {
		return aparser.PT
	}
}
