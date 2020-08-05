package ast

import (
	"reflect"
	"testing"
)

func TestBinaryOperationEvaluate(t *testing.T) {
	l := newDoubleOperand(3)
	r := newDoubleOperand(4)
	b := binaryOperation{left: l, right: r, operator: "+"}
	checkEvaluate(t, b, 7.0)
	sl := newStringOperand("hello")
	sr := newStringOperand("world")
	b = binaryOperation{left: sl, right: sr, operator: "+"}
	checkEvaluate(t, b, "helloworld")
}

func Test_evaluateNumeric2(t *testing.T) {
	r, err := evaluateNumeric(3, 4, "+")
	checkResult(t, r, err, 7.0)
	r, err = evaluateNumeric(3, 4, "-")
	checkResult(t, r, err, -1.0)
	r, err = evaluateNumeric(3, 4, "*")
	checkResult(t, r, err, 12.0)
}


func Test_evaluateBoolean(t *testing.T) {
	r, err := evaluateBoolean(true, true, "&&")
	checkResult(t, r, err, true)
	r, err = evaluateBoolean(true, false, "&&")
	checkResult(t, r, err, false)
}

func checkResult(t *testing.T, r interface{}, err error, expected interface{}) {
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Result should be %v, actual %v", expected, r)
	}
}