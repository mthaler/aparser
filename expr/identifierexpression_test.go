package expr

import (
	"aparser"
	"testing"
)

func TestIdentifierExpression_Read(t *testing.T) {
	e := identifier()
	b := aparser.CreateBuffer("test")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = aparser.CreateBuffer("1234")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
}