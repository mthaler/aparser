package expr

import (
	"testing"
)

func Test_sequenceExpression_Read(t *testing.T) {
	ws := optionalWhiteSpaces()
	number := and(ws, doubleLiteral(), ws)
	s := sequence(number, number)
	b := CreateBuffer("3.14")
	checkNotRead(t, b, s)
	checkPosition(t, b, 0)
	b = CreateBuffer("3.14 3.15")
	checkRead(t, b, s)
	checkPosition(t, b, 9)
}

func Test_zeroOrMoreExpression_Read(t *testing.T) {
	ws := optionalWhiteSpaces()
	number := and(ws, doubleLiteral(), ws)
	e := zeroOrMore(number)
	b := CreateBuffer("foo")
	checkRead(t, b, e)
	checkPosition(t, b, 0)
	b = CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = CreateBuffer("3.14 3.15 3.16")
	checkRead(t, b, e)
	checkPosition(t, b, 14)
}

func Test_oneOrMoreExpression_Read(t *testing.T) {
	ws := optionalWhiteSpaces()
	number := and(ws, doubleLiteral(), ws)
	e := oneOrMore(number)
	b := CreateBuffer("foo")
	checkNotRead(t, b, e)
	b = CreateBuffer("3.14")
	checkRead(t, b, e)
	checkPosition(t, b, 4)
	b = CreateBuffer("3.14 3.15 3.16")
	checkRead(t, b, e)
	checkPosition(t, b, 14)
}


func Test_and(t *testing.T) {
	ws := optionalWhiteSpaces()
	a := and(ws, stringLiteral("&&"), ws)
	b := CreateBuffer("&&")
	checkRead(t, b, a)
	checkPosition(t, b, 2)
	b = CreateBuffer(" && ")
	checkRead(t, b, a)
	checkPosition(t, b, 4)
}

func Test_or(t *testing.T) {
	e := or(stringLiteral("true"), stringLiteral("false"))
	b := CreateBuffer("foo")
	if Parse(e, b) != false {
		t.Error("Expression should not read foo")
	}
	checkPosition(t, b, 0)
	b = CreateBuffer("true")
	if Parse(e, b) != true {
		t.Error("Expression should be able to read true")
	}
	checkPosition(t, b, 4)
	b = CreateBuffer("false")
	if Parse(e, b) != true {
		t.Error("Expression should be able to read true")
	}
	checkPosition(t, b, 5)
}