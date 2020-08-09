package expr

import (
	"strings"
)

type quotedStringExpression struct {
	*abstractExpression
	quoteChars  string
	escapeChars string
}

func quotedString(qcs string, ecs string) quotedStringExpression {
	a := abstractExpression{}
	return quotedStringExpression{abstractExpression: &a, quoteChars: qcs, escapeChars: ecs}
}

func (q quotedStringExpression) parse(buffer *buffer) bool {
	var quoteChar rune
	if buffer.hasMoreChars() && q.isQuote(buffer.currentChar()) {
		quoteChar = buffer.currentChar()
		buffer.incrementCurrentPosition()
	} else {
		return false
	}
	for buffer.hasMoreChars() {
		if buffer.currentChar() == quoteChar {
			buffer.incrementCurrentPosition()
			return true
		}
		for buffer.hasMoreChars() && buffer.currentChar() != quoteChar && !q.isEscape(buffer.currentChar()) {
			buffer.incrementCurrentPosition()
		}
		if !buffer.hasMoreChars() {
			return false
		}
		if q.isEscape(buffer.currentChar()) {
			if len(buffer.rest()) < 2 {
				return false
			}
			buffer.setCurrentPosition(buffer.currentPosition + 2)
		}
	}
	return false
}

func (q quotedStringExpression) isQuote(c rune) bool {
	return strings.ContainsRune(q.quoteChars, c)
}

func (q quotedStringExpression) isEscape(c rune) bool {
	return strings.ContainsRune(q.escapeChars, c)
}
