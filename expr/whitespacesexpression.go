package expr

import (
	"unicode"
)

type whitespacesExpression struct {
	*abstractExpression
}

func whiteSpaces() whitespacesExpression {
	a := abstractExpression{}
	return whitespacesExpression{abstractExpression: &a}
}

func (w whitespacesExpression) parse(buffer *buffer) bool {
	pos := buffer.currentPosition
	for buffer.hasMoreChars() && unicode.IsSpace(buffer.currentChar()) {
		buffer.incrementCurrentPosition()
	}
	return buffer.currentPosition > pos
}

func optionalWhiteSpaces() Expression {
	return optional(whiteSpaces())
}
