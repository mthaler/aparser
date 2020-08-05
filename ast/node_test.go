package ast

import (
	"reflect"
	"testing"
)

func checkEvaluate(t *testing.T, o, expected interface{}) {
	n, ok := o.(Node)
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