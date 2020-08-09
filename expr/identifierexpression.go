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

func (i identifierExpression) parse(buffer *buffer) bool {
	count := 0
	for buffer.hasMoreChars() {
		if count == 0 {
			if unicode.IsLetter(buffer.currentChar()) || buffer.currentChar() == '_' {
				buffer.incrementCurrentPosition()
				count++
			} else {
				break
			}
		} else {
			if unicode.IsLetter(buffer.currentChar()) || unicode.IsDigit(buffer.currentChar()) || buffer.currentChar() == '_' {
				buffer.incrementCurrentPosition()
				count++
			} else {
				break
			}
		}
	}
	return count > 0
}
