package aparser

import (
	"reflect"
	"testing"
)

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
