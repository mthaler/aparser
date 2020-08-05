package expr

import (
	"aparser"
	"testing"
)

func TestWhitespacesExpression_Read(t *testing.T) {
	e := whiteSpaces()
	b := aparser.CreateBuffer("test")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
	b = aparser.CreateBuffer("   ")
	checkRead(t, b, e)
	checkPosition(t, b, 3)
}
