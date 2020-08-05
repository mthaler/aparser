package ast

import "testing"

func Test_ternaryOperation_Evaluate(t *testing.T) {
	l := newBoolOperand(true)
	m := newDoubleOperand(3)
	r := newDoubleOperand(4)
	o := ternaryOperation{left: l, middle: m, right: r, operator: "?:"}
	checkEvaluate(t, o, 3.0)
	l = newBoolOperand(false)
	o = ternaryOperation{left: l, middle: m, right: r, operator: "?:"}
	checkEvaluate(t, o, 4.0)
}
