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

func (w whitespacesExpression) parse(buffer *Buffer) bool {
	pos := buffer.CurrentPosition()
	for buffer.HasMoreChars() && unicode.IsSpace(buffer.CurrentChar()) {
		buffer.IncrementCurrentPosition()
	}
	return buffer.CurrentPosition() > pos
}

func optionalWhiteSpaces() Expression {
	return optional(whiteSpaces())
}
