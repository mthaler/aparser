package expr

import "testing"

func Test_charLiteralExpression_Read(t *testing.T) {
	e := charLiteral('s')
	b := CreateBuffer("test")
	checkNotRead(t, b, e)
	b.incrementCurrentPosition()
	checkNotRead(t, b, e)
	b.incrementCurrentPosition()
	checkRead(t, b, e)
}
