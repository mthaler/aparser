package expr

import (
	"aparser"
	"testing"
)

func TestLogicalExpression(t *testing.T) {
	e := logicalExpression()
	b := aparser.CreateBuffer("false || true")
	checkRead(t, b, e)
	checkPosition(t, b, 13)
	b = aparser.CreateBuffer("false || (true && false)")
	checkRead(t, b, e)
	checkPosition(t, b, 24)
}
