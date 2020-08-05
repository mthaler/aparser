package ast

import "testing"

func TestUnaryOperationEvaluate(t *testing.T) {
	d := doubleOperand{Value: 3.14}
	u := unaryOperation{node: d, funcName: "-"}
	checkEvaluate(t, u, -3.14)
	u = unaryOperation{node: d, funcName: "round"}
	checkEvaluate(t, u, 3.0)
}
