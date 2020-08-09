package ast

import (
	"reflect"
	"testing"
)

func TestToDouble(t *testing.T) {
	f, err := toDouble(3.14)
	if err != nil {
		t.Error("Could not convert 3.14 to double")
	}
	if f != 3.14 {
		t.Error("Converting 3.14 to double should be 1.0")
	}
}

func TestIsBool(t *testing.T) {
	if isBool(42) != false {
		t.Error("isBool(42) should be false")
	}
	if isBool(false) != true {
		t.Error("isBool(false) should be true")
	}
}

func TestIsString(t *testing.T) {
	if isString(42) != false {
		t.Error("isString(42) should be false")
	}
	if isString("foo") != true {
		t.Error("isString(foo) should be true")
	}
}

func Test_delete(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5}
	r := delete(a, 2)
	if !checkSliceEquals(r, []interface{}{0, 1, 3, 4, 5}) {
		t.Errorf("Result should be %v, actual %v", []interface{}{0, 1, 3, 4, 5}, r)
	}
}

func Test_deleteRange(t *testing.T) {
	a := []interface{}{0, 1, 2, 3, 4, 5}
	r := deleteRange(a, 1, 2)
	if !checkSliceEquals(r, []interface{}{0, 3, 4, 5}) {
		t.Errorf("Result should be %v, actual %v", []interface{}{0, 3, 4, 5}, r)
	}
}

func checkSliceEquals(x, y []interface{}) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if !reflect.DeepEqual(x[i], y[i]) {
			return false
		}
	}
	return true
}
