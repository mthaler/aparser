package expr

import (
	"testing"
)

func Test_quotedStringExpression_parse(t *testing.T) {
	e := quotedString("\"", "")
	b := CreateBuffer("test")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
	b = CreateBuffer("\"test\"")
	checkRead(t, b, e)
	checkPosition(t, b, 6)
}
