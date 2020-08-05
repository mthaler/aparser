package ast

import "testing"

func Test_createUnaryNode(t *testing.T) {
	u := unaryOperation{funcName: "-"}
	o := newDoubleOperand(3.14)
	code := []interface{}{u, o}
	n := createUnaryNode(code)
	v, err := n.Evaluate()
	if err != nil {
		t.Error(err)
	}
	if v != -3.14 {
		t.Error("Result should be -3.14")
	}
}

func Test_createBinaryNode(t *testing.T) {
	b := binaryOperation{operator: "+"}
	l := doubleOperand{Value: 3}
	r := doubleOperand{Value: 4}
	code := []interface{}{b, l, r}
	n := createBinaryNode(code)
	v, err := n.Evaluate()
	if err != nil {
		t.Error(err)
	}
	if v != 7.0 {
		t.Error("Result should be 7")
	}
}