package expr

type stringLiteralExpression struct {
	*abstractExpression
	str string
}

func stringLiteral(s string) stringLiteralExpression {
	a := abstractExpression{}
	return stringLiteralExpression{abstractExpression: &a, str: s}
}

func (s stringLiteralExpression) parse(buffer *buffer) bool {
	if !buffer.hasMoreChars() {
		return false
	} else if len(buffer.rest()) < len(s.str) {
		return false
	} else if string(buffer.rest()[:len(s.str)]) == s.str {
		buffer.incrementCurrentPositionBy(len(s.str))
		return true
	} else {
		return false
	}
}
