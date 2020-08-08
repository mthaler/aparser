package expr

import (
	"aparser"
	"testing"
)

func Test_sequenceExpression_Read(t *testing.T) {
	ws := optionalWhiteSpaces()
	number := and(ws, doubleLiteral(), ws)
	s := sequence(number, number)
	b := aparser.CreateBuffer("3.14")
	checkNotRead(t, b, s)
	checkPosition(t, b, 0)
	b = aparser.CreateBuffer("3.14 3.15")
	checkRead(t, b, s)
	checkPosition(t, b, 9)
}

func Test_zeroOrMoreExpression_Read(t *testing.T) {
	ws := optionalWhiteSpaces()
	number := and(ws, doubleLiteral(), ws)
	e := zeroOrMore(number)
	b := aparser.CreateBuffer("foo")
	checkRead(t, b, e)
	checkPosition(t, b, 0)
	b = aparser.CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = aparser.CreateBuffer("3.14 3.15 3.16")
	checkRead(t, b, e)
	checkPosition(t, b, 14)
}

func Test_oneOrMoreExpression_Read(t *testing.T) {
	ws := optionalWhiteSpaces()
	number := and(ws, doubleLiteral(), ws)
	e := oneOrMore(number)
	b := aparser.CreateBuffer("foo")
	checkNotRead(t, b, e)
	b = aparser.CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = aparser.CreateBuffer("3.14 3.15 3.16")
	checkRead(t, b, e)
	checkPosition(t, b, 14)
}
