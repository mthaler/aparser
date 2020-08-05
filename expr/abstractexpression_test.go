package expr

import (
	"aparser"
	"aparser/ast"
	"reflect"
	"testing"
)

func Test_and(t *testing.T) {
	ws := optionalWhiteSpaces()
	a := and(ws, stringLiteral("&&"), ws)
	b := aparser.CreateBuffer("&&")
	checkRead(t, b, a)
	checkPosition(t, b, 2)
	b = aparser.CreateBuffer(" && ")
	checkRead(t, b, a)
	checkPosition(t, b, 4)
}

func Test_or(t *testing.T) {
	e := or(stringLiteral("true"), stringLiteral("false"))
	b := aparser.CreateBuffer("foo")
	if Parse(e, b) != false {
		t.Error("Expression should not read foo")
	}
	checkPosition(t, b, 0)
	b = aparser.CreateBuffer("true")
	if Parse(e, b) != true {
		t.Error("Expression should be able to read true")
	}
	checkPosition(t, b, 4)
	b = aparser.CreateBuffer("false")
	if Parse(e, b) != true {
		t.Error("Expression should be able to read true")
	}
	checkPosition(t, b, 5)
}

func checkPosition(t *testing.T, b *aparser.Buffer, pos int) {
	if b.CurrentPosition() != pos {
		t.Errorf("Current position should be %d, actual %d", pos, b.CurrentPosition())
	}
}

func checkRead(t *testing.T, b *aparser.Buffer, e Expression) {
	if !Parse(e, b) {
		t.Errorf("Expression should be able to read %s", string(b.Rest()))
	}
}

func checkNotRead(t *testing.T, b *aparser.Buffer, e Expression) {
	if Parse(e, b) {
		t.Errorf("Expression should not be able to read %v", string(b.Rest()))
	}
}

func checkEvaluate(t *testing.T, o, expected interface{}) {
	n, ok := o.(ast.Node)
	if !ok {
		t.Errorf("Object %v not a node", o)
	}
	r, err := n.Evaluate()
	if err != nil {
		t.Errorf("Could not evalue node %v", n)
	}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Expected %v, actual %v", expected, r)
	}
}
