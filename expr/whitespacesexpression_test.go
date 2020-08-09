package expr

import (
	"testing"
)

func TestWhitespacesExpression_Read(t *testing.T) {
	e := whiteSpaces()
	b := CreateBuffer("test")
	checkNotRead(t, b, e)
	checkPosition(t, b, 0)
	b = CreateBuffer("   ")
	checkRead(t, b, e)
	checkPosition(t, b, 3)
}
