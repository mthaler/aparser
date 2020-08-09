package ast

import (
	"aparser"
	"reflect"
	"testing"
)

func Test_createDoubleOperand(t *testing.T) {
	b := aparser.CreateBuffer("3.14")
	b.SetCurrentPosition(len("3.14"))
	o := CreateDoubleOperand(b.CurrentMatch(), &b.Code)
	checkEvaluate(t, o, 3.14)
}

func Test_createBooleanOperand(t *testing.T) {
	b := aparser.CreateBuffer("false")
	b.SetCurrentPosition(len("false"))
	o := CreateBooleanOperand(b.CurrentMatch(), &b.Code)
	checkEvaluate(t, o, false)
	b = aparser.CreateBuffer("true")
	b.SetCurrentPosition(len("true"))
	o = CreateBooleanOperand(b.CurrentMatch(), &b.Code)
	checkEvaluate(t, o, true)
}

func Test_createStringOperand(t *testing.T) {
	b := aparser.CreateBuffer("\"foo\"")
	b.SetCurrentPosition(len("\"foo\""))
	o := CreateStringOperand(b.CurrentMatch(), &b.Code)
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
	b := aparser.CreateBuffer("3")
	b.SetCurrentPosition(len("3"))
	o0 := CreateDoubleOperand(b.CurrentMatch(), &b.Code)
	b = aparser.CreateBuffer("-")
	b.SetCurrentPosition(len("-"))
	o1 := CreateBinaryOperation(b.CurrentMatch(), &b.Code)
	b = aparser.CreateBuffer("4")
	b.SetCurrentPosition(len("4"))
	o2 := CreateDoubleOperand(b.CurrentMatch(), &b.Code)
	code := []interface{}{o0, o1, o2}
	n0 := createBinaryLeftAssocNode(code)
	n1 := createNode([]interface{}{n0})
	checkEvaluate(t, n1, -1.0)
}
