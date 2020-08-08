package expr

import (
	"aparser"
	"testing"
)

func Test_quotedStringExpression_parse(t *testing.T) {
	e := quotedString("\"", "")
	b := aparser.CreateBuffer("test")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
	b = aparser.CreateBuffer("\"test\"")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
}