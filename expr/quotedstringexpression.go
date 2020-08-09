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

func (q quotedStringExpression) parse(buffer *Buffer) bool {
	var quoteChar rune
	if buffer.HasMoreChars() && q.isQuote(buffer.CurrentChar()) {
		quoteChar = buffer.CurrentChar()
		buffer.IncrementCurrentPosition()
	} else {
		return false
	}
	for buffer.HasMoreChars() {
		if buffer.CurrentChar() == quoteChar {
			buffer.IncrementCurrentPosition()
			return true
		}
		for buffer.HasMoreChars() && buffer.CurrentChar() != quoteChar && !q.isEscape(buffer.CurrentChar()) {
			buffer.IncrementCurrentPosition()
		}
		if !buffer.HasMoreChars() {
			return false
		}
		if q.isEscape(buffer.CurrentChar()) {
			if len(buffer.Rest()) < 2 {
				return false
			}
			buffer.SetCurrentPosition(buffer.CurrentPosition() + 2)
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
