package ast

import (
	"aparser"
	"reflect"
	"testing"
)

func Test_createDoubleOperand(t *testing.T) {
	c := aparser.CreateCode()
	o := CreateDoubleOperand("3.14", c)
	checkEvaluate(t, o, 3.14)
}

func Test_createBooleanOperand(t *testing.T) {
	c := aparser.CreateCode()
	o := CreateBooleanOperand("false", c)
	checkEvaluate(t, o, false)
	o = CreateBooleanOperand("true", c)
	checkEvaluate(t, o, true)
}

func Test_createStringOperand(t *testing.T) {
	c := aparser.CreateCode()
	o := CreateStringOperand("\"foo\"", c)
	checkEvaluate(t, o, "foo")
}

func Test_swapElements(t *testing.T) {
	// swap first two elements
	a := []int{1, 2, 3}
	a[0], a[1] = a[1], a[0]
	ok := reflect.DeepEqual(a, []int{2, 1, 3})
	if !ok {
		t.Error("Result should be [2, 1, 3]")
	}
}

func Test_copyBinopLeft(t *testing.T) {
	code := []int{1, 2, 3, 4, 5}
	binop := make([]int, 3)
	copy(binop, code[:3])
	ok := reflect.DeepEqual(binop, []int{1, 2, 3})
	if !ok {
		t.Error("Result should be [1, 2, 3]")
	}
}

func Test_copyBinopRight(t *testing.T) {
	code := []int{1, 2, 3, 4, 5}
	binop := make([]int, 3)
	copy(binop, code[len(code)-3:])
	ok := reflect.DeepEqual(binop, []int{3, 4, 5})
	if !ok {
		t.Error("Result should be [3, 4, 5]")
	}
}

func Test_createResult(t *testing.T) {
	code := []int{1, 2, 3, 4, 5, 6}
	result := make([]int, len(code)-2)
	result[0] = 7
	copy(result[1:], code[3:])
	ok := reflect.DeepEqual(result, []int{7, 4, 5, 6})
	if !ok {
		t.Error("Result should be [7, 4, 5, 6]")
	}
}

func Test_createBinaryLeftAssocNode(t *testing.T) {
	c := aparser.CreateCode()
	o0 := CreateDoubleOperand("3", c)
	o1 := CreateBinaryOperation("-", c)
	o2 := CreateDoubleOperand("4", c)
	code := []interface{}{o0, o1, o2}
	n0 := createBinaryLeftAssocNode(code)
	n1 := createNode([]interface{}{n0})
	checkEvaluate(t, n1, -1.0)
}
