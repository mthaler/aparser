package expr

import (
	"aparser"
	"testing"
)

func Test_orderedChoiceExpression_Read(t *testing.T) {
	e := orderedChoice(stringLiteral("hello"), stringLiteral("world"))
	b := aparser.CreateBuffer("hello")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	b = aparser.CreateBuffer("world")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	b = aparser.CreateBuffer("foo")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
}
