package expr

import "fmt"

type charLiteralExpression struct {
	*abstractExpression
	char rune
}

func charLiteral(c rune) charLiteralExpression {
	a := abstractExpression{}
	return charLiteralExpression{abstractExpression: &a, char: c}
}

func (c charLiteralExpression) parse(buffer *buffer) bool {
	if buffer.hasMoreChars() && buffer.currentChar() == c.char {
		buffer.incrementCurrentPosition()
		return true
	} else {
		return false
	}
}

func (c charLiteralExpression) String() string {
	return fmt.Sprint(c.char)
}
