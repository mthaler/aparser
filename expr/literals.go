package expr

import (
	"aparser"
	"regexp"
	"strings"
)

type charLiteralExpression struct {
	*abstractExpression
	char rune
}

func charLiteral(c rune) charLiteralExpression {
	a := abstractExpression{}
	return charLiteralExpression{abstractExpression: &a, char: c}
}

func (c charLiteralExpression) parse(buffer *aparser.Buffer) bool {
	if buffer.HasMoreChars() && buffer.CurrentChar() == c.char {
		buffer.IncrementCurrentPosition()
		return true
	} else {
		return false
	}
}

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

func (d doubleLiteralExpression) parse(buffer *aparser.Buffer) bool {
	if buffer.HasMoreChars() {
		loc := d.regex.FindStringIndex(string(buffer.Rest()))
		text := make([]rune, 0)
		if loc != nil {
			text = buffer.Rest()[loc[0]:loc[1]]
			buffer.IncrementCurrentPositionBy(len(text))
		}
		if len(text) > 0 {
			return true
		}
	}
	return false
}

type stringLiteralExpression struct {
	*abstractExpression
	str string
}

func stringLiteral(s string) stringLiteralExpression {
	a := abstractExpression{}
	return stringLiteralExpression{abstractExpression: &a, str: s}
}

func (s stringLiteralExpression) parse(buffer *aparser.Buffer) bool {
	if !buffer.HasMoreChars() {
		return false
	} else if len(buffer.Rest()) < len(s.str) {
		return false
	} else if string(buffer.Rest()[:len(s.str)]) == s.str {
		buffer.IncrementCurrentPositionBy(len(s.str))
		return true
	} else {
		return false
	}
}

type caseInsensitiveStringLiteralExpression struct {
	*abstractExpression
	str string
}

func caseInsensitiveStringLiteral(s string) caseInsensitiveStringLiteralExpression {
	a := abstractExpression{}
	return caseInsensitiveStringLiteralExpression{abstractExpression: &a, str: strings.ToLower(s)}
}

func (c caseInsensitiveStringLiteralExpression) parse(buffer *aparser.Buffer) bool {
	if !buffer.HasMoreChars() {
		return false
	} else if len(buffer.Rest()) < len(c.str) {
		return false
	} else if strings.ToLower(string(buffer.Rest()[:len(c.str)])) == c.str {
		buffer.SetCurrentPosition(buffer.CurrentPosition() + len(c.str))
		return true
	} else {
		return false
	}
}
