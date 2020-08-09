package expr

import (
	"aparser/ast"
	"reflect"
	"testing"
)

func checkPosition(t *testing.T, b *buffer, pos int) {
	if b.currentPosition != pos {
		t.Errorf("Current position should be %d, actual %d", pos, b.currentPosition)
	}
}

func checkRead(t *testing.T, b *buffer, e Expression) {
	if !Parse(e, b) {
		t.Errorf("Expression should be able to read %s", string(b.rest()))
	}
}

func checkNotRead(t *testing.T, b *buffer, e Expression) {
	if Parse(e, b) {
		t.Errorf("Expression should not be able to read %v", string(b.rest()))
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
