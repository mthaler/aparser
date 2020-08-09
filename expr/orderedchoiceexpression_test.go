package expr

import (
	"testing"
)

func Test_orderedChoiceExpression_Read(t *testing.T) {
	e := orderedChoice(stringLiteral("hello"), stringLiteral("world"))
	b := CreateBuffer("hello")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	b = CreateBuffer("world")
	checkRead(t, b, e)
	checkPosition(t, b, 5)
	b = CreateBuffer("foo")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
}
