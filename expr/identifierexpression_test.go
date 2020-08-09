package expr

import (
	"testing"
)

func TestIdentifierExpression_Read(t *testing.T) {
	e := identifier()
	b := CreateBuffer("test")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = CreateBuffer("1234")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
}
