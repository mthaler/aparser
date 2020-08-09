package expr

import (
	"unicode"
)

type identifierExpression struct {
	*abstractExpression
}

func identifier() identifierExpression {
	a := abstractExpression{}
	return identifierExpression{abstractExpression: &a}
}

func (i identifierExpression) parse(buffer *Buffer) bool {
	count := 0
	for buffer.HasMoreChars() {
		if count == 0 {
			if unicode.IsLetter(buffer.CurrentChar()) || buffer.CurrentChar() == '_' {
				buffer.IncrementCurrentPosition()
				count++
			} else {
				break
			}
		} else {
			if unicode.IsLetter(buffer.CurrentChar()) || unicode.IsDigit(buffer.CurrentChar()) || buffer.CurrentChar() == '_' {
				buffer.IncrementCurrentPosition()
				count++
			} else {
				break
			}
		}
	}
	return count > 0
}
