package expr

import (
	"regexp"
	"strings"
)

var doubleLiteralRegex = "^\\d+(\\.\\d*)?([eE][+-]?\\d+)?"

type doubleLiteralExpression struct {
	*abstractExpression
	regex *regexp.Regexp
}

func doubleLiteral() doubleLiteralExpression {
	a := abstractExpression{}
	r := regexp.MustCompile(doubleLiteralRegex)
	return doubleLiteralExpression{abstractExpression: &a, regex: r}
}

func (d doubleLiteralExpression) parse(buffer *buffer) bool {
	if buffer.hasMoreChars() {
		loc := d.regex.FindStringIndex(string(buffer.rest()))
		text := make([]rune, 0)
		if loc != nil {
			text = buffer.rest()[loc[0]:loc[1]]
			buffer.incrementCurrentPositionBy(len(text))
		}
		if len(text) > 0 {
			return true
		}
	}
	return false
}

type caseInsensitiveStringLiteralExpression struct {
	*abstractExpression
	str string
}

func caseInsensitiveStringLiteral(s string) caseInsensitiveStringLiteralExpression {
	a := abstractExpression{}
	return caseInsensitiveStringLiteralExpression{abstractExpression: &a, str: strings.ToLower(s)}
}

func (c caseInsensitiveStringLiteralExpression) parse(buffer *buffer) bool {
	if !buffer.hasMoreChars() {
		return false
	} else if len(buffer.rest()) < len(c.str) {
		return false
	} else if strings.ToLower(string(buffer.rest()[:len(c.str)])) == c.str {
		buffer.setCurrentPosition(buffer.currentPosition + len(c.str))
		return true
	} else {
		return false
	}
}
